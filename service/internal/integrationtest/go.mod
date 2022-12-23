module oss-sdk-go/service/internal/integrationtest

require (
	oss-sdk-go v1.16.16
	oss-sdk-go/config v1.17.7
	oss-sdk-go/service/acm v1.14.18
	oss-sdk-go/service/apigateway v1.15.20
	oss-sdk-go/service/applicationautoscaling v1.15.18
	oss-sdk-go/service/applicationdiscoveryservice v1.14.9
	oss-sdk-go/service/appstream v1.17.11
	oss-sdk-go/service/athena v1.18.10
	oss-sdk-go/service/autoscaling v1.23.16
	oss-sdk-go/service/batch v1.18.16
	oss-sdk-go/service/cloudformation v1.22.10
	oss-sdk-go/service/cloudfront v1.20.5
	oss-sdk-go/service/cloudhsmv2 v1.13.18
	oss-sdk-go/service/cloudsearch v1.13.17
	oss-sdk-go/service/cloudtrail v1.18.1
	oss-sdk-go/service/cloudwatch v1.21.6
	oss-sdk-go/service/codebuild v1.19.17
	oss-sdk-go/service/codecommit v1.13.17
	oss-sdk-go/service/codedeploy v1.14.18
	oss-sdk-go/service/codepipeline v1.13.17
	oss-sdk-go/service/codestar v1.12.1
	oss-sdk-go/service/cognitoidentityprovider v1.20.1
	oss-sdk-go/service/configservice v1.26.1
	oss-sdk-go/service/costandusagereportservice v1.13.17
	oss-sdk-go/service/databasemigrationservice v1.21.12
	oss-sdk-go/service/devicefarm v1.13.17
	oss-sdk-go/service/directconnect v1.17.18
	oss-sdk-go/service/directoryservice v1.14.11
	oss-sdk-go/service/docdb v1.19.11
	oss-sdk-go/service/dynamodb v1.17.1
	oss-sdk-go/service/ec2 v1.58.0
	oss-sdk-go/service/ecr v1.17.18
	oss-sdk-go/service/ecs v1.18.22
	oss-sdk-go/service/efs v1.17.15
	oss-sdk-go/service/elasticache v1.22.10
	oss-sdk-go/service/elasticbeanstalk v1.14.18
	oss-sdk-go/service/elasticloadbalancing v1.14.18
	oss-sdk-go/service/elasticloadbalancingv2 v1.18.19
	oss-sdk-go/service/elasticsearchservice v1.16.10
	oss-sdk-go/service/elastictranscoder v1.13.17
	oss-sdk-go/service/emr v1.20.11
	oss-sdk-go/service/eventbridge v1.16.15
	oss-sdk-go/service/firehose v1.14.19
	oss-sdk-go/service/gamelift v1.15.5
	oss-sdk-go/service/glacier v1.13.17
	oss-sdk-go/service/glue v1.31.1
	oss-sdk-go/service/health v1.15.20
	oss-sdk-go/service/iam v1.18.19
	oss-sdk-go/service/inspector v1.12.17
	oss-sdk-go/service/iot v1.29.1
	oss-sdk-go/service/kinesis v1.15.19
	oss-sdk-go/service/kms v1.18.11
	oss-sdk-go/service/lambda v1.24.6
	oss-sdk-go/service/lightsail v1.22.12
	oss-sdk-go/service/marketplacecommerceanalytics v1.11.17
	oss-sdk-go/service/neptune v1.17.12
	oss-sdk-go/service/opsworks v1.13.17
	oss-sdk-go/service/pinpointemail v1.11.18
	oss-sdk-go/service/polly v1.17.9
	oss-sdk-go/service/rds v1.26.1
	oss-sdk-go/service/redshift v1.26.10
	oss-sdk-go/service/rekognition v1.20.5
	oss-sdk-go/service/route53 v1.22.1
	oss-sdk-go/service/route53domains v1.12.17
	oss-sdk-go/service/route53resolver v1.15.18
	oss-sdk-go/service/s3 v1.27.11
	oss-sdk-go/service/s3control v1.22.1
	oss-sdk-go/service/secretsmanager v1.16.1
	oss-sdk-go/service/servicecatalog v1.14.17
	oss-sdk-go/service/ses v1.14.18
	oss-sdk-go/service/sfn v1.14.1
	oss-sdk-go/service/shield v1.17.9
	oss-sdk-go/service/sms v1.12.18
	oss-sdk-go/service/snowball v1.15.18
	oss-sdk-go/service/sns v1.18.1
	oss-sdk-go/service/sqs v1.19.10
	oss-sdk-go/service/ssm v1.28.1
	oss-sdk-go/service/sts v1.16.19
	oss-sdk-go/service/support v1.13.17
	oss-sdk-go/service/timestreamwrite v1.14.1
	oss-sdk-go/service/transcribestreaming v1.6.19
	oss-sdk-go/service/waf v1.11.17
	oss-sdk-go/service/wafregional v1.12.18
	oss-sdk-go/service/wafv2 v1.22.8
	oss-sdk-go/service/workspaces v1.22.9
	github.com/aws/smithy-go v1.13.3
	github.com/google/go-cmp v0.5.8
)

