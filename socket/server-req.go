package main
//根据客户端不同请求返回不同结果
import (
	"fmt"
	"net"
	"os"
	"time"
	"strconv"
	"strings"
)

func main() {
	service := ":1200"
	tcpAddr, err := net.ResolveTCPAddr("tcp", service)
	checkError(err)
	listener, err := net.ListenTCP("tcp", tcpAddr)
	checkError(err)
	for {
		conn, err := listener.Accept()
		if err != nil {
			continue
		}
		go handleClient(conn)
	}
}

func handleClient(conn net.Conn) {
	conn.SetReadDeadline(time.Now().Add(2 * time.Minute)) // set 2 minutes timeout
	request := make([]byte, 128) // set maxium request length to 128B to prevent flood attack
	defer conn.Close()  // close connection before exit
	read_len, err := conn.Read(request)

	if err != nil {
		fmt.Println(err)
	}

	if read_len == 0 {
		fmt.Println("coon closed")
		//break // connection already closed by client
	} else if strings.TrimSpace(string(request[:read_len])) == "timestamp" {
		daytime := strconv.FormatInt(time.Now().Unix(), 10)
		conn.Write([]byte(daytime))
	} else {
		daytime := time.Now().String()
		conn.Write([]byte(daytime))
	}
}

func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}
}
