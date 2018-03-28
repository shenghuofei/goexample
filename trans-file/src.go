package main

import (
	"fmt"
	"net"
	"os"
    "io"
	"bufio"
_    "io/ioutil"
)

func main() {
    hname,err := os.Hostname()
	checkError(err)
	service := hname+":1200"
	tcpAddr, err := net.ResolveTCPAddr("tcp4", service)
	checkError(err)
	listener, err := net.ListenTCP("tcp", tcpAddr)
	checkError(err)
    fmt.Println("running on ",service)
	for {
		conn, err := listener.Accept()
		if err != nil {
			continue
		}
		go handleClient(conn)
	}
}

func handleClient(conn net.Conn) {
	defer conn.Close()
    fname_byte := make([]byte,10240)
    n,err := conn.Read(fname_byte)
	checkError(err)
    fname := string(fname_byte[:n]) 
    fi,err := os.Open(fname)
    //fi,err := os.Open("src.txt")
	checkError(err)
    defer fi.Close()
    r := bufio.NewReader(fi)
    //chunks := make([]byte,1024,1024)
    buf := make([]byte,1024)
    for{
        n,err := r.Read(buf)
	    checkError(err)
        if 0 ==n {break}
	    conn.Write(buf[:n]) // don't care about return value
        //chunks=append(chunks,buf[:n]...)
        fmt.Println(string(buf[:n]))
    }
    fmt.Println("done")
	//conn.Write([]byte(daytime)) // don't care about return value
	// we're finished with this client
}
func checkError(err error) {
	if err != nil && err != io.EOF {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}
}