go 1.15

replace oss-sdk-go => ../../../

replace oss-sdk-go/oss/protocol/eventstream => ../../../oss/protocol/eventstream/

replace oss-sdk-go/config => ../../../config/

replace oss-sdk-go/credentials => ../../../credentials/

replace oss-sdk-go/feature/ec2/imds => ../../../feature/ec2/imds/

replace oss-sdk-go/internal/configsources => ../../../internal/configsources/

replace oss-sdk-go/internal/endpoints/v2 => ../../../internal/endpoints/v2/

replace oss-sdk-go/internal/ini => ../../../internal/ini/

replace oss-sdk-go/internal/v4a => ../../../internal/v4a/

replace oss-sdk-go/service/acm => ../../../service/acm/

replace oss-sdk-go/service/apigateway => ../../../service/apigateway/

replace oss-sdk-go/service/applicationautoscaling => ../../../service/applicationautoscaling/

replace oss-sdk-go/service/applicationdiscoveryservice => ../../../service/applicationdiscoveryservice/

replace oss-sdk-go/service/appstream => ../../../service/appstream/

replace oss-sdk-go/service/athena => ../../../service/athena/

replace oss-sdk-go/service/autoscaling => ../../../service/autoscaling/

replace oss-sdk-go/service/batch => ../../../service/batch/

replace oss-sdk-go/service/cloudformation => ../../../service/cloudformation/

replace oss-sdk-go/service/cloudfront => ../../../service/cloudfront/

replace oss-sdk-go/service/cloudhsmv2 => ../../../service/cloudhsmv2/

replace oss-sdk-go/service/cloudsearch => ../../../service/cloudsearch/

replace oss-sdk-go/service/cloudtrail => ../../../service/cloudtrail/

replace oss-sdk-go/service/cloudwatch => ../../../service/cloudwatch/

replace oss-sdk-go/service/codebuild => ../../../service/codebuild/

replace oss-sdk-go/service/codecommit => ../../../service/codecommit/

replace oss-sdk-go/service/codedeploy => ../../../service/codedeploy/

replace oss-sdk-go/service/codepipeline => ../../../service/codepipeline/

replace oss-sdk-go/service/codestar => ../../../service/codestar/

replace oss-sdk-go/service/cognitoidentityprovider => ../../../service/cognitoidentityprovider/

replace oss-sdk-go/service/configservice => ../../../service/configservice/

replace oss-sdk-go/service/costandusagereportservice => ../../../service/costandusagereportservice/

replace oss-sdk-go/service/databasemigrationservice => ../../../service/databasemigrationservice/

replace oss-sdk-go/service/devicefarm => ../../../service/devicefarm/

replace oss-sdk-go/service/directconnect => ../../../service/directconnect/

replace oss-sdk-go/service/directoryservice => ../../../service/directoryservice/

replace oss-sdk-go/service/docdb => ../../../service/docdb/

replace oss-sdk-go/service/dynamodb => ../../../service/dynamodb/

replace oss-sdk-go/service/ec2 => ../../../service/ec2/

replace oss-sdk-go/service/ecr => ../../../service/ecr/

replace oss-sdk-go/service/ecs => ../../../service/ecs/

replace oss-sdk-go/service/efs => ../../../service/efs/

replace oss-sdk-go/service/elasticache => ../../../service/elasticache/

replace oss-sdk-go/service/elasticbeanstalk => ../../../service/elasticbeanstalk/

