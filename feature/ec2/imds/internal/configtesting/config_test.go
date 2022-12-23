package configtesting

import (
	"oss-sdk-go/config"
	imdsconfig "oss-sdk-go/feature/ec2/imds/internal/config"
)

var _ imdsconfig.EndpointModeResolver = (*config.LoadOptions)(nil)
var _ imdsconfig.EndpointModeResolver = (*config.EnvConfig)(nil)
var _ imdsconfig.EndpointModeResolver = (*config.SharedConfig)(nil)

var _ imdsconfig.EndpointResolver = (*config.LoadOptions)(nil)
var _ imdsconfig.EndpointResolver = (*config.EnvConfig)(nil)
var _ imdsconfig.EndpointResolver = (*config.SharedConfig)(nil)

var _ imdsconfig.ClientEnableStateResolver = (*config.LoadOptions)(nil)
var _ imdsconfig.ClientEnableStateResolver = (*config.EnvConfig)(nil)
