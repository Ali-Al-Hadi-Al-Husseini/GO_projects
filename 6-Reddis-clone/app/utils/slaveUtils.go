package utils

import (
	"bufio"
	"fmt"
	"net"
	"strings"

	"github.com/Ali-Al-Hadi-Al-Husseini/GO_projects/tree/main/6-Reddis-clone/app/storage"
)

func IntalizeSlavment(roleInfo map[string]string, slavePort int) {

	conn, err := net.Dial("tcp", formatAddr(roleInfo["addr"], roleInfo["port"])) // Replace with your host:port
	if err != nil {
		fmt.Println("Error connecting:", err)

	}

	conn.Write([]byte("*1\r\n$4\r\nPING\r\n")) // Send PING command in RESP format

	reader := bufio.NewReader(conn)
	line := readTcp(*reader)
	if line == "+PONG\r\n" {
		conn.Write([]byte(fmt.Sprintf("*3\r\n$8\r\nREPLCONF\r\n$14\r\nlistening-port\r\n$4\r\n%d\r\n", slavePort)))

		line = readTcp(*reader)
		ifOkWrite(conn, "*3\r\n$8\r\nREPLCONF\r\n$4\r\ncapa\r\n$6\r\npsync2\r\n", line, roleInfo)

		line = readTcp(*reader)
		ifOkWrite(conn, "*3\r\n$5\r\nPSYNC\r\n$1\r\n?\r\n$2\r\n-1\r\n", line, roleInfo)
		fmt.Printf("send psync %s\n", line)
	}

}
func readTcp(reader bufio.Reader) string {
	line, err := reader.ReadString('\n')
	if err != nil {
		return ""
	}
	return line
}
func ifOkWrite(conn net.Conn, content string, line string, roleInfo map[string]string) {
	if line == "+OK\r\n" {

		conn.Write([]byte(content))
	}
}
func PropagateCommand(command string) {

	for _, addrs := range storage.GlobalSlaveAddress {
		conn, err := net.Dial("tcp", addrs)
		if err != nil {
			continue
		}
		conn.Write([]byte(command))
	}
}

func formatAddr(ip string, port string) string {
	if strings.Contains(ip, ":") {
		// IPv6 needs brackets
		return fmt.Sprintf("[%s]:%s", ip, port)
	}
	return fmt.Sprintf("%s:%s", ip, port)
}
