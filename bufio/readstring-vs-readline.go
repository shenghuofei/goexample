package main

import (
    "bufio"
    "fmt"
    "os"
    "time"
)

func ReadString(filename string) {
    f,err := os.Open(filename)
    if err != nil {
        fmt.Println("open",filename,"fail",err)
        os.Exit(1)
    }
    defer f.Close()
    r := bufio.NewReader(f)
    for {
        c,err := r.ReadString('\n')
        if err != nil {
            break
        }
        fmt.Println(c)
    }
}

func ReadLine(filename string) {
    f,err := os.Open(filename)
    if err != nil {
        fmt.Println("open",filename,"fail",err)
        os.Exit(1)
    }
    r := bufio.NewReader(f)
    for {
        c,err := readLine(r)
        if err != nil {
            break
        } 
        fmt.Println(c)
    }
}

//此函数主要解决单行字节数大于4096的情况
func readLine(r *bufio.Reader) (string,error) {
    line, isprefix, err := r.ReadLine()
    for isprefix && err == nil {
        var bs []byte
        bs, isprefix, err = r.ReadLine()
        line = append(line,bs...)
    }
    return string(line),err
}

func main() {
    filename := "./bufio.go"
    s := time.Now()
    ReadString(filename)
    e1 := time.Now()
    fmt.Printf("readstring:%v\n",e1.Sub(s))
    ReadLine(filename)
    e2 := time.Now()
    fmt.Printf("readline:%v\n",e2.Sub(e1))
}
