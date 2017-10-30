package main

import (
    "fmt"
    "os"
    "os/signal"
    "syscall"
)

//for gracefull shutdown(net,http): https://github.com/shenghuofei/grace
//for gracefull restart(http): https://github.com/shenghuofei/endless

func wait_signal(){
    // Go signal notification works by sending `os.Signal`
    // values on a channel. We'll create a channel to
    // receive these notifications (we'll also make one to
    // notify us when the program can exit).
    sigs := make(chan os.Signal,1)
  
    // `signal.Notify` registers the given channel to
    // receive notifications of the specified signals.
    signal.Notify(sigs, os.Interrupt, syscall.SIGINT, syscall.SIGTERM, syscall.SIGKILL)
    for {
       sig := <-sigs
       _,err := fmt.Println("get signal:",sig)
       if (err != nil) {
            fmt.Printf("%v\n", err)
            fmt.Printf("unknown signal received: %v\n", sig)
            os.Exit(1)
       } else {
            fmt.Println("do some clean and exit")
            os.Exit(0)
       }
    }
}

func main(){
    fmt.Println("start")
    go wait_signal()
    fmt.Println("do anything")
    select {}
}
