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
	client := client.CreateClient()
	result, err := client.ListMultipartUploads(context.TODO(), &s3.ListMultipartUploadsInput{
		Bucket: oss.String(bucket),
	})
	if err != nil {
		fmt.Println(err.Error())
	} else {
		for _, upload := range result.Uploads {
			fmt.Println("Key:" + *upload.Key)
			fmt.Println("UploadId:" + *upload.UploadId)
			fmt.Println("Initiated:" + upload.Initiated.String())
		}
	}
}
