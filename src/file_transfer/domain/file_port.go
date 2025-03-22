package file_transfer

import "bytes"

type FilePort interface {
	Validate() (isValid bool, err error)
	GetName() string
	GetExtension() string
	GetSize() int64
	GetData() *bytes.Buffer
}
