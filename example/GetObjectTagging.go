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
	result, err := client.GetObjectTagging(context.TODO(), &s3.GetObjectTaggingInput{
		Bucket: oss.String(bucket),
		Key:    oss.String(key),
	})
	if err != nil {
		fmt.Println(err.Error())
	} else {
		for _, v := range result.TagSet {
			fmt.Println("key:" + *v.Key + ", value:" + *v.Value)
		}
	}
}
