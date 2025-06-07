package steganography

import (
	ports "Go-Hexagonal/src/steganography/ports"
	"errors"
	"image"
	"image/color"
	"image/draw"
)

type ImageCodecService struct {
	CodecService ports.CodecServicePort
}

func NewImageCodecService(s ports.CodecServicePort) *ImageCodecService {
	return &ImageCodecService{CodecService: s}
}

func (s *ImageCodecService) Encode(img image.Image, msg string) *image.RGBA {
	bounds := img.Bounds()
	minX, maxX, minY, maxY := bounds.Min.X, bounds.Max.X, bounds.Min.Y, bounds.Max.Y

	rgba := image.NewRGBA(bounds)
	draw.Draw(rgba, bounds, img, bounds.Min, draw.Src)

	msgBits := s.CodecService.StringToBits(msg, true)
	msgBits = append(msgBits, s.CodecService.ByteToBits(0, true)...)
	msgBitsLen := len(msgBits)

	msgBitIdx := 0
	for x := minX; x < maxX && msgBitIdx < msgBitsLen; x++ {
		for y := minY; y < maxY && msgBitIdx < msgBitsLen; y++ {
			r, g, b, a := s.getPixelBytes(rgba, x, y)

			// * LSB
			if msgBitIdx < msgBitsLen {
				r = (r & 0xFE) | msgBits[msgBitIdx]
				msgBitIdx++
			}
			if msgBitIdx < msgBitsLen {
				g = (g & 0xFE) | msgBits[msgBitIdx]
				msgBitIdx++
			}
			if msgBitIdx < msgBitsLen {
				b = (b & 0xFE) | msgBits[msgBitIdx]
				msgBitIdx++
			}

			rgba.Set(x, y, color.RGBA{r, g, b, a})
		}
	}

	return rgba
}

func (s *ImageCodecService) Decode(img image.Image) (string, error) {
	bounds := img.Bounds()
	minX, maxX, minY, maxY := bounds.Min.X, bounds.Max.X, bounds.Min.Y, bounds.Max.Y
	msgBits := []uint8{}

	for x := minX; x < maxX; x++ {
		for y := minY; y < maxY; y++ {
			r, g, b, _ := s.getPixelBytes(img, x, y)

			// * LSB
			msgBits = append(msgBits, r&1, g&1, b&1)
			msgBitsLen := len(msgBits)

			if msgBitsLen >= 8 && msgBitsLen%8 == 0 {
				lastByte := s.CodecService.BitsToByte(msgBits[msgBitsLen-8:], true)

				if lastByte == 0 {
					msg, err := s.CodecService.BitsToString(msgBits[:msgBitsLen-8], true)
					if err != nil {
						return "", err
					}

					return msg, nil
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
