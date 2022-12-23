module oss-sdk-go/service/kinesis/internal/testing

go 1.15

require (
	oss-sdk-go v1.16.16
	oss-sdk-go/oss/protocol/eventstream v1.4.8
	oss-sdk-go/service/internal/eventstreamtesting v1.0.33
	oss-sdk-go/service/kinesis v1.15.19
	github.com/aws/smithy-go v1.13.3
	github.com/google/go-cmp v0.5.8
)

replace oss-sdk-go => ../../../../

replace oss-sdk-go/oss/protocol/eventstream => ../../../../oss/protocol/eventstream/

replace oss-sdk-go/credentials => ../../../../credentials/

replace oss-sdk-go/feature/ec2/imds => ../../../../feature/ec2/imds/

replace oss-sdk-go/internal/configsources => ../../../../internal/configsources/

replace oss-sdk-go/internal/endpoints/v2 => ../../../../internal/endpoints/v2/

replace oss-sdk-go/service/internal/eventstreamtesting => ../../../../service/internal/eventstreamtesting/

replace oss-sdk-go/service/internal/presigned-url => ../../../../service/internal/presigned-url/

replace oss-sdk-go/service/kinesis => ../../../../service/kinesis/

replace oss-sdk-go/service/sso => ../../../../service/sso/

replace oss-sdk-go/service/ssooidc => ../../../../service/ssooidc/

replace oss-sdk-go/service/sts => ../../../../service/sts/
