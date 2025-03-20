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

var maxSize float64
var ErrSizeExceeded = fmt.Errorf("file size exceeded the limit (%.2f mB)", maxSize/(1024*1024))
var maxBufferSize float64
var ErrBufferExceeded = fmt.Errorf("buffer size exceeded the limit (%.2f mB)", maxBufferSize/(1024*1024))

func init() {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error loading .env file - %v", err)
	}

	maxBufferSizeMb, err := strconv.ParseFloat(os.Getenv("FT_BUFF_MB"), 64)
	if err != nil {
		log.Fatalf("Failed to parse ENV variable FT_BUFF_MB - %v", err)
	}
	maxBufferSize = maxBufferSizeMb * 1024 * 1024

	maxSizeGb, err := strconv.ParseFloat(os.Getenv("FT_MAX_GB"), 64)
	if err != nil {
		log.Fatalf("Failed to parse ENV variable FT_MAX_GB - %v", err)
	}
	maxSize = maxSizeGb * 1024 * 1024 * 1024

	govalidator.SetFieldsRequiredByDefault(true)
}

type File struct {
	Name      string        `valid:"-"`
	Extension string        `valid:"-"`
	Size      int           `valid:"-"`
	Buffer    *bytes.Buffer `valid:"-"`
}

func NewFile(name string, extension string) (*File, error) {
	file := &File{
		Name:      name,
		Extension: extension,
		Size:      0,
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

	// ? track if all active files summed up exceed max allowed
	if float64(f.Buffer.Len()) > maxBufferSize {
		return false, ErrBufferExceeded
	}

	if float64(f.Size) > maxSize {
		return false, ErrSizeExceeded
	}

	return true, nil
}

func (f *File) WriteBuffer(b []byte) (int, error) {
	n, err := f.Buffer.Write(b)
	if err != nil {
		return n, err
	}

	f.Size += n

	_, err = f.Validate()
	if err != nil {
		return n, err
	}

	return n, nil
}

func (f *File) ClearBuffer() error {
	f.Buffer.Reset()

	_, err := f.Validate()

	return err
}

func (f *File) GetName() string          { return f.Name }
func (f *File) GetExtension() string     { return f.Extension }
func (f *File) GetSize() int             { return f.Size }
func (f *File) GetBuffer() *bytes.Buffer { return f.Buffer }
