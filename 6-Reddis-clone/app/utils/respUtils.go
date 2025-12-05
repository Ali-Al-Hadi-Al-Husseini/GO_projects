package utils

import (
	"fmt"
	"strings"

	"github.com/Ali-Al-Hadi-Al-Husseini/GO_projects/tree/main/6-Reddis-clone/app/storage"
)

func ConvertToRESP(args string) string {
	// var resp string

	resp := fmt.Sprintf("$%d\r\n%s\r\n", len(args), args)

	return resp
}
func ArrayToRESP(args []string) string {
	var b strings.Builder
	// Write array header
	b.WriteString(fmt.Sprintf("*%d\r\n", len(args)))
	for _, arg := range args {
		b.WriteString(fmt.Sprintf("$%d\r\n%s\r\n", len(arg), arg))
	}
	return b.String()
}

func ConvertToRESPStream(keys []string, data map[string]map[string]string) string {
	result := fmt.Sprintf("*%d\r\n", len(keys))

	for _, id := range keys {
		fields, ok := data[id]
		if !ok {
			continue // skip if key not in map
		}

		result += "*2\r\n"
		result += fmt.Sprintf("$%d\r\n%s\r\n", len(id), id)

		// Add the key-value pairs inside the map
		result += fmt.Sprintf("*%d\r\n", len(fields)*2)
		for fieldKey, fieldVal := range fields {
			result += fmt.Sprintf("$%d\r\n%s\r\n", len(fieldKey), fieldKey)
			result += fmt.Sprintf("$%d\r\n%s\r\n", len(fieldVal), fieldVal)
		}
	}

	return result
}
func ConvertToRESPXread(streamKey string, keys []string, data map[string]map[string]string) string {
	result := "*1\r\n" // One stream
	result += "*2\r\n"
	result += fmt.Sprintf("$%d\r\n%s\r\n", len(streamKey), streamKey)

	result += fmt.Sprintf("*%d\r\n", len(keys)) // Number of entries for this stream

	for _, id := range keys {
		fields, ok := data[id]
		if !ok {
			continue
		}

		result += "*2\r\n"
		result += fmt.Sprintf("$%d\r\n%s\r\n", len(id), id)

		// key-value pairs
		result += fmt.Sprintf("*%d\r\n", len(fields)*2)
		for fieldKey, fieldVal := range fields {
			result += fmt.Sprintf("$%d\r\n%s\r\n", len(fieldKey), fieldKey)
			result += fmt.Sprintf("$%d\r\n%s\r\n", len(fieldVal), fieldVal)
		}
	}

	return result
}
func ConvertToRESPStreamMulti(
	streamKeys []string,
	streamIDs [][]string,
) string {
	result := fmt.Sprintf("*%d\r\n", len(streamKeys))

	for i, streamKey := range streamKeys {
		ids := streamIDs[i]
		result += "*2\r\n"
		result += fmt.Sprintf("$%d\r\n%s\r\n", len(streamKey), streamKey)
		result += fmt.Sprintf("*%d\r\n", len(ids)) // Number of entries in this stream

		for _, id := range ids {
			fieldsMap, ok := storage.GlobalStorageStream[streamKey]
			if !ok {
				continue // stream key not found
			}
			fields, ok := fieldsMap[id]
			if !ok {
				continue // ID not found in this stream
			}

			result += "*2\r\n"
			result += fmt.Sprintf("$%d\r\n%s\r\n", len(id), id)
			result += fmt.Sprintf("*%d\r\n", len(fields)*2)

			for fieldKey, fieldVal := range fields {
				result += fmt.Sprintf("$%d\r\n%s\r\n", len(fieldKey), fieldKey)
				result += fmt.Sprintf("$%d\r\n%s\r\n", len(fieldVal), fieldVal)
			}
		}
	}

	return result
}

func MapToRESPBulkString(data map[string]string) string {
	var content string
	for key, value := range data {
		content += key + ":" + value + "\n"
	}
	// Remove trailing newline
	if len(content) > 0 {
		content = content[:len(content)-1]
	}

	// Wrap in RESP Bulk String format
	return fmt.Sprintf("$%d\r\n%s\r\n", len(content), content)
}

// '*3\r\n$3\r\nGET\r\n$3\r\nfoo\r\n$3\r\nbar\r\n'
