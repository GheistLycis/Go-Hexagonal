package file_transfer

type ReceiverPort interface {
	Read(b []byte) (int, error)
}
