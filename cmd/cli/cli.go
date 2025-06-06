package cli

import (
	file_transfer "Go-Hexagonal/src/file_transfer/cmd/cli"
	steganography "Go-Hexagonal/src/steganography/cmd/cli"
	"log"
	"strconv"
)

func Init(args []string) {
	entry := args[0]

	args = args[1:]

	switch entry {
	case "tcp":
		port, err := strconv.Atoi(args[1])
		if err != nil {
			log.Fatalf("Error parsing arg for port - %v", err)
		}

		cmd := file_transfer.CommandDTO{
			Address:  args[0],
			Port:     int32(port),
			FilePath: args[2],
			Protocol: "tcp",
		}

		file_transfer.HandleCommands(cmd)
	case "steg":
		msg := ""

		if len(args) > 2 {
			msg = args[2]
		}

		cmd := steganography.CommandDTO{
			Operation: args[0],
			FilePath:  args[1],
			Message:   msg,
		}

		steganography.HandleCommands(cmd)
	default:
		log.Fatalf("Unexpected command \"%s\"", entry)
	}
}
