package constants

import "flag"

var (
	UP        = "up"
	GooseFlag = flag.NewFlagSet("goose", flag.ExitOnError)
	DIR       = GooseFlag.String("dir", ".", "directory with migration files")
)
