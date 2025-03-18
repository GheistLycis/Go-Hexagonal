package file_transfer

type Sender struct {
}

func NewSender(p string) SenderPort {
	return &Sender{}
}

func (c *Sender) Write(b []byte) (n int, err error) {
	return 0, nil
}
