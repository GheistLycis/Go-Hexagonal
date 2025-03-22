package file_transfer

import (
	"encoding/gob"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"path/filepath"

	domain "Go-Hexagonal/src/file_transfer/domain"
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

func (s *FileSenderService) getFile(fp string) (domain.FilePort, error) {
	osFile, err := os.Open(fp)
	if err != nil {
		return nil, err
	}

	defer osFile.Close()

	fileInfo, err := osFile.Stat()
	if err != nil {
		return nil, err
	}
	name := fileInfo.Name()
	extension := filepath.Ext(fp)
	file, err := domain.NewFile(
		name[:len(name)-len(extension)],
		extension,
		make([]byte, fileInfo.Size()),
	)
	if err != nil {
		return nil, err
	}

	_, err = io.ReadFull(osFile, file.GetData().Bytes())
	if err != nil && err != io.EOF {
		return nil, err
	}

	return file, nil
}

func (s *FileSenderService) upload(f domain.FilePort) error {
	fileContract := struct {
		Name, Extension string
		Size            int64
	}{f.GetName(), f.GetExtension(), f.GetSize()}

	if err := gob.NewEncoder(s.conn).Encode(fileContract); err != nil {
		return err
	}

	if _, err := io.Copy(s.conn, f.GetData()); err != nil {
		return err
	}

	return s.listenMessages()
}

func (s *FileSenderService) listenMessages() error {
	msgMaxSize := 100 * 1024
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
