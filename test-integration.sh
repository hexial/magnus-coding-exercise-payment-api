#!/bin/sh
mkdir -p tmp
source ./clean.sh
source ./build.sh
source ./run.sh
docker run --rm \
	--name backend_test_integration \
    -v $(pwd)/backend:$GOPATH/src/backend \
	-v $(pwd)/tmp/integration_go_pkg_mod:/go/pkg/mod \
    -w=$GOPATH/src/backend \
	--network="container:backend" \
	-e "DB_HOST=db" \
	-e "DB_NAME=payment-api" \
	-e "DB_USER=backend" \
	-e "DB_PASSWORD=password" \
    golang:1.12.1 \
    go test -tags=integration ./...
