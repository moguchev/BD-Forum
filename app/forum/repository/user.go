package repository

import (
	"errors"
	"fmt"

	"github.com/moguchev/BD-Forum/pkg/messages"
	. "github.com/moguchev/BD-Forum/pkg/models"
	"github.com/moguchev/BD-Forum/pkg/sql_queries"
)

func (r *Repository) CreateUser(u User) error {
	_, err := r.DbConn.Exec(sql_queries.InsertUser,
		u.About, u.Email, u.Fullname, u.Nickname)

	if err != nil {
		return fmt.Errorf(messages.UserAlreadyExists)
	}

	return nil
}

func (r *Repository) UpdateUser(u User) error {
	res, err := r.DbConn.Exec(sql_queries.UpdateUser,
		u.About, u.Email, u.Fullname, u.Nickname)

	if err != nil {
		return errors.New(messages.ConflictsInUserUpdate)
	}

	count, _ := res.RowsAffected()

	if count == 0 {
		return errors.New(messages.UserNotFound)
	}
	return nil
}

func (r *Repository) GetUserByNickname(nickname string) (User, error) {

	row := r.DbConn.QueryRowx(sql_queries.SelectUserByNickname, nickname)

	var user User
	err := row.StructScan(&user)
	if err != nil {
		fmt.Println(err)
	}
	return user, err
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

func (r *Repository) FindUsers(nickname string, email string) ([]User, error) {
	var users []User

	rows, err := r.DbConn.Queryx(sql_queries.SelectUsersByNicknameAndEmail, nickname, email)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		user := User{}
		err = rows.StructScan(&user)
		if err != nil {
			return users, err
		}
		users = append(users, user)
	}

	return users, nil
}
