module oss-sdk-go/service/internal/eventstreamtesting

go 1.15

require (
	golang.org/x/net v0.0.0-20220127200216-cd36cc0744dd
	oss-sdk-go v1.16.16
	oss-sdk-go/credentials v1.12.20
	oss-sdk-go/oss/protocol/eventstream v1.4.8
)

replace oss-sdk-go => ../../../

replace oss-sdk-go/oss/protocol/eventstream => ../../../oss/protocol/eventstream/

replace oss-sdk-go/credentials => ../../../credentials/

replace oss-sdk-go/feature/ec2/imds => ../../../feature/ec2/imds/

replace oss-sdk-go/internal/configsources => ../../../internal/configsources/

replace oss-sdk-go/internal/endpoints/v2 => ../../../internal/endpoints/v2/

replace oss-sdk-go/service/internal/presigned-url => ../../../service/internal/presigned-url/

replace oss-sdk-go/service/sso => ../../../service/sso/

replace oss-sdk-go/service/ssooidc => ../../../service/ssooidc/

replace oss-sdk-go/service/sts => ../../../service/sts/
