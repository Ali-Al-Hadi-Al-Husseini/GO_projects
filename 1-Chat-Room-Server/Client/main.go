package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
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
	scanner := bufio.NewScanner(os.Stdin)
	for {

		//scaning for txt to send
		fmt.Printf("=> ")
		scanner.Scan()

	}
}
