package main

import (
	"bufio"
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
	for {
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println(err)
		}

		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	defer conn.Close()
	reader := bufio.NewReader(conn)
	for {
		contnt, err := reader.ReadString('\n')
		if err != nil {
			return
		}

		fmt.Println("Got: ", contnt)
	}

}
