package main

import (
	"flag"
	"fmt"
	"os"

	log "github.com/golang/glog"
)

func main() {
	//Init the command-line flags.
	flag.Parse()

	// Will be ignored as the program has exited in Fatal().
	defer func() {
		fmt.Println("Message in defer")
	}()

	// Flushes all pending log I/O.
	defer log.Flush()

	fmt.Printf("Temp folder for log files: %s\n", os.TempDir())

	log.Info("log info")
	log.V(1).Info("L1 info")
	log.Error("error")
	log.Fatal("Fatal")

	//log.Error("Error after fatal")
}