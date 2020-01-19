package repository

import (
	"database/sql"
	"errors"
	"log"
	"strconv"
	"sync"
	"sync/atomic"
	"time"

	"github.com/moguchev/BD-Forum/pkg/messages"
	. "github.com/moguchev/BD-Forum/pkg/models"
	"github.com/moguchev/BD-Forum/pkg/mytools"
	"github.com/moguchev/BD-Forum/pkg/sql_queries"
)

var mutexMapMutex sync.Mutex

func init() {
	mutexMapMutex = sync.Mutex{}
}

func (r *Repository) GetThreadId(slugOrId string) (int64, error) {
	var id int64 = 0
	_, err := strconv.ParseInt(slugOrId, 10, 64)
	if err != nil {
		// slug
		err := r.DbConn.QueryRowx(sql_queries.SelectThreadIdBySlug, slugOrId).Scan(&id)
		if err != nil {
			return id, errors.New(messages.ThreadNotFound)
		}
	} else {
		// id
		err := r.DbConn.QueryRowx(sql_queries.SelectThreadIdById, slugOrId).Scan(&id)
		if err != nil {
			return id, errors.New(messages.ThreadNotFound)
		}
	}
	return id, nil
}

func (r *Repository) GetThreadForumSlug(threadId int64) (string, error) {
	var forum string
	err := r.DbConn.QueryRowx(`SELECT forum FROM threads WHERE id=$1`, threadId).Scan(&forum)
	if err != nil {
		return forum, errors.New(messages.ThreadNotFound)
	}
	return forum, nil
}

func (r *Repository) CreatePostsByPacket(threadId int64, forumSLug string, posts []Post, created time.Time) ([]Post, error) {
	var params []interface{}
	for _, post := range posts {
		var parent sql.NullInt64
		parent.Int64 = post.Parent
		if post.Parent != 0 {
			parent.Valid = true
		}
		params = append(params, post.Author, post.Message, parent, threadId, created, forumSLug)
	}

	query := `INSERT INTO posts (author, message, parent, thread, created, forum) VALUES `
	postfix := `RETURNING forum, id, created`

	query = mytools.CreatePacketQuery(query, 6, len(posts), postfix)

	rows, err := r.DbConn.Queryx(query, params...)
	defer rows.Close()

	if err != nil || (rows != nil && rows.Err() != nil) {
		return posts, err
	}
	i := 0
	for rows.Next() {
		var created time.Time
		err := rows.Scan(&(posts[i].Forum), &(posts[i].Id), &(created))
		if err != nil {
			return posts, err
		}
		posts[i].Created = created.Format(time.RFC3339Nano)
		posts[i].IsEdited = false
		posts[i].Thread = int32(threadId)

		log.Println(posts)
		i++
	}

	var cnt int64
	if i == 0 && len(posts) > 0 {
		// looking for exact error
		if row := r.DbConn.QueryRowx(`SELECT count(id) from threads WHERE id=$1;`, threadId); row.Scan(&cnt) != nil || cnt == 0 {
			return posts, errors.New(messages.ThreadNotFound)
		} else if row := r.DbConn.QueryRowx(`SELECT COUNT(nickname) FROM users WHERE nickname=$1`, posts[0].Author); row.Scan(&cnt) != nil || cnt == 0 {
			return posts, errors.New(messages.ThreadNotFound)
		} else {
			return posts, errors.New(messages.ParentNotFound)
		}
	}
	return posts, nil
}

func (r *Repository) UpdateForumPosts(forum string, numPosts int) error {
	query := `UPDATE ForumPosts SET posts = posts + $2 WHERE forum = $1;`
	_, err := r.DbConn.Exec(query, forum, numPosts)
	if err == nil {
		atomic.AddInt32(&accessToForumPosts.numberOfNewUpdates, 1)
	}
	return err
}

func (r *Repository) InsertUsersToUsersInForum(users map[string]bool, forum string) error {
	prefix := `INSERT INTO UsersInForum(nickname, forum) VALUES `
	postfix := `ON CONFLICT DO NOTHING`
	query := mytools.CreatePacketQuery(prefix, 2, len(users), postfix)

	params := make([]interface{}, 0, len(users))
	for key := range users {
		params = append(params, key, forum)
	}

	mutexMapMutex.Lock()
	defer mutexMapMutex.Unlock()

	_, err := r.DbConn.Exec(query, params...)
	return err
}
