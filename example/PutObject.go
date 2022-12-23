package main

import (
	"context"
	"example/client"
	"fmt"
	"oss-sdk-go/oss"
	"oss-sdk-go/service/s3"
)

func main() {
	bucket := "buc1"
	key := "a.txt"
	client := client.CreateClient()
	_, err := client.PutObject(context.TODO(), &s3.PutObjectInput{
		Bucket: oss.String(bucket),
		Key:    oss.String(key),
	})
	if err != nil {
		fmt.Println(err.Error())
	}
}
