package main

import (
	"fmt"
	"os"
	"runtime"
)

func main() {
	goos := runtime.GOOS
	fmt.Println("the operating system is : ", goos)

	path := os.Getenv("PATH")
	fmt.Println("Path is ", path)
}
