package tcp

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"strconv"
)

/*
Init creates a TCP server and listens to any incoming dials indefinitely.
*/
func Init() {
	serverPort, err := strconv.ParseInt(os.Getenv("TCP_PORT"), 10, 64)
	if err != nil {
		log.Fatalf("[WEB] Failed to parse ENV variable TCP_PORT - %v", err)
	}

	listener, err := net.Listen("tcp", fmt.Sprintf("0.0.0.0:%d", serverPort))
	if err != nil {
		log.Fatalf("Failed to init TCP server - %v", err)
	}

	defer listener.Close()
	fmt.Println("[TCP] Server on. Listening on port", serverPort)

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("[TCP] Connection error:", err)
			continue
		}

		go handleConnection(conn)
	}
}

func handleConnection(c net.Conn) {
	defer c.Close()

	reader := bufio.NewReader(c)
	for {
		message, err := reader.ReadString('\n')
		if err != nil {
			fmt.Printf("[TCP] Closing connection with %v.\nError reading incoming stream - %v", c.RemoteAddr(), err)
			return
		}

		fmt.Printf("[TCP] From %v: %s", c.RemoteAddr(), message)
	}
}
