// Code generated by smithy-go-codegen DO NOT EDIT.

package route53recoverycluster

import (
	"context"
	"oss-sdk-go/oss"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"io/ioutil"
	"net/http"
	"strings"
	"testing"
)

func TestClient_resolveRetryOptions(t *testing.T) {
	nopClient := smithyhttp.ClientDoFunc(func(_ *http.Request) (*http.Response, error) {
		return &http.Response{
			StatusCode: 200,
			Header:     http.Header{},
			Body:       ioutil.NopCloser(strings.NewReader("")),
		}, nil
	})

	cases := map[string]struct {
		defaultsMode            oss.DefaultsMode
		retryer                 oss.Retryer
		retryMaxAttempts        int
		opRetryMaxAttempts      *int
		retryMode               oss.RetryMode
		expectClientRetryMode   oss.RetryMode
		expectClientMaxAttempts int
		expectOpMaxAttempts     int
	}{
		"defaults": {
			defaultsMode:            oss.DefaultsModeStandard,
			expectClientRetryMode:   oss.RetryModeStandard,
			expectClientMaxAttempts: 3,
			expectOpMaxAttempts:     3,
		},
		"custom default retry": {
			retryMode:               oss.RetryModeAdaptive,
			retryMaxAttempts:        10,
			expectClientRetryMode:   oss.RetryModeAdaptive,
			expectClientMaxAttempts: 10,
			expectOpMaxAttempts:     10,
		},
		"custom op max attempts": {
			retryMode:               oss.RetryModeAdaptive,
			retryMaxAttempts:        10,
			opRetryMaxAttempts:      oss.Int(2),
			expectClientRetryMode:   oss.RetryModeAdaptive,
			expectClientMaxAttempts: 10,
			expectOpMaxAttempts:     2,
		},
		"custom op no change max attempts": {
			retryMode:               oss.RetryModeAdaptive,
			retryMaxAttempts:        10,
			opRetryMaxAttempts:      oss.Int(10),
			expectClientRetryMode:   oss.RetryModeAdaptive,
			expectClientMaxAttempts: 10,
			expectOpMaxAttempts:     10,
		},
		"custom op 0 max attempts": {
			retryMode:               oss.RetryModeAdaptive,
			retryMaxAttempts:        10,
			opRetryMaxAttempts:      oss.Int(0),
			expectClientRetryMode:   oss.RetryModeAdaptive,
			expectClientMaxAttempts: 10,
			expectOpMaxAttempts:     10,
		},
	}

	for name, c := range cases {
		t.Run(name, func(t *testing.T) {
			client := NewFromConfig(oss.Config{
				DefaultsMode: c.defaultsMode,
				Retryer: func() func() oss.Retryer {
					if c.retryer == nil {
						return nil
					}

					return func() oss.Retryer { return c.retryer }
				}(),
				HTTPClient:       nopClient,
				RetryMaxAttempts: c.retryMaxAttempts,
				RetryMode:        c.retryMode,
			})

			if e, a := c.expectClientRetryMode, client.options.RetryMode; e != a {
				t.Errorf("expect %v retry mode, got %v", e, a)
			}
			if e, a := c.expectClientMaxAttempts, client.options.Retryer.MaxAttempts(); e != a {
				t.Errorf("expect %v max attempts, got %v", e, a)
			}

			_, _, err := client.invokeOperation(context.Background(), "mockOperation", struct{}{},
				[]func(*Options){
					func(o *Options) {
						if c.opRetryMaxAttempts == nil {
							return
						}
						o.RetryMaxAttempts = *c.opRetryMaxAttempts
					},
				},
				func(s *middleware.Stack, o Options) error {
					s.Initialize.Clear()
					s.Serialize.Clear()
					s.Build.Clear()
					s.Finalize.Clear()
					s.Deserialize.Clear()

					if e, a := c.expectOpMaxAttempts, o.Retryer.MaxAttempts(); e != a {
						t.Errorf("expect %v op max attempts, got %v", e, a)
					}
					return nil
				})
			if err != nil {
				t.Fatalf("expect no operation error, got %v", err)
			}
		})
	}
}
