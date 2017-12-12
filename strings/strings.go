package main

import (
    "strings"
    "fmt"
)

/*

func EqualFold(s, t string) bool
判断两个utf-8编码字符串（将unicode大写、小写、标题三种格式字符视为相同）是否相同

func HasPrefix(s, prefix string) bool
判断s是否有前缀字符串prefix

func HasSuffix(s, suffix string) bool
判断s是否有后缀字符串suffix

func Contains(s, substr string) bool
判断字符串s是否包含子串substr

func ContainsAny(s, chars string) bool
判断字符串s是否包含字符串chars中的任一字符

func Count(s, sep string) int
返回字符串s中有几个不重复的sep子串

func Index(s, sep string) int
子串sep在字符串s中第一次出现的位置，不存在则返回-1

func LastIndex(s, sep string) int
子串sep在字符串s中最后一次出现的位置，不存在则返回-1

func Title(s string) string
返回s中每个单词的首字母都改为标题格式的字符串拷贝

func ToTitle(s string) string
返回将所有字母都转为对应的标题版本的拷贝

func ToLower(s string) string
返回将所有字母都转为对应的小写版本的拷贝

func ToUpper(s string) string
返回将所有字母都转为对应的大写版本的拷贝

func Repeat(s string, count int) string
返回count个s串联的字符串

func Replace(s, old, new string, n int) string
返回将s中前n个不重叠old子串都替换为new的新字符串，如果n<0会替换所有old子串

func Trim(s string, cutset string) string
返回将s前后端所有cutset包含的utf-8码值都去掉的字符串

func TrimSpace(s string) string
返回将s前后端所有空白（unicode.IsSpace指定）都去掉的字符串

func TrimLeft(s string, cutset string) string
返回将s前端所有cutset包含的utf-8码值都去掉的字符串

func TrimPrefix(s, prefix string) string
返回去除s可能的前缀prefix的字符串

func TrimRight(s string, cutset string) string
返回将s后端所有cutset包含的utf-8码值都去掉的字符串

func TrimSuffix(s, suffix string) string
返回去除s可能的后缀suffix的字符串

func Fields(s string) []string
返回将字符串按照空白（unicode.IsSpace确定，可以是一到多个连续的空白字符）分割的多个字符串。如果字符串全部是空白或者是空字符串的话，会返回空切片

func Split(s, sep string) []string
用去掉s中出现的sep的方式进行分割，会分割到结尾，并返回生成的所有片段组成的切片（每一个sep都会进行一次切割，即使两个sep相邻，也会进行两次切割）。如果sep为空字符，Split会将s切分成每一个unicode码值一个字符串

func SplitN(s, sep string, n int) []string
用去掉s中出现的sep的方式进行分割，会分割到结尾，并返回生成的所有片段组成的切片（每一个sep都会进行一次切割，即使两个sep相邻，也会进行两次切割）。如果sep为空字符，Split会将s切分成每一个unicode码值一个字符串。参数n决定返回的切片的数目

func SplitAfter(s, sep string) []string
用从s中出现的sep后面切断的方式进行分割，会分割到结尾，并返回生成的所有片段组成的切片

func SplitAfterN(s, sep string, n int) []string
用从s中出现的sep后面切断的方式进行分割，会分割到结尾，并返回生成的所有片段组成的切片

func Join(a []string, sep string) string
将一系列字符串连接为一个字符串，之间用sep来分隔

func NewReader(s string) *Reader
NewReader创建一个从s读取数据的Reader。本函数类似bytes.NewBufferString，但是更有效率，且为只读的

func NewReplacer(oldnew ...string) *Replacer
使用提供的多组old、new字符串对创建并返回一个*Replacer。替换是依次进行的，匹配时不会重叠

func (r *Replacer) Replace(s string) string
Replace返回s的所有替换进行完后的拷贝

func (r *Replacer) WriteString(w io.Writer, s string) (n int, err error)
WriteString向w中写入s的所有替换进行完后的拷贝

*/

func main() {
    fmt.Println(strings.EqualFold("Go", "go"))    
    fmt.Println(strings.HasPrefix("Go is good", "Go"))    
    fmt.Println(strings.Contains("seafood", "foo"))
    fmt.Println(strings.ContainsAny("failure", "d & u"))
    fmt.Println(strings.Count("cheese", "ee"))
    fmt.Println(strings.Index("chicken", "ken"))
    fmt.Println(strings.LastIndex("go gopher", "go"))
    fmt.Println(strings.Title("her royal highness"))
    fmt.Println(strings.ToLower("Gopher"))
    fmt.Println(strings.ToUpper("Gopher"))
    fmt.Println(strings.ToTitle("loud noises"))
    fmt.Println("ba" + strings.Repeat("na", 2))
    fmt.Println(strings.Replace("oink oink oink", "k", "ky", 2))
    fmt.Println(strings.Replace("oink oink oink", "oink", "moo", -1))
    fmt.Printf("[%q]\n", strings.Trim(" !!! Achtung! Achtung! !!! ", "! "))
    fmt.Println(strings.TrimSpace(" \t\n a lone gopher \n\t\r\n"))
    fmt.Println(strings.TrimPrefix("Goodbye,hello, world!", "Goodbye,"))
    fmt.Println(strings.TrimSuffix("Hello, goodbye, etc!", "goodbye, etc!"))
    fmt.Printf("Fields are: %q\n", strings.Fields("  foo bar  baz   "))
    fmt.Printf("%q\n", strings.Split("a,b,c", ","))
    fmt.Printf("%q\n", strings.Split("a man a plan a canal panama", "a "))
    fmt.Printf("%q\n", strings.Split(" xyz ", ""))
    fmt.Printf("%q\n", strings.Split("", "Bernardo O'Higgins"))
    fmt.Printf("%q\n", strings.SplitN("a,b,c", ",", 2))
    fmt.Printf("%q\n", strings.SplitAfter("a,b,c", ","))
    fmt.Printf("%q\n", strings.SplitAfterN("a,b,c", ",", 2))
    s := []string{"foo", "bar", "baz"}
    fmt.Println(strings.Join(s, ", "))
    rd := strings.NewReader("just test reader")
    fmt.Println(rd.Len())
    r := strings.NewReplacer("<", "&lt;", ">", "&gt;", "is", "Is")
    fmt.Println(r.Replace("This is <b>HTML</b>!"))
}
