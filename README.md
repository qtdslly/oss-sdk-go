# 适用于 Go 的云存储开发工具包

`oss-sdk-go` 是适用于 Go 编程语言的开发工具包。

SDK 需要 `Go 1.15` 的最低版本。

## 入门
要开始使用 SDK，请为 Go 模块设置您的项目，并使用 `go get` 检索 SDK 依赖项。
此示例说明如何使用开发工具包发出 API 请求。

###### 初始化项目
```
$ mkdir ~/hellooss
$ cd ~/hellooss
$ go mod init hellooss
```
###### 添加SDK依赖
```
$ go get oss-sdk-go/oss
$ go get oss-sdk-go/config
$ go get oss-sdk-go/service/dynamodb
```

###### 编写代码
在您喜欢的编辑器中，将以下内容添加到 `main.go`
```go
package main

import (
	"context"
	"oss-sdk-go/config"
	"oss-sdk-go/credentials"
	"oss-sdk-go/oss"
	"oss-sdk-go/service/s3"
)

func main() {
	AccessKeyId := ""
	SecretAccessKey := ""
	Endpoint := ""
	creds := credentials.NewStaticCredentialsProvider(AccessKeyId,
		SecretAccessKey, "")
	customResolver := oss.EndpointResolverWithOptionsFunc(func(service, region string,
		options ...interface{}) (oss.Endpoint, error) {
		return oss.Endpoint{
			URL: Endpoint,
		}, nil
	})
	cfg, err := config.LoadDefaultConfig(context.TODO(),
		config.WithCredentialsProvider(creds),
		config.WithEndpointResolverWithOptions(customResolver))
	if err != nil {
		println(err.Error())
	}
	return
	client := s3.NewFromConfig(cfg, func(o *s3.Options) {
		o.UsePathStyle = true
	})
	_, err = client.CreateBucket(context.TODO(), &s3.CreateBucketInput{
		Bucket: oss.String("bucket1"),
	})
	if err != nil {
		println(err.Error())
	} else {
		println("创建成功")
	}
}

```

###### 编译并执行
```sh
$ go run .
创建成功
```
