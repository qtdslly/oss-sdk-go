module oss-sdk-go/feature/dynamodbstreams/attributevalue

go 1.15

require (
	oss-sdk-go v1.16.16
	oss-sdk-go/service/dynamodb v1.17.1
	oss-sdk-go/service/dynamodbstreams v1.13.20
	github.com/aws/smithy-go v1.13.3
	github.com/google/go-cmp v0.5.8
)

replace oss-sdk-go => ../../../

replace oss-sdk-go/internal/configsources => ../../../internal/configsources/

replace oss-sdk-go/internal/endpoints/v2 => ../../../internal/endpoints/v2/

replace oss-sdk-go/service/dynamodb => ../../../service/dynamodb/

replace oss-sdk-go/service/dynamodbstreams => ../../../service/dynamodbstreams/

replace oss-sdk-go/service/internal/accept-encoding => ../../../service/internal/accept-encoding/

replace oss-sdk-go/service/internal/endpoint-discovery => ../../../service/internal/endpoint-discovery/
