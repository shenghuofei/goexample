package main

import (
    "github.com/garyburd/redigo/redis"
    "fmt"
    "time"
    "strconv"
)

var pool *redis.Pool

func newPool(addr string) {
    pool = &redis.Pool{
        MaxIdle: 3,
        IdleTimeout: 240 * time.Second,
        //Dial: func () (redis.Conn, error) { return redis.Dial("tcp", addr) },
        Dial: func () (redis.Conn, error) {
            c, err := redis.Dial("tcp", addr)
            if err != nil {
              return nil, err
            }
            /*
            if _, err := c.Do("AUTH", password); err != nil {
              c.Close()
              return nil, err
            }
            */
            //select db 3
            if _, err := c.Do("SELECT", 3); err != nil {
              c.Close()
              return nil, err
            }
            return c, nil
        },
        TestOnBorrow: func(c redis.Conn, t time.Time) error {
            if time.Since(t) < time.Minute {
                return nil
            }
            _, err := c.Do("PING")
            return err
        },

   }
}

func Lock() bool {
    timeout := time.Duration(3)*time.Second
    conn := pool.Get()
    defer conn.Close()
    for {
        lock_value := time.Now().Add(timeout).Unix() 
        // flag,err := redis.Int64(conn.Do("SETNX","lock",lock_value))
        flag,err := conn.Do("SET","lock",lock_value,"EX","3","NX")
        if err != nil { 
            fmt.Println("get lock setnx fail",err)
            return false 
        } 
        value,err := redis.String(conn.Do("GET","lock"))   
        if err != nil { 
            fmt.Println("get lock value fail",err)
            return false 
        } 
        now := strconv.FormatInt(time.Now().Unix(),10)
        // if flag == 1 {  //成功获得锁
        if flag == "OK" {
            fmt.Println("get lock success")
            return true
        } else if now > value {  //别人锁超时了，我也可以获得，防止加锁成功解锁失败的情况
            last_value, err := redis.String(conn.Do("GETSET","lock",lock_value))
            if (err != nil || now < last_value) {  //别人已经获得锁了,或者设置失败
                fmt.Println("get lock getset fail",err)
                return false 
            } else {
                fmt.Println("get lock success")
                return true
            }
        }else {
            fmt.Println("wait lock ...")
            time.Sleep(500*time.Millisecond)
        }
    }
    return false
}

func Unlock() bool {
    conn := pool.Get()
    defer conn.Close()
    _,err := conn.Do("DEL","lock")
    if err != nil{
        fmt.Println("del lock fail",err) 
    }
    fmt.Println("unlock success") 
    return true
}

func do_job(){
    Lock()
    fmt.Println("do job")
    time.Sleep(2*time.Second)
    fmt.Println("job done")
    Unlock()
}

func main(){
    newPool("localhost:6379")
    for i:=0;i<10;i++{
        go do_job()
    }
    select{}
}
