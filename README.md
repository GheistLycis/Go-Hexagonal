# Go-Hexagonal

Go-based implementation of the Hexagonal Architecture (Ports and Adapters), featuring both a HTTP server and a TCP server/client for P2P file transfers.

-   Web server implemented with Gin
-   Database handling with GORM and PostgreSQL
-   TCP file transfer service
-   PNG Steganography

## Installation 🔧

### Prerequisites

-   Go 1.24 (or Docker)

### Steps

```sh
git clone https://github.com/GheistLycis/Go-Hexagonal.git
cd Go-Hexagonal
go mod tidy
```

Set a `.env` according to the `env.example` in the root.

To compile:

```sh
./scripts/build.sh <CMD=web> <OUTPUT=main> <OS=user_os> <ARCH=user_os_arch>

# Or
make build ARGS="<CMD=web> <OUTPUT=main> <OS=user_os> <ARCH=user_os_arch>"
```

Where CMD is "web", "tcp" or "cli".

## Features 🌟

### Web Server

#### Prerequisites

-   PostgreSQL 16.x (or Docker)

If not using Docker, create two types in your database, user_gender and user_status, as described in `src/user/domain`.

#### Usage

Start the HTTP server with:

```sh
./scripts/web.sh

# Or
make web

# Or, with Docker
make docker-web
```

Access it at `http://localhost:<WEB_PORT>`

### TCP File Transfer

#### Usage

Start the TCP server to receive files with:

```sh
./scripts/tcp.sh

# Or
make tcp

# Or, with Docker
make docker-tcp
```

Start the CLI to send files to the TCP server with:

```sh
./scripts/cli.sh tcp <ADDRESS> <PORT> <FILE_PATH>

# Or
make cli tcp <ADDRESS> <PORT> <FILE_PATH>

# Or, with Docker
make docker-cli tcp <ADDRESS> <PORT> <FILE_PATH>
```

### Steganography

#### Usage

Start the CLI to encode messages within PNGs with:

```sh
./scripts/cli.sh steg encode <FILE_PATH> <MESSAGE>

# Or
make cli steg encode <FILE_PATH> <MESSAGE>

# Or, with Docker
make docker-cli steg encode <FILE_PATH> <MESSAGE>
```

The original file won't be altered. Instead, a copy will be saved in the same folder as the original (with "\_encoded" appended in the name).

Start the CLI to decode messages within PNGs with:

```sh
./scripts/cli.sh steg decode <FILE_PATH>

# Or
make cli steg decode <FILE_PATH>

# Or, with Docker
make docker-cli steg decode <FILE_PATH>
```

## DevOps 🔨

### Test

```sh
make test

# Or
make test-verbose
```

### Audit

```sh
make audit
```

## Roadmap 🚀

-   ✅ **Web server**: Hexagonal POC in Go
-   ✅ **TCP server**: Server to receive P2P file transfers
-   ✅ **CLI cmd**: Implement an entry point for the file_transfer module to serve as the client to the TCP server, removing the need to use external tools to send files (such as netcat)
-   **GUI**: Build an interface to interact with both CLI and TCP commands
-   ✅ **Dockerization**: Implement Docker to avoid any initial setup
-   **Unit tests**: Implement unit test coverage
