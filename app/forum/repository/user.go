package repository

import (
	"fmt"

	"github.com/moguchev/BD-Forum/pkg/messages"
	. "github.com/moguchev/BD-Forum/pkg/models"
	"github.com/moguchev/BD-Forum/pkg/sql_queries"
)

func (r *Repository) CreateUser(user NewUser, nickname string) error {
	var id int
	err := r.DbConn.QueryRow(sql_queries.InsertUser, user.About,
		user.Email, user.Fullname, nickname).Scan(&id)

	if err != nil {
		fmt.Println(err.Error())
		return fmt.Errorf(messages.UserAlreadyExists)
	}

	fmt.Println(id)
	return nil
}

func (r *Repository) GetUserByNickname(nickname string) (User, error) {

	row := r.DbConn.QueryRowx(sql_queries.SelectUserByNickname, nickname)

	var user User
	err := row.StructScan(&user)
	if err != nil {
		fmt.Println(err)
		return user, err
	}

	return user, nil
}

func (r *Repository) GetUserByEmail(email string) (User, error) {

	row := r.DbConn.QueryRowx(sql_queries.SelectUserByEmail, email)

	var user User
	err := row.StructScan(&user)
	if err != nil {
		fmt.Println(err)
		return user, err
	}

	return user, nil
}
