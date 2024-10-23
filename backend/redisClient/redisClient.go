package redisclient

import (
	"github.com/go-redis/redis/v8"
)

var redisclient *redis.Client

func RedisClient() *redis.Client {

	redisclient = redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})

	return redisclient
}
