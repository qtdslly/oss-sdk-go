package main

import (
	"context"
	"example/client"
	"fmt"
	"os"
	"oss-sdk-go/oss"
	"oss-sdk-go/service/s3"
)

func main() {
	bucket := "buc1"
	key := "a.exe"
	uploadId := "M6cWbCETGRHbSdUyeiurEerNPGLXMv43rEUnYYqtUeCaT9m6VZALWbLQVXA5KR49GCrmsYPuxUXTdqJZ6QbFN5mMaJWTzEu8ugGRVFafNJvZ38cgBa5DNxCsjfaBtU8e2nAhP9Z6JTzqf52e1Eto6AtMF2yL8X9t726QhzttgyEL6VEuP692Cpf2z3"
	partNumber := 2
	filePath := "D:\\xab"
	file, _ := os.Open(filePath)
	defer file.Close()
	client := client.CreateClient()
	_, err := client.UploadPart(context.TODO(), &s3.UploadPartInput{
		Bucket:     oss.String(bucket),
		Key:        oss.String(key),
		UploadId:   oss.String(uploadId),
		PartNumber: int32(partNumber),
		Body:       file,
	})
	if err != nil {
		fmt.Println(err.Error())
	}
}
