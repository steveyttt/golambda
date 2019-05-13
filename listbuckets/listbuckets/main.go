package main

import (
	"fmt"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

func handler() {

	//establish a session in ap-southeast-2
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String("ap-southeast-2")},
	)
	//create an s3 client
	s3svc := s3.New(sess)

	result, err := s3svc.ListBuckets(nil)
	if err != nil {
		fmt.Printf("Unable to list buckets, %v", err)
	}

	// fmt.Println("buckets:")
	fmt.Println(result.Buckets)

	// for _, b := range result.Buckets {
	// 	fmt.Printf("%s created on %s\n", aws.StringValue(b.Name), aws.TimeValue(b.CreationDate))
	// }

	input := &s3.GetBucketPolicyInput{
		Bucket: aws.String("tysonphotos"),
	}

	results, err := s3svc.GetBucketPolicy(input)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			default:
				fmt.Println(aerr.Error())
			}
		} else {
			// Print the error, cast err to awserr.Error to get the Code and
			// Message from an error.
			fmt.Println(err.Error())
		}
		return
	}

	fmt.Println(results)

}

func main() {
	lambda.Start(handler)
}
