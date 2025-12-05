package main

import (
	"net"
	"strings"

	"github.com/Ali-Al-Hadi-Al-Husseini/GO_projects/tree/main/6-Reddis-clone/app/utils"
)

func Parser(commands []string, conn net.Conn, roleInfo map[string]string) (string, bool) {
	if len(commands) == 0 {
		return "\r\n", true
	}
	commands[0] = strings.ToLower(commands[0])
	switch commands[0] {
	case "echo":
		return utils.ConvertToRESP(commands[1]), true

	case "ping":
		return "+PONG\r\n", true

	case "set":
		return handleSet(commands), true

	case "get":
		return handleGet(commands), true

	case "rpush":
		return handleRpush(commands), true

	case "lrange":
		return handleLrange(commands), true

	case "lpush":
		return handlelpush(commands), true

	case "llen":
		return handleLlen(commands), true

	case "lpop":
		return handleLpop(commands), true

	case "blpop":
		return handleBlpop(commands), true

	case "type":
		return handleType(commands), true

	case "xadd":
		return handleXadd(commands), true

	case "xrange":
		return handleXrange(commands), true

	case "xread":
		return handleXread(commands), true

	case "incr":
		return handleIncr(commands), true

	case "multi":
		return handleMulti(conn), false
	case "exec":
		return handleExec(conn, roleInfo), true

	case "discard":
		return handlerDiscard(conn), true
	case "info":
		return handleInfo(roleInfo), true

	case "replconf":
		return handleReplConf(conn), true
	case "psync":
		return handlePsync(conn), true
	}

	return "\r\n", true
}
