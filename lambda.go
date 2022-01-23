package main

import (
	"context"
	"fmt"
	"log"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

/*
** $ GOOS=linux GOARCH=amd64 go build -o main lambda.go && zip main.zip main
 */

type MyEvent struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
}

type Response struct {
	BucketName string `json:"bucketName"`
	Key        string `json:"key"`
	Size       int64  `json:"size"`
	StatusCode int    `json:"statusCode"`
}

func HandleRequest(ctx context.Context, name MyEvent) (string, error) {
	fmt.Println(name.FirstName)
	fmt.Println(name.LastName)
	return fmt.Sprintf("Hello %s %s!", name.FirstName, name.LastName), nil
}

func s3Lambda(ctx context.Context, event events.S3Event) (interface{}, error) {
	response := Response{}
	for _, record := range event.Records {
		response.BucketName = record.S3.Bucket.Name
		response.Key = record.S3.Object.Key
		response.Size = record.S3.Object.Size
	}
	response.StatusCode = 200
	log.Println("response: ", response)
	return response, nil
}

func main() {
	lambda.Start(HandleRequest)
	lambda.Start(s3Lambda)
}
