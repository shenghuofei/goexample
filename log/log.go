package main

import (
    "log"
    "bytes"
    "fmt"
    "os"
)

/*
const (
    // 字位共同控制输出日志信息的细节。不能控制输出的顺序和格式。
    // 在所有项目后会有一个冒号：2009/01/23 01:23:23.123123 /a/b/c/d.go:23: message
    Ldate         = 1 << iota     // 日期：2009/01/23
    Ltime                         // 时间：01:23:23
    Lmicroseconds                 // 微秒分辨率：01:23:23.123123（用于增强Ltime位）
    Llongfile                     // 文件全路径名+行号： /a/b/c/d.go:23
    Lshortfile                    // 文件无路径名+行号：d.go:23（会覆盖掉Llongfile）
    LstdFlags     = Ldate | Ltime // 标准logger的初始值
)

func New(out io.Writer, prefix string, flag int) *Logger
New创建一个Logger。参数out设置日志信息写入的目的地。参数prefix会添加到生成的每一条日志前面。参数flag定义日志的属性（时间、文件等等）

func (l *Logger) Flags() int
Flags返回logger的输出选项

func (l *Logger) SetFlags(flag int)
SetFlags设置logger的输出选项

func (l *Logger) Prefix() string
Prefix返回logger的输出前缀

func (l *Logger) SetPrefix(prefix string)
SetPrefix设置logger的输出前缀

func (l *Logger) Output(calldepth int, s string) error
Output写入输出一次日志事件

func (l *Logger) Printf(format string, v ...interface{})
Printf调用l.Output将生成的格式化字符串输出到logger，参数用和fmt.Printf相同的方法处理

func (l *Logger) Print(v ...interface{})
Print调用l.Output将生成的格式化字符串输出到logger，参数用和fmt.Print相同的方法处理

func (l *Logger) Println(v ...interface{})
Println调用l.Output将生成的格式化字符串输出到logger，参数用和fmt.Println相同的方法处理

func (l *Logger) Fatalf(format string, v ...interface{})
Fatalf等价于{l.Printf(v...); os.Exit(1)}

func (l *Logger) Fatal(v ...interface{})
Fatal等价于{l.Print(v...); os.Exit(1)}

func (l *Logger) Fatalln(v ...interface{})
Fatalln等价于{l.Println(v...); os.Exit(1)}

func (l *Logger) Panicf(format string, v ...interface{})
Panicf等价于{l.Printf(v...); panic(...)}

func (l *Logger) Panic(v ...interface{})
Panic等价于{l.Print(v...); panic(...)}

func (l *Logger) Panicln(v ...interface{})
Panicln等价于{l.Println(v...); panic(...)}

func Flags() int
Flags返回标准logger的输出选项

func SetFlags(flag int)
SetFlags设置标准logger的输出选项

func Prefix() string
Prefix返回标准logger的输出前缀

func SetPrefix(prefix string)
SetPrefix设置标准logger的输出前缀

func SetOutput(w io.Writer)
SetOutput设置标准logger的输出目的地，默认是标准错误输出

func Printf(format string, v ...interface{})
Printf调用Output将生成的格式化字符串输出到标准logger，参数用和fmt.Printf相同的方法处理

func Print(v ...interface{})
Print调用Output将生成的格式化字符串输出到标准logger，参数用和fmt.Print相同的方法处理

func Println(v ...interface{})
Println调用Output将生成的格式化字符串输出到标准logger，参数用和fmt.Println相同的方法处理

func Fatalf(format string, v ...interface{})
Fatalf等价于{Printf(v...); os.Exit(1)}

func Fatal(v ...interface{})
Fatal等价于{Print(v...); os.Exit(1)}

func Fatalln(v ...interface{})
Fatalln等价于{Println(v...); os.Exit(1)}

func Panicf(format string, v ...interface{})
Panicf等价于{Printf(v...); panic(...)}

func Panic(v ...interface{})
Panic等价于{Print(v...); panic(...)}

func Panicln(v ...interface{})
Panicln等价于{Println(v...); panic(...)}

*/


func main() {
    var buf bytes.Buffer
    logger := log.New(&buf, "logger: ", log.Lshortfile|log.LstdFlags)
    logger.Print("Hello, log file!")
    fmt.Print(&buf)
    fmt.Println(logger.Flags())
    logger.SetFlags(1<<3)
    logger.SetPrefix("[log] ")
    buf.Reset()
    logger.Print("Hello, log file 1!")
    fmt.Print(&buf)
    fmt.Println(logger.Flags())
    fmt.Println(logger.Prefix())
   
    //std info 
    fmt.Println(log.Flags())
    fmt.Println(log.Prefix())
    log.SetFlags(log.LstdFlags)
    log.SetPrefix("[logger] ")
    log.SetOutput(os.Stdout)
    log.Println("this is std log")
}
