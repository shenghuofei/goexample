package main

import (
    "fmt"
    "log"
    "net/rpc/jsonrpc"
    "os"
    "net"
)

type Args struct {
    A, B int
}

type Quotient struct {
    Quo, Rem int
}

func main() {
    if len(os.Args) != 3 {
        fmt.Println("Usage: ", os.Args[0], `"server:port"`, `"cmd"`)
        log.Fatal(1)
    }
    service := os.Args[1]
    cmd := os.Args[2]
     
    var err error
    //client, err := jsonrpc.Dial("tcp", service)
    client, err := net.DialTimeout("tcp", service, 1000*1000*1000*3) // 30秒超时时间
    clientRpc := jsonrpc.NewClient(client)
    defer client.Close()
    if err != nil {
        log.Fatal("dialing:", err)
    }
    // Synchronous call
    args := Args{17, 8}
    var reply int
    //err = client.Call("Arith.Multiply", args, &reply)
    err = clientRpc.Call("Arith.Multiply", args, &reply)
    if err != nil {
        log.Fatal("arith error:", err)
    }
    fmt.Printf("Arith: %d*%d=%d\n", args.A, args.B, reply)

    var quot Quotient
    //err = client.Call("Arith.Divide", args, &quot)
    err = clientRpc.Call("Arith.Divide", args, &quot)
    if err != nil {
        log.Fatal("arith error:", err)
    }
    fmt.Printf("Arith: %d/%d=%d remainder %d\n", args.A, args.B, quot.Quo, quot.Rem)
    //cmd := "ls"
    var replys []byte
    //err = client.Call("Arith.Exec", cmd, &replys)
    err = clientRpc.Call("Arith.Exec", cmd, &replys)
    if err != nil {
        log.Fatal("arith error:", err)
    }
    fmt.Printf("Arith: %s result %s\n", cmd, replys)
}
