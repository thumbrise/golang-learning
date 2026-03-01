#!/usr/bin/env bash
set -xe
go mod download
go tool goose up
exec dotenvx run $@
