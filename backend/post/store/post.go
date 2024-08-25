package store

import (
	"github.com/rs/zerolog"
	"gorm.io/gorm"

	"github.com/dark-vinci/wapp/backend/post/connections"
	"github.com/dark-vinci/wapp/backend/sdk/constants"
)

type Post struct {
	logger *zerolog.Logger
	db     *gorm.DB
}

//go:generate mockgen -source post.go -destination ./mock/post_mock.go -package mock PostDatabase
type PostDatabase interface{}

func NewPost(conn *connections.DBConn) *PostDatabase {
	logger := conn.Log.With().Str(constants.PackageStrHelper, packageName).Logger()

	post := &Post{
		logger: &logger,
		db:     conn.Connection,
	}

	postOperations := PostDatabase(post)

	logger.Info().Msg("successfully initialized post database")

	return &postOperations
}
