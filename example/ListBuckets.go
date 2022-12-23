package main

import (
	"context"
	"example/client"
	"fmt"
	"oss-sdk-go/service/s3"
)

func main() {
	client := client.CreateClient()
	result, err := client.ListBuckets(context.TODO(), &s3.ListBucketsInput{})
	if err != nil {
		fmt.Println(err.Error())
	} else {
		for _, object := range result.Buckets {
			fmt.Println("Name:" + *object.Name + ", CreationDate:" + object.CreationDate.String())
		}
	}
}
