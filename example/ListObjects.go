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
	result, err := client.ListObjects(context.TODO(), &s3.ListObjectsInput{
		Bucket: oss.String(bucket),
	})
	if err != nil {
		fmt.Println(err.Error())
	} else {
		for _, pre := range result.CommonPrefixes {
			fmt.Println("CommonPrefixes:" + *pre.Prefix)
		}
		for _, object := range result.Contents {
			fmt.Println("Contents:" + *object.Key)
		}
	}
}
