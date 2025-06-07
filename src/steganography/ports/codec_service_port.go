package steganography

type CodecServicePort interface {
	// ByteToBits converts given byte to a 8-bit slice. If msbFirst is true, uses MSB-first order; uses LSB-first otherwise
	ByteToBits(b byte, msbFirst bool) []uint8

	// BitsToByte converts given 8-bit slice to a byte. If msbFirst is true, uses MSB-first order; uses LSB-first otherwise
	BitsToByte(bits []uint8, msbFirst bool) byte

	// StringToBits converts given string to 8-bit slice. If msbFirst is true, uses MSB-first order; uses LSB-first otherwise
	StringToBits(str string, msbFirst bool) []uint8

	// BitsToString converts bit slice to string, ignoring any extra bits. If msbFirst is true, uses MSB-first order; uses LSB-first otherwise. Returns an error if len(bits) < 8.
	BitsToString(bits []uint8, msbFirst bool) (string, error)
}
