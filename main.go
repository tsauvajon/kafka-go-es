package main

import (
	"os"

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
	mainProducer()
}
