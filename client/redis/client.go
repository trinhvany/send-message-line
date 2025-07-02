package redis

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
)

var (
	ctx    = context.Background()
	client *redis.Client
)

func init() {
	client = redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
		Password: "",
		DB: 0,
	})

	if _, err := client.Ping(ctx).Result(); err != nil {
		panic(fmt.Sprintf("❌ Redis connect failed: %v", err))
	}

	fmt.Println("✅ Redis connected")
}

func GetClient() *redis.Client {
	return client
}

func GetContext() context.Context {
	return ctx
}