//http://www.jianshu.com/p/bee02c18b221
package main

import (
    "html/template"
    "net/http"
    "log"
    "fmt"
    "os"
)

type tpl string

func testf(name string)string{
    return fmt.Sprintf("%ss",name)
}

//模板渲染后写入文件
func alist(w http.ResponseWriter, req *http.Request){
    res := "res.html"
    fout,_ := os.Create(res)
    defer fout.Close()
    tmp := template.Must(template.ParseFiles("tpl.html"))
    tmp.Execute(fout, []string{"Han Meimei","Lilie"})
}

//http渲染模板并返回给客户端
func (t tpl)list(w http.ResponseWriter, req *http.Request){
    /*funcs := template.FuncMap{
        "multi": testf,
    }
    tmp.Funcs(funcs)
    tmp := template.Must(template.New("").Funcs(funcs).ParseFiles("tpl.html"))
    */
    tmp := template.Must(template.ParseFiles("tpl.html"))
    //tmp, _ := template.New("").Funcs(template.FuncMap{"multi":testf,}).ParseFiles("tpl.html")
    tmp.Execute(w, []string{"Han Meimei","Lilie"})
}

//读取模板文件直接返回给客户端
func ppp(w http.ResponseWriter, req *http.Request){
    //fmt.Fprintf(w,"this is just for %s","test")
    res := "res.html"
    var sres string //将文件内容读入sres变量
    fin,_ := os.Open(res)
    buf := make([]byte, 1024)
    for{
        n,_ := fin.Read(buf)
        if 0 == n {break}
        sres += string(buf[:n])
        fmt.Println(sres)
        fmt.Fprintf(w,string(buf[:n]))
        //os.Stdout.Write(buf[:n])
    }
}

func router(t tpl){
    http.HandleFunc("/", ppp)
    http.HandleFunc("/list", t.list)
    http.HandleFunc("/r", alist)
}

func main(){
    var t tpl
    router(t)
    log.Fatal(http.ListenAndServe("localhost:8000", nil))
}
