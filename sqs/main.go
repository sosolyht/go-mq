package main

import (
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/sqs"
	"log"
)

func main() {
	ctx := context.Background()
	// LoadDefaultConfig 로컬에 있는 aws config load
	cfg, err := config.LoadDefaultConfig(ctx, config.WithRegion("ap-northeast-2"))
	if err != nil {
		log.Print(err)
	}

	client := sqs.NewFromConfig(cfg)

	urlInput := &sqs.GetQueueUrlInput{
		QueueName: aws.String("test3"),
	}

	urlOutput, err := client.GetQueueUrl(ctx, urlInput)
	if err != nil {
		log.Println(err)
	}

	goodInput := &sqs.SendMessageInput{
		MessageBody:  aws.String("테스트"),
		QueueUrl:     urlOutput.QueueUrl,
		DelaySeconds: 0,
	}

	goodOutput, err := client.SendMessage(ctx, goodInput)
	if err != nil {
		log.Println(err)
	}

	fmt.Println(*goodOutput.MessageId)

}
