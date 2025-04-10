package file_transfer

import "os"

type FilePort interface {
	// Validate validates the file for business rules
	Validate() error

	GetName() string
	GetExtension() string
	GetSize() int64
	GetReference() *os.File
}
