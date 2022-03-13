package repository

import (
	"fmt"
	"github.com/globalsign/mgo"
)

const usersCollection = "users"

type Config struct {
	Host		string
	Port		string
	Username	string
	DBName		string
	Password	string
}

func NewMongoDB(cfg Config) (*mgo.Database, error) {
	session, err := mgo.Dial(fmt.Sprintf("mongodb://%s:%s@%s:%s/%s",
		cfg.Username, cfg.Password, cfg.Host, cfg.Port, cfg.DBName))
	if err != nil {
		return nil, err
	}

	session.SetMode(mgo.Monotonic, true)

	if err := session.Ping(); err != nil {
		return nil, err
	}

	db := session.DB(cfg.DBName)
	return db, nil
}
