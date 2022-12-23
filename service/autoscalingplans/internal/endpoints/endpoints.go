// Code generated by smithy-go-codegen DO NOT EDIT.

package endpoints

import (
	"oss-sdk-go/oss"
	endpoints "oss-sdk-go/internal/endpoints/v2"
	"github.com/aws/smithy-go/logging"
	"regexp"
)

// Options is the endpoint resolver configuration options
type Options struct {
	// Logger is a logging implementation that log events should be sent to.
	Logger logging.Logger

	// LogDeprecated indicates that deprecated endpoints should be logged to the
	// provided logger.
	LogDeprecated bool

	// ResolvedRegion is used to override the region to be resolved, rather then the
	// using the value passed to the ResolveEndpoint method. This value is used by the
	// SDK to translate regions like fips-us-east-1 or us-east-1-fips to an alternative
	// name. You must not set this value directly in your application.
	ResolvedRegion string

	// DisableHTTPS informs the resolver to return an endpoint that does not use the
	// HTTPS scheme.
	DisableHTTPS bool

	// UseDualStackEndpoint specifies the resolver must resolve a dual-stack endpoint.
	UseDualStackEndpoint oss.DualStackEndpointState

	// UseFIPSEndpoint specifies the resolver must resolve a FIPS endpoint.
	UseFIPSEndpoint oss.FIPSEndpointState
}

func (o Options) GetResolvedRegion() string {
	return o.ResolvedRegion
}

func (o Options) GetDisableHTTPS() bool {
	return o.DisableHTTPS
}

func (o Options) GetUseDualStackEndpoint() oss.DualStackEndpointState {
	return o.UseDualStackEndpoint
}

func (o Options) GetUseFIPSEndpoint() oss.FIPSEndpointState {
	return o.UseFIPSEndpoint
}

func transformToSharedOptions(options Options) endpoints.Options {
	return endpoints.Options{
		Logger:               options.Logger,
		LogDeprecated:        options.LogDeprecated,
		ResolvedRegion:       options.ResolvedRegion,
		DisableHTTPS:         options.DisableHTTPS,
		UseDualStackEndpoint: options.UseDualStackEndpoint,
		UseFIPSEndpoint:      options.UseFIPSEndpoint,
	}
}

// Resolver Auto Scaling Plans endpoint resolver
type Resolver struct {
	partitions endpoints.Partitions
}

// ResolveEndpoint resolves the service endpoint for the given region and options
func (r *Resolver) ResolveEndpoint(region string, options Options) (endpoint oss.Endpoint, err error) {
	if len(region) == 0 {
		return endpoint, &oss.MissingRegionError{}
	}

	opt := transformToSharedOptions(options)
	return r.partitions.ResolveEndpoint(region, opt)
}

// New returns a new Resolver
func New() *Resolver {
	return &Resolver{
		partitions: defaultPartitions,
	}
}

var partitionRegexp = struct {
	Aws      *regexp.Regexp
	AwsCn    *regexp.Regexp
	AwsIso   *regexp.Regexp
	AwsIsoB  *regexp.Regexp
	AwsUsGov *regexp.Regexp
}{

	Aws:      regexp.MustCompile("^(us|eu|ap|sa|ca|me|af)\\-\\w+\\-\\d+$"),
	AwsCn:    regexp.MustCompile("^cn\\-\\w+\\-\\d+$"),
	AwsIso:   regexp.MustCompile("^us\\-iso\\-\\w+\\-\\d+$"),
	AwsIsoB:  regexp.MustCompile("^us\\-isob\\-\\w+\\-\\d+$"),
	AwsUsGov: regexp.MustCompile("^us\\-gov\\-\\w+\\-\\d+$"),
}

