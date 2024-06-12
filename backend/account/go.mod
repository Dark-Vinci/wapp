module github.com/dark-vinci/linkedout/backend/account

replace github.com/dark-vinci/linkedout/backend/sdk => ../sdk

go 1.22.1

require (
	github.com/dark-vinci/isok v0.0.1
	github.com/dark-vinci/linkedout/backend/sdk v0.0.0-00010101000000-000000000000
	github.com/go-redis/redis/v8 v8.11.5
	github.com/google/uuid v1.6.0
	github.com/pressly/goose/v3 v3.20.0
	github.com/rs/zerolog v1.33.0
	github.com/segmentio/kafka-go v0.4.47
	google.golang.org/grpc v1.64.0
	gorm.io/driver/postgres v1.5.7
	gorm.io/gorm v1.25.10
)

require (
	github.com/cespare/xxhash/v2 v2.2.0 // indirect
	github.com/dgryski/go-rendezvous v0.0.0-20200823014737-9f7001d12a5f // indirect
	github.com/jackc/pgpassfile v1.0.0 // indirect
	github.com/jackc/pgservicefile v0.0.0-20240606120523-5a60cdf6a761 // indirect
	github.com/jackc/pgx/v5 v5.6.0 // indirect
	github.com/jackc/puddle/v2 v2.2.1 // indirect
	github.com/jinzhu/inflection v1.0.0 // indirect
	github.com/jinzhu/now v1.1.5 // indirect
	github.com/klauspost/compress v1.17.8 // indirect
	github.com/mattn/go-colorable v0.1.13 // indirect
	github.com/mattn/go-isatty v0.0.20 // indirect
	github.com/mfridman/interpolate v0.0.2 // indirect
	github.com/pierrec/lz4/v4 v4.1.21 // indirect
	github.com/sethvargo/go-retry v0.2.4 // indirect
	go.uber.org/multierr v1.11.0 // indirect
	golang.org/x/crypto v0.24.0 // indirect
	golang.org/x/net v0.26.0 // indirect
	golang.org/x/sync v0.7.0 // indirect
	golang.org/x/sys v0.21.0 // indirect
	golang.org/x/text v0.16.0 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20240610135401-a8a62080eff3 // indirect
	google.golang.org/protobuf v1.34.2 // indirect
)
