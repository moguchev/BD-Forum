package forum

import (
	. "github.com/moguchev/BD-Forum/pkg/models"
)

type ServiceInterface interface {
	CreateUser(User) ([]User, error)
	GetUser(string) (User, error)
	UpdateUser(User) error
}
