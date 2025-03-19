package file_transfer

import "bytes"

type FilePort interface {
	Validate() (bool, error)
	GetName() string
	GetExtension() string
	GetBuffer() *bytes.Buffer
}
