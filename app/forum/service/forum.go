package service

import (
	"github.com/moguchev/BD-Forum/pkg/messages"
	. "github.com/moguchev/BD-Forum/pkg/models"
)

func (s Service) CreateForum(nf NewForum) (Forum, error) {
	err := s.Repository.CreateForum(nf)

	f := Forum{}
	if err != nil {
		switch err.Error() {
		case messages.UserNotFound:
			return f, err
		case messages.ForumAlreadyExists:
			f, _ := s.Repository.GetForum(nf.Slug)
			return f, err
		}
	}

	f.Slug = nf.Slug
	f.Title = nf.Title
	f.User = nf.User
	f.Posts = 0
	f.Threads = 0

	return f, nil
}

func (s Service) GetForum(slug string) (Forum, error) {
	forum, err := s.Repository.GetForum(slug)
	return forum, err
}

func (s Service) CreateThread(t Thread) (Thread, error) {
	thread, err := s.Repository.CreateThread(t)

	if err != nil {
		switch err.Error() {
		case messages.UserNotFound:
			break
		case messages.ForumNotFound:
			break
		case messages.ForumAlreadyExists:
			thread, _ = s.Repository.GetThreadBySlug(thread.Slug)
			break
		}
	}
	return thread, err
}

func (s Service) GetThreads(forum string) ([]Thread, error) {
	ths, err := s.Repository.GetThreads(forum)
	return ths, err
}
