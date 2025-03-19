package file_transfer

import (
	"bytes"
	"fmt"
	"io"
	"net"
	"os"
	"slices"
	"strings"
	"time"

	domain "Go-Hexagonal/src/file_transfer/domain"
)

type FileTransferService struct {
	conn   net.Conn
	peerIp string
}

func NewFileTranserService(c net.Conn) *FileTransferService { // TODO: use generic interface for connection adapter
	return &FileTransferService{
		conn:   c,
		peerIp: strings.Split(c.RemoteAddr().String(), ":")[0],
	}
}

func (s *FileTransferService) HandleConnection() {
	defer s.shutConnection()

	if s.peerIsTrusted() {
		fmt.Printf("\nStablished connection with %s", s.peerIp)
	} else {
		fmt.Printf("\nDenied connection with %s", s.peerIp)
		return
	}

	buffer, err := s.download()
	if err != nil {
		fmt.Printf("\n(%s) Error downloading content: %v", s.peerIp, err)
		return
	}

	file, err := domain.NewFile("", "", buffer)
	if err != nil {
		fmt.Printf("\n(%s) Invalid content received: %v", s.peerIp, err)
		return
	}

	fmt.Printf("\n(%s) Content received (%.2f mB)", s.peerIp, float64(file.GetBuffer().Len())/(1024*1024))

	outDir, err := s.save(file, s.peerIp)
	if err != nil {
		fmt.Printf("\n(%s) Error saving content: %v", s.peerIp, err)
		return
	}

	fmt.Printf("\n(%s) Content saved sucessfully (%s)!", s.peerIp, outDir)
}

func (s *FileTransferService) shutConnection() {
	fmt.Printf("\nClosing connection with %s", s.peerIp)
	s.conn.Close()
}

func (s *FileTransferService) peerIsTrusted() bool {
	ipsWhitelist := strings.Split(os.Getenv("FT_IP_WHITELIST"), ",")

	return slices.Contains(ipsWhitelist, s.peerIp)
}

func (s *FileTransferService) download() (*bytes.Buffer, error) {
	var buffer bytes.Buffer

	if _, err := io.Copy(&buffer, s.conn); err != nil {
		return nil, err
	}

	return &buffer, nil
}

func (s *FileTransferService) save(fi domain.FilePort, fo string) (string, error) {
	sep := string(os.PathSeparator)
	workDir, err := os.Getwd()
	if err != nil {
		return "", err
	}

	outDir := workDir + sep + os.Getenv("FT_OUT_DIR") + sep + fo + sep

	if err := os.MkdirAll(outDir, os.ModePerm); err != nil {
		return "", err
	}

	fileName := time.Now().Format("2006-01-02T15:04:05")

	if name := fi.GetName(); name != "" {
		fileName = name
	}
	if ext := fi.GetExtension(); ext != "" {
		fileName += "." + ext
	}

	outPath := outDir + fileName
	outFile, err := os.Create(outPath)
	if err != nil {
		return "", err
	}

	defer outFile.Close()

	_, err = io.Copy(outFile, fi.GetBuffer())
	if err != nil {
		return "", err
	}

	return outPath, nil
}
