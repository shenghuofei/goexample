package main
//根据客户端不同请求返回不同结果，且是长连接
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
    i := 0
	for {
		read_len, err := conn.Read(request)

		if err != nil {
			fmt.Println(err)
			break
		}
        fmt.Println(read_len)
        i++
        fmt.Println("num",i)
		if read_len == 0 {
			break // connection already closed by client
		} else if strings.TrimSpace(string(request[:read_len])) == "timestamp" {
			daytime := strconv.FormatInt(time.Now().Unix(), 10)
			go func(daytime string){conn.Write([]byte(daytime))}(daytime)
		} else {
			daytime := time.Now().String()
			go func(daytime string){conn.Write([]byte(daytime))}(daytime)
		}
        //每次读取到请求处理完毕后，需要清理request，因为conn.Read()会将新读取到的内容append到原内容之后
		request = make([]byte, 128) // clear last read content
	}
}

func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}
}
