module oss-sdk-go/service/internal/benchmark

go 1.15

require (
	github.com/aws/aws-sdk-go v1.44.28
	github.com/aws/smithy-go v1.13.3
	oss-sdk-go v1.16.16
	oss-sdk-go/service/dynamodb v1.17.1
	oss-sdk-go/service/lexruntimeservice v1.12.17
	oss-sdk-go/service/s3 v1.27.11
	oss-sdk-go/service/schemas v1.14.17
)

replace oss-sdk-go => ../../../

replace oss-sdk-go/oss/protocol/eventstream => ../../../oss/protocol/eventstream/

replace oss-sdk-go/internal/configsources => ../../../internal/configsources/

replace oss-sdk-go/internal/endpoints/v2 => ../../../internal/endpoints/v2/

replace oss-sdk-go/internal/v4a => ../../../internal/v4a/

replace oss-sdk-go/service/dynamodb => ../../../service/dynamodb/

replace oss-sdk-go/service/internal/accept-encoding => ../../../service/internal/accept-encoding/

replace oss-sdk-go/service/internal/checksum => ../../../service/internal/checksum/

replace oss-sdk-go/service/internal/endpoint-discovery => ../../../service/internal/endpoint-discovery/

replace oss-sdk-go/service/internal/presigned-url => ../../../service/internal/presigned-url/

replace oss-sdk-go/service/internal/s3shared => ../../../service/internal/s3shared/

replace oss-sdk-go/service/lexruntimeservice => ../../../service/lexruntimeservice/

replace oss-sdk-go/service/s3 => ../../../service/s3/

replace oss-sdk-go/service/schemas => ../../../service/schemas/
