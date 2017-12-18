package main

import (
    "os"
    "log"
)

/*

const (
    O_RDONLY int = syscall.O_RDONLY // 只读模式打开文件
    O_WRONLY int = syscall.O_WRONLY // 只写模式打开文件
    O_RDWR   int = syscall.O_RDWR   // 读写模式打开文件
    O_APPEND int = syscall.O_APPEND // 写操作时将数据附加到文件尾部
    O_CREATE int = syscall.O_CREAT  // 如果不存在将创建一个新文件
    O_EXCL   int = syscall.O_EXCL   // 和O_CREATE配合使用，文件必须不存在
    O_SYNC   int = syscall.O_SYNC   // 打开文件用于同步I/O
    O_TRUNC  int = syscall.O_TRUNC  // 如果可能，打开时清空文件
)

const (
    SEEK_SET int = 0 // 相对于文件起始位置seek
    SEEK_CUR int = 1 // 相对于文件当前位置seek
    SEEK_END int = 2 // 相对于文件结尾位置seek
)

const (
    PathSeparator     = '/' // 操作系统指定的路径分隔符
    PathListSeparator = ':' // 操作系统指定的表分隔符
)

const DevNull = "/dev/null"

var (
    Stdin  = NewFile(uintptr(syscall.Stdin), "/dev/stdin")
    Stdout = NewFile(uintptr(syscall.Stdout), "/dev/stdout")
    Stderr = NewFile(uintptr(syscall.Stderr), "/dev/stderr")
)
Stdin、Stdout和Stderr是指向标准输入、标准输出、标准错误输出的文件描述符

const (
    // 单字符是被String方法用于格式化的属性缩写。
    ModeDir        FileMode = 1 << (32 - 1 - iota) // d: 目录
    ModeAppend                                     // a: 只能写入，且只能写入到末尾
    ModeExclusive                                  // l: 用于执行
    ModeTemporary                                  // T: 临时文件（非备份文件）
    ModeSymlink                                    // L: 符号链接（不是快捷方式文件）
    ModeDevice                                     // D: 设备
    ModeNamedPipe                                  // p: 命名管道（FIFO）
    ModeSocket                                     // S: Unix域socket
    ModeSetuid                                     // u: 表示文件具有其创建者用户id权限
    ModeSetgid                                     // g: 表示文件具有其创建者组id的权限
    ModeCharDevice                                 // c: 字符设备，需已设置ModeDevice
    ModeSticky                                     // t: 只有root/创建者能删除/移动文件
    // 覆盖所有类型位（用于通过&获取类型位），对普通文件，所有这些位都不应被设置
    ModeType = ModeDir | ModeSymlink | ModeNamedPipe | ModeSocket | ModeDevice
    ModePerm FileMode = 0777 // 覆盖所有Unix权限位（用于通过&获取类型位）
)

type FileInfo interface {
    Name() string       // 文件的名字（不含扩展名）
    Size() int64        // 普通文件返回值表示其大小；其他文件的返回值含义各系统不同
    Mode() FileMode     // 文件的模式位
    ModTime() time.Time // 文件的修改时间
    IsDir() bool        // 等价于Mode().IsDir()
    Sys() interface{}   // 底层数据来源（可以返回nil）
}
FileInfo用来描述一个文件对象

type ProcAttr struct {
    // 如果Dir非空，子进程会在创建进程前先进入该目录。（即设为当前工作目录）
    Dir string
    // 如果Env非空，它会作为新进程的环境变量。必须采用Environ返回值的格式。
    // 如果Env为空字符串，将使用Environ函数的返回值。
    Env []string
    // Files指定被新进程继承的活动文件对象。
    // 前三个绑定为标准输入、标准输出、标准错误输出。
    // 依赖底层操作系统的实现可能会支持额外的数据出入途径。
    // nil条目相当于在进程开始时关闭的文件对象。
    Files []*File
    // 操作系统特定的创建属性。
    // 注意设置本字段意味着你的程序可能会运作失常甚至在某些操作系统中无法通过编译。
    Sys *syscall.SysProcAttr
}
ProcAttr保管将被StartProcess函数用于一个新进程的属性

var Args []string
Args保管了命令行参数，第一个是程序名

func Hostname() (name string, err error)
Hostname返回内核提供的主机名

func Getpagesize() int
Getpagesize返回底层的系统内存页的尺寸

func Environ() []string
Environ返回表示环境变量的格式为"key=value"的字符串的切片拷贝

func Getenv(key string) string
Getenv检索并返回名为key的环境变量的值。如果不存在该环境变量会返回空字符串

func Setenv(key, value string) error
Setenv设置名为key的环境变量。如果出错会返回该错误

func Clearenv()
Clearenv删除所有环境变量

func Exit(code int)
Exit让当前程序以给出的状态码code退出。一般来说，状态码0表示成功，非0表示出错。程序会立刻终止，defer的函数不会被执行

func Expand(s string, mapping func(string) string) string
Expand函数替换s中的${var}或$var为mapping(var)

func ExpandEnv(s string) string
ExpandEnv函数替换s中的${var}或$var为名为var 的环境变量的值

func Getuid() int
Getuid返回调用者的用户ID

func Geteuid() int
Geteuid返回调用者的有效用户ID

func Getgid() int
Getgid返回调用者的组ID

func Getegid() int
Getegid返回调用者的有效组ID

func Getgroups() ([]int, error)
Getgroups返回调用者所属的所有用户组的组ID

func Getpid() int
Getpid返回调用者所在进程的进程ID

func Getppid() int
Getppid返回调用者所在进程的父进程的进程ID

func (m FileMode) IsDir() bool
IsDir报告m是否是一个目录

func (m FileMode) IsRegular() bool
IsRegular报告m是否是一个普通文件

func (m FileMode) Perm() FileMode
Perm方法返回m的Unix权限位

func (m FileMode) String() string

func Stat(name string) (fi FileInfo, err error)
Stat返回一个描述name指定的文件对象的FileInfo,如果指定的文件对象是一个符号链接，返回的FileInfo描述该符号链接指向的文件的信息，本函数会尝试跳转该链接

func Lstat(name string) (fi FileInfo, err error)
Lstat返回一个描述name指定的文件对象的FileInfo,如果指定的文件对象是一个符号链接，返回的FileInfo描述该符号链接的信息，本函数不会试图跳转该链接

func IsPathSeparator(c uint8) bool
IsPathSeparator返回字符c是否是一个路径分隔符

func IsExist(err error) bool
返回一个布尔值说明该错误是否表示一个文件或目录已经存在

func IsNotExist(err error) bool
返回一个布尔值说明该错误是否表示一个文件或目录不存在

func IsPermission(err error) bool
返回一个布尔值说明该错误是否表示因权限不足要求被拒绝

func Getwd() (dir string, err error)
Getwd返回一个对应当前工作目录的根路径

func Chdir(dir string) error
Chdir将当前工作目录修改为dir指定的目录

func Chmod(name string, mode FileMode) error
Chmod修改name指定的文件对象的mode

func Chown(name string, uid, gid int) error
Chown修改name指定的文件对象的用户id和组id,如果name指定的文件是一个符号链接，它会修改该链接的目的地文件的用户id和组id

func Lchown(name string, uid, gid int) error
Lchown修改name指定的文件对象的用户id和组id,如果name指定的文件是一个符号链接，它会修改该链接的目的地文件的用户id和组id

func Chtimes(name string, atime time.Time, mtime time.Time) error
Chtimes修改name指定的文件对象的访问时间和修改时间，类似Unix的utime()或utimes()函数

func Mkdir(name string, perm FileMode) error
Mkdir使用指定的权限和名称创建一个目录

func MkdirAll(path string, perm FileMode) error
MkdirAll使用指定的权限和名称创建一个目录，包括任何必要的上级目录，并返回nil，否则返回错误

func Rename(oldpath, newpath string) error
Rename修改一个文件的名字，移动一个文件。可能会有一些个操作系统特定的限制

func Truncate(name string, size int64) error
Truncate修改name指定的文件的大小,如清空一个文件

func Remove(name string) error
Remove删除name指定的文件或目录

func RemoveAll(path string) error
RemoveAll删除path指定的文件，或目录及它包含的任何下级对象

func Readlink(name string) (string, error)
Readlink获取name指定的符号链接文件指向的文件的路径

func Symlink(oldname, newname string) error
Symlink创建一个名为newname指向oldname的符号链接

func Link(oldname, newname string) error
Link创建一个名为newname指向oldname的硬链接

func SameFile(fi1, fi2 FileInfo) bool
SameFile返回fi1和fi2是否在描述同一个文件

func TempDir() string
TempDir返回一个用于保管临时文件的默认目录

func Create(name string) (file *File, err error)
Create采用模式0666（任何人都可读写，不可执行）创建一个名为name的文件，如果文件已存在会截断它（为空文件）。如果成功，返回的文件对象可用于I/O；对应的文件描述符具有O_RDWR模式

func Open(name string) (file *File, err error)
Open打开一个文件用于读取。如果操作成功，返回的文件对象的方法可用于读取数据；对应的文件描述符具有O_RDONLY模式

func OpenFile(name string, flag int, perm FileMode) (file *File, err error)
OpenFile是一个更一般性的文件打开函数，大多数调用者都应用Open或Create代替本函数

func NewFile(fd uintptr, name string) *File
NewFile使用给出的Unix文件描述符和名称创建一个文件

func Pipe() (r *File, w *File, err error)
Pipe返回一对关联的文件对象。从r的读取将返回写入w的数据。本函数会返回两个文件对象和可能的错误

func (f *File) Name() string
Name方法返回（提供给Open/Create等方法的）文件名称

func (f *File) Stat() (fi FileInfo, err error)
Stat返回描述文件f的FileInfo类型值

func (f *File) Fd() uintptr
Fd返回与文件f对应的整数类型的Unix文件描述符

func (f *File) Chdir() error
Chdir将当前工作目录修改为f，f必须是一个目录

func (f *File) Chmod(mode FileMode) error
Chmod修改文件的模式

func (f *File) Chown(uid, gid int) error
Chown修改文件的用户ID和组ID

func (f *File) Readdir(n int) (fi []FileInfo, err error)
Readdir读取目录f的内容，返回一个有n个成员的[]FileInfo，这些FileInfo是被Lstat返回的，采用目录顺序。对本函数的下一次调用会返回上一次调用剩余未读取的内容的信息

func (f *File) Readdirnames(n int) (names []string, err error)
Readdir读取目录f的内容，返回一个有n个成员的[]string，切片成员为目录中文件对象的名字，采用目录顺序。对本函数的下一次调用会返回上一次调用剩余未读取的内容的信息

func (f *File) Truncate(size int64) error
Truncate改变文件的大小，它不会改变I/O的当前位置。 如果截断文件，多出的部分就会被丢弃

func (f *File) Read(b []byte) (n int, err error)
Read方法从f中读取最多len(b)字节数据并写入b

func (f *File) ReadAt(b []byte, off int64) (n int, err error)
ReadAt从指定的位置（相对于文件开始位置）读取len(b)字节数据并写入b

func (f *File) Write(b []byte) (n int, err error)
Write向文件中写入len(b)字节数据

func (f *File) WriteString(s string) (ret int, err error)
WriteString类似Write，但接受一个字符串参数

func (f *File) WriteAt(b []byte, off int64) (n int, err error)
WriteAt在指定的位置（相对于文件开始位置）写入len(b)字节数据

func (f *File) Seek(offset int64, whence int) (ret int64, err error)
Seek设置下一次读/写的位置。offset为相对偏移量，而whence决定相对位置：0为相对文件开头，1为相对当前位置，2为相对文件结尾,它返回新的偏移量（相对开头）和可能的错误

func (f *File) Sync() (err error)
Sync递交文件的当前内容进行稳定的存储,一般来说，这表示将文件系统的最近写入的数据在内存中的拷贝刷新到硬盘中稳定保存

func (f *File) Close() error
Close关闭文件f，使文件不能用于读写

func FindProcess(pid int) (p *Process, err error)
FindProcess根据进程id查找一个运行中的进程。函数返回的进程对象可以用于获取其关于底层操作系统进程的信息

func StartProcess(name string, argv []string, attr *ProcAttr) (*Process, error)
StartProcess使用提供的属性、程序名、命令行参数开始一个新进程。StartProcess函数是一个低水平的接口。os/exec包提供了高水平的接口，应该尽量使用该包

func (p *Process) Signal(sig Signal) error
Signal方法向进程发送一个信号

func (p *Process) Kill() error
Kill让进程立刻退出

func (p *Process) Wait() (*ProcessState, error)
Wait方法阻塞直到进程退出，然后返回一个描述ProcessState描述进程的状态和可能的错误。Wait方法会释放绑定到进程p的所有资源。在大多数操作系统中，进程p必须是当前进程的子进程，否则会返回错误

func (p *Process) Release() error
Release释放进程p绑定的所有资源， 使它们（资源）不能再被（进程p）使用。只有没有调用Wait方法时才需要调用本方法

func (p *ProcessState) Pid() int
Pid返回一个已退出的进程的进程id

func (p *ProcessState) Exited() bool
Exited报告进程是否已退出

func (p *ProcessState) Success() bool
Success报告进程是否成功退出，如在Unix里以状态码0退出

func (p *ProcessState) SystemTime() time.Duration
SystemTime返回已退出进程及其子进程耗费的系统CPU时间

func (p *ProcessState) UserTime() time.Duration
UserTime返回已退出进程及其子进程耗费的用户CPU时间

func (p *ProcessState) Sys() interface{}
Sys返回该已退出进程系统特定的退出信息。需要将其类型转换为适当的底层类型，如Unix里转换为*syscall.WaitStatus类型以获取其内容

func (p *ProcessState) SysUsage() interface{}
SysUsage返回该已退出进程系统特定的资源使用信息。需要将其类型转换为适当的底层类型，如Unix里转换为*syscall.Rusage类型以获取其内容

func (p *ProcessState) String() string

*/

