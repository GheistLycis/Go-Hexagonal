package tcp

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

/*
Init creates a TCP server and listens to any incoming dials indefinitely.

-p: the server port.
*/
func Init(p int) {
	listener, err := net.Listen("tcp", fmt.Sprintf("0.0.0.0:%d", p))
	if err != nil {
		log.Fatalf("Failed to init TCP server: %v", err)
		return
	}

	defer listener.Close()
	fmt.Println("[TCP] Server on. Listening on port:", p)

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Connection error:", err)
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
			fmt.Printf("Closing connection with %v.\nError reading incoming stream - %v", c.RemoteAddr(), err)
			return
		}

		fmt.Printf("[TCP] %v: %s", c.RemoteAddr(), message)
	}
}
