package main

import (
	"fmt"
	"net"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		fmt.Println("Error")
	}
	defer conn.Close()

	go ListenForMsg()
	sendMsg()
}

func ListenForMsg() {

}

func sendMsg() {

}
