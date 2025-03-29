package file_transfer

import "os"

type FilePort interface {
	Validate() error
	GetName() string
	GetExtension() string
	GetSize() int64
	GetPointer() *os.File
}
