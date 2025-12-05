package utils

import (
	"github.com/Ali-Al-Hadi-Al-Husseini/GO_projects/tree/main/6-Reddis-clone/app/storage"
)

func intalizeKeyId(key string, id string) {
	_, exists := storage.GlobalStorageStream[key]
	if !exists {
		storage.GlobalStorageStream[key] = make(map[string]map[string]string, 0)
		storage.GlobalStorageStream[key]["low"] = make(map[string]string)
		storage.GlobalStorageStream[key]["low"]["low"] = "+inf-0"
		storage.GlobalStorageStream[key]["high"] = make(map[string]string)
		storage.GlobalStorageStream[key]["high"]["high"] = "inf-0"

	}
	_, exists = storage.GlobalStorageStream[key][id]
	if !exists {
		storage.GlobalStorageStream[key][id] = make(map[string]string)
	}

}

func poplulateFeilds(commands []string, key string, id string) {
	for i := 3; i < len(commands); i += 2 {
		field := commands[i]
		value := commands[i+1]

		storage.GlobalStorageStream[key][id][field] = value
	}
}

func IntalizeStreamMap(commands []string, key string, id string) {
	intalizeKeyId(key, id)
	storage.LastAddedStream = id
	poplulateFeilds(commands, key, id)
}

func CheckHighAndLow(key string, id string) {
	if compareIds(id, storage.GlobalStorageStream[key]["low"]["low"]) {
		storage.GlobalStorageStream[key]["low"]["low"] = id
	}
	if compareIds(storage.GlobalStorageStream[key]["high"]["high"], id) {
		storage.GlobalStorageStream[key]["high"]["high"] = id
	}
}

func CheckXaddErrors(id string) (string, error) {
	time, idNum := StringToID(id)
	// id = fmt.Sprintf("%d-%d", time, idNum)
	if time == 0 && idNum == 0 {
		return "-ERR The ID specified in XADD must be greater than 0-0\r\n", nil
	}

	if storage.LastAddedStream != "0-0" {

		if time == -1 || idNum == -1 {
			return refuseId()
		}
		// 1		2
		lastTime, lastIDnum := StringToID(storage.LastAddedStream)

		if lastTime > time {
			return refuseId()
		}

		if lastTime == time && lastIDnum >= idNum {
			return refuseId()
		}
	}
	return "", nil
}
