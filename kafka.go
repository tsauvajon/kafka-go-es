package main

import (
	"encoding/json"
	"fmt"

	"os"

	"github.com/Shopify/sarama"
)

var (
	brokers = []string{"127.0.0.1:9093", "127.0.0.1:9093"}
	topic   = "banku-transactions-2"
	topics  = []string{topic}
)

func newKafkaConfiguration() *sarama.Config {
	conf := sarama.NewConfig()
	conf.Producer.RequiredAcks = sarama.WaitForAll
	conf.Producer.Return.Successes = true
	conf.ChannelBufferSize = 1
	conf.Version = sarama.V0_10_1_0
	return conf
}

func newKafkaSyncProducer() sarama.SyncProducer {
	kafka, err := sarama.NewSyncProducer(brokers, newKafkaConfiguration())

	if err != nil {
		fmt.Printf("KafkaError: %s\n", err.Error())
		os.Exit(-1)
	}

	return kafka
}

func sendMsg(kafka sarama.SyncProducer, event interface{}) error {
	json, err := json.Marshal(event)

	if err != nil {
		return err
	}

	msgLog := &sarama.ProducerMessage{
		Topic: topic,
		Value: sarama.StringEncoder(string(json)),
	}

	partition, offset, err := kafka.SendMessage(msgLog)

	if err != nil {
		fmt.Printf("Kafka error : %s\n", err)
	}

	fmt.Printf("Message: %+v\n", event)
	fmt.Printf("Message is stored in partition %d, offset %d\n",
		partition, offset)

	return nil
}

func newKafkaConsumer() sarama.Consumer {
	consumer, err := sarama.NewConsumer(brokers, newKafkaConfiguration())

	if err != nil {
		fmt.Printf("Kafka error: %s\n", err)
		os.Exit(-1)
	}

	return consumer
}
