package defaults

import (
	"oss-sdk-go/oss"
	"runtime"
	"strings"
)

var getGOOS = func() string {
	return runtime.GOOS
}

// ResolveDefaultsModeAuto is used to determine the effective oss.DefaultsMode when the mode
// is set to oss.DefaultsModeAuto.
func ResolveDefaultsModeAuto(region string, environment oss.RuntimeEnvironment) oss.DefaultsMode {
	goos := getGOOS()
	if goos == "android" || goos == "ios" {
		return oss.DefaultsModeMobile
	}

	var currentRegion string
	if len(environment.EnvironmentIdentifier) > 0 {
		currentRegion = environment.Region
	}

	if len(currentRegion) == 0 && len(environment.EC2InstanceMetadataRegion) > 0 {
		currentRegion = environment.EC2InstanceMetadataRegion
	}

	if len(region) > 0 && len(currentRegion) > 0 {
		if strings.EqualFold(region, currentRegion) {
			return oss.DefaultsModeInRegion
		}
		return oss.DefaultsModeCrossRegion
	}

	return oss.DefaultsModeStandard
}
