package main

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sns"

	"fmt"
	"os"
)

func main() {
	// Initialize a session that the SDK will use to load
	sess, err := session.NewSession(&aws.Config{
		Endpoint: aws.String("http://localhost:4566"),
		Region:   aws.String("us-east-1")},
	)
	svc := sns.New(sess)
	snsTopicName := "sns-test"

	result, err := svc.CreateTopic(&sns.CreateTopicInput{
		Name: aws.String(snsTopicName),
	})
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	fmt.Println(*result.TopicArn)

	//List topics
	listResult, err := svc.ListTopics(nil)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	for _, t := range listResult.Topics {
		fmt.Println(*t.TopicArn)
	}
}
