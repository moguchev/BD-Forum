package service

import (
	. "github.com/moguchev/BD-Forum/pkg/models"
)

func (s Service) CreateUser(u User) ([]User, error) {
	existedUsers, _ := s.Repository.FindUsers(u.Nickname, u.Email)

	if len(existedUsers) > 0 {
		return existedUsers, nil
	}

	err := s.Repository.CreateUser(u)

	if err != nil {
		return nil, err
	}

	return nil, nil
}

func (s Service) UpdateUser(u User) error {
	err := s.Repository.UpdateUser(u)
	return err
}

func (s Service) GetUser(nickname string) (User, error) {
	user, err := s.Repository.GetUserByNickname(nickname)
	return user, err
}
