#!/bin/bash
# Wipes the database's data directory before delegating to the image's own
# entrypoint, so /docker-entrypoint-initdb.d scripts (schema + fake seed
# data) run again on every container start, not just on first creation.
set -e

data_dir="$1"
shift

rm -rf "${data_dir:?}"/*
exec docker-entrypoint.sh "$@"
