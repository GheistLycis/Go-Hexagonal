package steganography

import "errors"

type CodecService struct {
}

func NewCodecService() *CodecService {
	return &CodecService{}
}

func (s *CodecService) ByteToBits(b byte, msbFirst bool) []uint8 {
	bits := make([]uint8, 8)

	for i := range 8 {
		bit := (b >> i) & 1

		if msbFirst {
			bits[7-i] = bit
		} else {
			bits[i] = bit
		}
	}

	return bits
}

func (s *CodecService) BitsToByte(bits []uint8, msbFirst bool) byte {
	var b byte

	for i := range 8 {
		if msbFirst {
			b |= bits[i] << (7 - i)
		} else {
			b |= bits[i] << i
		}
	}

	return b
}

func (s *CodecService) StringToBits(str string, msbFirst bool) []uint8 {
	bits := []uint8{}

	for i := range len(str) {
		bits = append(bits, s.ByteToBits(str[i], msbFirst)...)
	}

	return bits
}

func (s *CodecService) BitsToString(bits []uint8, msbFirst bool) (string, error) {
	if len(bits) < 8 {
		return "", errors.New("provided slice has less than 8 bits")
	}
	if len(bits)%8 != 0 {
		bits = bits[:len(bits)-len(bits)%8]
	}

	bytes := make([]byte, len(bits)/8)

	for i := range len(bytes) {
		bytes[i] = s.BitsToByte(bits[i*8:(i+1)*8], msbFirst)
	}

	return string(bytes), nil
}
