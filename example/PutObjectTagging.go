package main

import (
	"context"
	"example/client"
	"fmt"
	"oss-sdk-go/oss"
	"oss-sdk-go/service/s3"
	"oss-sdk-go/service/s3/types"
)

func main() {
	bucket := "buc1"
	key := "a.txt"
	var tagging types.Tagging
	tagging.TagSet = append(tagging.TagSet, types.Tag{
		Key:   oss.String("key1"),
		Value: oss.String("value1"),
	}, types.Tag{
		Key:   oss.String("key2"),
		Value: oss.String("value2"),
	})
	client := client.CreateClient()
	_, err := client.PutObjectTagging(context.TODO(), &s3.PutObjectTaggingInput{
		Bucket: oss.String(bucket),
		Key:    oss.String(key),
		Tagging: &tagging,
	})
	if err != nil {
		fmt.Println(err.Error())
	}
}
