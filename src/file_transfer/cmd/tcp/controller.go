package file_transfer

import (
	"bytes"
	"fmt"
	"io"
	"net"
	"os"
	"slices"
	"strings"
)

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

		if connectionIsAllowed(ip) {
			fmt.Printf("\nStablished connection with %s", ip)
			go handleConnection(conn, ip)
		} else {
			fmt.Printf("\nDenied connection with %s", ip)
			conn.Close()
		}
	}
}

func connectionIsAllowed(ip string) bool {
	ipsWhitelist := strings.Split(os.Getenv("FT_IP_WHITELIST"), ",")

	return slices.Contains(ipsWhitelist, ip)
}

func handleConnection(c net.Conn, ip string) {
	defer closeConnection(c, ip)

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

		var buf bytes.Buffer
		bytes, err := io.Copy(&buf, c)
		if err != nil {
			fmt.Printf("\n(%s) Error copying from connection stream: %v", ip, err)
			return
		}

		fmt.Printf("\n(%s) Content received (%.2f mB). Working on it...", ip, float64(bytes)/(1024*1024))

		outFile, err := os.Create(dumpPath + fileName)
		if err != nil {
			fmt.Printf("\n(%s) Error generating destiny file: %v", ip, err)
			return
		}

		if _, err := io.Copy(outFile, &buf); err != nil {
			outFile.Close()
			fmt.Printf("\n(%s) Error copying from connection stream: %v", ip, err)
			return
		}

		outFile.Close()
		fmt.Printf("\n(%s) Content saved sucessfully!", ip)
	}
}

func closeConnection(c net.Conn, ip string) {
	fmt.Printf("\nClosing connection with %s", ip)
	c.Close()
}

func getDumpPath() (string, error) {
	separator := string(os.PathSeparator)
	workDir, err := os.Getwd()
	if err != nil {
		return "", err
	}

	dumpDir := os.Getenv("FT_DUMP_DIR")
	dumpPath := workDir + separator + dumpDir + separator

	if err := os.MkdirAll(dumpPath, os.ModePerm); err != nil {
		return "", err
	}

	return dumpPath, nil
}

func getFileName(c net.Conn) (string, error) { // TODO: get actual file name from received stream
	return "file.txt", nil
}
