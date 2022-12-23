package main

import (
	"context"
	"example/client"
	"fmt"
	"oss-sdk-go/oss"
	"oss-sdk-go/service/s3"
)

func main() {
	bucket := "buc2"
	key := "b.txt"
	client := client.CreateClient()
	_, err := client.DeleteObject(context.TODO(), &s3.DeleteObjectInput{
		Bucket: oss.String(bucket),
		Key:    oss.String(key),
	})
	if err != nil {
		fmt.Println(err.Error())
	}
}
