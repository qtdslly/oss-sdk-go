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
	_, err := client.CreateBucket(context.TODO(), &s3.CreateBucketInput{
		Bucket: oss.String(bucket),
	})
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("创建成功")
	}
}
