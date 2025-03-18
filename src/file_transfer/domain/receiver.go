package file_transfer

type Receiver struct {
}

func NewReceiver(p string) ReceiverPort {
	return &Receiver{}
}

func (c *Receiver) Read(b []byte) (n int, err error) {
	return 0, nil
}
