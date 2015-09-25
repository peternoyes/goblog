package main

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/defaults"
	"github.com/aws/aws-sdk-go/service/s3"
	"log"
)

func TestS3() {
	defaults.DefaultConfig.Region = aws.String("us-west-2")

	svc := s3.New(nil)
	result, err := svc.ListBuckets(&s3.ListBucketsInput{})

	if err != nil {
		log.Println("Failed to list buckets", err)
		return
	}
	for _, bucket := range result.Buckets {
		log.Printf("%s : %s\n", *bucket.Name, bucket.CreationDate)
	}

	input := s3.ListObjectsInput{}
	b := "blah"
	input.Bucket = &b

	result2, err2 := svc.ListObjects(&input)
	if err2 != nil {
		log.Println("Failed to list objects", err2)
		return
	}

	for _, object := range result2.Contents {
		log.Printf("%s : %s\n", *object.Key, object.LastModified)
	}
}
