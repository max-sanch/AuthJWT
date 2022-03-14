package service

import (
	"github.com/max-sanch/AuthJWT"
	"github.com/max-sanch/AuthJWT/pkg/repository"
)

type Authentication interface {
	GenerateTokens(string) (map[string]string, error)
	RefreshTokens(string, string) (map[string]string, error)
}

type User interface {
	CreateUser(auth_jwt.User) (string, error)
	DeleteUser(auth_jwt.User) (string, error)
}

type Service struct {
	Authentication
	User
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authentication: NewAuthService(repos.Authentication),
		User: NewUserService(repos.User),
	}
}
