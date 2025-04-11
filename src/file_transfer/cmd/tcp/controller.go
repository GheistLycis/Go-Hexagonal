package file_transfer

import (
	app "Go-Hexagonal/src/file_transfer/app"
	"fmt"
	"log"
	"net"
	"os"
	"slices"
	"strconv"
	"strings"
	"time"

	"github.com/joho/godotenv"
)

var ipsWhitelist []string
var timeOut time.Duration

func init() {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error loading .env file - %v", err)
	}

	ipsWhitelist = strings.Split(os.Getenv("FT_IP_WHITELIST"), ",")

	timeOutMins, err := strconv.Atoi(os.Getenv("FT_TIMEOUT_MINS"))
	if err != nil {
		log.Fatalf("Failed to parse ENV variable FT_TIMEOUT_MINS - %v", err)
	}
	timeOut = time.Duration(timeOutMins)
}

func HandleServer(l net.Listener) {
	defer l.Close()

	for {
		conn, err := l.Accept()
		if err != nil {
			fmt.Println("Connection error:", err)
			return
		}
		peerIp := strings.Split(conn.RemoteAddr().String(), ":")[0]

		defer shutConnection(conn, peerIp)

		if peerIsTrusted(peerIp) {
			fmt.Printf("\nEstablished connection with %s", peerIp)
			conn.SetReadDeadline(time.Now().Add(timeOut * time.Minute))
			go app.NewFileReceiverService(conn).HandleConnection()
		} else {
			fmt.Printf("\nDenied connection with %s", peerIp)
			return
		}
	}
}

func peerIsTrusted(ip string) bool {
	return slices.Contains(ipsWhitelist, ip)
}

func shutConnection(c net.Conn, ip string) {
	fmt.Printf("\nClosing connection with %s", ip)
	c.Close()
}
