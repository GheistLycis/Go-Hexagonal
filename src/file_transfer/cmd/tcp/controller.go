package file_transfer

import (
	"encoding/binary"
	"fmt"
	"io"
	"net"
	"os"
	"slices"
	"strings"
)

const pkgName = "file_transfer"

/*
SetListener handles active listener to receive incoming files from any allowed dials, in parallel.
*/
func HandleServer(l net.Listener) {
	defer l.Close()

	for {
		conn, err := l.Accept()
		if err != nil {
			fmt.Println("Connection error:", err)
		}

		if connectionIsAllowed(conn) {
			go handleConnection(conn)
		} else {
			conn.Close()
		}
	}
}

func connectionIsAllowed(c net.Conn) bool {
	ip := c.RemoteAddr().String()
	ipsWhitelist := strings.Split(os.Getenv("TCP_IP_WHITELIST"), ",")

	return slices.Contains(ipsWhitelist, ip)
}

func handleConnection(c net.Conn) {
	defer closeConnection(c)

	addr := c.RemoteAddr().String()
	dumpPath, err := getDumpPath()
	if err != nil {
		fmt.Printf("(%s) Error getting dump path: %v", addr, err)
		return
	}

	for {
		fileName, err := getFileName(c)
		if err != nil {
			fmt.Printf("(%s) Error reading file name: %v", addr, err)
			return
		}

		fmt.Printf("(%s) Receing content from...", addr)

		outFile, err := os.Create(dumpPath + fileName)
		if err != nil {
			fmt.Printf("(%s) Error generating destiny file: %v", addr, err)
			continue
		}

		if _, err = io.Copy(outFile, c); err != nil {
			outFile.Close()
			fmt.Printf("(%s) Error copying from connection stream: %v", addr, err)
			continue
		}

		outFile.Close()
		fmt.Printf("(%s) Received content sucessfully!", addr)
	}
}

func closeConnection(c net.Conn) {
	fmt.Printf("(%s) Closing connection.", c.RemoteAddr().String())
	c.Close()
}

func getDumpPath() (string, error) {
	separator := string(os.PathSeparator)
	workDir, err := os.Getwd()
	if err != nil {
		return "", err
	}

	currDir := strings.Split(workDir, pkgName)[0]
	dumpDir := os.Getenv("TCP_DUMP_DIR")
	dumpPath := currDir + pkgName + separator + dumpDir + separator

	return dumpPath, nil
}

func getFileName(c net.Conn) (string, error) {
	var nameLen int32
	if err := binary.Read(c, binary.LittleEndian, &nameLen); err != nil {
		return "", err
	}

	nameBuff := make([]byte, nameLen)
	if _, err := io.ReadFull(c, nameBuff); err != nil {
		return "", err
	}

	return string(nameBuff), nil
}
