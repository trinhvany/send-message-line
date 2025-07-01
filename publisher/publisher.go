package publisher

import (
	"encoding/json"
	"fmt"
	"small_demo_go/client/rds"
	"small_demo_go/client/sqs"
)

func StartPublisher() {
	rdsClient, err := rds.NewClient();

	if err != nil {
		fmt.Println("Connection error: ",err)
		return
	}

	messages, err := rdsClient.GetMessage()
	if err != nil {
		fmt.Println("Error while retrieving message: ", err)
		return
	}

	sqs.Init()
	queueName := "my-queue"

	url, err := sqs.CreateQueue(queueName)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("📦 Queue URL:", url)
	fmt.Println("message", messages)

	for _,message := range messages {
		fmt.Println(message);
		details, errDetail := rdsClient.GetMessageDetail(message.ID)
		if errDetail != nil {
			fmt.Println(errDetail)
		}
		var detailList []map[string]interface{}

		for _, detail := range details {
			fmt.Println("Details: ", detail.ID)
			var obj map[string]interface{}
			err := json.Unmarshal([]byte(detail.ContentJson), &obj)
			if err != nil {
				fmt.Println(fmt.Sprintf("JSON error: %v", err))
			}
			detailList = append(detailList, obj)
		}

		if len(details) > 0 {
			jsonDataIndent, err := json.MarshalIndent(detailList, "", "  ")
			if err != nil {
				fmt.Println(err)
			}
			err = sqs.SendMessage(url, string(jsonDataIndent))
		}
	}
}