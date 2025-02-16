#!/bin/sh
cd "$(dirname "$0")"

DDL_DIR=ddl
DB_URL="mysql://${MYSQL_USER}:${MYSQL_PASSWORD}@tcp(${MYSQL_HOST}:${MYSQL_PORT})/${MYSQL_DATABASE}"

migrate -path ${DDL_DIR} -database "${DB_URL}" $@
