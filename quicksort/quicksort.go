package main

import (
    "fmt"
)

func swap(arr []int,i,k int) {
    tmp := arr[i]
    arr[i] = arr[k]
    arr[k] = tmp
}

func sort(arr []int,s,e int) {
    if s >= e {
        return
    }
    pivot := arr[s] //选择一个基准元素
    i := s+1
    k := e
    for {
        for arr[i] < pivot && i <= e {   //从前往后找第一个大于基准的元素
           i++
        }
        for arr[k] > pivot && k >= s+1 {  //从后往前找第一个小于基准的元素
           k--
        }
        if i >= k {  //如果i >= k说明这个区间已经符合前面的比基准小，后面的比基准大
            break
        } 
        swap(arr,i,k)  //前面比基准大的和后面比基准小的元素互换
    }
    swap(arr,s,k)  //把基准元素换到正确的位置
    sort(arr,s,k-1) //对基准前的区间进行排序
    sort(arr,k+1,e) //对基准后的区间进行排序
}

func main() {
    arr := []int{1,6,3,7,2,34}
    arr_len := len(arr)
    fmt.Println("before:",arr)
    sort(arr,0,arr_len-1)
    fmt.Println("after:",arr)
}
