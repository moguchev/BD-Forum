package repository

import (
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/jackc/pgtype"
	"github.com/jackc/pgx"
	"github.com/moguchev/BD-Forum/pkg/codes"
	"github.com/moguchev/BD-Forum/pkg/messages"
	. "github.com/moguchev/BD-Forum/pkg/models"
	"github.com/moguchev/BD-Forum/pkg/sql_queries"
)

func (r *Repository) CreateForum(nf NewForum) error {
	_, err := r.DbConn.Exec(sql_queries.InsertForum,
		nf.Slug, nf.Title, nf.User)

	if err != nil {
		if e, ok := err.(pgx.PgError); ok {
			switch e.Code {
			case codes.NotNullViolation, codes.ForeignKeyViolation:
				err = errors.New(messages.UserNotFound)
				break
			case codes.UniqueViolation:
				err = errors.New(messages.ForumAlreadyExists)
				break
			default:
				log.Println(e.Code)
			}
		}
		return err
	}

	return nil
}

func (r *Repository) GetForum(slug string) (Forum, error) {
	row := r.DbConn.QueryRowx(sql_queries.SelectForum, slug)

	var forum Forum
	err := row.StructScan(&forum)
	if err != nil {
		fmt.Println(err)
	}

	return forum, err
}

func (r *Repository) CreateThread(t Thread) (Thread, error) {
	var slug pgtype.Text
	time, _ := time.Parse(time.RFC3339Nano, t.Created)
	err := r.DbConn.QueryRow(sql_queries.InsertThread,
		t.Author, t.Forum, t.Message, time, t.Title, t.Slug).
		Scan(&t.Id, &slug, &t.Votes, &t.Forum)

	t.Slug = slug.String

	if err != nil {
		if e, ok := err.(pgx.PgError); ok {
			switch e.Code {
			case codes.ForeignKeyViolation, codes.NotNullViolation:
				if e.ConstraintName == "thread_author_fkey" {
					err = errors.New(messages.UserNotFound)
				} else {
					err = errors.New(messages.ForumNotFound)
				}
				break
			case codes.UniqueViolation:
				err = errors.New(messages.ThreadAlreadyExists)
				break
			default:
				log.Println(e.Code)
				err = errors.New(messages.ForumNotFound)
			}
		} else {
			err = errors.New(messages.ForumNotFound)
		}
	}

	return t, err
}

func (r *Repository) GetThreadBySlug(slug string) (Thread, error) {
	row := r.DbConn.QueryRowx(sql_queries.SelectThreadBySlug, slug)

	var t Thread
	err := row.StructScan(&t)
	if err != nil {
		fmt.Println(err)
	}

	return t, err
}

func (r *Repository) GetThreads(forum string) ([]Thread, error) {
	var threads []Thread

	rows, err := r.DbConn.Queryx(sql_queries.SelectThreads, forum)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		t := Thread{}
		err = rows.StructScan(&t)

		if err != nil {
			return threads, nil
		}

		if t.Id < 1 {
			return nil, nil
		}

		threads = append(threads, t)
	}
	if len(threads) == 0 {
		return nil, errors.New(messages.ForumNotFound)
	}
	return threads, nil
}
