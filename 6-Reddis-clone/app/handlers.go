package main

import (
	"fmt"
	"net"
	"strconv"
	"strings"
	"time"

	"github.com/Ali-Al-Hadi-Al-Husseini/GO_projects/tree/main/6-Reddis-clone/app/storage"

	"github.com/Ali-Al-Hadi-Al-Husseini/GO_projects/tree/main/6-Reddis-clone/app/utils"
)

func handleSet(commands []string) string {
	storage.GlobalStorage[commands[1]] = commands[2]
	if len(commands) > 4 && strings.ToLower(commands[3]) == "px" {
		px, err := strconv.Atoi(commands[4])
		if err != nil {
			return "\r\n"
		}
		storage.GlobalStorageTimer[commands[1]] = time.Now().Add(time.Duration(px-8) * time.Millisecond).UnixMilli()

	}
	return "+OK\r\n"
}
func handleGet(commands []string) string {
	keyValue, exists := storage.GlobalStorage[commands[1]]
	if exists {
		value, exists := storage.GlobalStorageTimer[commands[1]]
		if exists && time.Now().UnixMilli() >= value {
			delete(storage.GlobalStorage, commands[1])
			return "$-1\r\n"
		}
		return utils.ConvertToRESP(keyValue)
	}
	return "$-1\r\n"

}

func handleRpush(commands []string) string {
	if len(commands) < 3 {
		return "-ERR wrong number of arguments for 'rpush' command"
	}

	key := commands[1]
	for i := 2; i < len(commands); i++ {
		value := commands[i]

		_, exists := storage.GlobalStorageArray[key]
		if !exists {
			storage.GlobalStorageArray[key] = []string{}
		}
		// if containsString(GlobalStorageArray[key], value) {
		// 	continue
		// }
		storage.GlobalStorageArray[key] = append(storage.GlobalStorageArray[key], value)
	}
	blocked, exists := storage.GlobalStorageArrayBlpop[key]
	if blocked && exists {
		storage.GlobalStorageArrayBlpop[key] = false
	}
	return fmt.Sprintf(":%d\r\n", len(storage.GlobalStorageArray[key]))

}
func handlelpush(commands []string) string {
	if len(commands) < 3 {
		return "-ERR wrong number of arguments for 'rpush' command"
	}

	key := commands[1]
	for i := 2; i < len(commands); i++ {
		value := commands[i]

		_, exists := storage.GlobalStorageArray[key]
		if !exists {
			storage.GlobalStorageArray[key] = []string{}
		}
		// if containsString(GlobalStorageArray[key], value) {
		// 	continue
		// }
		storage.GlobalStorageArray[key] = append([]string{value}, storage.GlobalStorageArray[key]...)
	}
	return fmt.Sprintf(":%d\r\n", len(storage.GlobalStorageArray[key]))

}
func handleLrange(commands []string) string {
	if len(commands) < 3 {
		return "-ERR wrong number of arguments for 'Lrange' command"
	}

	key := commands[1]
	start, _ := strconv.Atoi(commands[2])
	end, _ := strconv.Atoi(commands[3])
	array, exists := storage.GlobalStorageArray[key]

	if end < 0 {
		end = len(array) + end
		if end < 0 {
			end = 0
		}
	}
	if start < 0 {
		start = len(array) + start
		if start < 0 {
			start = 0
		}
	}

	if !exists || len(array) <= start || start >= end {
		return "*0\r\n"
	}
	if end >= len(array) {
		end = len(array) - 1
	}

	return utils.ArrayToRESP(array[start : end+1])
}

func handleLlen(commands []string) string {
	if len(commands) < 2 {
		return "-ERR wrong number of arguments for 'llen' command"
	}

	key := commands[1]
	array, exists := storage.GlobalStorageArray[key]

	if !exists {
		return ":0\r\n"
	}

	return fmt.Sprintf(":%d\r\n", len(array))
}

