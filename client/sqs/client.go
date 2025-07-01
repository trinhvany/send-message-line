package sqs

import (
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/sqs"
)

var (
	ctx    = context.Background()
	client *sqs.Client
)

func Init() {
	cfg, err := config.LoadDefaultConfig(ctx,
		config.WithRegion("us-east-1"),
		config.WithCredentialsProvider(credentials.NewStaticCredentialsProvider("dummy", "dummy", "")),
		config.WithEndpointResolverWithOptions(
			aws.EndpointResolverWithOptionsFunc(func(service, region string, _ ...interface{}) (aws.Endpoint, error) {
				return aws.Endpoint{
					URL:           "http://localhost:9324", // ElasticMQ local
					SigningRegion: "us-east-1",
				}, nil
			}),
		),
	)
	if err != nil {
		panic(fmt.Sprintf("❌ Connection error to SQS: %v", err))
	}

	client = sqs.NewFromConfig(cfg)
	fmt.Println("✅ SQS is ready")
}

func GetClient() *sqs.Client {
	return client
}

func GetContext() context.Context {
	return ctx
}

func CreateQueue(queueName string) (string, error) {
	out, err := client.CreateQueue(ctx, &sqs.CreateQueueInput{
		QueueName: &queueName,
	})
	if err != nil {
		return "", fmt.Errorf("Error creating queue: %v", err)
	}
	return *out.QueueUrl, nil
}

func SendMessage(queueUrl, message string) error {
	_, err := client.SendMessage(ctx, &sqs.SendMessageInput{
		QueueUrl: &queueUrl,
		MessageBody: &message,
	})
	if err != nil {
		return fmt.Errorf("Error sending message: %v", err)
	}
	fmt.Printf("Message sent to %s: %s\n", queueUrl, message)
	return nil
}

func ReceiveMessages(queueUrl string, maxMessages int32) ([]string, error) {
	out, err := client.ReceiveMessage(ctx, &sqs.ReceiveMessageInput{
		QueueUrl: &queueUrl,
		MaxNumberOfMessages: maxMessages,
		WaitTimeSeconds: 1,
	})
	if err != nil {
		return nil, fmt.Errorf("Error receiving message: %v", err)
	}

	var messages []string
	for _, m := range out.Messages {
		messages = append(messages, *m.Body)
	}
	return messages, nil
}
