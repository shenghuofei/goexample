package main

import (
    "fmt"
    "github.com/streadway/amqp"
    "os"
    "time"
)

var (
    do_chan = make(chan struct{},5) //并发控制
    conn *amqp.Connection
    ch *amqp.Channel
    q *amqp.Queue
    close_chan = make(chan *amqp.Error) // rabbitmq服务器关闭信号
    reset = make(chan int,1)
)

type cdata struct {
    Name string `json:"name"`
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
        return true
    }
    return false
}

func publish(body string){
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
   time.Sleep(20*time.Second)
   <- do_chan
}

func connect() bool {
    conn,err := amqp.Dial("amqp://guest:guest@localhost:5672/")
    if check_err(err,"Failed to connect to mq") { return false }
    //defer conn.Close()
    ch,err = conn.Channel()
    if check_err(err,"failed to open a channel") { return false }
    //defer ch.Close()
    //申明一个队列
    // https://godoc.org/github.com/streadway/amqp#Channel.QueueDeclare
    qtmp, err := ch.QueueDeclare(
            "task_queue", // name  有名字！
            true,         // durable  持久性的,如果事前已经声明了该队列，不能重复声明
            false,        // delete when unused
            false,        // exclusive 如果是真，连接一断开，队列删除
            false,        // no-wait
            nil,          // arguments
    )
    q = &qtmp
    if check_err(err, "Failed to declare a queue") { os.Exit(1) }
    conn.NotifyClose(close_chan)
    return true
}

func main() {
    if !connect() {
        os.Exit(1)
    }
    go func(){
        for {
            for s := range close_chan{
                if s != nil {
                    fmt.Println("info:",s)
                    reset <- 1
                }
            }
            time.Sleep(10*time.Second) 
        }
    }()
    go func(){
        for {
            <- reset
            for {
                if connect() {
                    fmt.Println("reconn success")
                    break
                } else {
                    fmt.Println("reconn fail")
                    //time.Sleep(1*time.Minute)
                    time.Sleep(3*time.Second)
                }
            }
        }
    }() 
    go func() {
        for i:=0;i<10;i++ {
            body := msg{Dtype: "create.docker", Data: cdata{Name: "test", Count: i}}
            sbody := fmt.Sprintf(`{"dtype":"%s","data":{"name":"%s","count":%d}}`,body.Dtype,body.Data.Name,body.Data.Count)
            //body := "hello.world." + fmt.Sprintf("%d",i)
            go publish(sbody)
        }
    }()
    fmt.Printf("To exit press CTRL+C\n")
    select {}
}
