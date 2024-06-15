package constants

import "flag"

var (
	UP        = "up"
	GooseFlag = flag.NewFlagSet("goose", flag.ExitOnError)
	DIR       = GooseFlag.String("dir", ".", "directory with migration files")
)

const (
	MethodStrHelper     = "METHOD_NME"
	PackageStrHelper    = "PACKAGE_NAME"
	FunctionInputHelper = "FUNCTION_INPUT"
	TimeZone            = "Africa/Lagos"
	FunctionNameHelper = "FUNCTION_NAME"
)

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
