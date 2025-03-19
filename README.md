# Go-Hexagonal

Go-based implementation of the Hexagonal Architecture (Ports and Adapters), featuring both a HTTP server and a TCP server for P2P file transfers.

## Features

-   Web server implemented with Gin
-   Database handling with GORM and PostgreSQL
-   TCP file transfer service

## Installation

### Prerequisites

-   Go 1.24
-   PostgreSQL (>= 16)

### Steps

```sh
git clone https://github.com/GheistLycis/Go-Hexagonal.git
cd Go-Hexagonal
go mod tidy
```

-   Set a .env according to the `env.example` in the root.

## Usage

### Web Server

Before running the web server, create two types in your database, user_gender and user_status, as described in `src/user/domain/user.go`.

Start the HTTP server with:

```sh
./scripts/run.sh
```

Access it at `http://localhost:8080`

### TCP Server

Start the TCP server with:

```sh
./scripts/run.sh tcp
```

Send a file:

```sh
nc localhost <TCP_PORT> < path/to/file
```

## Roadmap 🚀

-   **Dockerization**: Implement Docker
-   **CLI cmd**: Implement an entry point for the file_transfer module to serve as the client to the TCP server, removing the need to use external tools to send files (such as netcat)
-   **GUI**: Build an interface to interact with both CLI and TCP commands
