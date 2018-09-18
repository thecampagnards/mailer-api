package config

import (
	"os"

	"github.com/globalsign/mgo"
)

// DB
type DB struct {
	Session *mgo.Session
}

// DoDial
func (db *DB) DoDial() (s *mgo.Session, err error) {
	return mgo.Dial(DBUrl())
}

// Name
func (db *DB) Name() string {
	return "mailer"
}

// DBUrl
func DBUrl() string {
	dburl := os.Getenv("MONGOHQ_URL")

	if dburl == "" {
		dburl = "localhost"
	}

	return dburl
}
