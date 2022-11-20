#!/usr/bin/env bash
# uses podman to run a postgresql database for local testing/development
# db is exposed on the localhost:5432, the default database created is called ms_mvc

PG_PASS="postgres"
PG_USER="ms_mvc"
PG_ADDRESS="127.0.0.1:5432"
PG_DB="ms_mvc"
PG_IMAGE="postgres"

# use values above to spawn a container
podman run --name postgres  \
    -p ${PG_ADDRESS}:5432/tcp \
    -e POSTGRES_PASSWORD=${PG_PASSWORD} \ 
    -e POSTGRES_USER=${PG_USER}  \
    -e POSTGRES_DB=${PG_DB}  \
    -d ${PG_IMAGE}
