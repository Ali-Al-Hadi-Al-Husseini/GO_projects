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

	scanner := bufio.NewScanner(os.Stdin)
	fmt.Printf("name: ")
	scanner.Scan()
	name := scanner.Text()

	go ListenForMsg(conn)
	sendMsg(conn, name)
}

func ListenForMsg(conn net.Conn) {
	reader := bufio.NewReader(conn)
	for {
		content, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("errror")
		}

		fmt.Printf("\r\033[K%s=> ", content)

	}

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
		writer.WriteString(content)
		writer.Flush()

	}
}
