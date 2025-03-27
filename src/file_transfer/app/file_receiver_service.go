package file_transfer

import (
	"encoding/json"
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

func NewFileReceiverService(c net.Conn) *FileReceiverService { // TODO: use generic interface for connection adapter
	return &FileReceiverService{
		conn:   c,
		peerIp: strings.Split(c.RemoteAddr().String(), ":")[0],
	}
}

func (s *FileReceiverService) HandleConnection() {
	outPath, err := s.download(s.peerIp)
	if err != nil {
		fmt.Printf("\n(%s) Error downloading content: %v", s.peerIp, err)
		s.conn.Write([]byte("\nError downloading content"))
		return
	}

	fmt.Printf("\n(%s) Content downloaded sucessfully (%s)", s.peerIp, outPath)
	s.conn.Write([]byte("\nContent downloaded sucessfully"))
}

func (s *FileReceiverService) download(f string) (string, error) {
	var file domain.File
	if err := json.NewDecoder(s.conn).Decode(&file); err != nil {
		return "", err
	}
	if err := file.Validate(); err != nil {
		return "", err
	}
	msg := fmt.Sprintf("\nReceiving %s (%d mB)...", file.Name+file.Extension, file.Size/(1024*1024))
	fmt.Print(msg)
	s.conn.Write([]byte(msg))

	outDir := workDir + osSep + outFolder + osSep + f + osSep
	if err := os.MkdirAll(outDir, os.ModePerm); err != nil {
		return "", err
	}
	outPath := outDir + file.Name + file.Extension
	osFile, err := os.OpenFile(outPath, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)
	if err != nil {
		return "", err
	}

	defer osFile.Close()

	if _, err = io.CopyN(osFile, s.conn, file.Size); err != nil {
		return "", err
	}

	return outPath, nil
}
