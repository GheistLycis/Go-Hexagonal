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
.PHONY: compose-up-cli compose-build-cli compose-up-tcp compose-build-tcp compose-up-web compose-build-web

# CLI
compose-up-cli:
	docker compose up cli $(ARGS)

compose-build-cli:
	docker compose up --build cli $(ARGS)

# TCP
compose-up-tcp:
	docker compose up -d tcp

compose-build-tcp:
	docker compose up -d --build tcp

# WEB
compose-up-web:
	docker compose up -d web database

compose-build-web:
	docker compose up -d --build web database


# * ==TESTING==
.PHONY: test test-verbose

test:
	go test -cover -coverprofile=coverage.out -race ./src/...

test-verbose:
	go test -cover -coverprofile=coverage.out -race -v ./src/...