package store

import (
	"fmt"
	"os"

	"github.com/rs/zerolog"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/dark-vinci/wapp/backend/media/env"
	"github.com/dark-vinci/wapp/backend/sdk/constants"
)

const packageName = "media.store"

type Store struct {
	Connection *gorm.DB
	Logger     *zerolog.Logger
}

func New(z zerolog.Logger, e *env.Environment) *Store {
	logger := z.With().Str(constants.PackageStrHelper, packageName).Logger()
	db, err := gorm.Open(
		postgres.Open(
			fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=%v",
				e.PgMasterHost,
				e.PgMasterPort,
				e.PgMasterUser,
				e.PgMasterName,
				true,
			),
		),
	)

	if err != nil {
		logger.Fatal().Err(err).Msg("failed to connect to database")
		panic(err)
	}

	return &Store{
		Connection: db,
		Logger:     &logger,
	}
}

// NewFromDB created a new storage with just the database reference passed in
func NewFromDB(db *gorm.DB) *Store {
	logger := zerolog.New(os.Stderr).With().Timestamp().Logger()
	storeLog := logger.With().Str(constants.PackageStrHelper, packageName).Logger()

	return &Store{
		Connection: db,
		Logger:     &storeLog,
	}
}

func (s *Store) Close() {
	m, _ := s.Connection.DB()
	err := m.Close()

	if err != nil {
		return
	}
}
