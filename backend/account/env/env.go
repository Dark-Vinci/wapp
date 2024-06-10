package env

type Environment struct {
	AppPort        string
	AppEnvironment string
	ShouldMigrate  bool
	PgPassword     string
	PgHost         string
	PgUser         string
	PgExternalPort string
}

func NewEnv() *Environment {
	return &Environment{
		AppPort:        "2020",
		AppEnvironment: "development",
		ShouldMigrate:  false,
	}
}