var defaultPartitions = endpoints.Partitions{
	{
		ID: "aws",
		Defaults: map[endpoints.DefaultKey]endpoints.Endpoint{
			{
				Variant: endpoints.DualStackVariant,
			}: {
				Hostname:          "autoscaling-plans.{region}.api.aws",
				Protocols:         []string{"http", "https"},
				SignatureVersions: []string{"v4"},
			},
			{
				Variant: endpoints.FIPSVariant,
			}: {
				Hostname:          "autoscaling-plans-fips.{region}.amazonaws.com",
				Protocols:         []string{"http", "https"},
				SignatureVersions: []string{"v4"},
			},
			{
				Variant: endpoints.FIPSVariant | endpoints.DualStackVariant,
			}: {
				Hostname:          "autoscaling-plans-fips.{region}.api.aws",
				Protocols:         []string{"http", "https"},
				SignatureVersions: []string{"v4"},
			},
			{
				Variant: 0,
			}: {
				Hostname:          "autoscaling-plans.{region}.amazonaws.com",
				Protocols:         []string{"http", "https"},
				SignatureVersions: []string{"v4"},
			},
		},
		RegionRegex:    partitionRegexp.Aws,
		IsRegionalized: true,
		Endpoints: endpoints.Endpoints{
			endpoints.EndpointKey{
				Region: "af-south-1",
			}: endpoints.Endpoint{},
			endpoints.EndpointKey{
				Region: "ap-east-1",
			}: endpoints.Endpoint{},
			endpoints.EndpointKey{
				Region: "ap-northeast-1",
			}: endpoints.Endpoint{},
			endpoints.EndpointKey{
				Region: "ap-northeast-2",
			}: endpoints.Endpoint{},
			endpoints.EndpointKey{
				Region: "ap-northeast-3",
			}: endpoints.Endpoint{},
			endpoints.EndpointKey{
				Region: "ap-south-1",
			}: endpoints.Endpoint{},
			endpoints.EndpointKey{
				Region: "ap-southeast-1",
			}: endpoints.Endpoint{},
			endpoints.EndpointKey{
				Region: "ap-southeast-2",
			}: endpoints.Endpoint{},
			endpoints.EndpointKey{
				Region: "ap-southeast-3",
			}: endpoints.Endpoint{},
			endpoints.EndpointKey{
				Region: "ca-central-1",
			}: endpoints.Endpoint{},
			endpoints.EndpointKey{
				Region: "eu-central-1",
			}: endpoints.Endpoint{},
			endpoints.EndpointKey{
				Region: "eu-north-1",
			}: endpoints.Endpoint{},
			endpoints.EndpointKey{
				Region: "eu-south-1",
			}: endpoints.Endpoint{},
			endpoints.EndpointKey{
				Region: "eu-west-1",
			}: endpoints.Endpoint{},
			endpoints.EndpointKey{
				Region: "eu-west-2",
			}: endpoints.Endpoint{},
			endpoints.EndpointKey{
				Region: "eu-west-3",
			}: endpoints.Endpoint{},
			endpoints.EndpointKey{
				Region: "me-south-1",
			}: endpoints.Endpoint{},
			endpoints.EndpointKey{
				Region: "sa-east-1",
			}: endpoints.Endpoint{},
			endpoints.EndpointKey{
				Region: "us-east-1",
			}: endpoints.Endpoint{},
			endpoints.EndpointKey{
				Region: "us-east-2",
			}: endpoints.Endpoint{},
			endpoints.EndpointKey{
				Region: "us-west-1",
			}: endpoints.Endpoint{},
			endpoints.EndpointKey{
				Region: "us-west-2",
			}: endpoints.Endpoint{},
		},
	},
	{
		ID: "aws-cn",
		Defaults: map[endpoints.DefaultKey]endpoints.Endpoint{
			{
				Variant: endpoints.DualStackVariant,
			}: {
				Hostname:          "autoscaling-plans.{region}.api.amazonwebservices.com.cn",
				Protocols:         []string{"http", "https"},
				SignatureVersions: []string{"v4"},
			},
			{
				Variant: endpoints.FIPSVariant,
			}: {
				Hostname:          "autoscaling-plans-fips.{region}.amazonaws.com.cn",
				Protocols:         []string{"http", "https"},
				SignatureVersions: []string{"v4"},
			},
			{
				Variant: endpoints.FIPSVariant | endpoints.DualStackVariant,
			}: {
				Hostname:          "autoscaling-plans-fips.{region}.api.amazonwebservices.com.cn",
				Protocols:         []string{"http", "https"},
				SignatureVersions: []string{"v4"},
			},
			{
				Variant: 0,
			}: {
				Hostname:          "autoscaling-plans.{region}.amazonaws.com.cn",
				Protocols:         []string{"http", "https"},
				SignatureVersions: []string{"v4"},
			},
		},
		RegionRegex:    partitionRegexp.AwsCn,
		IsRegionalized: true,
		Endpoints: endpoints.Endpoints{
			endpoints.EndpointKey{
				Region: "cn-north-1",
			}: endpoints.Endpoint{},
			endpoints.EndpointKey{
				Region: "cn-northwest-1",
			}: endpoints.Endpoint{},
		},
	},
	{
		ID: "aws-iso",
		Defaults: map[endpoints.DefaultKey]endpoints.Endpoint{
			{
				Variant: endpoints.FIPSVariant,
			}: {
				Hostname:          "autoscaling-plans-fips.{region}.c2s.ic.gov",
				Protocols:         []string{"https"},
				SignatureVersions: []string{"v4"},
			},
			{
				Variant: 0,
			}: {
				Hostname:          "autoscaling-plans.{region}.c2s.ic.gov",
				Protocols:         []string{"https"},
				SignatureVersions: []string{"v4"},
			},
		},
		RegionRegex:    partitionRegexp.AwsIso,
		IsRegionalized: true,
	},
	{
		ID: "aws-iso-b",
		Defaults: map[endpoints.DefaultKey]endpoints.Endpoint{
			{
				Variant: endpoints.FIPSVariant,
			}: {
				Hostname:          "autoscaling-plans-fips.{region}.sc2s.sgov.gov",
				Protocols:         []string{"https"},
				SignatureVersions: []string{"v4"},
			},
			{
				Variant: 0,
			}: {
				Hostname:          "autoscaling-plans.{region}.sc2s.sgov.gov",
				Protocols:         []string{"https"},
				SignatureVersions: []string{"v4"},
			},
		},
		RegionRegex:    partitionRegexp.AwsIsoB,
		IsRegionalized: true,
	},
	{
		ID: "aws-us-gov",
		Defaults: map[endpoints.DefaultKey]endpoints.Endpoint{
			{
				Variant: endpoints.DualStackVariant,
			}: {
				Hostname:          "autoscaling-plans.{region}.api.aws",
				Protocols:         []string{"http", "https"},
				SignatureVersions: []string{"v4"},
			},
			{
				Variant: endpoints.FIPSVariant,
			}: {
				Hostname:          "autoscaling-plans-fips.{region}.amazonaws.com",
				Protocols:         []string{"http", "https"},
				SignatureVersions: []string{"v4"},
			},
			{
				Variant: endpoints.FIPSVariant | endpoints.DualStackVariant,
			}: {
				Hostname:          "autoscaling-plans-fips.{region}.api.aws",
				Protocols:         []string{"http", "https"},
				SignatureVersions: []string{"v4"},
			},
			{
				Variant: 0,
			}: {
				Hostname:          "autoscaling-plans.{region}.amazonaws.com",
				Protocols:         []string{"http", "https"},
				SignatureVersions: []string{"v4"},
			},
		},
		RegionRegex:    partitionRegexp.AwsUsGov,
		IsRegionalized: true,
		Endpoints: endpoints.Endpoints{
			endpoints.EndpointKey{
				Region: "us-gov-east-1",
			}: endpoints.Endpoint{
				Protocols: []string{"http", "https"},
			},
			endpoints.EndpointKey{
				Region: "us-gov-west-1",
			}: endpoints.Endpoint{
				Protocols: []string{"http", "https"},
			},
		},
	},
}