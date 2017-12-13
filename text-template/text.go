package main

import (
    "fmt"
    "text/template"
    "os"
)

/*

Actions:

{{pipeline}}
    pipeline的值的默认文本表示会被拷贝到输出里。
{{if pipeline}} T1 {{end}}
    如果pipeline的值为empty，不产生输出，否则输出T1执行结果。不改变dot的值。
    Empty值包括false、0、任意nil指针或者nil接口，任意长度为0的数组、切片、字典。
{{if pipeline}} T1 {{else}} T0 {{end}}
    如果pipeline的值为empty，输出T0执行结果，否则输出T1执行结果。不改变dot的值。
{{if pipeline}} T1 {{else if pipeline}} T0 {{end}}
    用于简化if-else链条，else action可以直接包含另一个if；等价于：
        {{if pipeline}} T1 {{else}}{{if pipeline}} T0 {{end}}{{end}}
{{range pipeline}} T1 {{end}}
    pipeline的值必须是数组、切片、字典或者通道。
    如果pipeline的值其长度为0，不会有任何输出；
    否则dot依次设为数组、切片、字典或者通道的每一个成员元素并执行T1；
    如果pipeline的值为字典，且键可排序的基本类型，元素也会按键的顺序排序。
{{range pipeline}} T1 {{else}} T0 {{end}}
    pipeline的值必须是数组、切片、字典或者通道。
    如果pipeline的值其长度为0，不改变dot的值并执行T0；否则会修改dot并执行T1。
{{template "name"}}
    执行名为name的模板，提供给模板的参数为nil，如模板不存在输出为""
{{template "name" pipeline}}
    执行名为name的模板，提供给模板的参数为pipeline的值。
{{with pipeline}} T1 {{end}}
    如果pipeline为empty不产生输出，否则将dot设为pipeline的值并执行T1。不修改外面的dot。
{{with pipeline}} T1 {{else}} T0 {{end}}
    如果pipeline为empty，不改变dot并执行T0，否则dot设为pipeline的值并执行T1。

Arguments:

- go语法的布尔值、字符串、字符、整数、浮点数、虚数、复数，视为无类型字面常数，字符串不能跨行
- 关键字nil，代表一个go的无类型的nil值
- 字符'.'（句点，用时不加单引号），代表dot的值
- 变量名，以美元符号起始加上（可为空的）字母和数字构成的字符串，如：$piOver2和$；
  执行结果为变量的值，变量参见下面的介绍
- 结构体数据的字段名，以句点起始，如：.Field；
  执行结果为字段的值，支持链式调用：.Field1.Field2；
  字段也可以在变量上使用（包括链式调用）：$x.Field1.Field2；
- 字典类型数据的键名；以句点起始，如：.Key；
  执行结果是该键在字典中对应的成员元素的值；
  键也可以和字段配合做链式调用，深度不限：.Field1.Key1.Field2.Key2；
  虽然键也必须是字母和数字构成的标识字符串，但不需要以大写字母起始；
  键也可以用于变量（包括链式调用）：$x.key1.key2；
- 数据的无参数方法名，以句点为起始，如：.Method；
  执行结果为dot调用该方法的返回值，dot.Method()；
  该方法必须有1到2个返回值，如果有2个则后一个必须是error接口类型；
  如果有2个返回值的方法返回的error非nil，模板执行会中断并返回给调用模板执行者该错误；
  方法可和字段、键配合做链式调用，深度不限：.Field1.Key1.Method1.Field2.Key2.Method2；
  方法也可以在变量上使用（包括链式调用）：$x.Method1.Field；
- 无参数的函数名，如：fun；
  执行结果是调用该函数的返回值fun()；对返回值的要求和方法一样；函数和函数名细节参见后面。
- 上面某一条的实例加上括弧（用于分组）
  执行结果可以访问其字段或者键对应的值：
      print (.F1 arg1) (.F2 arg2)
      (.StructValuedMethod "arg").Field

func HTMLEscape(w io.Writer, b []byte)
函数向w中写入b的HTML转义等价表示

func HTMLEscapeString(s string) string
返回s的HTML转义等价表示字符串

func JSEscape(w io.Writer, b []byte)
函数向w中写入b的JavaScript转义等价表示

func JSEscapeString(s string) string
返回s的JavaScript转义等价表示字符串

func Must(t *Template, err error) *Template
Must函数用于包装返回(*Template, error)的函数/方法调用，它会在err非nil时panic，一般用于变量初始化:var t = template.Must(template.New("name").Parse("render-data"))

func New(name string) *Template
创建一个名为name的模板

func ParseFiles(filenames ...string) (*Template, error)
ParseFiles函数创建一个模板并解析filenames指定的文件里的模板定义

func (t *Template) Name() string
返回模板t的名字

func (t *Template) Delims(left, right string) *Template
Delims方法用于设置action的分界字符串，应用于之后的Parse、ParseFiles、ParseGlob方法

func (t *Template) Funcs(funcMap FuncMap) *Template
Funcs方法向模板t的函数字典里加入参数funcMap内的键值对

func (t *Template) Clone() (*Template, error)
Clone方法返回模板的一个副本，包括所有相关联的模板

func (t *Template) Lookup(name string) *Template
Lookup方法返回与t关联的名为name的模板，如果没有这个模板返回nil

func (t *Template) Templates() []*Template
Templates方法返回与t相关联的模板的切片，包括t自己

func (t *Template) New(name string) *Template
New方法创建一个和t关联的名字为name的模板并返回它。这种可以传递的关联允许一个模板使用template action调用另一个模板

func (t *Template) Parse(text string) (*Template, error)
Parse方法将字符串text解析为模板

func (t *Template) ParseFiles(filenames ...string) (*Template, error)
ParseFiles方法解析filenames指定的文件里的模板定义并将解析结果与t关联

func (t *Template) ParseGlob(pattern string) (*Template, error)
ParseGlob方法解析匹配pattern的文件里的模板定义并将解析结果与t关联

func (t *Template) Execute(wr io.Writer, data interface{}) (err error)
Execute方法将解析好的模板应用到data上，并将输出写入wr

func (t *Template) ExecuteTemplate(wr io.Writer, name string, data interface{}) error
ExecuteTemplate方法类似Execute，但是使用名为name的t关联的模板产生输出

*/

