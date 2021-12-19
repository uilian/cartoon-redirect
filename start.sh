#!/bin/bash
# nodemon --exec go run . --ext go
# go test -race -v ./...
# go build -gcflags "-m -m"
# go build -gcflags "-bench bench_result.txt -m -m" ./cmd/web/api.go
go run ./cmd/api/api.go