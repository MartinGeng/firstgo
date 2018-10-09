package main

import (
"fmt"
	log "github.com/golang/glog"
	"gopkg.in/gcfg.v1"
)

func main() {
	cfgStr := `; Comment line
	[section]
	name=value # comment`
	cfg := struct {
	    Section struct {
	        Name string
	    }
	}{}
	err := gcfg.ReadStringInto(&cfg, cfgStr)
	if err != nil {
	    log.Fatalf("Failed to parse gcfg data: %s", err)
	}
	fmt.Println(cfg.Section.Name)
}