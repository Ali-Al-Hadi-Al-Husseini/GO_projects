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

func sendMsg(conn net.Conn) {
	scanner := bufio.NewScanner(os.Stdin)
	for {

		//scaning for msg to send
		fmt.Printf("=> ")
		scanner.Scan()
		content := scanner.Text() + "\n"

		//writing the txt to the response
		writer := bufio.NewWriter(conn)
		fmt.Println("got this msg")
		writer.WriteString(content)
		writer.Flush()

	}
}
