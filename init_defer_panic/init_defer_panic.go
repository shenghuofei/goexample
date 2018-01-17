package main

import "fmt"

func init() {
    fmt.Println("init1")
}

//同一个包可以有多个init函数，执行顺序从前到后
func init() {
    fmt.Println("init2")
}

//defer 执行顺序与语句顺序相反，且在return前执行，panic在return的时候执行
func defer_call() {
    defer func() { fmt.Println("before") }()
    defer func() { fmt.Println("run") }()
    defer func() { fmt.Println("after") }()

    panic("panic")
}

func main() {
    defer_call()
}
