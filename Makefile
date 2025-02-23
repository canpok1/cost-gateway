ifeq ($(ENV),test)
	ENV_FILE=.devcontainer/db.env.test
else ifeq ($(ENV),dev)
	ENV_FILE=.devcontainer/db.env.dev
endif

BUILD_OUTPUT_DIR=build/release
SERVER_MAIN_FILE=cmd/server/main.go
SERVER_BINARY=server

MIGRATE_SH=./db/migrate.sh
TBLS_CONFIG=./db/.tbls.yml
ifdef ENV_FILE
	MIGRATE_COMMAND=export $$(cat ${ENV_FILE} | xargs) && ${MIGRATE_SH}
	TBLS_COMMAND=export $$(cat ${ENV_FILE} | xargs) && tbls --config ${TBLS_CONFIG}
else
	MIGRATE_COMMAND=${MIGRATE_SH}
	TBLS_COMMAND=tbls --config ${TBLS_CONFIG}
endif

OPENAPI_YML_URL=https://raw.githubusercontent.com/canpok1/cost-gateway/refs/heads/main/openapi/openapi.yml

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

.PHONY: generate-code
generate-code:
	go generate ./...
	cd db && sqlc generate

.PHONY: generate-api-doc
generate-api-doc:
	mkdir -p docs/api
	curl -L https://github.com/swagger-api/swagger-ui/archive/refs/tags/v5.19.0.zip -o docs/api/swagger-ui.zip
	unzip docs/api/swagger-ui.zip -d docs/api
	cp -R docs/api/swagger-ui-5.19.0/dist docs/api/
	sed -i 's@url:.*@url: "${OPENAPI_YML_URL}",@g' ./docs/api/dist/swagger-initializer.js
	rm -r docs/api/swagger-ui-5.19.0
	rm docs/api/swagger-ui.zip

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
	${MIGRATE_COMMAND} create ${name}

.PHONY: db-doc
db-doc:
	${TBLS_COMMAND} doc --rm-dist

.PHONY: db-lint
db-lint:
	${TBLS_COMMAND} lint

.PHONY: db-diff
db-diff:
	${TBLS_COMMAND} diff
