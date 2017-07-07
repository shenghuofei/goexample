package main  
  
import (  
    "container/list"  
    "fmt"  
)  

type stu struct{
    name string
    age  int
}
  
func main() {  
    l := list.New() //创建一个新的list  
    for i := 0; i < 5; i++ {  
        stun := stu{name:"test"+fmt.Sprintf("%d",i),age:i+10}
        l.PushFront(stun)  
    }  
    for e := l.Front(); e != nil; e = e.Next() {  
        fmt.Println(e.Value) //输出list的值  
        fmt.Println((e.Value).(stu).name)   
    }  
    fmt.Println("===l list====")  
    fmt.Println("l front:",l.Front().Value) //输出首部元素的值  
    fmt.Println("l back:",l.Back().Value)  //输出尾部元素的值  
    l.InsertAfter(6, l.Front())  //首部元素之后插入一个元素  
    for e := l.Front(); e != nil; e = e.Next() {  
        fmt.Println(e.Value) //输出list的值  
    }  
    fmt.Println("===after insert l list=====")  
    l.MoveBefore(l.Front().Next(), l.Front()) //首部两个元素位置互换  
    for e := l.Front(); e != nil; e = e.Next() {  
        fmt.Println(e.Value) //输出list的值  
    }  
    fmt.Println("==after move berfore l list=======")  
    l.MoveToFront(l.Back()) //将尾部元素移动到首部  
    for e := l.Front(); e != nil; e = e.Next() {  
        fmt.Println(e.Value) //输出list的值  
    }  
    fmt.Println("===afert move front l list=====")  
    l2 := list.New()  
    l2.PushBackList(l) //将l中元素放在l2的末尾  
    for e := l2.Front(); e != nil; e = e.Next() {  
        fmt.Println(e.Value) //输出l2的值  
    }  
    fmt.Println("===l2 list======")  
    l.Init()           //清空l</span>  
    fmt.Println("l len:",l.Len()) //0  
    for e := l.Front(); e != nil; e = e.Next() {  
        fmt.Println(e.Value) //输出list的值,无内容  
    }  
  
}  
