package file_transfer

type ReceiverServicePort interface {
	Read(b []byte) error
}
