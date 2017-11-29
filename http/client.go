package  main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/url"
	"strings"
	"time"
    "log"
    "io/ioutil"
    "errors"
)

func request(method, requrl, contentType string, data map[string]interface{}, timeout int) (error,[]byte) {
	var req *http.Request
	var err error
    var res []byte
	switch method {
	case "GET":
		req, err = http.NewRequest(method, requrl, nil)
		if err != nil {
			log.Printf("Request url %s fail: new request error: %s", requrl, err)
			return err,res
		}
	case "POST":
		if contentType == "application/x-www-form-urlencoded" {
            v := url.Values{}
			for key, value := range data {
				vs, b := value.(string)
				if !b {
					log.Printf("Request url %s fail: post args key %s to value %v is not string",requrl, key, value)
					return err,res
				}
			    v.Set(key, vs)
			}
			req, err = http.NewRequest(method, requrl, bytes.NewBufferString(v.Encode()))
			if err != nil {
				log.Printf("Request url %s fail: new request error: %s", requrl, err)
				return err,res
			}
			req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
		} else if contentType == "application/json" {
			body, err := json.Marshal(data)
            //log.Println(string(body))
			if err != nil {
				log.Printf("Request url %s fail: parse body data %v error: %s", requrl, data, err)
				return err,res
			}
			req, err = http.NewRequest(method, requrl, strings.NewReader(string(body)))
			//req, err = http.NewRequest(method, requrl, bytes.NewBuffer(body))
			if err != nil {
				log.Printf("Request url %s fail: new request error: %s", requrl, err)
				return err,res
			}
			req.Header.Set("Content-Type", "application/json")
		} else {
			log.Printf("Request url %s fail: http Content-Type %s is not supported", requrl, method)
			return errors.New("content-type err"),res
		}
	default:
		log.Printf("Request url %s fail: http method %s is not supported", requrl, method)
		return errors.New("method not supported"),res
	}
	client := &http.Client{}
	if timeout != 0 {
		client.Timeout = time.Duration(timeout) * time.Second
	}
	resp, err := client.Do(req)
	if err != nil {
		log.Printf("Request url %s fail: query error: %s", requrl, err)
        return err,res
	}
    defer resp.Body.Close()
    body,err := ioutil.ReadAll(resp.Body)
    if err != nil {
        log.Printf("Request url %s get request body fail", requrl)
        return err,res
    }
    return nil,body
}


func request_get_test() bool{
    api := "http://127.0.0.1:1989/v1/get?endpoint=value&metric=cpu"
    args := map[string]interface{}{}
    err,res := request("GET",api,"application/json",args,3)
    if err != nil {
        log.Println("api fail")
        return false
    }
    log.Println(string(res))
    return true
}

func request_post_test() bool{
    api := "http://127.0.0.1:1989/v1/push"
    args := map[string]interface{}{"endpoint":"value","metric":"cpu"}
    err,res := request("POST",api,"application/json",args,3)
    if err != nil {
        log.Println("api fail")
        return false
    }
    log.Println(string(res))
    return true
}

func main(){
    request_get_test()
    request_post_test()
}
