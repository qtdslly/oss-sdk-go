module oss-sdk-go/service/inspector

go 1.15

require (
	oss-sdk-go v1.16.16
	oss-sdk-go/internal/configsources v1.1.23
	oss-sdk-go/internal/endpoints/v2 v2.4.17
	github.com/aws/smithy-go v1.13.3
)

replace oss-sdk-go => ../../

replace oss-sdk-go/internal/configsources => ../../internal/configsources/

replace oss-sdk-go/internal/endpoints/v2 => ../../internal/endpoints/v2/
