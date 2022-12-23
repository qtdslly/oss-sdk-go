module oss-sdk-go/feature/s3/manager

go 1.15

require (
	oss-sdk-go v1.16.16
	oss-sdk-go/config v1.17.7
	oss-sdk-go/service/s3 v1.27.11
	github.com/aws/smithy-go v1.13.3
	github.com/google/go-cmp v0.5.8
)

replace oss-sdk-go => ../../../

replace oss-sdk-go/oss/protocol/eventstream => ../../../oss/protocol/eventstream/

replace oss-sdk-go/config => ../../../config/

replace oss-sdk-go/credentials => ../../../credentials/

replace oss-sdk-go/feature/ec2/imds => ../../../feature/ec2/imds/

replace oss-sdk-go/internal/configsources => ../../../internal/configsources/

replace oss-sdk-go/internal/endpoints/v2 => ../../../internal/endpoints/v2/

replace oss-sdk-go/internal/ini => ../../../internal/ini/

replace oss-sdk-go/internal/v4a => ../../../internal/v4a/

replace oss-sdk-go/service/internal/accept-encoding => ../../../service/internal/accept-encoding/

replace oss-sdk-go/service/internal/checksum => ../../../service/internal/checksum/

replace oss-sdk-go/service/internal/presigned-url => ../../../service/internal/presigned-url/

replace oss-sdk-go/service/internal/s3shared => ../../../service/internal/s3shared/

replace oss-sdk-go/service/s3 => ../../../service/s3/

replace oss-sdk-go/service/sso => ../../../service/sso/

replace oss-sdk-go/service/ssooidc => ../../../service/ssooidc/

replace oss-sdk-go/service/sts => ../../../service/sts/
