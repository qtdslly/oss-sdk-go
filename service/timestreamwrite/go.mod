module oss-sdk-go/service/timestreamwrite

go 1.15

require (
	oss-sdk-go v1.16.16
	oss-sdk-go/internal/configsources v1.1.23
	oss-sdk-go/internal/endpoints/v2 v2.4.17
	oss-sdk-go/service/internal/endpoint-discovery v1.7.17
	github.com/aws/smithy-go v1.13.3
)

replace oss-sdk-go => ../../

replace oss-sdk-go/internal/configsources => ../../internal/configsources/

replace oss-sdk-go/internal/endpoints/v2 => ../../internal/endpoints/v2/

replace oss-sdk-go/service/internal/endpoint-discovery => ../../service/internal/endpoint-discovery/
