package main

import (
	"fmt"
	"strings"
)

func main() {
	var orig string = "Hey, how are you George?"
	var lower string
	var upper string

	fmt.Println("The original string is : ", orig)
	lower = strings.ToLower(orig)
	fmt.Println("The lower string is : ", lower)
	upper = strings.ToUpper(orig)
	fmt.Println("The upper string is : ", upper)
}
