module oss-sdk-go/service/s3

go 1.15

require (
	oss-sdk-go v1.16.16
	oss-sdk-go/oss/protocol/eventstream v1.4.8
	oss-sdk-go/internal/configsources v1.1.23
	oss-sdk-go/internal/endpoints/v2 v2.4.17
	oss-sdk-go/internal/v4a v1.0.14
	oss-sdk-go/service/internal/accept-encoding v1.9.9
	oss-sdk-go/service/internal/checksum v1.1.18
	oss-sdk-go/service/internal/presigned-url v1.9.17
	oss-sdk-go/service/internal/s3shared v1.13.17
	github.com/aws/smithy-go v1.13.3
	github.com/google/go-cmp v0.5.8
)

replace oss-sdk-go => ../../

replace oss-sdk-go/oss/protocol/eventstream => ../../oss/protocol/eventstream/

replace oss-sdk-go/internal/configsources => ../../internal/configsources/

replace oss-sdk-go/internal/endpoints/v2 => ../../internal/endpoints/v2/

replace oss-sdk-go/internal/v4a => ../../internal/v4a/

replace oss-sdk-go/service/internal/accept-encoding => ../../service/internal/accept-encoding/

replace oss-sdk-go/service/internal/checksum => ../../service/internal/checksum/

replace oss-sdk-go/service/internal/presigned-url => ../../service/internal/presigned-url/

replace oss-sdk-go/service/internal/s3shared => ../../service/internal/s3shared/
