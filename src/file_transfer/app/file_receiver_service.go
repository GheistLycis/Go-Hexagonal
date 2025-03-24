package file_transfer

import (
	"encoding/gob"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"strconv"
	"strings"

	domain "Go-Hexagonal/src/file_transfer/domain"

	"github.com/joho/godotenv"
)

var workDir string
var outFolder string
var osSep string
var maxBufferSize int64

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
	maxBufferSize = int64(maxBufferSizeMb * 1024 * 1024)
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
		fmt.Printf("\n(%s) Error downloading content: %v", s.peerIp, err)
		s.conn.Write([]byte("\nError downloading content"))
		return
	}

	fmt.Printf("\n(%s) Content downloaded sucessfully (%s)", s.peerIp, outDir)
	s.conn.Write([]byte("\nContent downloaded sucessfully"))
}

func (s *FileReceiverService) download(f string) (string, error) {
	var totalRead int64
	var outPath string
	file, err := domain.NewFile("", "", []byte{})
	if err != nil {
		return "", err
	}

	if err := gob.NewDecoder(s.conn).Decode(file); err != nil {
		return "", err
	}

	fmt.Printf("\n(%s) Receiving %s (%d mB)...", s.peerIp, file.GetName()+file.GetExtension(), file.GetSize()/(1024*1024))

	for {
		msg := fmt.Sprintf("\nDownloading data... (TOTAL = %d mB)", totalRead/(1024*1024))
		fmt.Print(msg)
		s.conn.Write([]byte(msg))

		n, err := io.CopyN(file.GetData(), s.conn, maxBufferSize)
		if err != nil && err != io.EOF {
			return "", err
		}

		if outPath, err = s.save(file, f); err != nil {
			return "", err
		}
		if totalRead += int64(n); totalRead == file.GetSize() {
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
	osFile, err := os.OpenFile(outPath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return "", err
	}

	defer osFile.Close()

	if _, err = io.Copy(osFile, fi.GetData()); err != nil {
		return "", err
	}
	if err := osFile.Sync(); err != nil {
		return "", err
	}

	return outPath, nil
}
