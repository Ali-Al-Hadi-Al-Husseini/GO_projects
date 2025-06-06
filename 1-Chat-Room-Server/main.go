package main

import (
	"bufio"
	"fmt"
	"net"
	"sync"
)

var (
	connections = make(map[net.Conn]bool)
	connMutex   = &sync.Mutex{}
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

		connMutex.Lock()
		connections[conn] = true
		connMutex.Unlock()

		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	defer func() {
		connMutex.Lock()
		delete(connections, conn)
		connMutex.Unlock()
		conn.Close()
	}()
	reader := bufio.NewReader(conn)
	for {
		msg, err := reader.ReadString('\n')
		if err != nil {
			return
		}
		broadCast(conn, msg)
	}

}

func broadCast(sender net.Conn, msg string) {}
