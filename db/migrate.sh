#!/bin/bash
cd "$(dirname "$0")"

DDL_DIR=ddl
DB_URL="mysql://${MYSQL_USER}:${MYSQL_PASSWORD}@tcp(${MYSQL_HOST}:${MYSQL_PORT})/${MYSQL_DATABASE}"

echo work dir: $(pwd)
echo ddl dir: ${DDL_DIR}

if [ "$1" = "create" ]; then
    migrate create -ext sql -dir "${DDL_DIR}" -seq ${@:2}
else
    migrate -path ${DDL_DIR} -database "${DB_URL}" $@
fi
