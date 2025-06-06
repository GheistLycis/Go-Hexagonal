package steganography

import (
	"errors"
	"image"
	"image/color"
)

type ImageCodecService struct {
	CodecService
}

func NewImageCodecService(c *CodecService) *ImageCodecService {
	return &ImageCodecService{
		CodecService: *c,
	}
}

func (s *ImageCodecService) Encode(img image.Image, msg string) *image.RGBA {
	bounds := img.Bounds()
	rgba := image.NewRGBA(bounds)
	msgBits := []uint8{}

	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			rgba.Set(x, y, img.At(x, y))
		}
	}

	for i := range len(msg) {
		msgBits = append(msgBits, s.ByteToBits(msg[i])...)
	}

	msgBits = append(msgBits, s.ByteToBits(0)...)

	bitI := 0
	for y := bounds.Min.Y; y < bounds.Max.Y && bitI < len(msgBits); y++ {
		for x := bounds.Min.X; x < bounds.Max.X && bitI < len(msgBits); x++ {
			r, g, b, a := s.getPixelBytes(rgba, x, y)

			// * LSB
			if bitI < len(msgBits) {
				r = (r & 0xFE) | msgBits[bitI]
				bitI++
			}
			if bitI < len(msgBits) {
				g = (g & 0xFE) | msgBits[bitI]
				bitI++
			}
			if bitI < len(msgBits) {
				b = (b & 0xFE) | msgBits[bitI]
				bitI++
			}

			rgba.Set(x, y, color.RGBA{r, g, b, a})
		}
	}

	return rgba
}

func (s *ImageCodecService) Decode(img image.Image) (string, error) {
	bounds := img.Bounds()
	bits := []uint8{}

	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			r, g, b, _ := s.getPixelBytes(img, x, y)

			// * LSB
			bits = append(bits, r&1, g&1, b&1)

			if len(bits) >= 8 && len(bits)%8 == 0 {
				lastByte := s.BitsToByte(bits[len(bits)-8:])

				if lastByte == 0 {
					message := []byte{}

					for i := 0; i < len(bits)-8; i += 8 {
						message = append(message, s.BitsToByte(bits[i:i+8]))
					}

					return string(message), nil
				}
			}
		}
	}

	return "", errors.New("no message detected")
}

func (s *ImageCodecService) getPixelBytes(img image.Image, x int, y int) (uint8, uint8, uint8, uint8) {
	r, g, b, a := img.At(x, y).RGBA()

	return uint8(r >> 8), uint8(g >> 8), uint8(b >> 8), uint8(a >> 8)
}
