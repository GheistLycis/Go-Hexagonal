package file_transfer

import (
	app "Go-Hexagonal/src/file_transfer/app"
	"fmt"
	"net"
)

/*
HandleServer handles active listener to receive incoming files from any allowed dials, in parallel.
*/
func HandleServer(l net.Listener) {
	defer l.Close()

	for {
		conn, err := l.Accept()
		if err != nil {
			fmt.Println("Connection error:", err)
			return
		}

		go app.NewFileTranserService(conn).HandleConnection()
	}
}
