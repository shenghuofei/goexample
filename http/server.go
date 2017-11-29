package main

import (
    "log"
    "net/http"
    "strings"
    "encoding/json"
_    "io/ioutil"
)

func configPushRoutes() {
    http.HandleFunc("/v1/push", func(w http.ResponseWriter, req *http.Request) {
        if req.ContentLength == 0 {
            http.Error(w, "body is blank", http.StatusBadRequest)
            return
        }
        rbody := req.Body // just can read req.Body once
                           
        /* method 1 */
        /*
        var metrics map[string]interface{}
        body, _ := ioutil.ReadAll(rbody)
        log.Println(string(body))
        json.Unmarshal(body, &metrics)
        log.Println("m1",metrics)
        */
        
        /* method 2 */
        var metrics1 map[string]interface{}
        decoder := json.NewDecoder(rbody)
        err := decoder.Decode(&metrics1)
        if err != nil {
                http.Error(w, "connot decode body", http.StatusBadRequest)
                return
        }
        log.Println("m2",metrics1)
        
        w.Write([]byte("success"))
    })
}

func configGetRoutes(){
    http.HandleFunc("/v1/get", func(w http.ResponseWriter, req *http.Request) {
        req.ParseForm()
        log.Println("method:", req.Method) //获取请求的方法

        log.Println("endpoint", req.Form["endpoint"])
        log.Println("metric", req.Form["metric"])

        for k, v := range req.Form {
            log.Print("key:", k, "; ")
            log.Println("val:", strings.Join(v, ""))
        }
        //endpoint = strings.Join(req.Form["endpoint"],"")
        endpoint := req.FormValue("endpoint")
        log.Println("endpoint", endpoint)
        w.Write([]byte("success"))
    })
}

func main() {
    configGetRoutes()
    configPushRoutes()
    addr := ":1989"
    s := &http.Server{
        Addr:           addr,
        MaxHeaderBytes: 1 << 30,
    }

    log.Println("listening", addr)
    log.Fatalln(s.ListenAndServe())
    select {}
}
