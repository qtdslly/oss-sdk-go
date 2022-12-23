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
	var objects []types.ObjectIdentifier
	bucket := "buc1"
	key1 := "a/a.txt"
	key2 := "b.txt"
	objects = append(objects, types.ObjectIdentifier{Key: oss.String(key1)}, types.ObjectIdentifier{Key: oss.String(key2)})
	client := client.CreateClient()
	_, err := client.DeleteObjects(context.TODO(), &s3.DeleteObjectsInput{
		Bucket: oss.String(bucket),
		Delete: &types.Delete{
			Objects: objects,
		},
	})
	if err != nil {
		fmt.Println(err.Error())
	}
}
