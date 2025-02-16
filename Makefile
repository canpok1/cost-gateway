ifeq ($(ENV),test)
	ENV_FILE=.devcontainer/db.env.test
else ifeq ($(ENV),dev)
	ENV_FILE=.devcontainer/db.env.dev
endif

BUILD_OUTPUT_DIR=build/release
SERVER_MAIN_FILE=cmd/server/main.go
SERVER_BINARY=server

ifdef ENV_FILE
	MIGRATE_COMMAND=export $$(cat ${ENV_FILE} | xargs) && ./db/migrate.sh
	TBLS_COMMAND=export $$(cat ${ENV_FILE} | xargs) && tbls --config db/.tbls.yml
else
	MIGRATE_COMMAND=./db/migrate.sh
	TBLS_COMMAND=tbls --config db/.tbls.yml
endif

.PHONY: setup
setup:
	go install github.com/sqlc-dev/sqlc/cmd/sqlc@latest
	go install github.com/oapi-codegen/oapi-codegen/v2/cmd/oapi-codegen@latest

.PHONY: run
run:
	export $$(cat ${ENV_FILE} | xargs) \
	&& go run ${SERVER_MAIN_FILE}

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
	cd db && sqlc generate

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
	${TBLS_COMMAND} doc --rm-dist

.PHONY: db-lint
db-lint:
	${TBLS_COMMAND} lint

.PHONY: db-diff
db-diff:
	${TBLS_COMMAND} diff
