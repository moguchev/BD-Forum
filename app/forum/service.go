package forum

import (
	. "github.com/moguchev/BD-Forum/pkg/models"
)

type ServiceInterface interface {
	CreateUser(NewUser, string) (User, error)
}
