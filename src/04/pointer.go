package main

import (
	"fmt"
)

func main() {
	var i1 = 5
	fmt.Println("i1 address : ", &i1)

	var intP *int
	intP = &i1
	fmt.Println("the value at memery location is ", *intP)
}