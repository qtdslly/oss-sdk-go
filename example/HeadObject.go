package main

import (
	"context"
	"example/client"
	"fmt"
	"oss-sdk-go/oss"
	"oss-sdk-go/service/s3"
	"strconv"
)

func main() {
	bucket := "buc1"
	key := "a.txt"
	client := client.CreateClient()
	result, err := client.HeadObject(context.TODO(), &s3.HeadObjectInput{
		Bucket: oss.String(bucket),
		Key:    oss.String(key),
	})
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("ETag:" + *result.ETag)
		fmt.Println("ContentLength:" + strconv.Itoa(int(result.ContentLength)))
		fmt.Println("LastModified:" + result.LastModified.String())
	}
}
