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

	dumpPath, err := getDumpPath(ip)
	if err != nil {
		fmt.Printf("\n(%s) Error getting dump path: %v", ip, err)
		return
	}

	var buffer bytes.Buffer

	if bytes, err := io.Copy(&buffer, c); err != nil {
		fmt.Printf("\n(%s) Error copying from connection stream: %v", ip, err)
		return
	} else {
		fmt.Printf("\n(%s) Content received (%.2f mB). Working on it...", ip, float64(bytes)/(1024*1024))
	}

	outFile, err := os.Create(dumpPath + time.Now().Format("2006-01-02"))
	if err != nil {
		fmt.Printf("\n(%s) Error generating destiny file: %v", ip, err)
		return
	}

	defer outFile.Close()

	if _, err := io.Copy(outFile, &buffer); err != nil {
		fmt.Printf("\n(%s) Error saving content: %v", ip, err)
		return
	}

	fmt.Printf("\n(%s) Content saved sucessfully!", ip)
}

func closeConnection(c net.Conn, ip string) {
	fmt.Printf("\nClosing connection with %s", ip)
	c.Close()
}

func getDumpPath(f string) (string, error) {
	sep := string(os.PathSeparator)
	workDir, err := os.Getwd()
	if err != nil {
		return "", err
	}

	dumpDir := os.Getenv("FT_DUMP_DIR")
	dumpPath := workDir + sep + dumpDir + sep + f + sep

	if err := os.MkdirAll(dumpPath, os.ModePerm); err != nil {
		return "", err
	}

	return dumpPath, nil
}
