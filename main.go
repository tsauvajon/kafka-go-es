package main

import (
	"os"

	"flag"

	"fmt"

	"strconv"

	"github.com/go-redis/redis"
)

var (
	// Redis :
	Redis = initRedis()
)

func initRedis() *redis.Client {
	redisURL := os.Getenv("REDIS_URL")

	if redisURL == "" {
		redisURL = "127.0.0.1:6379"
	}

	return redis.NewClient(&redis.Options{
		Addr:     redisURL,
		Password: "",
		DB:       0,
	})
}

func main() {
	// producer and consumer would be in != repos in real life, so
	// we use the "flog" lib to chose either one
	act := flag.String("act", "producer", "Either : producer or consumer")
	partition := flag.String("partition", "0", "Partition to which the consumer program will be subscribing")

	flag.Parse()

	fmt.Printf("Welcome to the service : %s\n\n", *act)

	switch *act {
	case "producer":
		mainProducer()
	case "consumer":
		if part32int, err := strconv.ParseInt(*partition, 10, 32); err == nil {
			mainConsumer(int32(part32int))
		}
	}
}
