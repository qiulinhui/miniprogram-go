package models

import (
	"context"

	"github.com/go-redis/redis/v8"
)

var (
	RDB *redis.Client
	Ctx = context.Background()
)

func ConnectRedis() {
	RDB = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})
}
