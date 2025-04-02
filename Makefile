.PHONY: build-tcp up-tcp build-web up-web

# CLI
build-cli:
	docker compose up --build cli
up-cli:
	docker compose up cli

# TCP
build-tcp:
	docker compose up -d --build tcp
up-tcp:
	docker compose up -d tcp

# WEB
build-web:
	docker compose up -d --build web database
up-web:
	docker compose up -d web database
