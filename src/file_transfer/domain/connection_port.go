package file_transfer

import "time"

type ConnectionPort interface {
	Close() error
	SetTimeOut(t time.Time) error
}
