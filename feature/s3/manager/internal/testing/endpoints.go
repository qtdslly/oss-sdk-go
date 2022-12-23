package testing

import (
	"oss-sdk-go/oss"
	"oss-sdk-go/service/s3"
)

// EndpointResolverFunc is a mock s3 endpoint resolver that wraps the given function
type EndpointResolverFunc func(region string, options s3.EndpointResolverOptions) (oss.Endpoint, error)

// ResolveEndpoint returns the results from the wrapped function.
func (m EndpointResolverFunc) ResolveEndpoint(region string, options s3.EndpointResolverOptions) (oss.Endpoint, error) {
	return m(region, options)
}
