package store

import (
	"fmt"
	"os"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/rs/zerolog"
	"github.com/stretchr/testify/require"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/plugin/dbresolver"

	"github.com/dark-vinci/wapp/backend/account/env"
	"github.com/dark-vinci/wapp/backend/sdk/constants"
	"github.com/dark-vinci/wapp/backend/sdk/utils/gorm_sqlmock"
)

const packageName = "account.store"

type Store struct {
	Connection *gorm.DB
	Log        *zerolog.Logger
}

func NewStore(z zerolog.Logger, e *env.Environment) *Store {
	log := z.With().Str(constants.PackageStrHelper, packageName).Logger()

	db, err := gorm.Open(
		postgres.Open(
			fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable TimeZone=Africa/Lagos",
				e.PgMasterHost,
				e.PgMasterPort,
				e.PgMasterUser,
				e.PgMasterName,
				e.PgMasterPassword,
			),
		),
		&gorm.Config{},
	)

	if err != nil {
		log.Fatal().Err(err).Msg("failed to connect to master databases")
		panic(err)
	}

	err = db.Use(dbresolver.Register(
		dbresolver.Config{
			Replicas: []gorm.Dialector{mysql.Open("db3_dsn"), mysql.Open("db4_dsn")},
			Policy:   dbresolver.RandomPolicy{},
		}))

	if err != nil {
		log.Fatal().Err(err).Msg("failed to connect to slave databases")
		panic(err)
	}

	return &Store{
		Connection: db,
		Log:        &z,
	}
}

// GetConnection helper for tests/mock
func GetConnection(t *testing.T) (sqlmock.Sqlmock, *Store) {
	var (
		mock sqlmock.Sqlmock
		db   *gorm.DB
		err  error
	)

	db, mock, err = gorm_sqlmock.New(gorm_sqlmock.Config{
		Config:     &gorm.Config{},
		DriverName: "postgres",
		DSN:        "mock",
	})

	require.NoError(t, err)

	return mock, NewFromDB(db)
}

// NewFromDB created a new storage with just the database reference passed in
func NewFromDB(db *gorm.DB) *Store {
	logger := zerolog.New(os.Stderr).With().Timestamp().Logger()
	storeLog := logger.With().Str(constants.PackageStrHelper, packageName).Logger()

	return &Store{
		Connection: db,
		Log:        &storeLog,
	}
}

func (db *Store) Close() {
	m, _ := db.Connection.DB()
	err := m.Close()

	if err != nil {
		return
	}
}
