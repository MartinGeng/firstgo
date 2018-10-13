package main

import (
	"fmt"
	"io"
	"net"
)

func main() {
	var (
		host          = "www.apache.org"
		port          = "80"
		remote        = host + ":" + port
		msg    string = "GET /\n"
		data          = make([]uint8, 4096)
		read          = true
		count         = 0
	)

	//create a socket
	conn, err := net.Dial("tcp", remote)

	//send a get request
	io.WriteString(conn, msg)

	for read {
		count, err = conn.Read(data)
		read = (err == nil)
		fmt.Println(string(data[:count]))
	}

	conn.Close()
}
