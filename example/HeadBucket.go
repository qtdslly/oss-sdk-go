package main

import (
	"context"
	"example/client"
	"fmt"
	"oss-sdk-go/oss"
	"oss-sdk-go/service/s3"
)

func main() {
	bucket := "buc11"
	client := client.CreateClient()
	_, err := client.HeadBucket(context.TODO(), &s3.HeadBucketInput{
		Bucket: oss.String(bucket),
	})
	if err != nil {
		fmt.Println(err.Error())
	}
}
