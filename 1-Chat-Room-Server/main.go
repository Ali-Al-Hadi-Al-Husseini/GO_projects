package main

import (
	"fmt"
	"net"
)

func main() {
	fmt.Println(("server Started"))
	listen, err := net.Listen("tcp", ":8080")
	if err != nil {
		fmt.Println(err)
	}
	defer listen.Close()

}
