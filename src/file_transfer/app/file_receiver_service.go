package file_transfer

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"strings"

	domain "Go-Hexagonal/src/file_transfer/domain"

	"github.com/joho/godotenv"
)

const ackMsg string = "_EOF_ACK_"

var workDir string
var outFolder string
var osSep string

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
	osFile, outPath, err := s.createFile(s.peerIp)
	if err != nil {
		fmt.Printf("\n(%s) Error downloading content: %v", s.peerIp, err)
		return
	}

	if err := s.download(osFile); err != nil {
		fmt.Printf("\n(%s) Error downloading content: %v", s.peerIp, err)
		return
	}

	fmt.Printf("\n(%s) Content downloaded sucessfully (%s)", s.peerIp, outPath)
}

func (s *FileReceiverService) createFile(f string) (*os.File, string, error) {
	var file domain.File
	if err := json.NewDecoder(s.conn).Decode(&file); err != nil {
		return nil, "", err
	}
	if err := file.Validate(); err != nil {
		return nil, "", err
	}
	fmt.Printf("\nReceiving %s (%d mB)...", file.Name+file.Extension, file.Size/(1024*1024))

	outDir := workDir + osSep + outFolder + osSep + f
	if err := os.MkdirAll(outDir, os.ModePerm); err != nil {
		return nil, "", err
	}
	outPath := outDir + osSep + file.Name + file.Extension
	osFile, err := os.Create(outPath)
	if err != nil {
		return nil, "", err
	}

	return osFile, outPath, err
}

func (s *FileReceiverService) download(f *os.File) error {
	defer f.Close()

	fileInfo, err := f.Stat()
	if err != nil {
		return err
	}

	if _, err = io.CopyN(f, s.conn, fileInfo.Size()); err != nil {
		return err
	}

	if _, err := s.conn.Write([]byte(ackMsg)); err != nil {
		return err
	}

	return nil
}
