package file_transfer

import (
	"fmt"
	"io"
	"net"
	"os"
	"slices"
	"strconv"
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

	file, err := domain.NewFile("", "")
	if err != nil {
		fmt.Printf("\n(%s) Error creating file: %v", s.peerIp, err)
		return
	}
	err = s.download(file)
	if err != nil {
		fmt.Printf("\n(%s) Error downloading content: %v", s.peerIp, err)
		return
	}

	outDir, err := s.save(file, s.peerIp)
	if err != nil {
		fmt.Printf("\n(%s) Error saving content: %v", s.peerIp, err)
		return
	}

	fmt.Printf("\n(%s) Content saved sucessfully (%s)", s.peerIp, outDir)
}

func (s *FileTransferService) shutConnection() {
	fmt.Printf("\nClosing connection with %s", s.peerIp)
	s.conn.Close()
}

func (s *FileTransferService) peerIsTrusted() bool {
	ipsWhitelist := strings.Split(os.Getenv("FT_IP_WHITELIST"), ",")

	return slices.Contains(ipsWhitelist, s.peerIp)
}

func (s *FileTransferService) download(f domain.FilePort) error { // TODO: use parallelism for faster transfer
	chunkSize, err := strconv.ParseInt(os.Getenv("FT_CHUNK_MB"), 10, 64)
	if err != nil {
		return err
	} else {
		fmt.Println("CHUNK:", chunkSize)
		chunkSize = chunkSize * 1024 * 1024
		fmt.Println("CHUNK:", chunkSize)
	}

	for {
		fmt.Printf("\nReceiving file... (TOTAL = %.2f mB)", float64(f.GetBuffer().Len())/(1024*1024))

		chunk := make([]byte, chunkSize)
		n, err := s.conn.Read(chunk)

		if n > 0 {
			if _, err := f.WriteBuffer(chunk[:n]); err != nil {
				return err
			}
		}
		if err != nil {
			if err == io.EOF {
				break
			}

			return err
		}
	}

	return nil
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
