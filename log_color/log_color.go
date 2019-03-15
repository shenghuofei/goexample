package main
import (
    "fmt"
    "time"
)

func log_color(level string, msg interface{}) {
    now := time.Now().Format("2006-01-02 15:04:05")
    red := func(msg interface{}) {
        fmt.Printf("\x1b[31m[%v] %v\n\x1b[0m", now, msg)
    }
    green := func(msg interface{}) {
        fmt.Printf("\x1b[32m[%v] %v\n\x1b[0m", now, msg)
    }
    blue := func(msg interface{}) {
        fmt.Printf("\x1b[34m[%v] %v\n\x1b[0m", now, msg)
    }
    yellow := func(msg interface{}) {
        fmt.Printf("\x1b[33m[%v] %v\n\x1b[0m", now, msg)
    }
    level_color_map := map[string]func(msg interface{}){"error": red, "success": green, "notice": blue, "warning": yellow}
    if color, ok := level_color_map[level]; ok {
        color(msg)
    } else {
        fmt.Printf("[%v] %v\n", now, msg)
    }
}

func main() {
    log_color("error", "asdfasf")
    log_color("success", "asdfasf")
    log_color("warning", "asdfasf")
    log_color("notice", "asdfasf")
    log_color("default", "asdfasf")
}
