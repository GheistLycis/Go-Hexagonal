package file_transfer

type SenderServicePort interface {
	Write(b []byte) error
}
