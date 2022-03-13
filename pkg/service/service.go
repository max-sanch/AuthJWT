package service

import (
	"github.com/max-sanch/AuthJWT"
	"github.com/max-sanch/AuthJWT/pkg/repository"
)

type Authentication interface {

}

type User interface {
	CreateUser(user auth_jwt.User) (string, error)
	DeleteUser(user auth_jwt.User) (string, error)
}

type Service struct {
	Authentication
	User
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		User: NewUserService(repos.User),
	}
}
