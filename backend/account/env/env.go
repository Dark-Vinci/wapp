package env

import (
	"os"
	"strconv"

	"github.com/dark-vinci/linkedout/backend/sdk/constants"
	"github.com/dark-vinci/linkedout/backend/sdk/utils"
)

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

func (e *Environment) MigrationConfig() *utils.MigrationConfig {
	return &utils.MigrationConfig{
		PgUser:         e.PgMasterUser,
		PgPassword:     e.PgMasterPassword,
		PgHost:         e.PgMasterHost,
		PgPort:         e.PgMasterPort,
		PgExternalPort: e.PgMasterPort,
	}
}

func NewEnv() *Environment {
	p := os.Getenv(constants.ShouldMigrate)
	shouldMigrate, _ := strconv.ParseBool(p)

	return &Environment{
		AppPort:          os.Getenv(constants.AppPort),
		AppEnvironment:   os.Getenv(constants.AppPort),
		ShouldMigrate:    shouldMigrate,
		KafkaURL:         os.Getenv(constants.KafkaURL),
		PgSlaveHost:      os.Getenv(constants.PgSlaveHost),
		PgMasterHost:     os.Getenv(constants.PgMasterHost),
		PgMasterName:     os.Getenv(constants.PgMasterName),
		PgSlaveName:      os.Getenv(constants.PgSlaveName),
		PgMasterPort:     os.Getenv(constants.PgMasterPort),
		PgSlavePort:      os.Getenv(constants.PgSlavePort),
		PgMasterUser:     os.Getenv(constants.PgMasterUser),
		PgSlaveUser:      os.Getenv(constants.PgSlaveUser),
		PgMasterPassword: os.Getenv(constants.PgMasterPassword),
		PgSlavePassword:  os.Getenv(constants.PgSlavePassword),
	}
}
