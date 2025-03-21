package tcp

import (
	file_transfer "Go-Hexagonal/src/file_transfer/cmd/tcp"
	"fmt"
	"log"
	"net"
	"os"
	"strconv"
)

/*
Init starts a TCP server.
*/
func Init() {
	serverPort, err := strconv.ParseInt(os.Getenv("TCP_PORT"), 10, 64)
	if err != nil {
		log.Fatalf("[TCP] Failed to parse ENV variable TCP_PORT - %v", err)
	}

	listener, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", serverPort))
	if err != nil {
		log.Fatalf("Failed to init TCP server - %v", err)
	}

	fmt.Println("[TCP] Server on. Listening on port", serverPort)

	file_transfer.HandleServer(listener)
}
