module oss-sdk-go/service/neptune

go 1.15

require (
	oss-sdk-go v1.16.16
	oss-sdk-go/internal/configsources v1.1.23
	oss-sdk-go/internal/endpoints/v2 v2.4.17
	oss-sdk-go/service/internal/presigned-url v1.9.17
	github.com/aws/smithy-go v1.13.3
	github.com/jmespath/go-jmespath v0.4.0
)

replace oss-sdk-go => ../../

replace oss-sdk-go/internal/configsources => ../../internal/configsources/

replace oss-sdk-go/internal/endpoints/v2 => ../../internal/endpoints/v2/

replace oss-sdk-go/service/internal/presigned-url => ../../service/internal/presigned-url/
