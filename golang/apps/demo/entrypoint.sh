#!/usr/bin/env bash
set -xe
go mod tidy
go mod download
go tool goose up
exec dotenvx run $@
