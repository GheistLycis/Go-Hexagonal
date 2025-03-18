package file_transfer

import (
	"net"
	"time"

	domain "Go-Hexagonal/src/file_transfer/domain"
)

type Connection struct {
	conn net.Conn
}

func NewConnection(c net.Conn) domain.ConnectionPort {
	return &Connection{
		conn: c,
	}
}

func (c *Connection) Close() error {
	return c.Close()
}

func (c *Connection) SetTimeOut(t time.Time) error {
	return c.conn.SetReadDeadline(t)
}
