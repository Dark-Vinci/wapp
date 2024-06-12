package env

import (
	"os"
	"strconv"
)

const packageName = "account.env"

type Environment struct {
	AppPort          string
	AppEnvironment   string
	ShouldMigrate    bool
	PgMasterPassword string
	PgMasterHost     string
	PgMasterUser     string
	PgMasterPort     string
	PgSlavePassword  string
	PgSlaveHost      string
	PgSlaveUser      string
	PgSlavePort      string
	PgMasterName     string
	PgSlaveName      string
	KafkaURL         string
}

const (
	AppPort          string = "APP_PORT"
	KafkaURL         string = "KAFKA_URL"
	PgMasterHost     string = "PG_MASTER_HOST"
	PgMasterPort     string = "PG_MASTER_PORT"
	PgMasterUser     string = "PG_MASTER_USER"
	PgMasterPassword string = "PG_MASTER_PASSWORD"
	PgMasterName     string = "PG_MASTER_NAME"
	PgSlaveHost      string = "PG_SLAVE_HOST"
	PgSlavePort      string = "PG_SLAVE_PORT"
	PgSlaveUser      string = "PG_SLAVE_USER"
	PgSlavePassword  string = "PG_SLAVE_PASSWORD"
	PgSlaveName      string = "PG_SLAVE_NAME"
	ShouldMigrate    string = "SHOULD_MIGRATE"
)

func NewEnv() *Environment {
	p := os.Getenv(ShouldMigrate)
	shouldMigrate, _ := strconv.ParseBool(p)

	return &Environment{
		AppPort:          os.Getenv(AppPort),
		AppEnvironment:   os.Getenv(AppPort),
		ShouldMigrate:    shouldMigrate,
		KafkaURL:         os.Getenv(KafkaURL),
		PgSlaveHost:      os.Getenv(PgSlaveHost),
		PgMasterHost:     os.Getenv(PgMasterHost),
		PgMasterName:     os.Getenv(PgMasterName),
		PgSlaveName:      os.Getenv(PgSlaveName),
		PgMasterPort:     os.Getenv(PgMasterPort),
		PgSlavePort:      os.Getenv(PgSlavePort),
		PgMasterUser:     os.Getenv(PgMasterUser),
		PgSlaveUser:      os.Getenv(PgSlaveUser),
		PgMasterPassword: os.Getenv(PgMasterPassword),
		PgSlavePassword:  os.Getenv(PgSlavePassword),
	}
}
