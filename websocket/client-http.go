package main

import (
    "net/http"
    "log"
    "fmt"
    "os"
)

//读取模板文件直接返回给客户端
func ppp(w http.ResponseWriter, req *http.Request){
    res := "index.html"
    //var sres string //将文件内容读入sres变量
    fin,_ := os.Open(res)
    buf := make([]byte, 1024)
    for{
        n,_ := fin.Read(buf)
        if 0 == n {break}
        //sres += string(buf[:n])
        //fmt.Println(sres)
        fmt.Fprintf(w,string(buf[:n]))
        //os.Stdout.Write(buf[:n])
    }
}

func main(){
    http.HandleFunc("/", ppp)
    log.Fatal(http.ListenAndServe("localhost:8000", nil))
}