replace oss-sdk-go/service/elasticloadbalancing => ../../../service/elasticloadbalancing/

replace oss-sdk-go/service/elasticloadbalancingv2 => ../../../service/elasticloadbalancingv2/

replace oss-sdk-go/service/elasticsearchservice => ../../../service/elasticsearchservice/

replace oss-sdk-go/service/elastictranscoder => ../../../service/elastictranscoder/

replace oss-sdk-go/service/emr => ../../../service/emr/

replace oss-sdk-go/service/eventbridge => ../../../service/eventbridge/

replace oss-sdk-go/service/firehose => ../../../service/firehose/

replace oss-sdk-go/service/gamelift => ../../../service/gamelift/

replace oss-sdk-go/service/glacier => ../../../service/glacier/

replace oss-sdk-go/service/glue => ../../../service/glue/

replace oss-sdk-go/service/health => ../../../service/health/

replace oss-sdk-go/service/iam => ../../../service/iam/

replace oss-sdk-go/service/inspector => ../../../service/inspector/

replace oss-sdk-go/service/internal/accept-encoding => ../../../service/internal/accept-encoding/

replace oss-sdk-go/service/internal/checksum => ../../../service/internal/checksum/

replace oss-sdk-go/service/internal/endpoint-discovery => ../../../service/internal/endpoint-discovery/

replace oss-sdk-go/service/internal/presigned-url => ../../../service/internal/presigned-url/

replace oss-sdk-go/service/internal/s3shared => ../../../service/internal/s3shared/

replace oss-sdk-go/service/iot => ../../../service/iot/

replace oss-sdk-go/service/kinesis => ../../../service/kinesis/

replace oss-sdk-go/service/kms => ../../../service/kms/

replace oss-sdk-go/service/lambda => ../../../service/lambda/

replace oss-sdk-go/service/lightsail => ../../../service/lightsail/

replace oss-sdk-go/service/marketplacecommerceanalytics => ../../../service/marketplacecommerceanalytics/

replace oss-sdk-go/service/neptune => ../../../service/neptune/

replace oss-sdk-go/service/opsworks => ../../../service/opsworks/

replace oss-sdk-go/service/pinpointemail => ../../../service/pinpointemail/

replace oss-sdk-go/service/polly => ../../../service/polly/

replace oss-sdk-go/service/rds => ../../../service/rds/

replace oss-sdk-go/service/redshift => ../../../service/redshift/

replace oss-sdk-go/service/rekognition => ../../../service/rekognition/

replace oss-sdk-go/service/route53 => ../../../service/route53/

replace oss-sdk-go/service/route53domains => ../../../service/route53domains/

replace oss-sdk-go/service/route53resolver => ../../../service/route53resolver/

replace oss-sdk-go/service/s3 => ../../../service/s3/

replace oss-sdk-go/service/s3control => ../../../service/s3control/

replace oss-sdk-go/service/secretsmanager => ../../../service/secretsmanager/

replace oss-sdk-go/service/servicecatalog => ../../../service/servicecatalog/

replace oss-sdk-go/service/ses => ../../../service/ses/

replace oss-sdk-go/service/sfn => ../../../service/sfn/

replace oss-sdk-go/service/shield => ../../../service/shield/

replace oss-sdk-go/service/sms => ../../../service/sms/

replace oss-sdk-go/service/snowball => ../../../service/snowball/

replace oss-sdk-go/service/sns => ../../../service/sns/

replace oss-sdk-go/service/sqs => ../../../service/sqs/

replace oss-sdk-go/service/ssm => ../../../service/ssm/

replace oss-sdk-go/service/sso => ../../../service/sso/

replace oss-sdk-go/service/ssooidc => ../../../service/ssooidc/

replace oss-sdk-go/service/sts => ../../../service/sts/

replace oss-sdk-go/service/support => ../../../service/support/

replace oss-sdk-go/service/timestreamwrite => ../../../service/timestreamwrite/

replace oss-sdk-go/service/transcribestreaming => ../../../service/transcribestreaming/

replace oss-sdk-go/service/waf => ../../../service/waf/

replace oss-sdk-go/service/wafregional => ../../../service/wafregional/

replace oss-sdk-go/service/wafv2 => ../../../service/wafv2/

replace oss-sdk-go/service/workspaces => ../../../service/workspaces/
