package main

import (
	"bufio"
	"fmt"
	"os"

	"encoding/json"

	"github.com/Shopify/sarama"
	cluster "github.com/bsm/sarama-cluster"
)

func consumeEvents(consumer *cluster.Consumer) {
	var msgVal []byte
	var log interface{}
	var logMap map[string]interface{}
	var bankAccount *BankAccount
	var err error

	for {
		select {
		case err, more := <-consumer.Errors():
			if more {
				fmt.Printf("Kafka error : %s\n", err)
			}
		case msg := <-consumer.Messages():
			consumer.MarkOffset(msg, "")
			msgVal = msg.Value
			if err = json.Unmarshal(msgVal, &log); err != nil {
				fmt.Printf("Failed parsing: %s", err)
			} else {
				logMap = log.(map[string]interface{})
				logType := logMap["Type"]
				fmt.Printf("Processing %s : \n%s\n", logType, string(msgVal))

				switch logType {
				case "CreateEvent":
					event := new(CreateEvent)
					if err = json.Unmarshal(msgVal, &event); err == nil {
						bankAccount, err = event.Process()
					}
				case "DepositEvent":
					event := new(DepositEvent)
					if err = json.Unmarshal(msgVal, &event); err == nil {
						bankAccount, err = event.Process()
					}
				case "WithdrawEvent":
					event := new(WithdrawEvent)
					if err = json.Unmarshal(msgVal, &event); err == nil {
						bankAccount, err = event.Process()
					}
				case "TransferEvent":
					event := new(TransferEvent)
					if err = json.Unmarshal(msgVal, &event); err == nil {
						if bankAccount, err = event.Process(); err == nil {
							if target, err := FetchAccount(event.TargetID); err == nil {
								fmt.Printf("Target: %#v\n", *target)
							}
						}
					}
				default:
					fmt.Println("Unkown command : ", logType)
				}

				if err != nil {
					fmt.Printf("Error processing: %s\n", err)
				} else {
					fmt.Printf("Bank account : %#v\n\n", *bankAccount)
				}
			}
		}
	}
}

func mainConsumer(partition int32) {
	config := cluster.NewConfig()
	config.Consumer.Offsets.Initial = sarama.OffsetNewest
	consumer, err := cluster.NewConsumer(brokers, "consumer", topics, config)

	kafka := newKafkaConsumer()
	defer kafka.Close()

	/*
		Note that we are using sarama.OffsetOldest, which means that Kafka will
		be sending a log all the way from the first message ever created. This may
		be good for development mode since we don't need to write message after
		message to test out features. In production, we definitely would want to
		change it with sarama.OffsetNewest, which will only ask for the newest
		messages that haven't been sent to us.
	*/
	// consumer, err := kafka.ConsumePartition(topic, partition, sarama.OffsetOldest)

	if err != nil {
		fmt.Printf("Kafka error: %s\n", err)
		os.Exit(-1)
	}

	go consumeEvents(consumer)

	fmt.Println("Press [enter] to exit consumer")
	bufio.NewReader(os.Stdin).ReadString('\n')
	fmt.Println("Terminating...")
}
