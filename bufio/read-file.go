package main  
  
import(  
    "fmt"  
    "os"  
    "flag"  
    "io"  
    "io/ioutil"  
    "bufio"  
    "time"  
)  
  
func read1(path string)string{  
    fi,err := os.Open(path)  
    if err != nil{  
        panic(err)  
    }  
    defer fi.Close()  
  
    chunks := make([]byte,1024,1024)  
    buf := make([]byte,1024)  
    for{  
        n,err := fi.Read(buf)  
        if err != nil && err != io.EOF{panic(err)}  
        if 0 ==n {break}  
        chunks=append(chunks,buf[:n]...)  
        // fmt.Println(string(buf[:n]))  
    }  
    return string(chunks)  
}  
  
func read2(path string)string{  
    fi,err := os.Open(path)  
    if err != nil{panic(err)}  
    defer fi.Close()  
    r := bufio.NewReader(fi)  
      
    chunks := make([]byte,1024,1024)  
       
    buf := make([]byte,1024)  
    for{  
        n,err := r.Read(buf)  
        if err != nil && err != io.EOF{panic(err)}  
        if 0 ==n {break}  
        chunks=append(chunks,buf[:n]...)  
        // fmt.Println(string(buf[:n]))  
    }  
    return string(chunks)  
}  
  
func read3(path string)string{  
    fi,err := os.Open(path)  
    if err != nil{panic(err)}  
    defer fi.Close()  
    fd,err := ioutil.ReadAll(fi)  
    // fmt.Println(string(fd))  
    return string(fd)  
}  

func read4(path string){
    _,err := ioutil.ReadFile(path)
    if err != nil{panic(err)}
}
  
func main(){  
     
    flag.Parse()  
    file := flag.Arg(0)  
    /*f,err := ioutil.ReadFile(file)  
    if err != nil{  
        fmt.Printf("%s\n",err)  
        panic(err)  
    }  
    fmt.Println(string(f))  */
    start := time.Now()  
    read1(file)  
    t1 := time.Now()  
    fmt.Printf("File-Read Cost time %v\n",t1.Sub(start))  
    read2(file)  
    t2 := time.Now()  
    fmt.Printf("Buffio-Read Cost time %v\n",t2.Sub(t1))  
    read3(file)  
    t3 := time.Now()  
    fmt.Printf("Ioutil-ReadAll Cost time %v\n",t3.Sub(t2))  
    read4(file) 
    t4 := time.Now()  
    fmt.Printf("Ioutil-ReadFile Cost time %v\n",t4.Sub(t3))  
}  

