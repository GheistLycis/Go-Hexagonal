package steganography

import (
	ports "Go-Hexagonal/src/steganography/ports"

	"image/png"
	"log"
	"os"
	"path/filepath"
	"strings"
)

type CodecImageService struct {
	ImageCodecService ports.ImageCodecServicePort
}

func NewCodecImageService(s ports.ImageCodecServicePort) *CodecImageService {
	return &CodecImageService{
		ImageCodecService: s,
	}
}

func (s *CodecImageService) Encode(fp string, msg string) {
	file, err := os.Open(fp)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	img, err := png.Decode(file) // TODO: auto detect image type
	if err != nil {
		log.Fatal(err)
	}

	encodedImg := s.ImageCodecService.Encode(img, msg)

	outFile, err := os.Create(s.appendToFileName(fp, "_encoded"))
	if err != nil {
		log.Fatal(err)
	}
	defer outFile.Close()

	png.Encode(outFile, encodedImg)
}

func (s *CodecImageService) Decode(fp string) {
	file, err := os.Open(fp)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	img, err := png.Decode(file)
	if err != nil {
		log.Fatal(err)
	}

	decodedMessage, err := s.ImageCodecService.Decode(img)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Recovered message:", decodedMessage)
}

func (s *CodecImageService) appendToFileName(fp string, str string) string {
	dir := filepath.Dir(fp)
	base := filepath.Base(fp)
	ext := filepath.Ext(base)
	name := strings.TrimSuffix(base, ext)
	newName := name + str + ext

	return filepath.Join(dir, newName)
}
