package main

import (
    "github.com/garyburd/redigo/redis"
    "fmt"
    "time"
    "encoding/json"
)

type Aa struct {
	Id   int
	Prio int
}

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

func reconnect(){
    for {
        conn := pool.Get()
        _, err := conn.Do("PING")
        if err != nil {
            fmt.Println("reconnect...")
            newPool("localhost:6379")
        } else {
            fmt.Println("alive...")
        }
        time.Sleep(3*time.Second)
        conn.Close()
    }
}

func main(){
    newPool("localhost:6379")
    conn := pool.Get()
    defer conn.Close()
    go reconnect()
    v,err := conn.Do("SET","name","test")
    fmt.Println(v,err)
    v,err = redis.String(conn.Do("GET","name"))
    fmt.Println(v,err)
    
    exist,err := redis.Bool(conn.Do("EXISTS","name"))
    fmt.Println(exist,err) 

    _, err = conn.Do("HSET", "norecovery", "user1", "asdfasf")
    if err != nil {
	    fmt.Printf("push norecovery alarm from redis fail: %s", err)
    }
  
    eventsMap, err := redis.StringMap(conn.Do("HGETALL", "norecovery"))
    if err != nil {
	    fmt.Printf("get norecovery alarm from redis fail: %s", err)
    }
    fmt.Printf("%v\n",eventsMap)

    domain := Aa{Id: 1, Prio: 1}
    bs, _ := json.Marshal(domain)
    _,err = conn.Do("LPUSH","lowqueue",string(bs))
    if err != nil {
	    fmt.Printf("push lowqueue alarm from redis fail: %s", err)
    }
    v,err = conn.Do("LRANGE","lowqueue",0,10)
    if err != nil {
	    fmt.Printf("get lowqueue alarm from redis fail: %s", err)
    }else {
        res,_ := redis.Strings(v,nil)
        fmt.Println(res)
    }
    v,err = conn.Do("LPOP","lowqueue")
    if err != nil {
        fmt.Printf("pop lowqueue alarm from redis fail: %s", err)
    }else {
        res1 := Aa{}
        res,_ := redis.String(v,nil)
        fmt.Println(res)
        res2,_ := redis.Bytes(v,nil)
        json.Unmarshal(res2,&res1)
        fmt.Println(res1)
    }
    select {}
}
