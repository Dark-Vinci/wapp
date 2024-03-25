package models

type Env struct {
	PgUser         string
	PgPassword     string
	PgHost         string
	PgExternalPort string
	ShouldMigrate  bool
	AppPort        string
}
