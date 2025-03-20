package file_transfer

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/asaskevich/govalidator"
	"github.com/joho/godotenv"
)

var maxSizeBytes float64

func init() {
	govalidator.SetFieldsRequiredByDefault(true)

	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error loading .env file - %v", err)
	}

	maxSizeGb, err := strconv.ParseFloat(os.Getenv("FT_MAX_GB"), 64)
	if err != nil {
		log.Fatalf("Failed to parse ENV variable FT_MAX_GB - %v", err)
	}
	maxSizeBytes = maxSizeGb * 1024 * 1024 * 1024
}

type File struct {
	Name      string        `valid:"-"`
	Extension string        `valid:"-"`
	Buffer    *bytes.Buffer `valid:"-"`
}

func NewFile(name string, extension string) (*File, error) {
	file := &File{
		Name:      name,
		Extension: extension,
		Buffer:    &bytes.Buffer{},
	}

	if _, err := file.Validate(); err != nil {
		return nil, err
	}

	return file, nil
}

func (f *File) Validate() (bool, error) {
	_, err := govalidator.ValidateStruct(f)
	if err != nil {
		return false, err
	}

	if float64(f.Buffer.Len()) > maxSizeBytes {
		return false, fmt.Errorf("o arquivo não pode ser maior que %.2f gB", maxSizeBytes/(1024*1024*1024))
	}

	return true, nil
}

func (f *File) WriteBuffer(b []byte) (int, error) {
	n, err := f.Buffer.Write(b)
	if err != nil {
		return n, err
	}

	_, err = f.Validate()
	if err != nil {
		return n, err
	}

	return n, nil
}

func (f *File) GetName() string          { return f.Name }
func (f *File) GetExtension() string     { return f.Extension }
func (f *File) GetBuffer() *bytes.Buffer { return f.Buffer }
