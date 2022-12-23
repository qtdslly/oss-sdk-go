module oss-sdk-go/service/s3/internal/configtesting

go 1.15

require (
	oss-sdk-go/config v1.17.7
	oss-sdk-go/service/internal/s3shared v1.13.17
)

replace oss-sdk-go => ../../../../

replace oss-sdk-go/config => ../../../../config/

replace oss-sdk-go/credentials => ../../../../credentials/

replace oss-sdk-go/feature/ec2/imds => ../../../../feature/ec2/imds/

replace oss-sdk-go/internal/configsources => ../../../../internal/configsources/

replace oss-sdk-go/internal/endpoints/v2 => ../../../../internal/endpoints/v2/

replace oss-sdk-go/internal/ini => ../../../../internal/ini/

replace oss-sdk-go/service/internal/presigned-url => ../../../../service/internal/presigned-url/

replace oss-sdk-go/service/internal/s3shared => ../../../../service/internal/s3shared/

replace oss-sdk-go/service/sso => ../../../../service/sso/

replace oss-sdk-go/service/ssooidc => ../../../../service/ssooidc/

replace oss-sdk-go/service/sts => ../../../../service/sts/
