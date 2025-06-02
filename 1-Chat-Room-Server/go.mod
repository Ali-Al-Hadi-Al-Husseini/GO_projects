module github.com/Ali-Al-Hadi-Al-Husseini/GO_projects/tree/main/1-Chat-Room-Server

go 1.22.3

	defer conn.Close()
	reader := bufio.NewReader(conn)
	for {
		contnt, err := reader.ReadString('\n')
		if err != nil {
			return
		}

		fmt.Println("Got: ", contnt)
	}

