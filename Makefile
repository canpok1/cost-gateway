ifeq ($(ENV),test)
	include .devcontainer/db.env.test
else
	include .devcontainer/db.env.dev
endif

BUILD_OUTPUT_DIR=build/release
SERVER_MAIN_FILE=cmd/server/main.go
SERVER_BINARY=server
DB_DDL_DIR=db/ddl
DB_DOC_DIR=docs/database
DB_URL_FOR_MIGRATE=mysql://${MYSQL_USER}:${MYSQL_PASSWORD}@tcp(${MYSQL_HOST}:${MYSQL_PORT})/${MYSQL_DATABASE}
DB_URL_FOR_TBLS=mysql://${MYSQL_USER}:${MYSQL_PASSWORD}@${MYSQL_HOST}:${MYSQL_PORT}/${MYSQL_DATABASE}

MIGRATE_COMMAND=migrate -path ${DB_DDL_DIR} -database "${DB_URL_FOR_MIGRATE}"

.PHONY: setup
setup:
	go install github.com/sqlc-dev/sqlc/cmd/sqlc@latest
	go install -tags 'mysql' github.com/golang-migrate/migrate/v4/cmd/migrate@latest
	go install github.com/k1LoW/tbls@latest
	go install github.com/oapi-codegen/oapi-codegen/v2/cmd/oapi-codegen@latest

.PHONY: run
run:
	go run ${SERVER_MAIN_FILE}

.PHONY: build
build:
	go build -o ${BUILD_OUTPUT_DIR}/${SERVER_BINARY} ${SERVER_MAIN_FILE}

.PHONY: clean
clean:
	go clean
	rm -f ${BUILD_OUTPUT_DIR}/${SERVER_BINARY}

.PHONY: test
test:
	go test -v ./...

.PHONY: generate
generate:
	go generate ./...

.PHONY: migrate
migrate:
	${MIGRATE_COMMAND} ${options}

.PHONY: migrate-version
migrate-version:
	${MIGRATE_COMMAND} version

.PHONY: migrate-up-all
migrate-up-all:
	${MIGRATE_COMMAND} up

.PHONY: migrate-up-one
migrate-up-one:
	${MIGRATE_COMMAND} up 1

.PHONY: migrate-down-one
migrate-down-one:
	${MIGRATE_COMMAND} down 1

.PHONY: migrate-force-v
migrate-force-v:
	${MIGRATE_COMMAND} force ${v}

.PHONY: migrate-create
migrate-create:
	${MIGRATE_COMMAND} create -ext sql -dir "${DB_DDL_DIR}" -seq ${name}

.PHONY: db-doc
db-doc:
	tbls doc "${DB_URL_FOR_TBLS}" "${DB_DOC_DIR}" --rm-dist

.PHONY: db-lint
db-lint:
	tbls lint "${DB_URL_FOR_TBLS}"

.PHONY: db-diff
db-diff:
	tbls diff "${DB_URL_FOR_TBLS}" "${DB_DOC_DIR}"
