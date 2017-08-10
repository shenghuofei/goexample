package main 
import (
    "fmt"
    "time"
)

func do_something(i *int) {
   fmt.Println(*i)
   (*i)++
}

func main(){
   i := 0
   c := time.Tick(10*time.Second)
   for {
       select {
           case <-c:
               go do_something(&i)
       }

   }
}
