package file_transfer

import (
	"net"
	"time"
)

type ListenerPort interface {
	Accept() (net.Conn, error)
	Close() error
	SetTimeOut(c net.Conn, t time.Time) error
}
