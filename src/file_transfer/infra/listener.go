package file_transfer

import (
	"net"
	"time"

	domain "Go-Hexagonal/src/file_transfer/domain"
)

type Listener struct {
	host net.Listener
}

func NewListener(l net.Listener) domain.ListenerPort {
	return &Listener{host: l}
}

func (s *Listener) Accept() (net.Conn, error) {
	return s.host.Accept()
}

func (s *Listener) Close() error {
	return s.host.Close()
}

func (l *Listener) SetTimeOut(c net.Conn, t time.Time) error {
	return c.SetWriteDeadline(t)
}
