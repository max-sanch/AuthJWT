package repository

import (
	"errors"
	"github.com/globalsign/mgo"
	"github.com/max-sanch/AuthJWT"
)

type UserMongo struct {
	db *mgo.Database
}

func NewUserMongo(db *mgo.Database) *UserMongo {
	return &UserMongo{db: db}
}

func (u *UserMongo) CreateUser(user auth_jwt.User) (string, error) {
	var result interface{}
	if err := u.db.C(usersCollection).Find(&user).One(&result); err != nil && err.Error() != "not found" {
		return "", err
	}

	if result != nil {
		return "", errors.New("user with this guid exists")
	}

	if err := u.db.C(usersCollection).Insert(&user); err != nil {
		return "", err
	}

	return user.GUID, nil
}

func (u *UserMongo) DeleteUser(user auth_jwt.User) (string, error) {
	if err := u.db.C(usersCollection).Remove(&user); err != nil {
		return "", err
	}

	return user.GUID, nil
}
