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

func NewFile(name string, extension string, buffer *bytes.Buffer) (*File, error) {
	file := &File{
		Name:      name,
		Extension: extension,
		Buffer:    buffer,
	}

	if _, err := file.Validate(); err != nil {
		return nil, err
	}

	return file, nil
}

func (u *File) Validate() (bool, error) {
	_, err := govalidator.ValidateStruct(u)
	if err != nil {
		return false, err
	}

	if float64(u.Buffer.Len()) > maxSizeBytes {
		return false, fmt.Errorf("o arquivo não pode ser maior que %.2f gB", maxSizeBytes/(1024*1024*1024))
	}

	return true, nil
}

func (u *File) GetName() string          { return u.Name }
func (u *File) GetExtension() string     { return u.Extension }
func (u *File) GetBuffer() *bytes.Buffer { return u.Buffer }
