package main

import (
    "fmt"
    "os"
    "bufio"
    "strings"
    "strconv"
)

/*
func NewReader(rd io.Reader) *Reader
NewReader创建一个具有默认大小缓冲、从rd读取的*Reader

func NewReaderSize(rd io.Reader, size int) *Reader
NewReaderSize创建一个具有最少有size尺寸的缓冲、从rd读取的*Reader

func (b *Reader) Reset(r io.Reader)
Reset丢弃缓冲中的数据，清除任何错误，将b重设为其下层从r读取数据

func (b *Reader) Buffered() int
Buffered返回缓冲中现有的可读取的字节数

func (b *Reader) Peek(n int) ([]byte, error)
Peek返回输入流的下n个字节，而不会移动读取位置

func (b *Reader) Read(p []byte) (n int, err error)
Read读取数据写入p

func (*Reader) ReadByte
ReadByte读取并返回一个字节。如果没有可用的数据，会返回错误

func (b *Reader) UnreadByte() error
UnreadByte吐出最近一次读取操作读取的最后一个字节

func (b *Reader) ReadLine() (line []byte, isPrefix bool, err error)
ReadLine是一个低水平的行数据读取原语。大多数调用者应使用ReadBytes('\n')或ReadString('\n')代替，或者使用Scanner

func (b *Reader) ReadSlice(delim byte) (line []byte, err error)
ReadSlice读取直到第一次遇到delim字节，返回缓冲里的包含已读取的数据和delim字节的切片

func (b *Reader) ReadString(delim byte) (line string, err error)
ReadString读取直到第一次遇到delim字节，返回一个包含已读取的数据和delim字节的字符串

func NewWriter(w io.Writer) *Writer
NewWriter创建一个具有默认大小缓冲、写入w的*Writer

func NewWriterSize(w io.Writer, size int) *Writer
NewWriterSize创建一个具有最少有size尺寸的缓冲、写入w的*Writer

func (b *Writer) Reset(w io.Writer)
Reset丢弃缓冲中的数据，清除任何错误，将b重设为将其输出写入w

func (b *Writer) Buffered() int
Buffered返回缓冲中已使用的字节数

func (b *Writer) Available() int
Available返回缓冲中还有多少字节未使用

func (b *Writer) Write(p []byte) (nn int, err error)
Write将p的内容写入缓冲

func (b *Writer) WriteString(s string) (int, error)
WriteString写入一个字符串

func (b *Writer) WriteByte(c byte) error
WriteByte写入单个字节

func (b *Writer) Flush() error
Flush方法将缓冲中的数据写入下层的io.Writer接口

func (b *Writer) ReadFrom(r io.Reader) (n int64, err error)
ReadFrom实现了io.ReaderFrom接口

func NewReadWriter(r *Reader, w *Writer) *ReadWriter
NewReadWriter申请创建一个新的、将读写操作分派给r和w 的ReadWriter

func ScanBytes(data []byte, atEOF bool) (advance int, token []byte, err error)
ScanBytes是用于Scanner类型的分割函数（符合SplitFunc），本函数会将每个字节作为一个token返回

func ScanWords(data []byte, atEOF bool) (advance int, token []byte, err error)
ScanRunes是用于Scanner类型的分割函数（符合SplitFunc），本函数会将空白（参见unicode.IsSpace）分隔的片段（去掉前后空白后）作为一个token返回

func NewScanner(r io.Reader) *Scanner
NewScanner创建并返回一个从r读取数据的Scanner，默认的分割函数是ScanLines

func (s *Scanner) Split(split SplitFunc)
Split设置该Scanner的分割函数。本方法必须在Scan之前调用

func (s *Scanner) Scan() bool
Scan方法获取当前位置的token（该token可以通过Bytes或Text方法获得），并让Scanner的扫描位置移动到下一个token

func (s *Scanner) Bytes() []byte
Bytes方法返回最近一次Scan调用生成的token

func (s *Scanner) Text() string
Text方法返回最近一次Scan调用生成的token，会申请创建一个字符串保存token并返回该字符串

*/

func scanner_custom(){
    // An artificial input source.
    const input = "1234 5678 1234567901234567890"
    scanner := bufio.NewScanner(strings.NewReader(input))
    // Create a custom split function by wrapping the existing ScanWords function.
    split := func(data []byte, atEOF bool) (advance int, token []byte, err error) {
        advance, token, err = bufio.ScanWords(data, atEOF)
        if err == nil && token != nil {
            _, err = strconv.ParseInt(string(token), 10, 32)
        }
        return
    }
    // Set the split function for the scanning operation.
    scanner.Split(split)
    // Validate the input
    for scanner.Scan() {
        fmt.Printf("scanner custom %s\n", scanner.Text())
    }
    if err := scanner.Err(); err != nil {
        fmt.Printf("scanner custon Invalid input: %s\n", err)
    }
}

func scanner_lines(){
    fmt.Println("input sumething:")
    scanner := bufio.NewScanner(os.Stdin)
    count := 0
    for scanner.Scan() {
        fmt.Println("scanner lines:",scanner.Text()) // Println will add back the final '\n'
        count += 1
        if count >= 1 { break }
    }
    if err := scanner.Err(); err != nil {
        fmt.Fprintln(os.Stderr, "scanner lines reading standard input:", err)
    }
}

func scanner_words(){
    // An artificial input source.
    const input = "Now is the winter of our discontent,\nMade glorious summer by this sun of York.\n"
    scanner := bufio.NewScanner(strings.NewReader(input))
    // Set the split function for the scanning operation.
    scanner.Split(bufio.ScanWords)
    // Count the words.
    count := 0
    for scanner.Scan() {
        count++
    }
    if err := scanner.Err(); err != nil {
        fmt.Fprintln(os.Stderr, "scanner words reading input:", err)
    }
    fmt.Printf("scanner words %d\n", count)
}

func main() {
    scanner_custom()
    scanner_lines()
    scanner_words()
    fd,_ := os.Open("r.txt")
    reader := bufio.NewReader(fd)
    reader.Reset(fd)
    p,_ := reader.Peek(4)
    fmt.Println("peek:",string(p))
    fmt.Println("reader bufferd:",reader.Buffered())
    n,_ := reader.Read(p)
    fmt.Println("reader read:",n,string(p))
    c,_ := reader.ReadByte()
    fmt.Println("reader readbyte:",c)
    l,ispre,_ := reader.ReadLine()
    fmt.Println("reader readline:",string(l),ispre)
    l,_ = reader.ReadSlice('\n')
    fmt.Println("reader readslice:",string(l))
    fmt.Println(reader.ReadString('\n'))

    wd,_ := os.OpenFile("w.txt", os.O_RDWR|os.O_CREATE, 0755) 
    w := bufio.NewWriterSize(wd,10) //buffer size 10
    fmt.Fprint(w, "Hello, ")
    fmt.Fprint(w, "world!\n")
    w.Flush() // Don't forget to flush!
    w.Write([]byte("7777\n"))  //buffer use 4
    fmt.Println("write buffered:",w.Buffered())  //userd 4
    fmt.Println("write buffered available:",w.Available()) //able to use 6
    w.WriteString("8888\n")
    fmt.Println(w.ReadFrom(fd))
    w.Flush() // Don't forget to flush!
    
}
