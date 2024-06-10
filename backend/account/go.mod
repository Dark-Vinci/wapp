module github.com/dark-vinci/linkedout/backend/account

replace github.com/dark-vinci/linkedout/backend/sdk => ../sdk

go 1.22.1

require (
	github.com/dark-vinci/isok v0.0.0-20240610125516-bfbad745e1f9
	github.com/dark-vinci/linkedout/backend/sdk v0.0.0-00010101000000-000000000000
	github.com/pressly/goose/v3 v3.19.2
	github.com/rs/zerolog v1.32.0
	google.golang.org/grpc v1.62.1
)

require (
	github.com/golang/protobuf v1.5.3 // indirect
	github.com/mattn/go-colorable v0.1.13 // indirect
	github.com/mattn/go-isatty v0.0.20 // indirect
	github.com/mfridman/interpolate v0.0.2 // indirect
	github.com/sethvargo/go-retry v0.2.4 // indirect
	go.uber.org/multierr v1.11.0 // indirect
	golang.org/x/net v0.20.0 // indirect
	golang.org/x/sync v0.6.0 // indirect
	golang.org/x/sys v0.18.0 // indirect
	golang.org/x/text v0.14.0 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20240318140521-94a12d6c2237 // indirect
	google.golang.org/protobuf v1.33.0 // indirect
)
