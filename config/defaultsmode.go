package config

import (
	"context"
	"os"

	"oss-sdk-go/oss"
	"oss-sdk-go/feature/ec2/imds"
)

const execEnvVar = "AWS_EXECUTION_ENV"

// DefaultsModeOptions is the set of options that are used to configure
type DefaultsModeOptions struct {
	// The SDK configuration defaults mode. Defaults to legacy if not specified.
	//
	// Supported modes are: auto, cross-region, in-region, legacy, mobile, standard
	Mode oss.DefaultsMode

	// The EC2 Instance Metadata Client that should be used when performing environment
	// discovery when oss.DefaultsModeAuto is set.
	//
	// If not specified the SDK will construct a client if the instance metadata service has not been disabled by
	// the AWS_EC2_METADATA_DISABLED environment variable.
	IMDSClient *imds.Client
}

func resolveDefaultsModeRuntimeEnvironment(ctx context.Context, envConfig *EnvConfig, client *imds.Client) (oss.RuntimeEnvironment, error) {
	getRegionOutput, err := client.GetRegion(ctx, &imds.GetRegionInput{})
	// honor context timeouts, but if we couldn't talk to IMDS don't fail runtime environment introspection.
	select {
	case <-ctx.Done():
		return oss.RuntimeEnvironment{}, err
	default:
	}

	var imdsRegion string
	if err == nil {
		imdsRegion = getRegionOutput.Region
	}

	return oss.RuntimeEnvironment{
		EnvironmentIdentifier:     oss.ExecutionEnvironmentID(os.Getenv(execEnvVar)),
		Region:                    envConfig.Region,
		EC2InstanceMetadataRegion: imdsRegion,
	}, nil
}
