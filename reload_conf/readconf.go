package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type conf struct {
	Debug bool   `json:"debug"`
	Name  string `json:"name"`
}

var cfg = &conf{}

func get_conf() *conf {
    new_cfg := &conf{}
	cfg_file, err := ioutil.ReadFile("cfg.json")
	if err != nil {
		fmt.Print(err)
	}
	fmt.Println(string(cfg_file))
	err = json.Unmarshal(cfg_file,&new_cfg)
	if err != nil {
		fmt.Println(err)
	}
	return new_cfg
}

func reload() {
    cfg = get_conf()
    fmt.Println(*cfg)
}

/*
func main() {
    reload()
}
*/
