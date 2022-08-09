package redis

import (
	"time"

	"github.com/go-redis/redis"
)

var client = ConnectionToRedis()

func ConnectionToRedis() *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})
}

func Set(key, value string) error {
	_, err := client.Set(key, value, time.Hour).Result()
	return err
}

func Get(key string) (string, error) {
	return client.Get(key).Result()
}
