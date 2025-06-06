package steganography

type CodecServicePort interface {
	// ByteToBits converts given byte to a 8-bit slice
	ByteToBits(b byte) []uint8

	// BitsToByte converts given 8-bit slice to a byte
	BitsToByte(bits []uint8) byte
}
