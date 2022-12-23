package configsources

import (
	"context"
	"oss-sdk-go/oss"
)

// EnableEndpointDiscoveryProvider is an interface for retrieving external configuration value
// for Enable Endpoint Discovery
type EnableEndpointDiscoveryProvider interface {
	GetEnableEndpointDiscovery(ctx context.Context) (value oss.EndpointDiscoveryEnableState, found bool, err error)
}

// ResolveEnableEndpointDiscovery extracts the first instance of a EnableEndpointDiscoveryProvider from the config slice.
// Additionally returns a oss.EndpointDiscoveryEnableState to indicate if the value was found in provided configs,
// and error if one is encountered.
func ResolveEnableEndpointDiscovery(ctx context.Context, configs []interface{}) (value oss.EndpointDiscoveryEnableState, found bool, err error) {
	for _, cfg := range configs {
		if p, ok := cfg.(EnableEndpointDiscoveryProvider); ok {
			value, found, err = p.GetEnableEndpointDiscovery(ctx)
			if err != nil || found {
				break
			}
		}
	}
	return
}

// UseDualStackEndpointProvider is an interface for retrieving external configuration values for UseDualStackEndpoint
type UseDualStackEndpointProvider interface {
	GetUseDualStackEndpoint(context.Context) (value oss.DualStackEndpointState, found bool, err error)
}

// ResolveUseDualStackEndpoint extracts the first instance of a UseDualStackEndpoint from the config slice.
// Additionally returns a boolean to indicate if the value was found in provided configs, and error if one is encountered.
func ResolveUseDualStackEndpoint(ctx context.Context, configs []interface{}) (value oss.DualStackEndpointState, found bool, err error) {
	for _, cfg := range configs {
		if p, ok := cfg.(UseDualStackEndpointProvider); ok {
			value, found, err = p.GetUseDualStackEndpoint(ctx)
			if err != nil || found {
				break
			}
		}
	}
	return
}

// UseFIPSEndpointProvider is an interface for retrieving external configuration values for UseFIPSEndpoint
type UseFIPSEndpointProvider interface {
	GetUseFIPSEndpoint(context.Context) (value oss.FIPSEndpointState, found bool, err error)
}

// ResolveUseFIPSEndpoint extracts the first instance of a UseFIPSEndpointProvider from the config slice.
// Additionally, returns a boolean to indicate if the value was found in provided configs, and error if one is encountered.
func ResolveUseFIPSEndpoint(ctx context.Context, configs []interface{}) (value oss.FIPSEndpointState, found bool, err error) {
	for _, cfg := range configs {
		if p, ok := cfg.(UseFIPSEndpointProvider); ok {
			value, found, err = p.GetUseFIPSEndpoint(ctx)
			if err != nil || found {
				break
			}
		}
	}
	return
}
