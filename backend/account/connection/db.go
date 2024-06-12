package connection

import (
	"fmt"
	"sync"

	"github.com/dark-vinci/isok"
	"github.com/rs/zerolog"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/dark-vinci/linkedout/backend/account/env"
)

const packageName string = "app.connection"

const SLAVE_COUNT = 4

type DBConn struct {
	Master *gorm.DB
	Slaves []*gorm.DB
	Log    *zerolog.Logger
}

func (db *DBConn) Close() {
	m, _ := db.Master.DB()
	m.Close()

	for _, v := range db.Slaves {
		m, _ := v.DB()
		m.Close()
	}
}

func NewDBConn(z zerolog.Logger, e *env.Environment) *DBConn {
	log := z.With().Str("KEY", packageName).Logger()

	c := make(chan *gorm.DB, 3)
	m := make(chan *gorm.DB)
	var wg sync.WaitGroup

	wg.Add(1)

	go func(wg *sync.WaitGroup, m chan<- *gorm.DB) {
		defer wg.Done()

		mRes := isok.ResultFun1(
			gorm.Open(
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
			),
		)

		if mRes.IsErr() {
			log.Fatal().Err(mRes.UnwrapErr())
			panic(mRes.UnwrapErr())
		}

		m <- mRes.Unwrap()
	}(&wg, m)

	for range SLAVE_COUNT {
		wg.Add(1)

		go func(wg *sync.WaitGroup, e *env.Environment, c chan<- *gorm.DB, log zerolog.Logger) {
			defer wg.Wait()

			mRes := isok.ResultFun1(
				gorm.Open(
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
				),
			)

			if mRes.IsErr() {
				log.Fatal().Err(mRes.UnwrapErr())
				panic(mRes.UnwrapErr())
			}

			c <- mRes.Unwrap()
		}(&wg, e, c, log)
	}

	go func() {
		wg.Wait()
		close(c)
	}()

	slaves := make([]*gorm.DB, 0)

	for r := range c {
		slaves = append(slaves, r)
	}

	ma := <-m

	return &DBConn{
		Slaves: slaves,
		Master: ma,
		Log:    &z,
	}
}
