package mock

import (
	"context"

	"oss-sdk-go/oss"
)

// LoadDefaultConfig is a mock for config.LoadDefaultConfig
func LoadDefaultConfig(context.Context, ...func()) (cfg oss.Config, err error) {
	return cfg, err
}

// WithRetryer is a mock for config.WithRetryer
func WithRetryer(v func() oss.Retryer) (f func()) {
	return f
}
