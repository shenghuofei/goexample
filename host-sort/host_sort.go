package main

import (
    "fmt"
    "regexp"
    "sort"
)

type host []string

func (h host) Len() int {
	return len(h)
}
func (h host) Less(i, j int) bool {
	reg := regexp.MustCompile(`(\d+)-(online|test|dev|vm|pm)$`)
	id := reg.FindAllStringSubmatch(h[i], -1)
	jd := reg.FindAllStringSubmatch(h[j], -1)
	//fmt.Println("less", id[0][1])
	//fmt.Println("less", jd[0])
	return id[0][1] < jd[0][1]
}
func (h host) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func main() {
	h_list := host{"fidc-pro-a3-online", "fidc-pro-a2-vm", "fidc-pro-a1-pm", "sidc-pro-a2-test","sidc-pro-a1-dev"}
	fmt.Println("before sort:", h_list)
	sort.Sort(h_list)
    sort.Strings(h_list)
	fmt.Println("after sort:", h_list)
}
