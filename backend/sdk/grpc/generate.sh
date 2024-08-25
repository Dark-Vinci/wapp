#!/usr/bin/env bash
set -euxo pipefail

# get gRPC dependencies
# setting to v1.57.0 to avert deployment failure
go get google.golang.org/grpc@v1.57.0

# install tools
go install github.com/golang/protobuf/protoc-gen-go

# gen for the models
#protoc --go_out=plugins=grpc:. --go_opt=paths=source_relative ./models/*.proto
# gen for the account service
protoc --go_out=plugins=grpc:. --go_opt=paths=source_relative ./account/*.proto
# gen for the deposit service
protoc --go_out=plugins=grpc:. --go_opt=paths=source_relative ./chat/*.proto
# gen for the withdrawal service
protoc --go_out=plugins=grpc:. --go_opt=paths=source_relative ./media/*.proto
## gen for the banking service
protoc --go_out=plugins=grpc:. --go_opt=paths=source_relative ./post/*.proto
