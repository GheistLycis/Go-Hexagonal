package file_transfer

import "bytes"

type FilePort interface {
	Validate() error
	GetName() string
	GetExtension() string
	GetSize() int64
	GetData() *bytes.Buffer
}
