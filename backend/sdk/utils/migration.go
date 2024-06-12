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
	// "github.com/dark-vinci/isok"
	
	"github.com/dark-vinci/linkedout/backend/account/env"
	"github.com/dark-vinci/linkedout/backend/sdk/constants"
	"github.com/dark-vinci/linkedout/backend/sdk/sdkerror"
	"github.com/dark-vinci/linkedout/backend/sdk/isok"
)

func Migration(ctx context.Context, logger *zerolog.Logger, env *env.Environment, service string) error {
	_ = constants.GooseFlag.Parse(os.Args[1:])

	args := constants.GooseFlag.Args()

	if len(args) < 3 {
		constants.GooseFlag.Usage()
		return sdkerror.ErrNotEnoughArguments
	}

	command := args[1]

	connection := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
		os.Getenv(env.PgUser),
		os.Getenv(env.PgPassword),
		os.Getenv(env.PgHost),
		os.Getenv(env.PgExternalPort),
		service)
	
	dbR := isok.ResultFun1(sql.Open("postgres", connection))

	if dbR.IsErr() {
		logger.Fatal().Err(dbR.UnwrapErr()).Msgf(fmt.Sprintf("goose %v: %v", command, dbR.UnwrapErr()))
		return sdkerror.ErrUnableToConnectToDB
	}

	defer func() {
		if res := isok.ResultFun0(dbR.Unwrap().Close()); res.IsErr() {
			log.Fatalf("goose: failed to close DB: %v\n", res.UnwrapErr())
		}
	}()

	currentDir, _ := os.Getwd()

	var arguments []string
	if len(args) > 3 {
		arguments = append(arguments, args[3:]...)
	}

	logger.Info().Msgf(fmt.Sprintf("service(%s)::: running goose %s %v : args=%d", service, command, arguments, len(arguments)))

	migrationDirectory := fmt.Sprintf("%s/migrations", currentDir)

	gRes := isok.ResultFun0(goose.RunContext(ctx, command, dbR.Unwrap(), migrationDirectory, arguments...))
	if gRes.IsErr() {
		logger.Fatal().Err(gRes.UnwrapErr()).Msgf(fmt.Sprintf("goose %v: %v", command, gRes.UnwrapErr()))
	}

	// we want to grant required permissions and privileges after every up - run
	if command == constants.UP {
		resRes := isok.ResultFun0(runUpMigrationHook(dbR.Unwrap(), os.Getenv(env.PgUser)))
		if resRes.IsErr() {
			logger.Err(resRes.UnwrapErr()).Msg("runUpMigrationHook failed")
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

	cRes := isok.ResultFun1(io.Copy(buf, r))
	if cRes.IsErr() {
		return fmt.Errorf("failed to read SQL script: %v", cRes.UnwrapErr())
	}

	eRes := isok.ResultFun1(db.Exec(buf.String()))
	if eRes.IsErr() {
		return fmt.Errorf("failed to execute SQL script: %v", eRes.UnwrapErr())
	}

	return nil
}
