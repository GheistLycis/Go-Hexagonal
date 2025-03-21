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
func HandleCommands(cmd CommandsDTO) {
	// TODO: handle IPv6 addresses
	conn, err := net.Dial(cmd.Protocol, fmt.Sprintf("%s:%d", cmd.Address, cmd.Port))
	if err != nil {
		log.Fatalf("Failed to stablish connection - %v", err)
	}

	defer conn.Close()

	app.NewFileSenderService(conn).HandleConnection(cmd.FilePath)
}

type CommandsDTO struct {
	Address  string
	Port     int32
	FilePath string
	Protocol string
}
