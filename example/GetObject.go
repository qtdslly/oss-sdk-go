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
	result, err := client.GetObject(context.TODO(), &s3.GetObjectInput{
		Bucket: oss.String(bucket),
		Key:    oss.String(key),
	})

	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("etag:" + *result.ETag)
	}
}
