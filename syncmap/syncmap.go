package main
import (
    "fmt"
    "sync"
    "time"
)

type syncmap struct {
    m map[string]string
    sync.RWMutex
}

var smap syncmap  //define
var done chan bool //define

func wr() {
    keys := []string{"a","b","c"}
    for _,k := range keys {
        smap.Lock()
        smap.m[k] = k
        smap.Unlock()
        time.Sleep(1*time.Second)
    }
    done <- true
}
func wr1() {
    keys := []string{"aa","bb","cc"}
    for _,k := range keys {
        smap.Lock()
        smap.m[k] = k
        smap.Unlock()
        time.Sleep(1*time.Second)
    }
    done <- true
}
func rd() {
   smap.RLock()
   fmt.Println("rlock")
   for k,v := range smap.m {
       fmt.Println(k,v)
   } 
   smap.RUnlock()
}
func main() {
    smap = syncmap{m:make(map[string]string)} //init
    done = make(chan bool,2) //init
    go wr()
    go wr1()
    for {
        rd()
        if len(done) == 2 {
            fmt.Println(smap.m)
            for k,v := range smap.m {
                fmt.Println(k,v)
            }
            break
        } else {
            time.Sleep(1*time.Second)
        }
    }
}
