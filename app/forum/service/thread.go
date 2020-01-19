package service

import (
	"errors"
	"log"
	"math"
	"time"

	"github.com/moguchev/BD-Forum/pkg/messages"
	. "github.com/moguchev/BD-Forum/pkg/models"
)

const packageSize = 30

func (s Service) CreatePosts(slugOrId string, posts []Post) ([]Post, error) {
	id, err := s.Repository.GetThreadId(slugOrId)

	if err != nil {
		return nil, errors.New(messages.ThreadNotFound)
	}

	if len(posts) == 0 {
		return posts, nil
	}

	forum, err := s.Repository.GetThreadForumSlug(id)
	if err != nil {
		log.Printf(err.Error())
		return posts, err
	}

	userList := make(map[string]bool)
	postPacketSize := packageSize
	created := time.Now()
	for i := 0; i < len(posts); i += postPacketSize {
		currentPacket := posts[i:int(math.Min(float64(i+postPacketSize), float64(len(posts))))]
		currentPacket, err = s.Repository.CreatePostsByPacket(id, forum, currentPacket, created)
		if err != nil {
			return posts, err
		}
		for j, post := range currentPacket {
			posts[i+j] = post
			userList[post.Author] = true
		}
	}

	err = s.Repository.UpdateForumPosts(forum, len(posts))
	if err != nil {
		log.Println("UpdateForumPosts", err.Error())
		return posts, err
	}

	err = s.Repository.InsertUsersToUsersInForum(userList, forum)
	if err != nil {
		log.Println("InsertUsersToUsersInForum", err.Error())
		return posts, errors.New(messages.ThreadNotFound)
	}
	return posts, nil
}
