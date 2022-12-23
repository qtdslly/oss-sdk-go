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
	key := "a.exe"
	uploadId := "M6cWbCETGRHbSdUyeiurEerNPGLXMv43rEUnYYqtUeCaT9m6VZALWbLQVXA5KR49GCrmsYPuxUXTdqJZ6QbFN5mMaJWTzEu8ugGRVFafNJvZ38cgBa5DNxCsjfaBtU8e2nAhP9Z6JTzqf52e1Eto6AtMF2yL8X9t726QhzttgyEL6VEuP692Cpf2z3"
	partNumber1 := 1
	etag1 := "a84196d8618498781f7e6d464b928cb0"
	partNumber2 := 2
	etag2 := "93252f429fa8f02078282c9caaeb6ace"
	part1 := types.CompletedPart{
		PartNumber: int32(partNumber1),
		ETag:       oss.String(etag1),
	}
	part2 := types.CompletedPart{
		PartNumber: int32(partNumber2),
		ETag:       oss.String(etag2),
	}
	var parts []types.CompletedPart
	parts = append(parts, part1, part2)
	client := client.CreateClient()
	_, err := client.CompleteMultipartUpload(context.TODO(), &s3.CompleteMultipartUploadInput{
		Bucket:   oss.String(bucket),
		Key:      oss.String(key),
		UploadId: oss.String(uploadId),
		MultipartUpload: &types.CompletedMultipartUpload{
			Parts: parts,
		},
	})
	if err != nil {
		fmt.Println(err.Error())
	}
}
