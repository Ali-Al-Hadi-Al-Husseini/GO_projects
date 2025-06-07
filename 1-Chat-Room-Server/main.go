package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"sync"
	"time"
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
		logMsg(msg)
		broadCast(conn, msg)
	}

}

func broadCast(sender net.Conn, msg string) {
	connMutex.Lock()
	defer connMutex.Unlock()

	for conn := range connections {
		if conn != sender {
			conn.Write([]byte(msg))
		}
	}

}

func logMsg(msg string) {
	f, err := os.OpenFile("chat.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return
	}

	defer f.Close()

	// Add a timestamp
	timestamp := time.Now().Format("2006-01-02 15:04:05")
	logLine := fmt.Sprintf("[%s] %s", timestamp, msg)

	// Write the log line
	f.WriteString(logLine)
}
