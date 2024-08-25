package connections

import (
	"fmt"

	"github.com/rs/zerolog"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/plugin/dbresolver"

	"github.com/dark-vinci/wapp/backend/post/env"
	"github.com/dark-vinci/wapp/backend/sdk/constants"
)

type DBConn struct {
	Connection *gorm.DB
	Log        *zerolog.Logger
}

//func GetMockConnection(t *testing.T) (sqlmock.Sqlmock, *DBConn) {
//	var (
//		mock sqlmock.Sqlmock
//		db   *gorm.DB
//		err  error
//	)
//
//	db, mock, err := gorm_sqlmock
//}

func NewDBConn(z zerolog.Logger, e *env.Environment) *DBConn {
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

	return &DBConn{
		Connection: db,
		Log:        &z,
	}
}

func (db *DBConn) Close() {
	m, _ := db.Connection.DB()
	err := m.Close()

	if err != nil {
		return
	}
}
