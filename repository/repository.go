package repository

import (
	"fmt"
	"time"

	"github.com/go-redis/redis"
)

var redisInstance *redisRepository

// Repository is...
type RedisRepository interface {
	Set(key, value string, exp time.Duration) error
	Get(key string) (string, error)
}

// GetRedis is ...
func GetRedis() RedisRepository {
	return redisInstance
}

// SetupRedis is...
func SetupRedis() {
	client := redis.NewClient(&redis.Options{
		Addr: "address",
	})

	_, err := client.Ping().Result()
	if err != nil {
		fmt.Println("error")
	}

	redisInstance = &redisRepository{client}
}

type redisRepository struct {
	client *redis.Client
}

func (r *redisRepository) Set(key, value string, exp time.Duration) error {
	return nil
}

func (r *redisRepository) Get(key string) (string, error) {
	return "", nil
}