func defer_test_exit() {
    log.Println("with exit,defer will not run")
}

func main() {
    file, err := os.Open("file.txt") // For read access.
    if err != nil {
        log.Fatal(err)
    }
    data := make([]byte, 100)
    count, err := file.Read(data)
    if err != nil {
        log.Fatal(err)
    }
    log.Printf("read %d bytes: %q\n", count, data[:count])
   
    defer file.Close()   
    defer defer_test_exit() 

    hname,_ := os.Hostname()
    log.Println("hostname:",hname)
    log.Println("pagesize:",os.Getpagesize())
    log.Println("ENV:",os.Environ())
    log.Println("GOPATH:",os.Getenv("GOPATH"))
    log.Println("uid:",os.Getuid())
    log.Println("euid:",os.Geteuid())
    log.Println("gid:",os.Getgid())
    log.Println("egid:",os.Getegid())
    log.Println(os.Getgroups())
    log.Println("pid:",os.Getpid())
    log.Println("ppid:",os.Getppid())
   
    finfo,_ := os.Stat("./file.txt")
    log.Println(finfo.Name(),finfo.Size(),finfo.Mode(),finfo.ModTime(),finfo.IsDir(),finfo.Sys())
    uint8str := []uint8(",")
    log.Println("uint8str:",uint8str)
    log.Println("uint8str is path separator:",os.IsPathSeparator(uint8str[0]))
    path,_ := os.Getwd()
    log.Println("pwd:",path)
    log.Println(os.Chdir(path))
    log.Println(os.Chmod("file.txt",0755)) //must be absolute path
    log.Println(os.Mkdir("aaa",0755))
    log.Println(os.MkdirAll("./ccc/bbb",0755)) //create with parents dir
    log.Println(os.Remove("aaa"))
    log.Println(os.RemoveAll("ccc")) //delete with child dir/file
    log.Println(os.Rename("file.txt","file"))
    log.Println(os.Rename("file","file.txt"))
    log.Println(os.TempDir())
    log.Println(os.Pipe())
    os.Exit(1)  //去掉，defer_test_exit将会执行
}
