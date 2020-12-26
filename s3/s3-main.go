package main

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

func main() {
	//sess, err := session.NewSession()
	// Equivalent to session.New
	//sess, err := session.NewSessionWithOptions(session.Options{})
	// Specify profile to load for the session's config
	//sess, err := session.NewSessionWithOptions(session.Options{
	//	Profile: "profile_name",
	//})
	// Specify profile for config and region for requests
	//sess, err := session.NewSessionWithOptions(session.Options{
	//	Config: aws.Config{Region: aws.String("us-east-2")},
	//	Profile: "profile_name",
	//})
	// Force enable Shared Config support
	//sess, err := session.NewSessionWithOptions(session.Options{
	//	SharedConfigState: SharedConfigEnable,
	//})
	// Assume an IAM role with MFA prompting for token code on stdin
	//sess := session.Must(session.NewSessionWithOptions(session.Options{
	//	AssumeRoleTokenProvider: stscreds.StdinTokenProvider,
	//	SharedConfigState: SharedConfigEnable,
	//}))
	// Initialize a session in us-west-2 that the SDK will use to load credentials
	// from the shared credentials file. (~/.aws/credentials).
	sess, err := session.NewSession(&aws.Config{
		Endpoint: aws.String("http://localhost:4566"),
		Region:   aws.String("us-east-1")},
	)
	if err != nil {
		fmt.Println(err)
	}
	svc := s3.New(sess)

	// Pre-defined values
	bucket := "MyBucket"
	tagName1 := "Cost Center"
	tagValue1 := "123456"
	tagName2 := "Stack"
	tagValue2 := "MyTestStack"

	_, err = svc.CreateBucket(&s3.CreateBucketInput{
		Bucket: aws.String(bucket),
	})
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	// Create input for PutBucket method
	putInput := &s3.PutBucketTaggingInput{
		Bucket: aws.String(bucket),
		Tagging: &s3.Tagging{
			TagSet: []*s3.Tag{
				{
					Key:   aws.String(tagName1),
					Value: aws.String(tagValue1),
				},
				{
					Key:   aws.String(tagName2),
					Value: aws.String(tagValue2),
				},
			},
		},
	}
	_, err = svc.PutBucketTagging(putInput)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	// Now show the tags
	// Create input for GetBucket method
	getInput := &s3.GetBucketTaggingInput{
		Bucket: aws.String(bucket),
	}

	result, err := svc.GetBucketTagging(getInput)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	numTags := len(result.TagSet)

	if numTags > 0 {
		fmt.Println("Found", numTags, "Tag(s):")
		fmt.Println("")

		for _, t := range result.TagSet {
			fmt.Println("  Key:  ", *t.Key)
			fmt.Println("  Value:", *t.Value)
			fmt.Println("")
		}
	} else {
		fmt.Println("Did not find any tags")
	}
}
