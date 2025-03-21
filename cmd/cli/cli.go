package cli

import (
	file_transfer "Go-Hexagonal/src/file_transfer/cmd/cli"
	"log"
	"strconv"
)

/*
Init starts CLI program, handling args and passing them to the file_transfer module.
*/
func Init(args []string) {
	port, err := strconv.Atoi(args[1])
	if err != nil {
		log.Fatalf("Error parsing arg for port - %v", err)
	}

	cmd := file_transfer.CommandsDTO{
		Address:  args[0],
		Port:     int32(port),
		FilePath: args[2],
		Protocol: "tcp",
	}

	file_transfer.HandleCommands(cmd)
}
