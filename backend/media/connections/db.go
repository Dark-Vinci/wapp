package connections

import (
	"fmt"

	"github.com/rs/zerolog"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/plugin/dbresolver"

	"github.com/dark-vinci/wapp/backend/media/env"
	"github.com/dark-vinci/wapp/backend/sdk/constants"
)

type DBConn struct {
	Connection *gorm.DB
	Logger     *zerolog.Logger
}

func NewDBConn(z zerolog.Logger, e *env.Environment) *DBConn {
	log := z.With().Str(constants.PackageStrHelper, packageName).
		Str(constants.FunctionNameHelper, "NewDBConn").
		Logger()

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
		log.Fatal().Err(err).Msg("Failed to connect to master database")
		panic(err)
	}

	err = db.Use(dbresolver.Register(
		dbresolver.Config{
			Replicas: []gorm.Dialector{mysql.Open("db3_dsn"), mysql.Open("db4_dsn")},
			Policy:   dbresolver.RandomPolicy{},
		}))

	if err != nil {
		log.Fatal().Err(err).Msg("Failed to connect to master database")
		panic(err)
	}

	return &DBConn{
		Connection: db,
		Logger:     &z,
	}
}

func (db *DBConn) Close() {
	m, _ := db.Connection.DB()
	err := m.Close()

	if err != nil {
		return
	}
}