func example() {
    // Define a template.
    const letter = `
    Dear {{.Name}},
    {{if .Attended}}
        It was a pleasure to see you at the wedding.
    {{else}}
        It is a shame you couldn't make it to the wedding.
    {{end}}
    
    {{with .Gift}} 
        Thank you for the lovely {{.}}. 
    {{end}}
    Best wishes,
    Josie
    `
    // Prepare some data to insert into the template.
    type Recipient struct {
        Name, Gift string
        Attended   bool
    }
    var recipients = []Recipient{
        {"Aunt Mildred", "bone china tea set", true},
        {"Uncle John", "moleskin pants", false},
        {"Cousin Rodney", "", false},
    }
    // Create a new template and parse the letter into it.
    t := template.Must(template.New("lettertpl").Parse(letter))
    // Execute the template for each recipient.
    for _, r := range recipients {
        err := t.Execute(os.Stdout, r)
        if err != nil {
            fmt.Println("executing template:", err)
        }
    }
}

func simple_use() {
    type Inventory struct {
        Material string
        Count    uint
    }
    sweaters := Inventory{"wool", 17}
    tmpl, err := template.New("test").Parse("{{.Count}} of {{.Material}}\n") //test 为模板名
    if err != nil { panic(err) }
    err = tmpl.Execute(os.Stdout, sweaters)
    if err != nil { panic(err) }
}

func just_print(c string) bool {
    fmt.Println("this is my",c)
    if c == "online" {
        return true
    } else {
        return false
    }
}

func main() {
    simple_use() 
    example() 

    s := "<p>after html asdf</p>\n"
    template.HTMLEscape(os.Stdout,[]byte(s))
    hs := template.HTMLEscapeString(s)
    fmt.Println(hs)
    j := "<script> function(){alert('asd')}</script>"
    template.JSEscape(os.Stdout,[]byte(j))
    js := template.JSEscapeString(j)
    fmt.Println(js)

    res := "tpl/res.txt"
    fout,_ := os.Create(res)
    defer fout.Close()
   
    //注册自定义的模板方法 
    funcmap := template.FuncMap{"just_print":just_print}
    
    //New中的名字tpl.txt(不要路径)必须和ParseFiles中的文件名相同
    tmp := template.Must(template.New("tpl.txt").Funcs(funcmap).ParseFiles("tpl/tpl.txt"))
    
    err := tmp.Execute(fout,[]string{"test","online"})
    fmt.Println(err)
}
