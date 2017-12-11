package main

import (
    "strconv"
    "fmt"
)

/*

func ParseInt(s string, base int, bitSize int) (i int64, err error) 
base指定进制（2到36），如果base为0，则会从字符串前置判断，"0x"是16进制，"0"是8进制，否则是10进制
bitSize指定结果必须能无溢出赋值的整数类型，0、8、16、32、64 分别代表 int、int8、int16、int32、int64

func ParseUint(s string, base int, bitSize int) (n uint64, err error)
ParseUint类似ParseInt但不接受正负号，用于无符号整型

func ParseFloat(s string, bitSize int) (f float64, err error)
如果s合乎语法规则，函数会返回最为接近s表示值的一个浮点数（使用IEEE754规范舍入）。bitSize指定了期望的接收类型，32是float32（返回值可以不改变精确值的赋值给float32），64是float64

func FormatBool(b bool) string
根据b的值返回"true"或"false"

func FormatInt(i int64, base int) string
返回i的base进制的字符串表示。base 必须在2到36之间，结果中会使用小写字母'a'到'z'表示大于10的数字

func FormatUint(i uint64, base int) string
是FormatInt的无符号整数版本

func FormatFloat(f float64, fmt byte, prec, bitSize int) string
函数将浮点数表示为字符串并返回
bitSize表示f的来源类型（32：float32、64：float64），会据此进行舍入
fmt表示格式：'f'（-ddd.dddd）、'b'（-ddddp±ddd，指数为二进制）、'e'（-d.dddde±dd，十进制指数）、'E'（-d.ddddE±dd，十进制指数）、'g'（指数很大时用'e'格式，否则'f'格式）、'G'（指数很大时用'E'格式，否则'f'格式）。
prec控制精度（排除指数部分）：对'f'、'e'、'E'，它表示小数点后的数字个数；对'g'、'G'，它控制总的数字个数。如果prec 为-1，则代表使用最少数量的、但又必需的数字来表示f

func Atoi(s string) (i int, err error)
Atoi是ParseInt(s, 10, 0)的简写

func Itoa(i int) string
Itoa是FormatInt(i, 10) 的简写

func AppendBool(dst []byte, b bool) []byte
等价于append(dst, FormatBool(b)...),bool -> []byte

func AppendInt(dst []byte, i int64, base int) []byte
等价于append(dst, FormatInt(I, base)...),int -> []byte

func AppendUint(dst []byte, i uint64, base int) []byte
等价于append(dst, FormatUint(I, base)...), uint -> []byte

func AppendFloat(dst []byte, f float64, fmt byte, prec int, bitSize int) []byte
等价于append(dst, FormatFloat(f, fmt, prec, bitSize)...),float -> []byte

func AppendQuote(dst []byte, s string) []byte
等价于append(dst, Quote(s)...),string -> []byte
*/

func main() {
    fmt.Println(strconv.ParseInt("-23",10,64))
    fmt.Println(strconv.ParseUint("23",10,64))
    fmt.Println(strconv.ParseFloat("23.4",64))
    fmt.Println(strconv.FormatBool(true))
}
