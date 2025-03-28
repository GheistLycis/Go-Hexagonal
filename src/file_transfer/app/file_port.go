package file_transfer

type FilePort interface {
	Validate() error
	GetName() string
	GetExtension() string
	GetSize() int64
}
