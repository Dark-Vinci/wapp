package store

import (
	"github.com/rs/zerolog"

	"github.com/dark-vinci/linkedout/backend/account/connection"
)

type UserDatabase interface{}

type User struct {
	log *zerolog.Logger
	db  *connection.DBConn
}

func NewUser(db *connection.DBConn) *UserDatabase {
	u := User{
		db:  db,
		log: db.Log,
	}

	uu := UserDatabase(u)

	return &uu
}
