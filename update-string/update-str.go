package main

import "fmt"

// 修改字符串的错误示例
/*
func main() {
    x := "text"
    x[0] = "T"        // error: cannot assign to x[0]
    fmt.Println(x)
}
*/


//更新字串的正确姿势：将 string 转为 rune slice（此时 1 个 rune 可能占多个 byte），直接更新 rune 中的字符

func main() {
    x := "text"
    fmt.Println("before update:",x)
    xRunes := []rune(x)
    xRunes[0] = '我'
    x = string(xRunes)
    fmt.Println("after update:",x)    // 我ext
}
