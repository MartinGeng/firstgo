package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("in main()")

	go longWait()
	go shortWait()

	fmt.Println("About to sleep in main()")
	time.Sleep(10*1e9)
	fmt.Println("At the end of main()")
}

func longWait() {
	fmt.Println("beginning longWait()")
	time.Sleep(5*1e9)
	fmt.Println("end of longWait()")
}

func shortWait() {
	fmt.Println("beginning shortWait()")
	time.Sleep(2*1e9)
	fmt.Println("end of shortWait()")
}