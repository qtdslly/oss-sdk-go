package defaults

import (
	"oss-sdk-go/oss"
	"strconv"
	"testing"
)

func TestDetermineDefaultsMode(t *testing.T) {
	cases := []struct {
		Region      string
		GOOS        string
		Environment oss.RuntimeEnvironment
		Expected    oss.DefaultsMode
	}{
		{
			Region: "us-east-1",
			GOOS:   "ios",
			Environment: oss.RuntimeEnvironment{
				EnvironmentIdentifier: oss.ExecutionEnvironmentID("AWS_Lambda_java8"),
				Region:                "us-east-1",
			},
			Expected: oss.DefaultsModeMobile,
		},
		{
			Region: "us-east-1",
			GOOS:   "android",
			Environment: oss.RuntimeEnvironment{
				EnvironmentIdentifier: oss.ExecutionEnvironmentID("AWS_Lambda_java8"),
				Region:                "us-east-1",
			},
			Expected: oss.DefaultsModeMobile,
		},
		{
			Region: "us-east-1",
			Environment: oss.RuntimeEnvironment{
				EnvironmentIdentifier: oss.ExecutionEnvironmentID("AWS_Lambda_java8"),
				Region:                "us-east-1",
			},
			Expected: oss.DefaultsModeInRegion,
		},
		{
			Region: "us-east-1",
			Environment: oss.RuntimeEnvironment{
				EnvironmentIdentifier: oss.ExecutionEnvironmentID("AWS_Lambda_java8"),
				Region:                "us-west-2",
			},
			Expected: oss.DefaultsModeCrossRegion,
		},
		{
			Region: "us-east-1",
			Environment: oss.RuntimeEnvironment{
				Region:                    "us-east-1",
				EC2InstanceMetadataRegion: "us-east-1",
			},
			Expected: oss.DefaultsModeInRegion,
		},
		{
			Region: "us-east-1",
			Environment: oss.RuntimeEnvironment{
				EC2InstanceMetadataRegion: "us-west-2",
			},
			Expected: oss.DefaultsModeCrossRegion,
		},
		{
			Region:      "us-west-2",
			Environment: oss.RuntimeEnvironment{},
			Expected:    oss.DefaultsModeStandard,
		},
	}

	for i, tt := range cases {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			if len(tt.GOOS) > 0 {
				orig := getGOOS
				getGOOS = func() string {
					return tt.GOOS
				}
				defer func() {
					getGOOS = orig
				}()
			}
			got := ResolveDefaultsModeAuto(tt.Region, tt.Environment)
			if got != tt.Expected {
				t.Errorf("expect %v, got %v", tt.Expected, got)
			}
		})
	}
}
