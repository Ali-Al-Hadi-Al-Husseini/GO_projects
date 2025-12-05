package storage

import "net"

var GlobalStorage = map[string]string{}
var GlobalStorageTimer = map[string]int64{}
var GlobalStorageArray = map[string][]string{}
var GlobalStorageArrayBlpop = map[string]bool{}
var GlobalStorageStream = map[string]map[string]map[string]string{}
var GlovalStreamIDExists = map[string]bool{}
var GlobalStorageStreamBlock = map[string]map[string]map[string]string{}
var GlobalMutliStorage = map[net.Conn][][]string{}
var GlobalSlaveAddress = []string{}
var LastAddedStream string = "0-0"
var NonWriteCommands = map[string]bool{
	"ping":     true,
	"echo":     true,
	"replconf": true,
	"psync":    true,
	"info":     true,
}