func handleLpop(commands []string) string {
	key := commands[1]
	_, exists := storage.GlobalStorageArray[key]
	if !exists {
		return ":0\r\n"
	}
	var popValue int
	if len(commands) <= 2 {
		popValue = 1
		returnValue := storage.GlobalStorageArray[key][0]
		storage.GlobalStorageArray[key] = append([]string(nil), storage.GlobalStorageArray[key][popValue:]...)
		return utils.ConvertToRESP(returnValue)

	} else {
		popValue, _ = strconv.Atoi(commands[2])
	}
	returnValue := storage.GlobalStorageArray[key][:popValue]
	storage.GlobalStorageArray[key] = append([]string(nil), storage.GlobalStorageArray[key][popValue:]...)
	return utils.ArrayToRESP(returnValue)

}

// TODO
// could just keep the connection not closed until the key is not empty
// and then return the value
// but need to refactor the code to handle this
func handleBlpop(commands []string) string {
	key := commands[1]
	timeOut, _ := strconv.ParseFloat(commands[2], 64)
	timeOut *= 1000
	storage.GlobalStorageArrayBlpop[key] = true
	if timeOut == 0 {
		return blpopHelper(key)
	}
	time.Sleep(time.Duration(timeOut) * time.Millisecond)
	if len(storage.GlobalStorageArray[key]) > 0 {
		return popAndReturnarray(key)
	}
	return "$-1\r\n"
}

func blpopHelper(key string) string {

	for {
		_, exists := storage.GlobalStorageArrayBlpop[key]
		if !exists || !storage.GlobalStorageArrayBlpop[key] {
			return popAndReturnarray(key)
		}
		time.Sleep(10 * time.Millisecond)
	}
}
func popAndReturnarray(key string) string {
	if len(storage.GlobalStorageArray[key]) == 0 {
		return ""
	}
	returnValue := storage.GlobalStorageArray[key][0]
	storage.GlobalStorageArray[key] = append([]string(nil), storage.GlobalStorageArray[key][1:]...)
	return utils.ArrayToRESP([]string{key, returnValue})
}

func handleType(commands []string) string {
	if len(commands) < 2 {
		return "-ERR wrong number of arguments for 'type' command"
	}

	key := commands[1]
	val, exists := storage.GlobalStorage[key]
	if exists {

		_, isInt := strconv.Atoi(val)
		if isInt == nil {
			return "+int\r\n"
		}
		return "+string\r\n"
	}
	_, exists = storage.GlobalStorageStream[key]
	if exists {
		return "+stream\r\n"
	}

	return "+none\r\n"
}

func handleXadd(commands []string) string {
	key := commands[1]
	id := commands[2]
	// if id == "*"
	if !utils.IsID(id) {
		id = "*"
		commands = append(commands[:2], append([]string{"*"}, commands[2:]...)...)

	}
	err, _ := utils.CheckXaddErrors(id)
	if err != "" {
		return err
	}
	time, idNum := utils.StringToID(id)
	_, exists := storage.GlobalStorageStreamBlock[key]
	if exists {
		storage.GlobalStorageStreamBlock[key][id] = make(map[string]string)
		utils.PoplulateFeildsBlock(commands, key, id)
	}
	go func() {
		utils.IntalizeStreamMap(commands, key, id)
		storage.GlovalStreamIDExists[id] = true
		utils.CheckHighAndLow(key, id)
	}()
	id = fmt.Sprintf("%d-%d", time, idNum)
	return utils.ConvertToRESP(id)
}

func handleXrange(commands []string) string {
	key := commands[1]
	idMap := storage.GlobalStorageStream[key]
	start := commands[2]
	end := commands[3]

	start, end = utils.CheckForEndStart(start, end, key)
	ids := utils.GetAllIds(start, end, idMap)

	return utils.ConvertToRESPStream(ids, idMap)
}

