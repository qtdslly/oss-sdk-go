package oss

import (
	"context"
	"fmt"
	"time"

	"oss-sdk-go/internal/sdk"
)

// AnonymousCredentials provides a sentinel CredentialsProvider that should be
// used to instruct the SDK's signing middleware to not sign the request.
//
// Using `nil` credentials when configuring an API client will achieve the same
// result. The AnonymousCredentials type allows you to configure the SDK's
// external config loading to not attempt to source credentials from the shared
// config or environment.
//
// For example you can use this CredentialsProvider with an API client's
// Options to instruct the client not to sign a request for accessing public
// S3 bucket objects.
//
// The following example demonstrates using the AnonymousCredentials to prevent
// SDK's external config loading attempt to resolve credentials.
//
//	cfg, err := config.LoadDefaultConfig(context.TODO(),
//	     config.WithCredentialsProvider(oss.AnonymousCredentials{}),
//	)
//	if err != nil {
//	     log.Fatalf("failed to load config, %v", err)
//	}
//
//	client := s3.NewFromConfig(cfg)
//
// Alternatively you can leave the API client Option's `Credential` member to
// nil. If using the `NewFromConfig` constructor you'll need to explicitly set
// the `Credentials` member to nil, if the external config resolved a
// credential provider.
//
//	client := s3.New(s3.Options{
//	     // Credentials defaults to a nil value.
//	})
//
// This can also be configured for specific operations calls too.
//
//	cfg, err := config.LoadDefaultConfig(context.TODO())
//	if err != nil {
//	     log.Fatalf("failed to load config, %v", err)
//	}
//
//	client := s3.NewFromConfig(config)
//
//	result, err := client.GetObject(context.TODO(), s3.GetObject{
//	     Bucket: oss.String("example-bucket"),
//	     Key: oss.String("example-key"),
//	}, func(o *s3.Options) {
//	     o.Credentials = nil
//	     // Or
//	     o.Credentials = oss.AnonymousCredentials{}
//	})
type AnonymousCredentials struct{}

// Retrieve implements the CredentialsProvider interface, but will always
// return error, and cannot be used to sign a request. The AnonymousCredentials
// type is used as a sentinel type instructing the AWS request signing
// middleware to not sign a request.
func (AnonymousCredentials) Retrieve(context.Context) (Credentials, error) {
	return Credentials{Source: "AnonymousCredentials"},
		fmt.Errorf("the AnonymousCredentials is not a valid credential provider, and cannot be used to sign AWS requests with")
}

// A Credentials is the AWS credentials value for individual credential fields.
type Credentials struct {
	// AWS Access key ID
	AccessKeyID string

	// AWS Secret Access Key
	SecretAccessKey string

	// AWS Session Token
	SessionToken string

	// Source of the credentials
	Source string

	// States if the credentials can expire or not.
	CanExpire bool

	// The time the credentials will expire at. Should be ignored if CanExpire
	// is false.
	Expires time.Time
}

// Expired returns if the credentials have expired.
func (v Credentials) Expired() bool {
	if v.CanExpire {
		// Calling Round(0) on the current time will truncate the monotonic
		// reading only. Ensures credential expiry time is always based on
		// reported wall-clock time.
		return !v.Expires.After(sdk.NowTime().Round(0))
	}

	return false
}

// HasKeys returns if the credentials keys are set.
func (v Credentials) HasKeys() bool {
	return len(v.AccessKeyID) > 0 && len(v.SecretAccessKey) > 0
}

// A CredentialsProvider is the interface for any component which will provide
// credentials Credentials. A CredentialsProvider is required to manage its own
// Expired state, and what to be expired means.
//
// A credentials provider implementation can be wrapped with a CredentialCache
// to cache the credential value retrieved. Without the cache the SDK will
// attempt to retrieve the credentials for every request.
type CredentialsProvider interface {
	// Retrieve returns nil if it successfully retrieved the value.
	// Error is returned if the value were not obtainable, or empty.
	Retrieve(ctx context.Context) (Credentials, error)
}

// CredentialsProviderFunc provides a helper wrapping a function value to
// satisfy the CredentialsProvider interface.
type CredentialsProviderFunc func(context.Context) (Credentials, error)

// Retrieve delegates to the function value the CredentialsProviderFunc wraps.
func (fn CredentialsProviderFunc) Retrieve(ctx context.Context) (Credentials, error) {
	return fn(ctx)
}
