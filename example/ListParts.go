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
	key := "a.exe"
	uploadId := "M6cWbCETGRHbSdUyeiurEerNPGLXMv43rEUnYYqtUeCaT9m6VZALWbLQVXA5KR49GCrmsYPuxUXTdqJZ6QbFN5mMaJWTzEu8ugGRVFafNJvZ38cgBa5DNxCsjfaBtU8e2nAhP9Z6JTzqf52e1Eto6AtMF2yL8X9t726QhzttgyEL6VEuP692Cpf2z3"
	client := client.CreateClient()
	result, err := client.ListParts(context.TODO(), &s3.ListPartsInput{
		Bucket: oss.String(bucket),
		Key:    oss.String(key),
		UploadId: oss.String(uploadId),
	})
	if err != nil {
		fmt.Println(err.Error())
	} else {
		for _, object := range result.Parts {
			fmt.Println("PartNumber:" + strconv.Itoa(int(object.PartNumber)))
			fmt.Println("ETag:" + *object.ETag)
		}
	}
}
