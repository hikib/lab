#! /usr/bin/bash

# sudo apt install postgresql
# podman pull docker.io/library/postgres
# podman pull docker.io/library/pgadmin4

podman run -d --name postgres \
    -e POSTGRES_USER=hikib \
    -e POSTGRES_PASSWORD=somepw \
    -e POSTGRES_DB=someapp \
    -p 5432:5432 \
    -v ./pgdata:/var/lib/postgresql/data \
    postgres

# psql --host=localhost --username=hikib --dbname=someapp
