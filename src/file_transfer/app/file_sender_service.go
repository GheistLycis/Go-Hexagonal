package file_transfer

import (
	"net"
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
}
