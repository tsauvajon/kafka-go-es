package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func mainProducer() {
	var err error
	reader := bufio.NewReader(os.Stdin)
	kafka := newKafkaSyncProducer()

	for {
		fmt.Print("-> ")
		text, _ := reader.ReadString('\n')
		text = strings.Replace(text, "\n", "", -1)
		args := strings.Split(text, "###")
		cmd := args[0]

		switch cmd {
		case "create":
			if len(args) == 2 {
				accName := args[1]
				event := NewCreateAccountEvent(accName)
				sendMsg(kafka, event)
			} else {
				fmt.Println("Only specify create###Account Name")
			}
		default:
			fmt.Printf("Unknown command %s, only: create, deposit, withdraw, transfer\n", cmd)
		}

		if err != nil {
			fmt.Printf("Error: %s\n", err)
			err = nil
		}
	}
}
