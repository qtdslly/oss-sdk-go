package retry

import (
	"testing"

	"oss-sdk-go/oss"
)

var _ oss.RetryerV2 = (*withIsErrorRetryable)(nil)
var _ oss.RetryerV2 = (*withMaxAttempts)(nil)
var _ oss.RetryerV2 = (*withMaxBackoffDelay)(nil)

func TestAddWithErrorCodes(t *testing.T) {
	cases := map[string]struct {
		Err    error
		Expect bool
	}{
		"retryable": {
			Err:    &mockErrorCodeError{code: "Error1"},
			Expect: true,
		},
		"not retryable": {
			Err:    &mockErrorCodeError{code: "Error3"},
			Expect: false,
		},
	}

	r := AddWithErrorCodes(oss.NopRetryer{}, "Error1", "Error2")

	for name, c := range cases {
		t.Run(name, func(t *testing.T) {
			if e, a := c.Expect, r.IsErrorRetryable(c.Err); e != a {
				t.Errorf("expect %v, got %v", e, a)
			}
		})
	}
}
