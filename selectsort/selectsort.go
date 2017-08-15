package main
import (
    "fmt"
)

func select_sort(list []int64){
    list_len := len(list)
    for i:=0;i<list_len;i++ {
        k := i
        for j:=i+1;j<list_len;j++{  //选择最小的元素
            if list[k] > list[j] {
                k = j
            }
        }
        if k != i { //如果最小元素不是当前元素，则交换
            t := list[k]
            list[k] = list[i]
            list[i] = t
        }
    }
}

func main() {
    list := []int64{243,123,343,35,9,10,1,3,2}
    fmt.Println("before sort",list)
    select_sort(list)
    fmt.Println("after sort",list)
}
