package file_transfer

import (
	app "Go-Hexagonal/src/file_transfer/app"
	"fmt"
	"log"
	"net"
)

/*
HandleCommands handles commands in file_transfer context to send files.
*/
func HandleCommands(cmd CommandDTO) {
	conn, err := net.Dial(cmd.Protocol, net.JoinHostPort(cmd.Address, fmt.Sprintf("%d", cmd.Port)))
	if err != nil {
		log.Fatalf("Failed to establish connection - %v", err)
	}

	defer conn.Close()

	app.NewFileSenderService(conn).HandleConnection(cmd.FilePath)
}

type CommandDTO struct {
	Address  string
	Port     int32
	FilePath string
	Protocol string
}
