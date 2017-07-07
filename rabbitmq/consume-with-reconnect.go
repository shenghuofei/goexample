package main

import (
    "fmt"
    "os"
    "time"
    "bytes"
    "encoding/json"
    "github.com/streadway/amqp"
)

var (
    conn *amqp.Connection
    ch *amqp.Channel
//    msgs *(<- chan amqp.Delivery)
    close_chan = make(chan *amqp.Error) // rabbitmq服务器关闭信号
    reset = make(chan int,1)
)

type cdata struct {
    Nmae string `json:"name"`
    Count int   `json:"count"`
}

type msg struct {
    Dtype string `json:"dtype"`
    Data cdata   `json:"data"`
}

func check_err(err error,msg string) bool {
    if err != nil {
        fmt.Println(msg)
        fmt.Println(err)
        return false 
    }
    return true
}

func do_job(d *amqp.Delivery) {
    fmt.Println("start job")
    var json_data msg
    err := json.Unmarshal(d.Body,&json_data)
    if !check_err(err,"failed to convert body to json") { os.Exit(1) }
    fmt.Println(json_data)
    
    dot_count := bytes.Count(d.Body, []byte("."))
    t := time.Duration(dot_count)
    time.Sleep(t * time.Second)
    fmt.Println("Done")
    
    // 确认消息被收到！！如果为真的，那么同在一个channel,在该消息之前未确认的消息都会确认，适合批量处理
    // 真时场景：每十条消息确认一次，类似
    d.Ack(false)
}

func connect() bool {
    conn,err := amqp.Dial("amqp://guest:guest@localhost:5672/")
    if !check_err(err,"failed to connet to mq") { return false }
    //defer conn.Close()
    ch,err := conn.Channel()
    if !check_err(err,"failed to open a channel") { return false }
    //defer ch.Close()
    
    // 指定队列！
    q, err := ch.QueueDeclare(
        "task_queue", // name
        true,         // durable
        false,        // delete when unused
        false,        // exclusive
        false,        // no-wait
        nil,          // arguments
    )
    if !check_err(err,"failed to declare a queue") { return false }
    
    // Fair dispatch 预取，每个工作方每次拿一个消息，确认后才拿下一次，缓解压力
    err = ch.Qos(
        1,     // prefetch count,1:确保消费完一条消息之后再去获取下一条,0:直接获取下一条
        // 待解释
        0,     // prefetch size
        false, // global
    )
    if !check_err(err,"failed to set QoS") { return false }
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
    go func() {
        for d := range msgs {
            fmt.Printf("Received a message: %s\n", d.Body)
            go do_job(&d) 
        }
    }()
    if !check_err(err,"failed to register a consumer") { return false }
    conn.NotifyClose(close_chan) //如果连接断开给close_chan发送信号
    return true 
}

func main(){
    if !connect() { 
        os.Exit(1)
    }
    go func(){
        for {
            for s := range close_chan {//收到连接断开的信号发送给reset进行重连
                if s != nil { 
                    fmt.Println("info:",s)
                    reset <- 1
                }else {
                    fmt.Println("info:",s)
                }
            }
            time.Sleep(10*time.Second) 
        }
    }()
    go func(){
        for {
            <- reset //收到重连信号，进行重连
            for {  //不停重试直到连接成功
                if connect() {
                    fmt.Println("reconn success")
                    break
                }else { 
                    fmt.Println("reconn fail")
                    //time.Sleep(1*time.Minute) 
                    time.Sleep(3*time.Second) 
                }
            }
        }
    }()
    
    forever := make(chan bool)
    fmt.Printf(" [*] Waiting for messages. To exit press CTRL+C\n")
    <-forever
}
