package service

import (
	"github.com/max-sanch/AuthJWT"
	"github.com/max-sanch/AuthJWT/pkg/repository"
)

type UserService struct {
	repo repository.User
}

func NewUserService(repo repository.User)	*UserService {
	return &UserService{repo: repo}
}

func (u *UserService) CreateUser(user auth_jwt.User) (string, error) {
	return u.repo.CreateUser(user)
}

func (u *UserService) DeleteUser(user auth_jwt.User) (string, error) {
	return u.repo.DeleteUser(user)
}
