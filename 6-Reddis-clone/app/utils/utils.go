package utils

import (
	"crypto/rand"
	"flag"
	"fmt"
	"math/big"
	"os"
	"strconv"
	"strings"

	"github.com/Ali-Al-Hadi-Al-Husseini/GO_projects/tree/main/6-Reddis-clone/app/storage"
)

func RemoveResp(input string) string {
	var result []string
	lines := strings.Split(input, "\r\n")

	for i := 0; i < len(lines); i++ {
		line := lines[i]

		// Skip empty or metadata lines
		if line == "" || line[0] == '*' || line[0] == '$' {
			continue
		}

		// Remove + or : from beginning of actual values
		if line[0] == '+' || line[0] == ':' {
			result = append(result, line[1:])
		} else {
			result = append(result, line)
		}
	}

	return strings.Join(result, " ")
}

func PoplulateFeildsBlock(commands []string, key string, id string) {
	for i := 3; i < len(commands); i += 2 {
		field := commands[i]
		value := commands[i+1]

		storage.GlobalStorageStreamBlock[key][id][field] = value
	}
}

// ConvertToRESPExecResult generates RESP array with correct types (bulk string or integer)

func EncodeRESPArray(results []interface{}) string {
	resp := fmt.Sprintf("*%d\r\n", len(results))

	for _, val := range results {
		switch v := val.(type) {
		case int:
			resp += fmt.Sprintf(":%d\r\n", v)
		case string:
			// ✅ Handle RESP error
			if strings.HasPrefix(v, "-ERR") {
				resp += fmt.Sprintf("%s\r\n", v) // Don't wrap as bulk string
			} else {
				resp += fmt.Sprintf("$%d\r\n%s\r\n", len(v), v)
			}
		default:
			s := fmt.Sprint(v)
			resp += fmt.Sprintf("$%d\r\n%s\r\n", len(s), s)
		}
	}

	return resp
}

type FlagDef struct {
	Name        string
	Type        string // "int", "bool", "string"
	Default     string
	Description string
}

func ParseFlags(defs []FlagDef) map[string]interface{} {
	values := make(map[string]interface{})

	for _, def := range defs {
		switch def.Type {
		case "int":
			defVal, _ := strconv.Atoi(def.Default)
			values[def.Name] = flag.Int(def.Name, defVal, def.Description)
		case "bool":
			defVal, _ := strconv.ParseBool(def.Default)
			values[def.Name] = flag.Bool(def.Name, defVal, def.Description)
		case "string":
			values[def.Name] = flag.String(def.Name, def.Default, def.Description)
		default:
			fmt.Fprintf(os.Stderr, "Unsupported flag type: %s\n", def.Type)
			os.Exit(1)
		}
	}

	flag.Parse()

	// Extract final parsed values into a flat map
	result := make(map[string]interface{})
	for k, v := range values {
		switch val := v.(type) {
		case *int:
			result[k] = *val
		case *bool:
			result[k] = *val
		case *string:
			result[k] = *val
		}
	}
	return result
}

func GetFlags() map[string]interface{} {
	flagDefs := []FlagDef{
		{"port", "int", "6379", "Port number"},

		{"replicaof", "string", "master", "Master address and port"},
	}

	flags := ParseFlags(flagDefs)

	return flags
}
func GetServerRole(replica string) map[string]string {
	result := make(map[string]string)

	if replica == "master" {
		result["role"] = "master"
	}
	if strings.Contains(replica, " ") {
		parts := strings.SplitN(replica, " ", 2)
		if len(parts) != 2 {
			result["err"] = "-ERR invalid replicaof format\r\n"
		}
		address := parts[0]
		port, err := strconv.Atoi(parts[1])
		if err != nil {
			result["err"] = "-ERR invalid port number\r\n"
		}
		result["addr"] = address
		result["port"] = strconv.Itoa(port)
		result["role"] = "slave"
	}
	return result

}

const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
const idLength = 40

func GenerateID() string {
	result := make([]byte, idLength)
	for i := range result {
		num, _ := rand.Int(rand.Reader, big.NewInt(int64(len(charset))))
		result[i] = charset[num.Int64()]
	}
	return string(result)
}
