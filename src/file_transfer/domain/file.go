package file_transfer

import (
	"errors"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/joho/godotenv"
)

var validate = validator.New() // TODO: implement singleton validator
var maxSize int64

func init() {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error loading .env file - %v", err)
	}

	maxSizeGb, err := strconv.ParseFloat(os.Getenv("FT_MAX_GB"), 64)
	if err != nil {
		log.Fatalf("Failed to parse ENV variable FT_MAX_GB - %v", err)
	}
	maxSize = int64(maxSizeGb * 1024 * 1024 * 1024)

}

type File struct {
	Name      string
	Extension string
	Size      int64
	Reference *os.File
}

func NewFile(name string, extension string, size int64, reference *os.File) (*File, error) {
	file := &File{
		Name:      filepath.Clean(name),
		Extension: extension,
		Size:      size,
		Reference: reference,
	}

	if name == "" {
		file.Name = time.Now().Format("2006-01-02T15:04:05")
	}

	if err := file.Validate(); err != nil {
		return nil, err
	}

	return file, nil
}

func (f *File) Validate() error {
	if err := validate.Struct(f); err != nil {
		return err
	}

	if f.Size > maxSize {
		return fmt.Errorf("file size exceeds the limit of %d mB", maxSize/(1024*1024))
	}

	if filepath.IsAbs(f.Name) ||
		strings.Contains(f.Name, "..") ||
		strings.Contains(f.Name, "~") ||
		strings.Contains(f.Name, ":") {
		return errors.New("file contains invalid path string")
	}

	return nil
}

func (f *File) GetName() string        { return f.Name }
func (f *File) GetExtension() string   { return f.Extension }
func (f *File) GetSize() int64         { return f.Size }
func (f *File) GetReference() *os.File { return f.Reference }
