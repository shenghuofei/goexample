package main

import (
	"fmt"
_	"io/ioutil"
    "io"
	"net"
	"os"
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
    handlleServer(conn,"dest.txt")
}

func handlleServer(conn *net.TCPConn,fname string) {
    fd,err := os.OpenFile(fname,os.O_RDWR|os.O_CREATE|os.O_APPEND,0644)
	checkError(err)
    _,err = conn.Write([]byte("src.txt"))
	checkError(err)
    buf := make([]byte, 1024)
    for {
        n,err := conn.Read(buf)
	    checkError(err)
        if n == 0 {
            break 
        } else {
            //ioutil.WriteFile("dest.txt",buf[:n],0644)
            fd.Write(buf[:n])
        }
    }
}

func checkError(err error) {
	if err != nil && err != io.EOF {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}
}
