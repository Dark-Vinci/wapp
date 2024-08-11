package utils

import (
	"bytes"
	"context"
	"database/sql"
	"fmt"
	"io"
	"log"
	"os"
	"strings"

	"github.com/pressly/goose/v3"
	"github.com/rs/zerolog"

	"github.com/dark-vinci/wapp/backend/sdk/constants"
	"github.com/dark-vinci/wapp/backend/sdk/sdkerror"
)

type MigrationConfig struct {
	PgUser         string
	PgPassword     string
	PgHost         string
	PgPort         string
	PgExternalPort string
}

func Migration(ctx context.Context, logger *zerolog.Logger, e MigrationConfig, service string) error {
	_ = constants.GooseFlag.Parse(os.Args[1:])

	args := constants.GooseFlag.Args()

	if len(args) < 3 {
		constants.GooseFlag.Usage()
		return sdkerror.ErrNotEnoughArguments
	}

	command := args[1]

	connection := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
		e.PgUser,
		e.PgPassword,
		e.PgHost,
		e.PgExternalPort,
		service)

	db, err := sql.Open("postgres", connection)

	if err != nil {
		logger.Fatal().Err(err).Msgf(fmt.Sprintf("goose %v: %v", command, err))
		return sdkerror.ErrUnableToConnectToDB
	}

	defer func() {
		if err = db.Close(); err != nil {
			log.Fatalf("goose: failed to close DB: %v\n", err)
		}
	}()

	currentDir, _ := os.Getwd()

	var arguments []string
	if len(args) > 3 {
		arguments = append(arguments, args[3:]...)
	}

	logger.Info().Msgf(fmt.Sprintf("service(%s)::: running goose %s %v : args=%d", service, command, arguments, len(arguments)))

	migrationDirectory := fmt.Sprintf("%s/migrations", currentDir)

	err = goose.RunContext(ctx, command, db, migrationDirectory, arguments...)
	if err != nil {
		logger.Fatal().Err(err).Msgf(fmt.Sprintf("goose %v: %v", command, err))
	}

	// we want to grant required permissions and privileges after every up - run
	if command == constants.UP {
		err = runUpMigrationHook(db, os.Getenv(e.PgUser))
		if err != nil {
			logger.Err(err).Msg("runUpMigrationHook failed")
			return nil
		}

		logger.Info().Msg("runUpMigrationHook ran successfully")
	}

	return nil
}

func upMigrationHookScript(dbUser string) string {
	a := fmt.Sprintf("GRANT ALL PRIVILEGES ON ALL TABLES IN SCHEMA public TO %s;", dbUser)
	b := fmt.Sprintf("GRANT USAGE, SELECT ON SEQUENCE goose_db_version_id_seq TO %s;", dbUser)
	resp := fmt.Sprintf("%s\n%s", a, b)
	return resp
}

func runUpMigrationHook(db *sql.DB, dbUser string) error {
	script := upMigrationHookScript(dbUser)

	buf := bytes.NewBuffer(nil)
	r := strings.NewReader(script)

	_, err := io.Copy(buf, r)
	if err != nil {
		return fmt.Errorf("failed to read SQL script: %v", err)
	}

	_, err = db.Exec(buf.String())
	if err != nil {
		return fmt.Errorf("failed to execute SQL script: %v", err)
	}

	return nil
}
