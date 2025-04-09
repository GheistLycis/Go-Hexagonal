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


# * ==DOCKER==
.PHONY: docker-cli docker-build-cli docker-tcp docker-build-tcp docker-web docker-build-web

# CLI
docker-cli:
	docker compose up cli $(ARGS)

docker-build-cli:
	docker compose up --build cli $(ARGS)

# TCP
docker-tcp:
	docker compose up -d tcp

docker-build-tcp:
	docker compose up -d --build tcp

# WEB
docker-web:
	docker compose up -d web database

docker-build-web:
	docker compose up -d --build web database


# * ==DEV OPS==
.PHONY: test test-verbose audit

test:
	go test -cover -coverprofile=coverage.out -race ./src/...

test-verbose:
	go test -cover -coverprofile=coverage.out -race -v ./src/...

audit:
	govulncheck ./...