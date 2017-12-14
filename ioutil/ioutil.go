package main

import (
    "io/ioutil"
    "fmt"
    "os"
)

/*

func NopCloser(r io.Reader) io.ReadCloser
NopCloser用一个无操作的Close方法包装r返回一个ReadCloser接口

func ReadAll(r io.Reader) ([]byte, error)
ReadAll从r读取数据直到EOF或遇到error，返回读取的数据和遇到的错误

func ReadFile(filename string) ([]byte, error)
ReadFile 从filename指定的文件中读取数据并返回文件的内容

func WriteFile(filename string, data []byte, perm os.FileMode) error
函数向filename指定的文件中写入数据。如果文件不存在将按给出的权限创建文件，否则在写入数据之前清空文件

func ReadDir(dirname string) ([]os.FileInfo, error)
返回dirname指定的目录的目录信息的有序列表

func TempDir(dir, prefix string) (name string, err error)
在dir目录里创建一个新的、使用prfix作为前缀的临时文件夹，并返回文件夹的路径

func TempFile(dir, prefix string) (f *os.File, err error)
在dir目录下创建一个新的、使用prefix为前缀的临时文件，以读写模式打开该文件并返回os.File指针

type FileInfo interface {
        Name() string       // base name of the file
        Size() int64        // length in bytes for regular files; system-dependent for others
        Mode() FileMode     // file mode bits
        ModTime() time.Time // modification time
        IsDir() bool        // abbreviation for Mode().IsDir()
        Sys() interface{}   // underlying data source (can return nil)
}

*/

func main() {
    con,err := ioutil.ReadFile("./a")
    fmt.Println(string(con),err)
    ioutil.WriteFile("./b",[]byte("test"),0755)
    finfo,err := ioutil.ReadDir("./")
    for _,f := range finfo {
        fmt.Println(f)
        fmt.Println(f.Name(),f.Size(),f.Mode(),f.ModTime())
    }
    tdir,err := ioutil.TempDir("./", "tmpdir")
    fmt.Println(tdir)
    os.Remove(tdir)
    tfile,err := ioutil.TempFile("./", "tmpfile")
    fmt.Println(tfile)
    os.Remove(tfile.Name())
}
