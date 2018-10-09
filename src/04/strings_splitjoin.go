package main

import (
	"fmt"
	"strings"
)

func main() {
	str := "The quick brown fox jumps over the lazy dog"
	sl := strings.Fields(str)
	fmt.Println("Splitted in slice : ", sl)

	for _, val := range sl {
		fmt.Println("s = ", val)
	}

	str2 := "G01|The ABC of GO|25"
	sl2 := strings.Split(str2, "|")
	fmt.Println("Splitted in slice : ", sl2)
	for _, val := range sl2{
		fmt.Println(val)
	}

	str3 := strings.Join(sl2, ";")
	fmt.Println("sl2 joined by; :", str3)
}
