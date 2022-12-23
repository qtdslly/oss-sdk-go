package client

import (
	"context"
	"oss-sdk-go/config"
	"oss-sdk-go/credentials"
	"oss-sdk-go/oss"
	"oss-sdk-go/service/s3"
)

func CreateClient() *s3.Client {
	accessKeyId := "jvv4i43jx5jjszcjwxn4xvlf72kq"
	secretAccessKey := "j27ksdw63rbqmalfge23ss4aqg5tbfmkz7qeanrbguvyysp7xokh2"
	endpoint := "https://gateway.99265.net"
	creds := credentials.NewStaticCredentialsProvider(accessKeyId, secretAccessKey, "")
	customResolver := oss.EndpointResolverWithOptionsFunc(func(service, region string, options ...interface{}) (oss.Endpoint, error) {
		return oss.Endpoint{
			URL: endpoint,
		}, nil
	})
	cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithCredentialsProvider(creds), config.WithEndpointResolverWithOptions(customResolver))
	if err != nil {
		println("failed to load SDK configuration, %v", err)
	}
	return s3.NewFromConfig(cfg, func(o *s3.Options) {
		o.UsePathStyle = true
	})
}
