package main

import (
	"context"
	"example/client"
	"fmt"
	"oss-sdk-go/oss"
	"oss-sdk-go/service/s3"
)

func main() {
	bucket := "buc3"
	client := client.CreateClient()
	_, err := client.DeleteBucket(context.TODO(), &s3.DeleteBucketInput{
		Bucket: oss.String(bucket),
	})
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("删除成功")
	}
}
