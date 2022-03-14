package repository

import (
	"github.com/globalsign/mgo"
	"github.com/max-sanch/AuthJWT"
)

type Authentication interface {
	GetUser(string) (auth_jwt.User, error)
	SetRefresh(string, string) error
	GetRefresh(string) (string, error)
}

type User interface {
	CreateUser(auth_jwt.User) (string, error)
	DeleteUser(auth_jwt.User) (string, error)
}

type Repository struct {
	Authentication
	User
}

func NewRepository(db *mgo.Database) *Repository {
	return &Repository{
		Authentication: NewAuthMongo(db),
		User: NewUserMongo(db),
	}
}