package main

import (
	"small_demo_go/publisher"
	"small_demo_go/subscriber"
	"time"
)

func main() {
	// redisClient := redis.GetClient()
	// ctx := redis.GetContext()
	// val, err := redisClient.Get(ctx, "user:list").Result()
	// if err != nil {
	// 	panic(err)
	// }

	// var users []redis.User
	// if err := json.Unmarshal([]byte(val), &users); err != nil {
	// 	panic(err)
	// }

	// fmt.Println("📦 Danh sách user từ Redis:")
	// for _, user := range users {
	// 	fmt.Printf("- ID: %d, Name: %s\n", user.ID, user.Name)
	// }

	publisher.StartPublisher()
	subcriber.StartSubcriber()

	
	time.Sleep(7 * time.Second)
}