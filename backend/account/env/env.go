package env

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
	KafkaURL         string
}

func NewEnv() *Environment {
	return &Environment{
		AppPort:        "2020",
		AppEnvironment: "development",
		ShouldMigrate:  false,
	}
}
