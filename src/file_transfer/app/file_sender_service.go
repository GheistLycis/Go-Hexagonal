package file_transfer

import (
	"encoding/json"
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

func NewFileSenderService(c net.Conn) *FileSenderService { // TODO: use generic interface for connection adapter
	return &FileSenderService{
		conn: c,
	}
}

func (s *FileSenderService) HandleConnection(fp string) {
	file, osFile, err := s.getFile(fp)
	if err != nil {
		log.Fatalf("Error reading file - %v", err)
		return
	}

	if err := s.upload(file, osFile); err != nil {
		log.Fatalf("Error sending file - %v", err)
		return
	}

	s.waitForConfirmation()
}

func (s *FileSenderService) getFile(fp string) (FilePort, *os.File, error) {
	osFile, err := os.Open(fp)
	if err != nil {
		return nil, nil, err
	}
	defer func() {
		if err != nil {
			osFile.Close()
		}
	}()

	fileInfo, err := osFile.Stat()
	if err != nil {
		return nil, nil, err
	}
	name := fileInfo.Name()
	extension := filepath.Ext(fp)
	file, err := domain.NewFile(
		name[:len(name)-len(extension)],
		extension,
		fileInfo.Size(),
	)
	if err != nil {
		return nil, nil, err
	}

	return file, osFile, nil
}

func (s *FileSenderService) upload(f FilePort, osF *os.File) error {
	defer osF.Close()

	if err := json.NewEncoder(s.conn).Encode(f); err != nil {
		return err
	}

	if _, err := io.Copy(s.conn, osF); err != nil {
		return err
	}

	return nil
}

func (s *FileSenderService) waitForConfirmation() {
	buffer := make([]byte, len(ackMsg))
	if _, err := s.conn.Read(buffer); err != nil {
		log.Fatal(err.Error())
	}
	if string(buffer) != ackMsg {
		log.Fatal("failed to receive download confirmation from receiver")
	}
}
