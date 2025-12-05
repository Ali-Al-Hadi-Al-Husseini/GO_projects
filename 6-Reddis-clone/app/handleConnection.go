package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"net"
	"os"
	"strings"

	"github.com/Ali-Al-Hadi-Al-Husseini/GO_projects/tree/main/6-Reddis-clone/app/storage"

	"github.com/Ali-Al-Hadi-Al-Husseini/GO_projects/tree/main/6-Reddis-clone/app/utils"
)

func handleConnection(reader *bufio.Reader, conn net.Conn, roleInfo map[string]string) {
	defer conn.Close()
	for {

		commands, err := readCommands(reader)
		if err != nil {
			// client closed or network error
			if errors.Is(err, io.EOF) {
				// clean shutdown by client
				return
			}
			// log other errors then close this connection
			fmt.Println("readCommands error:", err)
			return
		}

		raw := utils.ArrayToRESP(commands)
		if len(commands) == 0 {

			continue
		}
		commands[0] = strings.ToLower(commands[0])
		executeCommand(commands, conn, roleInfo, raw)

	}
}

func executeCommand(commands []string, conn net.Conn, roleInfo map[string]string, raw string) {

	if roleInfo["role"] == "master" {
		_, isNonwrite := storage.NonWriteCommands[commands[0]]
		if !isNonwrite {
			utils.PropagateCommand(raw)
		}

	}

	_, isTransaction := storage.GlobalMutliStorage[conn]
	if !isTransaction || commands[0] == "exec" || commands[0] == "discard" {
		//execute the command immediately or discard prev transaction
		res, finished := Parser(commands, conn, roleInfo)
		if finished {
			_, exits := storage.GlobalMutliStorage[conn]
			if exits {
				delete(storage.GlobalMutliStorage, conn)
			}
		}

		fmt.Println(res)

		conn.Write([]byte(res))

	} else {
		// If in transaction mode, store the command
		storage.GlobalMutliStorage[conn] = append(storage.GlobalMutliStorage[conn], commands)
		conn.Write([]byte("+QUEUED\r\n"))

	}

}

func readCommands(reader *bufio.Reader) ([]string, error) {
	var args []string

	line, err := reader.ReadString('\n')
	if err != nil {
		return nil, err
	}
	line = strings.TrimSuffix(line, "\r\n")

	var numArgs int
	if _, err = fmt.Sscanf(line, "*%d", &numArgs); err != nil {
		return nil, fmt.Errorf("invalid multibulk header: %w", err)
	}

	for i := 0; i < numArgs; i++ {
		lenLine, err := reader.ReadString('\n')
		if err != nil {
			return nil, err
		}
		lenLine = strings.TrimSuffix(lenLine, "\r\n")

		var argLen int
		if _, err = fmt.Sscanf(lenLine, "$%d", &argLen); err != nil {
			return nil, fmt.Errorf("invalid bulk string length: %w", err)
		}

		// read exactly argLen bytes + \r\n
		buf := make([]byte, argLen+2)
		if _, err = io.ReadFull(reader, buf); err != nil {
			return nil, fmt.Errorf("failed to read bulk string: %w", err)
		}
		args = append(args, string(buf[:argLen]))
	}

	return args, nil
}

// ParseRESPCommand parses a RESP command from a string input.
// It expects the input to be in the format of a RESP array.
// It returns a slice of strings representing the command and its arguments.
// If the input is not a valid RESP command, it returns nil.
// It also handles the case where the first line starts with '*'.
func ParseRESPCommand(input string) []string {
	lines := strings.Split(input, "\r\n")
	if len(lines) < 3 || !strings.HasPrefix(lines[0], "*") {
		return nil
	}

	var result []string
	for i := 2; i < len(lines); i += 2 {
		if lines[i] == "" {
			continue
		}
		arg := lines[i]
		if len(result) == 0 {
			arg = strings.ToLower(arg)
		}
		result = append(result, arg)
	}
	return result
}

func listenForConnections(l net.Listener, roleInfo map[string]string) {
	for {

		conn, err := l.Accept()
		if err != nil {
			fmt.Println("Error accepting connection: ", err.Error())
			os.Exit(1)
		}

		reader := bufio.NewReader(conn)
		go handleConnection(reader, conn, roleInfo)
	}
}
