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
	file, outPath, err := s.createFile(s.peerIp)
	if err != nil {
		fmt.Printf("\n(%s) Error creating destiny file: %v", s.peerIp, err)
		return
	}

	fmt.Printf("\nReceiving %s (%d mB)...", file.Name+file.Extension, file.Size/(1024*1024))

	if err = s.download(file); err != nil {
		fmt.Printf("\n(%s) Error downloading content: %v", s.peerIp, err)
		return
	}

	fmt.Printf("\n(%s) Content downloaded sucessfully (%s)", s.peerIp, outPath)
}

func (s *FileReceiverService) createFile(f string) (*domain.File, string, error) {
	var file domain.File
	if err := json.NewDecoder(s.conn).Decode(&file); err != nil {
		return nil, "", err
	}
	if err := file.Validate(); err != nil {
		return nil, "", err
	}

	outDir := workDir + osSep + outFolder + osSep + f
	if err := os.MkdirAll(outDir, os.ModePerm); err != nil {
		return nil, "", err
	}
	outPath := outDir + osSep + file.Name + file.Extension
	osFile, err := os.Create(outPath)
	if err != nil {
		return nil, "", err
	}

	file.Reference = osFile

	return &file, outPath, nil
}

func (s *FileReceiverService) download(f domain.FilePort) error {
	defer f.GetReference().Close()

	if _, err := io.CopyN(f.GetReference(), s.conn, f.GetSize()); err != nil {
		return err
	}

	if _, err := s.conn.Write([]byte(ackMsg)); err != nil {
		return err
	}

	return nil
}
