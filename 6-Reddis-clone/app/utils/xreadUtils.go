package utils

import "github.com/Ali-Al-Hadi-Al-Husseini/GO_projects/tree/main/6-Reddis-clone/app/storage"

func GetXreadData(commands []string, start int) string {

	var ids [][]string
	var keys []string
	startIdx := calcuclateStartIndex(2, commands)
	for i := start; startIdx < len(commands); i++ {
		key := commands[i]
		idMap := storage.GlobalStorageStream[key]

		start := commands[startIdx]
		end := "+"
		start, end = CheckForEndStart(start, end, key)
		ids = append(ids, GetAllIds(start, end, idMap))
		keys = append(keys, key)
		startIdx++
	}
	return ConvertToRESPStreamMulti(keys, ids)
}
