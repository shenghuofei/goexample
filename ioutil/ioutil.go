package main

import (
    "io/ioutil"
    "fmt"
    "os"
)

/*

func ReadFile(filename string) ([]byte, error)
func WriteFile(filename string, data []byte, perm os.FileMode) error
func ReadDir(dirname string) ([]os.FileInfo, error)
func TempDir(dir, prefix string) (name string, err error)
func TempFile(dir, prefix string) (f *os.File, err error)

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
    }
    tdir,err := ioutil.TempDir("./", "tmpdir")
    fmt.Println(tdir)
    os.Remove(tdir)
    tfile,err := ioutil.TempFile("./", "tmpfile")
    fmt.Println(tfile)
    os.Remove(tfile.Name())
}
