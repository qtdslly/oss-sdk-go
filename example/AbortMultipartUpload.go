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
	uploadId := "sf"
	client := client.CreateClient()
	_, err := client.AbortMultipartUpload(context.TODO(), &s3.AbortMultipartUploadInput{
		Bucket:   oss.String(bucket),
		Key:      oss.String(key),
		UploadId: oss.String(uploadId),
	})
	if err != nil {
		fmt.Println(err.Error())
	}
}
