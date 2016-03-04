#!/bin/sh -eux

goimports -w ./*.go ./cmd/jwg/*.go
go tool vet .
golint .
go test ./...
