package main

import (
	"fmt"
_	"io/ioutil"
	"net"
	"os"
    "time"
    "strings"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "Usage: %s host:port ", os.Args[0])
		os.Exit(1)
	}
	service := os.Args[1]
	tcpAddr, err := net.ResolveTCPAddr("tcp", service)
	checkError(err)
	conn, err := net.DialTCP("tcp", nil, tcpAddr)
	checkError(err)
    request := make([]byte, 128)
	/*_, err = conn.Write([]byte("HEAD / HTTP/1.0\r\n\r\n"))
	checkError(err)
	result, err := ioutil.ReadAll(conn)
	checkError(err)
	fmt.Println(string(result))*/
    for {
	    go func(){
            _, err = conn.Write([]byte("timestamp"))
	        checkError(err) 
        }()
        read_len, err := conn.Read(request)
	    checkError(err)
        res := strings.TrimSpace(string(request[:read_len]))
        fmt.Println(res)
	    /*go func(conn net.Conn){
            result, err := ioutil.ReadAll(conn)
	        checkError(err)
	        fmt.Println(string(result))
        }(conn)*/
        //time.Sleep(3*time.Minute)
        time.Sleep(3*time.Second)
    }
    os.Exit(0)
}
func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}
}
