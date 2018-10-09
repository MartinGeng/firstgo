package main

import (
	"bufio"
	"fmt"
	"os"
)

var (
	inputReader *bufio.Reader
	input       string
	err         error
)

func main() {
	inputReader := bufio.NewReader(os.Stdin)
	fmt.Println("please enter some input:")
	input, err = inputReader.ReadString("\n")
	if err == nil {
		fmt.Println("The input was: %s\n", input)
	}

}
