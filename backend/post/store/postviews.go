package store

import (
	"github.com/rs/zerolog"
	"gorm.io/gorm"

	"github.com/dark-vinci/wapp/backend/post/connections"
	"github.com/dark-vinci/wapp/backend/sdk/constants"
)

type PostView struct {
	logger *zerolog.Logger
	db     *gorm.DB
}

//go:generate mockgen -source postviews.go -destination ./mock/postviews_mock.go -package mock PostViewDatabase
type PostViewDatabase interface{}

func NewPostView(conn *connections.DBConn) *PostViewDatabase {
	logger := conn.Log.With().Str(constants.PackageStrHelper, packageName).Logger()

	postView := &PostView{
		logger: &logger,
		db:     conn.Connection,
	}

	postViewOperations := PostViewDatabase(postView)

	logger.Info().Msg("successfully initialized post views database")

	return &postViewOperations
}
