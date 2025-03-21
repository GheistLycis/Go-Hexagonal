package file_transfer

import (
	"encoding/binary"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"strconv"
	"strings"
	"time"

	domain "Go-Hexagonal/src/file_transfer/domain"

	"github.com/joho/godotenv"
)

var workDir string
var outFolder string
var osSep string
var maxBufferSize float64

func init() {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error loading .env file - %v", err)
	}

	wd, err := os.Getwd()
	if err != nil {
		log.Fatalf("Failed to get working directory - %v", err)
	}
	workDir = wd

	outFolder = os.Getenv("FT_OUT_DIR")

	osSep = string(os.PathSeparator)

	maxBufferSizeMb, err := strconv.ParseFloat(os.Getenv("FT_BUFF_MB"), 64)
	if err != nil {
		log.Fatalf("Failed to parse ENV variable FT_BUFF_MB - %v", err)
	}
	maxBufferSize = maxBufferSizeMb * 1024 * 1024
}

type FileReceiverService struct {
	conn   net.Conn
	peerIp string
}

func NewFileReceiverService(c net.Conn) FileReceiverServicePort { // TODO: use generic interface for connection adapter
	return &FileReceiverService{
		conn:   c,
		peerIp: strings.Split(c.RemoteAddr().String(), ":")[0],
	}
}

func (s *FileReceiverService) HandleConnection() {
	outDir, err := s.download(s.peerIp)
	if err != nil {
		s.conn.Write([]byte("\nError downloading content"))
		fmt.Printf("\n(%s) Error downloading content: %v", s.peerIp, err)
		return
	}

	s.conn.Write([]byte("\nContent downloaded sucessfully"))
	fmt.Printf("\n(%s) Content downloaded sucessfully (%s)", s.peerIp, outDir)
}

func (s *FileReceiverService) download(f string) (string, error) {
	var totalRead int64
	var outPath string
	file, err := domain.NewFile(
		time.Now().Format("2006-01-02T15:04:05"),
		"",
		0,
	)
	if err != nil {
		fmt.Printf("\n(%s) Error creating file: %v", s.peerIp, err)
		return "", err
	}

	binary.Read(s.conn, binary.LittleEndian, file.GetSize())
	fmt.Printf("\n(%s) Total file size to be received: %d mB", s.peerIp, *file.GetSize()/(1024*1024))
	time.Sleep(2 * time.Second)

	for {
		if _, err := file.Validate(); err != nil {
			return "", err
		}

		msg := fmt.Sprintf("\nReceiving file... (TOTAL = %d mB)", totalRead/(1024*1024))
		fmt.Print(msg)
		s.conn.Write([]byte(msg))

		n, err := io.CopyN(file.GetBuffer(), s.conn, int64(maxBufferSize))
		if err != nil && err != io.EOF {
			return "", err
		}

		if outPath, err = s.save(file, f); err != nil {
			return "", err
		}

		file.GetBuffer().Reset()
		totalRead += n

		if totalRead == *file.GetSize() {
			break
		}
	}

	return outPath, nil
}

func (s *FileReceiverService) save(fi domain.FilePort, fo string) (string, error) {
	outDir := workDir + osSep + outFolder + osSep + fo + osSep

	if err := os.MkdirAll(outDir, os.ModePerm); err != nil {
		return "", err
	}

	outPath := outDir + fi.GetName() + fi.GetExtension()
	outFile, err := os.OpenFile(outPath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return "", err
	}

	defer outFile.Close()

	_, err = io.Copy(outFile, fi.GetBuffer())
	if err != nil {
		return "", err
	}

	if err := outFile.Sync(); err != nil {
		return "", err
	}

	return outPath, nil
}
