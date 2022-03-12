package service

import "github.com/max-sanch/AuthJWT/pkg/repository"

type Authentication interface {

}

type User interface {

}

type Service struct {
	Authentication
	User
}

func NewService(repos *repository.Repository) *Service {
	return &Service{}
}
