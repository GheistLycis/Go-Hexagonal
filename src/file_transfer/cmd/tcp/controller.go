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
			return
		}

		ip := strings.Split(conn.RemoteAddr().String(), ":")[0]

		if connectionIsAllowed(conn) {
			fmt.Printf("\nStablished connection with %s", ip)
			go handleConnection(conn)
		} else {
			fmt.Printf("\nDenied connection with %s", ip)
			conn.Close()
		}
	}
}

func connectionIsAllowed(c net.Conn) bool {
	ip := strings.Split(c.RemoteAddr().String(), ":")[0]
	ipsWhitelist := strings.Split(os.Getenv("FT_IP_WHITELIST"), ",")

	return slices.Contains(ipsWhitelist, ip)
}

func handleConnection(c net.Conn) {
	defer closeConnection(c)

	ip := strings.Split(c.RemoteAddr().String(), ":")[0]

	dumpPath, err := getDumpPath()
	if err != nil {
		fmt.Printf("\n(%s) Error getting dump path: %v", ip, err)
		return
	}

	for {
		fileName, err := getFileName(c)
		if err != nil {
			fmt.Printf("\n(%s) Error reading file name: %v", ip, err)
			return
		}

		fmt.Printf("\n(%s) Receiving content from...", ip)

		outFile, err := os.Create(dumpPath + fileName)
		if err != nil {
			fmt.Printf("\n(%s) Error generating destiny file: %v", ip, err)
			return
		}

		if _, err = io.Copy(outFile, c); err != nil {
			outFile.Close()
			fmt.Printf("\n(%s) Error copying from connection stream: %v", ip, err)
			return
		}

		outFile.Close()
		fmt.Printf("\n(%s) Received content sucessfully!", ip)
	}
}

func closeConnection(c net.Conn) {
	ip := strings.Split(c.RemoteAddr().String(), ":")[0]

	fmt.Printf("\nClosing connection with %s", ip)
	c.Close()
}

func getDumpPath() (string, error) {
	separator := string(os.PathSeparator)
	workDir, err := os.Getwd()
	if err != nil {
		return "", err
	}

	currDir := strings.Split(workDir, pkgName)[0]
	dumpDir := os.Getenv("FT_DUMP_DIR")
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
