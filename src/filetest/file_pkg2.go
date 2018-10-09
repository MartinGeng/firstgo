package main

import (
	"fmt"
	"gopkg.in/gcfg.v1"
)

type MyConfig struct{
	Name string
}

var (
	ServerConfigFile string = "./server_cfg.conf"
)

func main() {

	var cfg MyConfig
	err := gcfg.ReadFileInto(&cfg, ServerConfigFile)
	if err != nil {
	    fmt.Println("err: ", err)
	}
	fmt.Println(cfg.Name)
}