func handleXread(commands []string) string {
	if len(commands) <= 4 {
		key := commands[2]
		start := commands[3]
		idMap := storage.GlobalStorageStream[key]

		end := "+"
		start, end = utils.CheckForEndStart(start, end, key)
		ids := utils.GetAllIds(start, end, idMap)
		return utils.ConvertToRESPXread(key, ids, idMap)
	}

	if commands[1] != "block" {
		return utils.GetXreadData(commands, 2)
	}

	key := commands[4]

	storage.GlobalStorageStreamBlock[key] = make(map[string]map[string]string)
	timeToSleep, _ := strconv.Atoi(commands[2])

	defer delete(storage.GlobalStorageStreamBlock, key)
	if timeToSleep > 0 {
		time.Sleep(time.Duration(timeToSleep) * time.Millisecond)

		var ids []string
		for k := range storage.GlobalStorageStreamBlock[key] {
			ids = append(ids, k)
		}
		if len(ids) != 0 {
			return utils.ConvertToRESPXread(key, ids, storage.GlobalStorageStreamBlock[key])
		}
		time.Sleep(time.Duration(5) * time.Millisecond)
		return "$-1\r\n"
	}
	var ids []string
	for len(storage.GlobalStorageStreamBlock[key]) <= 1 {

		for k := range storage.GlobalStorageStreamBlock[key] {
			ids = append(ids, k)
		}
		time.Sleep(10 * time.Millisecond)
		if len(ids) != 0 {
			break
		}

	}
	return utils.ConvertToRESPXread(key, ids, storage.GlobalStorageStreamBlock[key])
}

func handleIncr(Commands []string) string {

	key := Commands[1]
	_, exists := storage.GlobalStorage[key]
	if !exists {
		storage.GlobalStorage[key] = "1"
		return ":1\r\n"
	}
	value, err := strconv.Atoi(storage.GlobalStorage[key])
	if err != nil {
		return "-ERR value is not an integer or out of range\r\n"
	}
	storage.GlobalStorage[key] = fmt.Sprintf("%d", value+1)
	return fmt.Sprintf(":%d\r\n", value+1)

}
func handleMulti(conn net.Conn) string {
	storage.GlobalMutliStorage[conn] = [][]string{}
	return "+OK\r\n"
}

// change to commit to test
func handleExec(conn net.Conn, roleInfo map[string]string) string {
	_, exists := storage.GlobalMutliStorage[conn]
	if exists {

		if !(len(storage.GlobalMutliStorage[conn]) < 1) {
			var result []interface{}
			for _, Currcommands := range storage.GlobalMutliStorage[conn] {
				res, _ := Parser(Currcommands, conn, roleInfo)
				res = utils.RemoveResp(res)
				switch Currcommands[0] {
				case "set":

					result = append(result, res)
				case "incr":
					currRes, err := strconv.Atoi(res) // integer
					if err != nil {
						result = append(result, res)
					} else {
						result = append(result, currRes)
					}

				case "get":
					currRes := res // bulk string
					result = append(result, currRes)
				}

				delete(storage.GlobalMutliStorage, conn)

			}

			return utils.EncodeRESPArray(result)
		}
		return "*0\r\n"
	}

	return "-ERR EXEC without MULTI\r\n"

}
func handlerDiscard(conn net.Conn) string {
	_, exists := storage.GlobalMutliStorage[conn]
	if exists {
		delete(storage.GlobalMutliStorage, conn)
		return "+OK\r\n"
	}
	return "-ERR DISCARD without MULTI\r\n"

}

func handleInfo(roleInfo map[string]string) string {
	info := make(map[string]string)

	info["role"] = roleInfo["role"]

	switch info["role"] {
	case "master":
		info["master_replid"] = "8371b4fb1155b71f4a04d3e1bc3e18c4a990aeeb"
		info["master_repl_offset"] = "0"

	case "slave":

	}
	return utils.MapToRESPBulkString(info)
}
func handleReplConf(conn net.Conn) string {

	return "+OK\r\n"
}

func handlePsync(conn net.Conn) string {
	storage.GlobalSlaveAddress = append(storage.GlobalSlaveAddress, conn.RemoteAddr().String())
	fmt.Println(storage.GlobalSlaveAddress)
	conn.Write([]byte("+FULLRESYNC 8371b4fb1155b71f4a04d3e1bc3e18c4a990aeeb 0\r\n"))
	return "$66\r\nREDIS0011\xfa\tredis-ver\x057.2.0\xfa\nredis-bits\xc0@\xfa\x05ctime\xc2m\b\xbce\xfa\busd-mem\xc2\xb0\xc4\x10\x00"

}
