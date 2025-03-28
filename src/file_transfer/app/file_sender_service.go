package file_transfer

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"path/filepath"

	domain "Go-Hexagonal/src/file_transfer/domain"
)

var msgMaxSize = 100 * 1024

type FileSenderService struct {
	conn net.Conn
}

func NewFileSenderService(c net.Conn) *FileSenderService { // TODO: use generic interface for connection adapter
	return &FileSenderService{
		conn: c,
	}
}

func (s *FileSenderService) HandleConnection(fp string) {
	go s.listenForMessages()

	file, err := s.getFile(fp)
	if err != nil {
		log.Fatalf("Error reading file - %v", err)
		return
	}

	if err := s.upload(file); err != nil {
		log.Fatalf("Error sending file - %v", err)
		return
	}
}

func (s *FileSenderService) listenForMessages() error {
	buffer := make([]byte, msgMaxSize)

	for {
		n, err := s.conn.Read(buffer)
		if err != nil {
			if err == io.EOF {
				break
			}
			return err
		}

		fmt.Println(string(buffer[:n]))
	}

	return nil
}

func (s *FileSenderService) getFile(fp string) (FilePort, error) {
	osFile, err := os.Open(fp)
	if err != nil {
		return nil, err
	}

	defer osFile.Close()

	name := filepath.Base(fp)
	extension := filepath.Ext(fp)
	buffer := &bytes.Buffer{}
	if _, err = io.Copy(buffer, osFile); err != nil {
		return nil, err
	}
	data := buffer.Bytes()

	file, err := domain.NewFile(
		name[:len(name)-len(extension)],
		extension,
		&data,
	)
	if err != nil {
		return nil, err
	}

	return file, nil
}

func (s *FileSenderService) upload(f FilePort) error {
	fileContract := struct {
		Name, Extension string
		Size            int64
	}{f.GetName(), f.GetExtension(), f.GetSize()}

	if err := json.NewEncoder(s.conn).Encode(fileContract); err != nil {
		return err
	}

	if _, err := io.Copy(s.conn, f.GetData()); err != nil {
		return err
	}

	return nil
}
