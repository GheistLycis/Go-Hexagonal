package steganography

type CodecService struct {
}

func NewCodecService() *CodecService {
	return &CodecService{}
}

func (s *CodecService) ByteToBits(b byte) []uint8 {
	bits := make([]uint8, 8)

	for i := 7; i >= 0; i-- {
		bits[7-i] = (b >> i) & 1
	}

	return bits
}

func (s *CodecService) BitsToByte(bits []uint8) byte {
	var b byte

	for i := range 8 {
		b |= bits[i] << (7 - i)
	}

	return b
}
