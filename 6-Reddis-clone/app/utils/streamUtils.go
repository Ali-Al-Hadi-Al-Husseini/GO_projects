package utils

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/Ali-Al-Hadi-Al-Husseini/GO_projects/tree/main/6-Reddis-clone/app/storage"
)

func StringToID(str string) (int, int) {

	if str == "*" {
		newTime := int(time.Now().UnixMilli())
		newTimeStr := fmt.Sprintf("%d", newTime)

		_, exists := storage.GlovalStreamIDExists[newTimeStr+"-0"]
		if exists {
			for i := 1; i < 10; i++ {
				newID := fmt.Sprintf("%s-%d", newTimeStr, i)

				_, exists = storage.GlovalStreamIDExists[newID]
				if !exists {
					return newTime, i
				}
			}
		}

		return newTime, 0

	}
	splited := strings.Split(str, "-")
	currTime, err := strconv.Atoi(splited[0])

	lastTime, LastId := stringToIDHelper(storage.LastAddedStream)

	if err != nil {
		return -1, -1
	}

	if splited[1] == "*" {
		if currTime == 0 {
			return currTime, 1
		}

		if lastTime == currTime {
			return currTime, LastId + 1
		}
		return currTime, 0
	}
	id, err := strconv.Atoi(splited[1])
	if err != nil {
		return -1, -1
	}
	return currTime, id
}
func refuseId() (string, error) {
	return "-ERR The ID specified in XADD is equal or smaller than the target stream top item\r\n", nil
}

func IsID(str string) bool {
	if str == "*" {
		return true
	}
	if strings.Contains(str, "-") {
		splited := strings.Split(str, "-")
		if len(splited) == 2 {
			if _, err := strconv.Atoi(splited[0]); err == nil {
				if _, err := strconv.Atoi(splited[1]); err == nil {
					return true
				}
				if splited[1] == "*" {
					return true
				}
			}
			if splited[0] == "*" {
				if _, err := strconv.Atoi(splited[1]); err == nil {
					return true
				}
			}

		}
	}
	return false
}
func stringToIDHelper(str string) (int, int) {
	splited := strings.Split(str, "-")
	int1, _ := strconv.Atoi(splited[0])
	int2, _ := strconv.Atoi(splited[1])
	return int1, int2
}

// func convertToIdString(time int, idNum int) string {
// 	return fmt.Sprintf("%d-%d", time, idNum)
// }

func incrementID(idStr string) string {
	if strings.Contains(idStr, "-") {
		parts := strings.Split(idStr, "-")
		idNum, _ := strconv.Atoi(parts[1])
		time := parts[0]

		return fmt.Sprintf("%s-%d", time, idNum+1)
	}
	time, _ := strconv.Atoi(idStr)
	return fmt.Sprintf("%d-0", time)
}
func incrementTime(idStr string) string {
	if strings.Contains(idStr, "-") {
		parts := strings.Split(idStr, "-")
		time, _ := strconv.Atoi(parts[0])
		idNum := parts[1]

		return fmt.Sprintf("%d-%s", time+1, idNum)
	}
	time, _ := strconv.Atoi(idStr)
	return fmt.Sprintf("%d-0", time+1)
}

// checks if id1 is smaller than id2
func compareIds(id1 string, id2 string) bool {
	var time1, time2, idNum1, idNum2 int
	if strings.Contains(id1, "-") {
		parts := strings.Split(id1, "-")
		if parts[0] == "+inf" {
			return false
		}
		if parts[0] == "inf" {
			return true
		}
		time1, _ = strconv.Atoi(parts[0])
		idNum1, _ = strconv.Atoi(parts[1])

	}
	if strings.Contains(id2, "-") {
		parts := strings.Split(id2, "-")
		if parts[0] == "+inf" {
			return true
		}
		if parts[0] == "inf" {
			return false
		}
		time2, _ = strconv.Atoi(parts[0])
		idNum2, _ = strconv.Atoi(parts[1])

	}
	if time1 <= time2 {
		if time1 == time2 {
			return idNum1 <= idNum2
		}
		return true
	}
	return false
}

func GetAllIds(start string, end string, idMap map[string]map[string]string) []string {
	var ids []string

	for compareIds(start, end) {

		_, exists := idMap[start]
		if !exists {

			for i := 0; i < 5; i++ {
				_, exists = idMap[start]
				if exists {
					break
				}
				start = incrementID(start)
			}
			if !exists {
				start = incrementTime(start)
			}
			continue
		}

		ids = append(ids, start)
		start = incrementID(start)
	}
	return ids
}

func CheckForEndStart(start string, end string, key string) (string, string) {
	if start == "-" {
		start = storage.GlobalStorageStream[key]["low"]["low"]
	}
	if end == "+" {
		end = storage.GlobalStorageStream[key]["high"]["high"]
	}
	return start, end
}

// func isID(str string) bool {
// 	hasDash := strings.Contains(str, "-")
// 	if !hasDash {
// 		return false
// 	}
// 	parts := strings.Split(str, "-")
// 	for _, val := range parts {
// 		_, err := strconv.Atoi(val)
// 		if err != nil {
// 			return false
// 		}
// 	}
// }

func calcuclateStartIndex(keyIdx int, command []string) int {
	for !IsID(command[keyIdx]) {
		keyIdx++
	}
	return keyIdx
}

// func getXreadInfo(commands []string) map[string]string {
// 	result := make(map[string]string)

// 	return result
// }
