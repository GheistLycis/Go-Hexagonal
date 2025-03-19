# Go-Hexagonal

Go-based implementation of the Hexagonal Architecture (Ports and Adapters), featuring both a HTTP server and a TCP server for P2P file transfers.

-   Web server implemented with Gin
-   Database handling with GORM and PostgreSQL
-   TCP file transfer service

## Installation 🔧

### Prerequisites

-   Go 1.24

### Steps

```sh
git clone https://github.com/GheistLycis/Go-Hexagonal.git
cd Go-Hexagonal
go mod tidy
```

Set a `.env` according to the `env.example` in the root.

To compile:

```sh
./scripts/build.sh <CMD="web"> <OUTPUT="main"> <OS=user_os> <ARCH=user_os_arch>
```

## Features 💻

### Web Server

#### Prerequisites

-   PostgreSQL (>= 16)

Create two types in your database, user_gender and user_status, as described in `src/user/domain/user.go`.

#### Usage

Start the HTTP server with:

```sh
./scripts/run.sh
```

Access it at `http://localhost:<WEB_PORT>`

### TCP Server

#### Usage

Start the TCP server with:

```sh
./scripts/run.sh tcp
```

Send a file:

```sh
nc localhost <TCP_PORT> < path/to/file
```

## Roadmap 🚀

-   ✅ **Web server**: Hexagonal POC in Go
-   ✅ **TCP server**: Server to receive P2P file transfers
-   **CLI cmd**: Implement an entry point for the file_transfer module to serve as the client to the TCP server, removing the need to use external tools to send files (such as netcat)
-   **GUI**: Build an interface to interact with both CLI and TCP commands
-   **Dockerization**: Implement Docker to avoid any initial setup
