package forum

import (
	. "github.com/moguchev/BD-Forum/pkg/models"
)

type ServiceInterface interface {
	// user section
	CreateUser(User) ([]User, error)
	GetUser(string) (User, error)
	UpdateUser(User) error

	// service section
	Clear() error

	// forum section
	CreateForum(NewForum) (Forum, error)
	GetForum(string) (Forum, error)
	CreateThread(Thread) (Thread, error)
	GetThreads(string) ([]Thread, error)
}
