package main

import (
    "errors"
    "fmt"
    "io/ioutil"
    "net"
    "net/rpc"
    "net/rpc/jsonrpc"
    "os"
    "os/exec"
    "time"
)

type Args struct {
    A, B int
}

type Quotient struct {
    Quo, Rem int
}

type Arith int

func (t *Arith) Multiply(args *Args, reply *int) error {
    *reply = args.A * args.B
    return nil
}

func (t *Arith) Exec(c string, reply *[]byte) error {
    cmd := exec.Command("/bin/sh", "-c", c)
    timeout := time.Duration(5)*time.Second
    tmouterr := errors.New("timeout")
    
    stdout, err := cmd.StdoutPipe()
    if err != nil {
        return err
    }
    
    if err := cmd.Start(); err != nil {
        fmt.Println("Start: ", err.Error())
        return err
    }
    
    *reply, _ = ioutil.ReadAll(stdout)
    fmt.Printf("%s",*reply)
    
    done := make(chan error)
    go func() {
        done <- cmd.Wait()
    }()
    select {
        case <-time.After(timeout):
            if err := cmd.Process.Kill(); err != nil {
                fmt.Printf("failed to kill: %s \"%s\", error: %s \n", cmd.Path, c, err)
            }
            go func() {
                <-done // allow goroutine to exit
            }()
            fmt.Printf("process:%s killed", cmd.Path)
            return tmouterr
        case err := <-done:
            if err != nil {
                fmt.Printf("exec \"%s\" error\n", c)
                return err
            } else {
                fmt.Printf("exec \"%s\" success\n", c)
            }
    }
    return nil 
}

func (t *Arith) Divide(args *Args, quo *Quotient) error {
    if args.B == 0 {
        return errors.New("divide by zero")
    }
    quo.Quo = args.A / args.B
    quo.Rem = args.A % args.B
    return nil
}

func main() {

    arith := new(Arith)
    rpc.Register(arith)

    tcpAddr, err := net.ResolveTCPAddr("tcp", ":1234")
    //tcpAddr, err := net.DialTimeout("tcp", "localhost:1234", 1000*1000*1000*3)
    checkError(err)

    listener, err := net.ListenTCP("tcp", tcpAddr)
    defer listener.Close()
    checkError(err)

    for {
        conn, err := listener.Accept()
        if err != nil {
            continue
        }
        go jsonrpc.ServeConn(conn)
    }

}

func checkError(err error) {
    if err != nil {
        fmt.Println("Fatal error ", err.Error())
        os.Exit(1)
    }
}

