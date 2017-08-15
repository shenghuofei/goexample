package main
import (
    "fmt"
)

//search_list must be a sorted list
func binsearch(search_list []int64,element int64) bool {
    low := 0
    middle := 0
    high := len(search_list)
    for ;low<=high; {
        middle = (low+high)/2
        if search_list[middle] == element {
            return true
        }else if search_list[middle] > element {
            high = middle-1
        }else {
            low = middle+1
        }
    }
    return false
}

func main() {
    search_list := []int64{1,3,4,6,8,9}
    var element int64 = 2
    if binsearch(search_list,element) {
        fmt.Println(element,"in search_list")
    }else {
        fmt.Println(element,"not in search_list")
    }
}
