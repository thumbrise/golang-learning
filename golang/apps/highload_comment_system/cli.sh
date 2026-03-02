#!/usr/bin/env bash
set -xe
exec dotenvx run go run . $@
