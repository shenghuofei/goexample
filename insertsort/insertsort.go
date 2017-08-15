package main
import (
    "fmt"
)

func insert_sort(list []int64){
    list_len := len(list)
    for i:=1;i<list_len;i++ {  //若第i个元素大于i-1元素，直接插入。小于的话，移动有序表后插入 
        if list[i] < list[i-1] {
            j := i-1
            x := list[i]    //复制为哨兵，即存储待排序元素 
            list[i] = list[i-1]  //先后移一个元素
            for ;x<list[j]; {   //查找在有序表的插入位置 
                list[j+1] = list[j]  //元素后移
                j--
                if j <= -1 {   //最多到-1，此时待排序的元素会插入在最前面(j+1)=0的位置
                    break
                }
            }
            list[j+1] = x  //插入到正确位置 
        }
    }
}

func main() {
    list := []int64{243,123,343,35,9,10,1,3,2}
    fmt.Println("before sort",list)
    insert_sort(list)
    fmt.Println("after sort",list)
}
