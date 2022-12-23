module oss-sdk-go/credentials

go 1.15

require (
	oss-sdk-go v1.16.16
	oss-sdk-go/feature/ec2/imds v1.12.17
	oss-sdk-go/service/sso v1.11.23
	oss-sdk-go/service/ssooidc v1.13.5
	oss-sdk-go/service/sts v1.16.19
	github.com/aws/smithy-go v1.13.3
	github.com/google/go-cmp v0.5.8
)

replace oss-sdk-go => ../

replace oss-sdk-go/feature/ec2/imds => ../feature/ec2/imds/

replace oss-sdk-go/internal/configsources => ../internal/configsources/

replace oss-sdk-go/internal/endpoints/v2 => ../internal/endpoints/v2/

replace oss-sdk-go/service/internal/presigned-url => ../service/internal/presigned-url/

replace oss-sdk-go/service/sso => ../service/sso/

replace oss-sdk-go/service/ssooidc => ../service/ssooidc/

replace oss-sdk-go/service/sts => ../service/sts/
