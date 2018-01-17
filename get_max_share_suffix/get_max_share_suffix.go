package main

import "fmt"

func get_max_suffix(list []string) string {
    suffix := ""
    list_len := len(list)
    if list_len == 0 {   //list长度为0 直接返回空
        return suffix
    }
    if list_len == 1 {  //list长度为1直接返回唯一的那个元素
        return list[0]  
    }
    max_suffix := list[0] //取第一个字符串为标准进行判断
    max_suffix_len := len(max_suffix)
    if max_suffix_len == 0 {   //如果第一个字符串为空，则直接返回空
        return suffix
    }
    for i,_ := range max_suffix {  //第一个字符串长度依次减一判断
        tmp_len := max_suffix_len - i //当前待判断的公共后缀的长度
        flag := true
        for _,v := range list[1:] {  //挨个判断各字符串是否有当前判断的后缀
            if v == "" {  //如果有空字符串，直接返回空
                return suffix
            }
            if len(v) < tmp_len || max_suffix[i:] != v[len(v)-tmp_len:] { //如果当前字符串的长度比当前判断的后缀长度还小那当前后缀肯定不是公共后缀；或者当前字符串同长度后缀字符与当前判断的后缀不同
                flag = false
                break
            }
        }
        if flag {
            return max_suffix[i:]
        }
    }
    return suffix
}

func main() {
    //str_list := []string{"abcc","abccc","abcccc"}
    str_list := []string{"abcccc","abccc","abcc"}
    res := get_max_suffix(str_list)
    fmt.Println(str_list,"max suffix: ",res)
}
