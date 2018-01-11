package main

import (
	"fmt"
	"hash/crc32"
	"sort"
)

type rindex []uint32  //hash环索引数组类型

//hash 环结构体
type ring struct {
	rmap      map[uint32]string  //{hashindex:nodename}
	rindexarr rindex  //hashindexarr
}


//hash环索引类型实现排序接口，以进行排序和查找node
func (this rindex) Less(i, j int) bool {
	return this[i] < this[j]
}

func (this rindex) Len() int {
	return len(this)
}

func (this rindex) Swap(i, j int) {
	this[i], this[j] = this[j], this[i]
}


func (this *ring) AddNode(nodename string) {
    index := crc32.ChecksumIEEE([]byte(nodename))
    if _,ok := this.rmap[index]; ok {
        return
    }
    this.rmap[index] = nodename
    this.rindexarr = append(this.rindexarr, index)
    sort.Sort(this.rindexarr)
}

func (this *ring) RemoveNode(nodename string) {
    index := crc32.ChecksumIEEE([]byte(nodename))
    if _,ok := this.rmap[index]; !ok {
        return
    }
    delete(this.rmap,index)
    this.rindexarr = rindex{}
    for k := range this.rmap {
        this.rindexarr = append(this.rindexarr,k)
    }
    sort.Sort(this.rindexarr)
}


func main() {
	host := []string{"h1", "h2", "h3"} //node数组

    //实例化hash环ring
	hashmap := &ring{
		rmap:      map[uint32]string{},
		rindexarr: rindex{},
	}
    
    //使用hash算法获取各node的唯一hash值，并更新hash环
	for _, v := range host {
		index := crc32.ChecksumIEEE([]byte(v))
		hashmap.rmap[index] = v
		hashmap.rindexarr = append(hashmap.rindexarr, index)
	}
    
    //对hash环索引数组进行排序
	sort.Sort(hashmap.rindexarr)
	fmt.Println("hash ring: ",hashmap)

    //增加node
	hashmap.AddNode("h4")
	fmt.Println("add node h4 to hash ring: ",hashmap)
    
    //删除node
	hashmap.RemoveNode("h3")
	fmt.Println("del node h3 from hash ring: ",hashmap)

    //用相同的hash算法 获取自己key的hash值
	key := "my key"
	hash := crc32.ChecksumIEEE([]byte(key))

    //根据key分配node
    //sort.Search二分查找hash环索引数组中第一个大于自己key hash值的元素下标,进而获取所分配的node
	i := sort.Search(len(hashmap.rindexarr), func(i int) bool { return hashmap.rindexarr[i] >= hash })
    node := hashmap.rmap[hashmap.rindexarr[i]] //分配的node即为hash环数组中索引为i的元素对应的ramp中的key的value值
	fmt.Println("key: ", key, "hash value: ", hash, "ring index:", i, "node:", node)
}

