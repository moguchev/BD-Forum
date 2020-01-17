package forum

import (
	. "github.com/moguchev/BD-Forum/pkg/models"
)

type RepositoryInterface interface {
	// User section
	CreateUser(User) error
	UpdateUser(User) error
	GetUserByNickname(string) (User, error)
	GetUserByEmail(string) (User, error)
	FindUsers(string, string) ([]User, error)

	InitDBSQL() error
}
