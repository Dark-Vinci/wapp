package store

import (
	"context"
	"errors"

	"github.com/rs/zerolog"
	"gorm.io/gorm"

	"github.com/dark-vinci/wapp/backend/account/connection"
	"github.com/dark-vinci/wapp/backend/sdk/constants"
	"github.com/dark-vinci/wapp/backend/sdk/models/account"
	"github.com/dark-vinci/wapp/backend/sdk/sdkerror"
	"github.com/dark-vinci/wapp/backend/sdk/utils"
)

type LastSeen struct {
	logger *zerolog.Logger
	db     *gorm.DB
}

//go:generate mockgen -source lastseen.go -destination ./mock/lastseen_mock.go -package mock LastSeenDatabase
type LastSeenDatabase interface {
	Create(ctx context.Context, lastSeen account.LastSeen) (*account.LastSeen, error)
}

func NewLastSeen(conn *connection.DBConn) *LastSeenDatabase {
	logger := conn.Log.With().Str(constants.PackageStrHelper, packageName).Logger()

	lastSeen := &LastSeen{
		logger: &logger,
		db:     conn.Connection,
	}

	lastSeenOperation := LastSeenDatabase(lastSeen)

	return &lastSeenOperation
}

func (ls *LastSeen) Create(ctx context.Context, lastSeen account.LastSeen) (*account.LastSeen, error) {
	log := ls.logger.With().
		Str(constants.MethodStrHelper, "lastSeen.Create").
		Str(constants.RequestID, utils.GetRequestID(ctx)).Logger()

	log.Info().Msg("Got a request to create a users last seen")

	if err := ls.db.WithContext(ctx).Model(&account.LastSeen{}).Create(&lastSeen).Error; err != nil {
		log.Err(err).Msg("Failed to create last seen")

		if errors.Is(err, gorm.ErrDuplicatedKey) {
			return nil, sdkerror.ErrDuplicateKey
		}

		return nil, sdkerror.ErrRecordCreation
	}

	return &lastSeen, nil
}
