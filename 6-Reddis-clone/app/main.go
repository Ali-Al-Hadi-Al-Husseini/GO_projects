package main

import (
	"fmt"
	"net"
	"os"

	"github.com/Ali-Al-Hadi-Al-Husseini/GO_projects/tree/main/6-Reddis-clone/app/utils"
)

func main() {
	flags := utils.GetFlags()

	port := flags["port"].(int)

	// Uncomment this block to pass the first stage

	l, err := net.Listen("tcp", fmt.Sprintf("0.0.0.0:%d", port))
	if err != nil {
		fmt.Printf("Failed to bind to port %d\n", port)
		os.Exit(1)
	}

	fmt.Printf("Server is running on %d\n", port)
	roleInfo := utils.GetServerRole(flags["replicaof"].(string))

	if roleInfo["role"] == "slave" {
		utils.IntalizeSlavment(roleInfo, port)

	}
	listenForConnections(l, roleInfo)

}
