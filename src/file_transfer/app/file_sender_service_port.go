package file_transfer

type FileSenderServicePort interface {
	HandleConnection(filePath string)
}
