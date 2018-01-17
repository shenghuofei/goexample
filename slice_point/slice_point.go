package main

import "fmt"

type student struct {
    name string
    age  int
}

func main() {
    student_map := map[string](*student){}
    stu_list := []student{{"wang",33},{"li",23}}

    //因为在for...range中v是值拷贝，而非元素本身，当使用&获取元素的地址时，实际上只是取到了v这个临时变量的地址，而不是真正被遍历到的元素的地址，而v这个临时变量会被重复使用，所以map中的value都是最后一次循环中v的地址，要正确获取每个元素的地址，临时变量就不能被重复使用(比如使用索引来取)
    for _,value := range stu_list { 
        student_map[value.name] = &value
    }
    for k,v := range student_map {
        fmt.Println(k,v)
    }


    //正确做法
    fmt.Println("\nthe right method\n")
    for i,value := range stu_list {
        student_map[value.name] = &stu_list[i]
    }
    for k,v := range student_map {
        fmt.Println(k,v)
    }
}
