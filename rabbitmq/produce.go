package main

import (
    "fmt"
    "github.com/streadway/amqp"
    "os"
    "time"
)

type cdata struct {
    Name string `json:"name"`
    Count int   `json:"count"`
}

type msg struct {
    Dtype string `json:"dtype"`
    Data cdata   `json:"data"`
}

func check_err(err error,msg string) {
    if err != nil {
        fmt.Println(msg)
        os.Exit(1)
    }
}

var do_chan = make(chan struct{},5) //并发控制

func publish(ch *amqp.Channel,q amqp.Queue,body string){
    do_chan <- struct{}{}
    //发布
    err := ch.Publish(
        "",  //exchange,默认模式，exchange为空
        q.Name,  //routing key 默认模式路由到同名队列，即是task_queue
        false,   //mandatory
        false,   //immediate
        amqp.Publishing{
            // 持久性的发布，因为队列被声明为持久的，发布消息必须加上这个（可能不用），但消息还是可能会丢，如消息到缓存但MQ挂了来不及持久化。
            DeliveryMode: amqp.Persistent,
            ContentType:  "text/plain",
            Body:         []byte(body),
        })
   check_err(err,"failed to publish a msg")
   fmt.Printf("[x] sent %s\n",body)
   time.Sleep(2*time.Second)
   <- do_chan
}

func main() {
    conn,err := amqp.Dial("amqp://guest:guest@localhost:5672/")
    check_err(err,"Failed to connect to mq")
    defer conn.Close()
    ch,err := conn.Channel()
    check_err(err,"failed to open a channel")
    defer ch.Close()
    //申明一个队列
    // https://godoc.org/github.com/streadway/amqp#Channel.QueueDeclare
    q, err := ch.QueueDeclare(
            "task_queue", // name  有名字！
            true,         // durable  持久性的,如果事前已经声明了该队列，不能重复声明
            false,        // delete when unused
            false,        // exclusive 如果是真，连接一断开，队列删除
            false,        // no-wait
            nil,          // arguments
    )
    check_err(err, "Failed to declare a queue")
    
    go func() {
        for i:=0;i<10;i++ {
            body := msg{Dtype: "create.docker", Data: cdata{Name: "test", Count: i}}
            sbody := fmt.Sprintf(`{"dtype":"%s","data":{"name":"%s","count":%d}}`,body.Dtype,body.Data.Name,body.Data.Count)
            //body := "hello.world." + fmt.Sprintf("%d",i)
            go publish(ch,q,sbody)
        }
    }()
    fmt.Printf("To exit press CTRL+C\n")
    select {}
}
