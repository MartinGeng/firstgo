package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {
	//open connection
	conn, err := net.Dial("tcp", "localhost:50000")
	if err != nil {
		fmt.Println("Error dialing", err.Error())
		return
	}

	inputReader := bufio.NewReader(os.Stdin)
	fmt.Println("First, what is your name?")
	clientName, _ := inputReader.ReadString('\n')
	fmt.Println("clientName:", clientName)

	trimmedClient := strings.Trim(clientName, "\n")

	//send data until quit
	for {
		fmt.Println("what will send to server? Type Q to quit.")
		input, _ := inputReader.ReadString('\n')
		trimmedInput := strings.Trim(input, "\n")

		if trimmedInput == "Q" {
			return
		}

		_, err = conn.Write([]byte(trimmedClient + " say : " + trimmedInput))
	}
}
