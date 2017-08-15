package main
import (
    "fmt"
)

func bubble_sort_desc(list []int64){
    list_len := len(list)
    for i:=0;i<list_len;i++ { //需要冒泡的次数
        for j:=0;j<list_len-i-1;j++ {  //已有序的元素不需再判断了
            if list[j+1] > list[j] {   //每个元素与相邻的元素比较，顺序不对则交换
                t := list[j]
                list[j] = list[j+1]
                list[j+1] = t
            }
        }
    }
}

func bubble_sort_esc(list []int64){
    list_len := len(list)
    for i:=0;i<list_len;i++ {
        for j:=0;j<list_len-i-1;j++ {
            if list[j+1] < list[j] {
                t := list[j]
                list[j] = list[j+1]
                list[j+1] = t
            }
        }
    }
}
func main() {
    list := []int64{34,456,12,4,7,1,3,99}
    fmt.Println("before sort",list)
    bubble_sort_desc(list)
    fmt.Println("after sort desc",list)
    list = []int64{34,456,12,4,7,1,3,99}
    bubble_sort_esc(list)
    fmt.Println("after sort esc",list)
}
