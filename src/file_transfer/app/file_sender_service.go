package file_transfer

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"io"
	"log"
	"net"
	"os"
)

type FileSenderService struct {
	conn net.Conn
}

func NewFileSenderService(c net.Conn) FileSenderServicePort { // TODO: use generic interface for connection adapter
	return &FileSenderService{
		conn: c,
	}
}

func (s *FileSenderService) HandleConnection(fp string) {
	data, err := os.ReadFile(fp)
	if err != nil {
		log.Fatalf("Error reading file - %v", err)
		return
	}
	size := int64(len(data))

	go s.listenMessages()

	// TODO: also send file name and extension
	binary.Write(s.conn, binary.LittleEndian, size)

	if _, err := io.CopyN(s.conn, bytes.NewReader(data), size); err != nil {
		log.Fatalf("Error sending file - %v", err)
		return
	}
}

func (s *FileSenderService) listenMessages() {
	msgMaxSize := 100 * 1024 // * 100kB - meant only for incoming messages
	buffer := make([]byte, msgMaxSize)

	for {
		n, err := s.conn.Read(buffer)
		if err != nil && err != io.EOF {
			log.Fatalf("Error reading message from server - %v", err)
			return
		}

		fmt.Println(string(buffer[:n]))
	}
}
