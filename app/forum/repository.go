package forum

import (
	. "github.com/moguchev/BD-Forum/pkg/models"
)

type RepositoryInterface interface {
	// User section
	CreateUser(NewUser, string) error
	GetUserByNickname(string) (User, error)
	GetUserByEmail(string) (User, error)

	InitDBSQL() error
}
