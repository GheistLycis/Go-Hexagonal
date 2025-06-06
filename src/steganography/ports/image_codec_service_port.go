package steganography

import "image"

type ImageCodecServicePort interface {
	// Encode encodes msg within img pixels. It does not alter the original image, but returns an encoded version.
	Encode(img image.Image, msg string) *image.RGBA

	// Decode decodes any hidden messge within img pixels. It returns an error if none is found.
	Decode(img image.Image) (string, error)
}
