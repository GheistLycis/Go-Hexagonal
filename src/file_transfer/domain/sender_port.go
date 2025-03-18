package file_transfer

type SenderPort interface {
	Write(b []byte) (int, error)
}
