
package main

import (
	"flag"
	"fmt"
	"os"
	"time"
)

//usage: go run flagset.go -ver 123 -time 12m -a asdf asf asdfasf

/*

func NewFlagSet(name string, errorHandling ErrorHandling) *FlagSet
NewFlagSet创建一个新的、名为name，采用errorHandling为错误处理策略的FlagSet

func (f *FlagSet) Init(name string, errorHandling ErrorHandling)
Init设置flag集合f的名字和错误处理属性。FlagSet零值没有名字，默认采用ContinueOnError错误处理策略

func (f *FlagSet) NFlag() int
NFlag返回解析时进行了设置的flag的数量

func (f *FlagSet) Lookup(name string) *Flag
返回f中已注册flag的Flag结构体指针；如果flag不存在的话，返回nil

func (f *FlagSet) NArg() int
NArg返回解析flag之后剩余参数的个数

func (f *FlagSet) Args() []string
返回解析之后剩下的非flag参数

func (f *FlagSet) Arg(i int) string
返回解析之后剩下的第i个参数，从0开始索引

func (f *FlagSet) PrintDefaults()
PrintDefault打印集合中所有注册好的flag的默认值

func (f *FlagSet) SetOutput(output io.Writer)
设置使用信息和错误信息的输出流，如果output为nil，将使用os.Stderr

func (f *FlagSet) Bool(name string, value bool, usage string) *bool
Bool用指定的名称、默认值、使用信息注册一个bool类型flag。返回一个保存了该flag的值的指针

func (f *FlagSet) BoolVar(p *bool, name string, value bool, usage string)
BoolVar用指定的名称、默认值、使用信息注册一个bool类型flag，并将flag的值保存到p指向的变量

func (f *FlagSet) Int(name string, value int, usage string) *int
Int用指定的名称、默认值、使用信息注册一个int类型flag。返回一个保存了该flag的值的指针

func (f *FlagSet) IntVar(p *int, name string, value int, usage string)
IntVar用指定的名称、默认值、使用信息注册一个int类型flag，并将flag的值保存到p指向的变量

func (f *FlagSet) Int64(name string, value int64, usage string) *int64
Int64用指定的名称、默认值、使用信息注册一个int64类型flag。返回一个保存了该flag的值的指针

func (f *FlagSet) Int64Var(p *int64, name string, value int64, usage string)
Int64Var用指定的名称、默认值、使用信息注册一个int64类型flag，并将flag的值保存到p指向的变量

func (f *FlagSet) Uint(name string, value uint, usage string) *uint
Uint用指定的名称、默认值、使用信息注册一个uint类型flag。返回一个保存了该flag的值的指针

func (f *FlagSet) UintVar(p *uint, name string, value uint, usage string)
UintVar用指定的名称、默认值、使用信息注册一个uint类型flag，并将flag的值保存到p指向的变量

func (f *FlagSet) Uint64(name string, value uint64, usage string) *uint64
Uint64用指定的名称、默认值、使用信息注册一个uint64类型flag。返回一个保存了该flag的值的指针

func (f *FlagSet) Uint64Var(p *uint64, name string, value uint64, usage string)
Uint64Var用指定的名称、默认值、使用信息注册一个uint64类型flag，并将flag的值保存到p指向的变量

func (f *FlagSet) Float64(name string, value float64, usage string) *float64
Float64用指定的名称、默认值、使用信息注册一个float64类型flag。返回一个保存了该flag的值的指针

func (f *FlagSet) Float64Var(p *float64, name string, value float64, usage string)
Float64Var用指定的名称、默认值、使用信息注册一个float64类型flag，并将flag的值保存到p指向的变量

func (f *FlagSet) String(name string, value string, usage string) *string
String用指定的名称、默认值、使用信息注册一个string类型flag。返回一个保存了该flag的值的指针

func (f *FlagSet) StringVar(p *string, name string, value string, usage string)
StringVar用指定的名称、默认值、使用信息注册一个string类型flag，并将flag的值保存到p指向的变量

func (f *FlagSet) Duration(name string, value time.Duration, usage string) *time.Duration
Duration用指定的名称、默认值、使用信息注册一个time.Duration类型flag。返回一个保存了该flag的值的指针

func (f *FlagSet) DurationVar(p *time.Duration, name string, value time.Duration, usage string)
DurationVar用指定的名称、默认值、使用信息注册一个time.Duration类型flag，并将flag的值保存到p指向的变量

func (f *FlagSet) Var(value Value, name string, usage string)
Var方法使用指定的名字、使用信息注册一个flag。该flag的类型和值由第一个参数表示，该参数应实现了Value接口

func (f *FlagSet) Set(name, value string) error
设置已注册的flag的值

func (f *FlagSet) Parse(arguments []string) error
从arguments中解析注册的flag

func (f *FlagSet) Parsed() bool
返回是否f.Parse已经被调用过

func (f *FlagSet) Visit(fn func(*Flag))
按照字典顺序遍历标签，并且对每个标签调用fn。 这个函数只遍历解析时进行了设置的标签

func (f *FlagSet) VisitAll(fn func(*Flag))
按照字典顺序遍历标签，并且对每个标签调用fn。 这个函数会遍历所有标签，不管解析时有无进行设置

*/


var (
/*
	参数解析出错时错误处理方式
	switch f.errorHandling {
		case ContinueOnError:
			return err
		case ExitOnError:
			os.Exit(2)
		case PanicOnError:
			panic(err)
	} 
*/

	//flagSet = flag.NewFlagSet(os.Args[0],flag.PanicOnError) 
	flagSet = flag.NewFlagSet(os.Args[0],flag.ExitOnError) 
	//flagSet = flag.NewFlagSet("xcl",flag.ExitOnError) 
	verFlag = flagSet.String("ver", "", "version")
	xtimeFlag  = flagSet.Duration("time", 10*time.Minute, "time Duration")

	addrFlag = StringArray{}
)

func init() {
	flagSet.Var(&addrFlag, "a", "b")
}

func main() {
	fmt.Println("os.Args[0]:", os.Args[0])
	flagSet.Parse(os.Args[1:]) //flagSet.Parse(os.Args[0:])
    flagSet.PrintDefaults()
    fmt.Println("lookup ver:",flagSet.Lookup("ver"))
	fmt.Println("当前命令行参数类型个数:", flagSet.NFlag())  
    fmt.Println("参数值:")
	fmt.Println("ver:", *verFlag)
	fmt.Println("xtimeFlag:", *xtimeFlag)
	fmt.Println("addrFlag:",addrFlag.String())

    fmt.Println("after parse flagSet args:",flagSet.NArg())

    //非flagSet参数
	for i,param := range flagSet.Args(){
        fmt.Printf("---#%d :%s\n",i,param)
    }
}


type StringArray []string

func (s *StringArray) String() string {
	return fmt.Sprint([]string(*s))
}

func (s *StringArray) Set(value string) error {
	*s = append(*s, value)
	return nil
}
