module oss-sdk-go/internal/protocoltest/query

go 1.15

require (
	oss-sdk-go v1.16.16
	oss-sdk-go/internal/configsources v1.1.23
	oss-sdk-go/internal/endpoints/v2 v2.4.17
	github.com/aws/smithy-go v1.13.3
	github.com/google/go-cmp v0.5.8
)

replace oss-sdk-go => ../../../

replace oss-sdk-go/internal/configsources => ../../../internal/configsources/

replace oss-sdk-go/internal/endpoints/v2 => ../../../internal/endpoints/v2/
