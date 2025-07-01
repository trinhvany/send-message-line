package redis

import (
	"context"
	"encoding/json"
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

	preloadUsers()
}

func GetClient() *redis.Client {
	return client
}

func GetContext() context.Context {
	return ctx
}

func preloadUsers() {
	exists, err := client.Exists(ctx, "user:list").Result()
	if err != nil {
		fmt.Println("Key does not exist: " + err.Error())
	}

	if exists == 0 {
		data, _ := json.Marshal(defaultUsers)
		err = client.Set(ctx, "user:list", data, 0).Err()
		if err != nil {
			panic("Error during data input: " + err.Error())
		}

		fmt.Println("Initialized list into Redis")
	} else {
		fmt.Println("already exists in Redis")
	}
}