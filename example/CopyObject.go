package main

import (
	"context"
	"example/client"
	"fmt"
	"oss-sdk-go/oss"
	"oss-sdk-go/service/s3"
)

func main() {
	sourceBucket := "buc1"
	sourceKey := "a/a.txt"
	destinationBucket := "buc2"
	destinationKey := "b.txt"
	client := client.CreateClient()
	_, err := client.CopyObject(context.TODO(), &s3.CopyObjectInput{
		CopySource: oss.String(sourceBucket + "/" + sourceKey),
		Bucket:     oss.String(destinationBucket),
		Key:        oss.String(destinationKey),
	})
	if err != nil {
		fmt.Println(err.Error())
	}
}
