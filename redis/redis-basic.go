package main

import (
    "github.com/garyburd/redigo/redis"
    "fmt"
    "time"
)

func main() {
    //Connect
    //conn, err := redis.Dial("tcp", "localhost:6379")
    conn, err := redis.DialTimeout("tcp", "localhost:6379",60*time.Second,60*time.Second,60*time.Second)
    if err != nil {
        panic(err)
    }
    defer conn.Close()

    _, err = conn.Do("ping")
    if err != nil {
	fmt.Println("[ERROR] ping redis fail", err)
    }

    /*
    response, err := conn.Do("AUTH", "YOUR_PASSWORD")

    if err != nil {
        panic(err)
    }

    fmt.Printf("Connected! ", response)
    */

    v,err := conn.Do("SET", "best_car_ever", "Tesla Model S")
    fmt.Printf("%v,%v\n",v,err)
    v,err = conn.Do("GET", "best_car_ever")
    fmt.Printf("%s,%v\n",v,err)
    vs,err := redis.String(v,nil)
    fmt.Println(vs,err)

    exist,err := redis.Bool(conn.Do("EXISTS","best_car_ever"))
    fmt.Println(exist,err)
}

