package forum

import (
	. "github.com/moguchev/BD-Forum/pkg/models"
)

type RepositoryInterface interface {
	// user section
	CreateUser(User) error
	UpdateUser(User) error
	GetUserByNickname(string) (User, error)
	GetUserByEmail(string) (User, error)
	FindUsers(string, string) ([]User, error)

	// service section
	Clear() error

	// forum section
	CreateForum(NewForum) error
	GetForum(string) (Forum, error)

	CreateThread(Thread) (Thread, error)
	GetThreadBySlug(string) (Thread, error)
	GetThreads(string) ([]Thread, error)

	InitDBSQL() error
}
