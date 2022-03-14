package repository

import (
	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
	"github.com/max-sanch/AuthJWT"
)

type AuthMongo struct {
	db *mgo.Database
}

func NewAuthMongo(db *mgo.Database) *AuthMongo {
	return &AuthMongo{db: db}
}

func (a *AuthMongo) GetUser(guid string) (auth_jwt.User, error) {
	var user auth_jwt.User
	err := a.db.C(usersCollection).Find(bson.M{"guid": guid}).One(&user)
	return user, err
}

func (a *AuthMongo) SetRefresh(guid string, refreshToken string) error {
	refresh := auth_jwt.Refresh{
		GUID: guid,
		RefreshToken: refreshToken,
	}
	err := a.db.C(refreshCollection).Update(bson.M{"guid": guid}, &refresh)

	if err != nil && err.Error() != notFoundMessage {
		return err
	}

	if err != nil && err.Error() == notFoundMessage {
		if err := a.db.C(refreshCollection).Insert(&refresh); err != nil {
			return err
		}

		return nil
	}

	return nil
}

func (a *AuthMongo) GetRefresh(guid string) (string, error) {
	var refresh auth_jwt.Refresh
	err := a.db.C(refreshCollection).Find(bson.M{"guid": guid}).One(&refresh)
	return refresh.RefreshToken, err
}
