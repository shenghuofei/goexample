package main

import (
    "fmt"
    "net"
)

func main(){
    ip := "256.0.0.0"
    if res := net.ParseIP(ip); res == nil {
        fmt.Println("not ip format")
    } else {
        fmt.Println("ip ok")
    }
}
