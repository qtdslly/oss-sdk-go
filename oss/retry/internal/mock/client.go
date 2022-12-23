package mock

import (
	"context"

	"oss-sdk-go/oss"
)

// Options is a mock client Options
type Options struct {
	Retryer oss.Retryer
}

// Client is a mock service client
type Client struct{}

// GetObjectInput is mock input
type GetObjectInput struct {
	Bucket *string
	Key    *string
}

// GetObjectOutput is mock output
type GetObjectOutput struct{}

// NewFromConfig is a mock client constructor
func NewFromConfig(oss.Config, ...func(options *Options)) Client {
	return Client{}
}

// GetObject is a mock GetObject API
func (Client) GetObject(context.Context, *GetObjectInput, ...func(*Options)) (o *GetObjectOutput, err error) {
	return o, err
}
