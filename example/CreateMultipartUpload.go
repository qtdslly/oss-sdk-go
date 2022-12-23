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
	key := "a.exe"
	client := client.CreateClient()
	result, err := client.CreateMultipartUpload(context.TODO(), &s3.CreateMultipartUploadInput{
		Bucket:   oss.String(bucket),
		Key:      oss.String(key),
	})
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("uploadId:" + *result.UploadId)
	}
}
