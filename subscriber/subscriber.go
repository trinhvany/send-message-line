package subcriber

import (
	"fmt"
	"small_demo_go/client/redis"
	"small_demo_go/client/s3"
	"small_demo_go/client/sqs"
	"small_demo_go/system"
	"small_demo_go/utils"
	"strconv"
	"strings"
)
var redisClient = redis.GetClient()
var redisCtx = redis.GetContext()

func StartSubcriber() {
	s3.Init()
	s3.PutObject( "my-bucket", "userid.txt", []byte(system.UserIDFake))

	userIDsData, errS3 := s3.GetObject("my-bucket", "userid.txt")
	if errS3 != nil {
		panic(errS3)
	}
	userIDs := strings.Split(string(userIDsData), ",")

	sqs.Init()
	queueName := "my-queue"
	queueUrl, err := sqs.CreateQueue(queueName)
	if err != nil {
		panic(err)
	}

	messages, err := sqs.ReceiveMessages(queueUrl, 10)
	if err != nil {
		panic(err)
	}

	for index, msg := range messages {
		key := "messageSentToUser_" + strconv.Itoa(index)
		for _,uid := range userIDs {
			added, err := redisClient.SAdd(redisCtx, key, strings.TrimSpace(uid)).Result()
			if err != nil {
				fmt.Println("Add uid to Redis failed: ", err)
				continue
			}
			if added == 1 {
				body := strings.ReplaceAll(system.LinePushTemplate, "<<TO>>", strings.TrimSpace(uid))
				body = strings.ReplaceAll(body, "<<MESSAGES>>", msg)
				utils.PushLineMessage(body)
			} else {
				fmt.Printf("skipped (duplicate) %s", uid)
			}
		}
	}

	deleteSentUserKeys("messageSentToUser_")
}

func deleteSentUserKeys(prefix string) {
	keys, err := redisClient.Keys(redisCtx, prefix +"*").Result()
	if err != nil {
		fmt.Println("Error when getting key:", err)
		return
	}

	if len(keys) == 0 {
		fmt.Println("No key to delete")
		return
	}

	_, err = redisClient.Del(redisCtx, keys...).Result()
	if err != nil {
		fmt.Println("Delete key failed:", err)
	} else {
		fmt.Printf("Deleted %d key(s): %v\n", len(keys), keys)
	}
}