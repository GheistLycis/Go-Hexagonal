package file_transfer

import "bytes"

type FilePort interface {
	Validate() (bool, error)
	WriteBuffer(b []byte) (int, error)
	ClearBuffer() error
	GetName() string
	GetExtension() string
	GetBuffer() *bytes.Buffer
}
