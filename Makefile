# * ==SCRIPTS==
.PHONY: build cli tcp web

build:
	./scripts/build.sh $(ARGS)

cli:
	./scripts/cli.sh $(ARGS)

tcp:
	./scripts/tcp.sh

web:
	./scripts/web.sh


# * ==DOCKER COMPOSE==
.PHONY: compose-cli compose-build-cli compose-tcp compose-build-tcp compose-web compose-build-web

# CLI
compose-cli:
	docker compose up cli $(ARGS)

compose-build-cli:
	docker compose up --build cli $(ARGS)

# TCP
compose-tcp:
	docker compose up -d tcp

compose-build-tcp:
	docker compose up -d --build tcp

# WEB
compose-web:
	docker compose up -d web database

compose-build-web:
	docker compose up -d --build web database


# * ==DEV OPS==
.PHONY: test test-verbose audit

test:
	go test -cover -coverprofile=coverage.out -race ./src/...

test-verbose:
	go test -cover -coverprofile=coverage.out -race -v ./src/...

audit:
	govulncheck ./...