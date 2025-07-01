package s3

import (
	"bytes"
	"context"
	"fmt"
	"io"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

var (
	ctx = context.Background()
	client *s3.Client
)

func Init() {
	cfg, err := config.LoadDefaultConfig(ctx,
		config.WithRegion("us-east-1"),
		config.WithCredentialsProvider(credentials.NewStaticCredentialsProvider("fake", "fake", "")),
		config.WithEndpointResolverWithOptions(
			aws.EndpointResolverWithOptionsFunc(func(service, region string, options ...interface{}) (aws.Endpoint, error) {
				return aws.Endpoint{
					URL: "http://localhost:4567",
					SigningRegion: "us-east-1",
				}, nil
			}),
		),
	)

	if err != nil {
		panic(fmt.Sprintf("❌ Connection error to S3: %v", err))
	}

	client = s3.NewFromConfig(cfg, func(o *s3.Options) {
		o.UsePathStyle = true
	})

	fmt.Println("✅ S3 is ready")
}

func GetClient() *s3.Client {
	return client
}

func GetContext() context.Context {
	return ctx
}

func PutObject(bucket string, key string, content []byte) error {
	cli := GetClient()

	_, err := cli.PutObject(ctx, &s3.PutObjectInput{
		Bucket: &bucket,
		Key:    &key,
		Body:   bytes.NewReader(content),
	})
	if err != nil {
		return fmt.Errorf("PutObject failed: %v", err)
	}

	fmt.Printf("File written [%s] into bucket [%s]\n", key, bucket)
	return nil
}

// GetObject đọc nội dung từ 1 object trong S3 (dạng []byte)
func GetObject(bucket string, key string) ([]byte, error) {
	cli := GetClient()

	resp, err := cli.GetObject(ctx, &s3.GetObjectInput{
		Bucket: &bucket,
		Key:    &key,
	})
	if err != nil {
		return nil, fmt.Errorf("GetObject failed: %v", err)
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("Error reading content: %v", err)
	}

	return data, nil
}
