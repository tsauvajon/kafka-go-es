package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
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
		case "deposit":
			if len(args) == 3 {
				accName := args[1]
				trimmedAmount := strings.Replace(args[2], "\r", "", -1)
				if amount, err := strconv.Atoi(trimmedAmount); err != nil {
					fmt.Printf("invalid number : %s\n", err)
				} else {
					event := NewDepositEvent(accName, amount)
					sendMsg(kafka, event)
				}
			} else {
				fmt.Println("deposit###Account Name###Amount")
			}
		case "withdraw":
			if len(args) == 3 {
				accName := args[1]
				trimmedAmount := strings.Replace(args[2], "\r", "", -1)
				if amount, err := strconv.Atoi(trimmedAmount); err != nil {
					fmt.Printf("invalid number : %s\n", err)
				} else {
					event := NewWithdrawEvent(accName, amount)
					sendMsg(kafka, event)
				}
			} else {
				fmt.Println("withdraw###Account Name###Amount")
			}
		case "transfer":
			if len(args) == 4 {
				accName := args[1]
				target := args[2]
				trimmedAmount := strings.Replace(args[3], "\r", "", -1)
				if amount, err := strconv.Atoi(trimmedAmount); err != nil {
					fmt.Printf("invalid number : %s\n", err)
				} else {
					event := NewTransferEvent(accName, target, amount)
					sendMsg(kafka, event)
				}
			} else {
				fmt.Println("transfer###Account Name###Target Name###Amount")
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
