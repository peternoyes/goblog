// Copyright (c) 2015 Peter Noyes

package main

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/defaults"
	"github.com/aws/aws-sdk-go/service/s3"
	"strings"
	"io"
	"io/ioutil"
	"fmt"
	"time"
)

type DriverS3 struct {
	Bucket string
	Region string
	Svc *s3.S3
}

func (d *DriverS3) New() {
	fmt.Println("driverS3.New")
	defaults.DefaultConfig.Region = aws.String(d.Region)
	d.Svc = s3.New(nil)
}

func (d *DriverS3) GetConfig() ([]byte, error) {
	fmt.Println("driverS3.GetConfig")
	key := "config.json"
	input := s3.GetObjectInput{
		Bucket: aws.String(d.Bucket),
		Key: aws.String(key),
	}

	result, err := d.Svc.GetObject(&input)

	if err != nil {
		return nil, err
	}

	defer result.Body.Close()

	return ioutil.ReadAll(result.Body)
}

func (d *DriverS3) GlobMarkdown() ([]*PostStub, error) {
	input := s3.ListObjectsInput{}
	input.Bucket = &d.Bucket

	result, err := d.Svc.ListObjects(&input)
	if err != nil {
		return nil, err
	}

	ret := make([]*PostStub, 0)

	for _, object := range result.Contents {
		key := *object.Key
		fmt.Println("Key: ", key)
		if strings.HasSuffix(key, ".md") {
			var title string
			var date time.Time
			title, date, err = GetDateAndTitleFromFile(key)
			if err != nil {
				return nil, err
			}
			mod := *object.LastModified

			ret = append(ret, &PostStub{key, title, date, mod, nil})
		}
	}

	return ret, nil
}

func (d *DriverS3) Open(file string) (io.ReadCloser, error) {
	fmt.Println("driverS3.Open")
	
	input := s3.GetObjectInput{
		Bucket: aws.String(d.Bucket),
		Key: aws.String(file),
	}

	result, err := d.Svc.GetObject(&input)

	if err != nil {
		return nil, err
	}

	return result.Body, nil
}

