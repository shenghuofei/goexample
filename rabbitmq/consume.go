package main

import (
    "fmt"
    "os"
    "time"
    "bytes"
    "encoding/json"
    "github.com/streadway/amqp"
)

type cdata struct {
    Nmae string `json:"name"`
    Count int   `json:"count"`
}

type msg struct {
    Dtype string `json:"dtype"`
    Data cdata   `json:"data"`
}

func check_err(err error,msg string){
    if err != nil {
        fmt.Println(msg)
        fmt.Println(err)
        os.Exit(1)
    }
}

func main(){
    conn,err := amqp.Dial("amqp://guest:guest@localhost:5672/")
    check_err(err,"failed to connet to mq")
    defer conn.Close()
    ch,err := conn.Channel()
    check_err(err,"failed to open a channel")
    defer ch.Close()
    
    // 指定队列！
    q, err := ch.QueueDeclare(
        "task_queue", // name
        true,         // durable
        false,        // delete when unused
        false,        // exclusive
        false,        // no-wait
        nil,          // arguments
    )
    check_err(err,"failed to declare a queue")
    
    // Fair dispatch 预取，每个工作方每次拿一个消息，确认后才拿下一次，缓解压力
    err = ch.Qos(
        1,     // prefetch count, 1:确保消费完一条消息之后再去获取下一条, 0:直接获取下一条
        // 待解释
        0,     // prefetch size, 消息预取窗口大小, 0:无特殊限制, 如果设置了no-ack选项，则预取大小将被忽略
        false, // global
    )
    check_err(err,"failed to set QoS")

    // 消费根据队列名
    msgs, err := ch.Consume(
        q.Name, // queue
        "",     // consumer
        false,  // auto-ack   设置为真自动确认消息
        false,  // exclusive
        false,  // no-local
        false,  // no-wait
        nil,    // args
    )
    check_err(err,"failed to register a consumer")

    forever := make(chan bool)
    
    go func() {
        for d := range msgs {
            fmt.Printf("Received a message: %s\n", d.Body)
           
            var json_data msg
            err = json.Unmarshal(d.Body,&json_data)
            check_err(err,"failed to convert body to json")
            fmt.Println(json_data)
            
            dot_count := bytes.Count(d.Body, []byte("."))
            t := time.Duration(dot_count)
            time.Sleep(t * time.Second)
            fmt.Println("Done")
            
            // 确认消息被收到！！如果为真的，那么同在一个channel,在该消息之前未确认的消息都会确认，适合批量处理
            // 真时场景：每十条消息确认一次，类似
            d.Ack(false)
        }
    }()
    fmt.Printf(" [*] Waiting for messages. To exit press CTRL+C\n")
    <-forever
}
