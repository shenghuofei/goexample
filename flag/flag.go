package main

import (
    "fmt"
    "flag"
)

//usage:go run flag.go -flagname 1 asdf 

/*
ErrorHandling定义如何处理flag解析错误
const (
    ContinueOnError ErrorHandling = iota
    ExitOnError
    PanicOnError
)

Value接口是用于将动态的值保存在一个flag里
type Value interface {
    String() string
    Set(string) error
}


func NFlag() int
NFlag返回已被设置的flag的数量

func Lookup(name string) *Flag
返回已经已注册flag的Flag结构体指针；如果flag不存在的话，返回nil

func NArg() int
NArg返回解析flag之后剩余参数的个数

func Args() []string
返回解析之后剩下的非flag参数

func Arg(i int) string
返回解析之后剩下的第i个参数，从0开始索引

func PrintDefaults()
PrintDefault会向标准错误输出写入所有注册好的flag的默认值

func Bool(name string, value bool, usage string) *bool
Bool用指定的名称、默认值、使用信息注册一个bool类型flag。返回一个保存了该flag的值的指针

func BoolVar(p *bool, name string, value bool, usage string)
BoolVar用指定的名称、默认值、使用信息注册一个bool类型flag，并将flag的值保存到p指向的变量

func Int(name string, value int, usage string) *int
Int用指定的名称、默认值、使用信息注册一个int类型flag。返回一个保存了该flag的值的指针

func IntVar(p *int, name string, value int, usage string)
IntVar用指定的名称、默认值、使用信息注册一个int类型flag，并将flag的值保存到p指向的变量

func Int64(name string, value int64, usage string) *int64
Int64用指定的名称、默认值、使用信息注册一个int64类型flag。返回一个保存了该flag的值的指针

func Int64Var(p *int64, name string, value int64, usage string)
Int64Var用指定的名称、默认值、使用信息注册一个int64类型flag，并将flag的值保存到p指向的变量

func Uint(name string, value uint, usage string) *uint
Uint用指定的名称、默认值、使用信息注册一个uint类型flag。返回一个保存了该flag的值的指针

func UintVar(p *uint, name string, value uint, usage string)
UintVar用指定的名称、默认值、使用信息注册一个uint类型flag，并将flag的值保存到p指向的变量

func Uint64(name string, value uint64, usage string) *uint64
Uint64用指定的名称、默认值、使用信息注册一个uint64类型flag。返回一个保存了该flag的值的指针

func Uint64Var(p *uint64, name string, value uint64, usage string)
Uint64Var用指定的名称、默认值、使用信息注册一个uint64类型flag，并将flag的值保存到p指向的变量

func Float64(name string, value float64, usage string) *float64
Float64用指定的名称、默认值、使用信息注册一个float64类型flag。返回一个保存了该flag的值的指针

func Float64Var(p *float64, name string, value float64, usage string)
Float64Var用指定的名称、默认值、使用信息注册一个float64类型flag，并将flag的值保存到p指向的变量

func String(name string, value string, usage string) *string
String用指定的名称、默认值、使用信息注册一个string类型flag。返回一个保存了该flag的值的指针

func StringVar(p *string, name string, value string, usage string)
StringVar用指定的名称、默认值、使用信息注册一个string类型flag，并将flag的值保存到p指向的变量

func Duration(name string, value time.Duration, usage string) *time.Duration
Duration用指定的名称、默认值、使用信息注册一个time.Duration类型flag。返回一个保存了该flag的值的指针

func DurationVar(p *time.Duration, name string, value time.Duration, usage string)
DurationVar用指定的名称、默认值、使用信息注册一个time.Duration类型flag，并将flag的值保存到p指向的变量

func Var(value Value, name string, usage string)
Var方法使用指定的名字、使用信息注册一个flag。该flag的类型和值由第一个参数表示，该参数应实现了Value接口

func Set(name, value string) error
设置已注册的flag的值

func Parse()
从os.Args[1:]中解析注册的flag。必须在所有flag都注册好而未访问其值时执行。未注册却使用flag -help时，会返回ErrHelp

func Parsed() bool
返回是否Parse已经被调用过

func Visit(fn func(*Flag))
按照字典顺序遍历标签，并且对每个标签调用fn。 这个函数只遍历解析时进行了设置的标签

func VisitAll(fn func(*Flag))
按照字典顺序遍历标签，并且对每个标签调用fn。 这个函数会遍历所有标签，不管解析时有无进行设置
*/

//for -h|--help message
func usage() {
    fmt.Println("usage:go run flag.go -flagname 1 asdf") 
}

func main() {
    flag.Usage = usage
    //使用flag.String(), Bool(), Int()等函数注册flag
    var ip = flag.Int("flagname", 1234, "help message for flagname")
   
    //也可以将flag绑定到一个变量，使用Var系列函数
    var flagvar int
    flag.IntVar(&flagvar, "flagname1", 1234, "help message for flagname")
   
    //在所有flag都注册之后，调用flag.Parse()
    flag.Parse()
    flag.PrintDefaults()
    fmt.Println("seted flag num:",flag.NFlag())
    fmt.Println("lookup flagname:",flag.Lookup("flagname"))

    fmt.Println("ip:",*ip)
    fmt.Println("flagvar:",flagvar)
    
    //获取非flag参数 
    fmt.Println("after flag parse num:",flag.NArg()) //非flag参数
    fmt.Println(flag.Args())
    fmt.Println(flag.Arg(0))
}
