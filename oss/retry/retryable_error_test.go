package retry

import (
	"fmt"
	"net"
	"net/url"
	"testing"

	"oss-sdk-go/oss"
)

type mockTemporaryError struct{ b bool }

func (m mockTemporaryError) Temporary() bool { return m.b }
func (m mockTemporaryError) Error() string {
	return fmt.Sprintf("mock temporary %t", m.b)
}

type mockTimeoutError struct{ b bool }

func (m mockTimeoutError) Timeout() bool { return m.b }
func (m mockTimeoutError) Error() string {
	return fmt.Sprintf("mock timeout %t", m.b)
}

type mockRetryableError struct{ b bool }

func (m mockRetryableError) RetryableError() bool { return m.b }
func (m mockRetryableError) Error() string {
	return fmt.Sprintf("mock retryable %t", m.b)
}

type mockCanceledError struct{ b bool }

func (m mockCanceledError) CanceledError() bool { return m.b }
func (m mockCanceledError) Error() string {
	return fmt.Sprintf("mock canceled %t", m.b)
}

type mockStatusCodeError struct{ code int }

func (m mockStatusCodeError) HTTPStatusCode() int { return m.code }
func (m mockStatusCodeError) Error() string {
	return fmt.Sprintf("status code error, %v", m.code)
}

type mockConnectionError struct{ err error }

func (m *mockConnectionError) ConnectionError() bool {
	return true
}
func (m *mockConnectionError) Error() string {
	return fmt.Sprintf("request error: %v", m.err)
}
func (m *mockConnectionError) Unwrap() error {
	return m.err
}

type mockErrorCodeError struct {
	code string
	err  error
}

func (m *mockErrorCodeError) ErrorCode() string { return m.code }
func (m *mockErrorCodeError) Error() string {
	return fmt.Sprintf("%v: mock error", m.code)
}
func (m *mockErrorCodeError) Unwrap() error {
	return m.err
}

func TestRetryConnectionErrors(t *testing.T) {
	cases := map[string]struct {
		Err       error
		Retryable oss.Ternary
	}{
		"nested connection reset": {
			Retryable: oss.TrueTernary,
			Err: fmt.Errorf("serialization error, %w",
				fmt.Errorf("connection reset")),
		},
		"top level connection reset": {
			Retryable: oss.TrueTernary,
			Err:       fmt.Errorf("connection reset"),
		},
		"wrapped connection reset": {
			Retryable: oss.TrueTernary,
			Err:       fmt.Errorf("some error: %w", fmt.Errorf("connection reset")),
		},
		"url.Error connection refused": {
			Retryable: oss.TrueTernary,
			Err: fmt.Errorf("some error, %w", &url.Error{
				Err: fmt.Errorf("connection refused"),
			}),
		},
		"other connection refused": {
			Retryable: oss.UnknownTernary,
			Err:       fmt.Errorf("connection refused"),
		},
		"nil error connection reset": {
			Retryable: oss.UnknownTernary,
		},
		"some other error": {
			Retryable: oss.UnknownTernary,
			Err:       fmt.Errorf("some error: %w", fmt.Errorf("something bad")),
		},
		"request send error": {
			Retryable: oss.TrueTernary,
			Err: fmt.Errorf("some error: %w", &mockConnectionError{err: &url.Error{
				Err: fmt.Errorf("another error"),
			}}),
		},
		"temporary error": {
			Retryable: oss.TrueTernary,
			Err:       &mockErrorCodeError{code: "SomeCode", err: mockTemporaryError{b: true}},
		},
		"timeout error": {
			Retryable: oss.TrueTernary,
			Err:       fmt.Errorf("some error: %w", mockTimeoutError{b: true}),
		},
		"timeout false error": {
			Retryable: oss.UnknownTernary,
			Err:       fmt.Errorf("some error: %w", mockTimeoutError{b: false}),
		},
		"net.OpError dial": {
			Retryable: oss.TrueTernary,
			Err: &net.OpError{
				Op:  "dial",
				Err: mockTimeoutError{b: false},
			},
		},
		"net.OpError nested": {
			Retryable: oss.TrueTernary,
			Err: &net.OpError{
				Op:  "read",
				Err: fmt.Errorf("some error %w", mockTimeoutError{b: true}),
			},
		},
	}

	for name, c := range cases {
		t.Run(name, func(t *testing.T) {
			var r RetryableConnectionError

			retryable := r.IsErrorRetryable(c.Err)
			if e, a := c.Retryable, retryable; e != a {
				t.Errorf("expect %v retryable, got %v", e, a)
			}
		})
	}
}

func TestRetryHTTPStatusCodes(t *testing.T) {
	cases := map[string]struct {
		Err    error
		Expect oss.Ternary
	}{
		"top level": {
			Err:    &mockStatusCodeError{code: 500},
			Expect: oss.TrueTernary,
		},
		"nested": {
			Err:    fmt.Errorf("some error, %w", &mockStatusCodeError{code: 500}),
			Expect: oss.TrueTernary,
		},
		"response error": {
			Err: fmt.Errorf("some error, %w", &mockErrorCodeError{
				code: "SomeCode",
				err:  &mockStatusCodeError{code: 502},
			}),
			Expect: oss.TrueTernary,
		},
	}

	r := RetryableHTTPStatusCode{Codes: map[int]struct{}{
		500: {},
		502: {},
	}}

	for name, c := range cases {
		t.Run(name, func(t *testing.T) {
			if e, a := c.Expect, r.IsErrorRetryable(c.Err); e != a {
				t.Errorf("expect %v, got %v", e, a)
			}
		})
	}
}

func TestRetryErrorCodes(t *testing.T) {
	cases := map[string]struct {
		Err    error
		Expect oss.Ternary
	}{
		"retryable code": {
			Err: &MaxAttemptsError{
				Err: &mockErrorCodeError{code: "ErrorCode1"},
			},
			Expect: oss.TrueTernary,
		},
		"not retryable code": {
			Err: &MaxAttemptsError{
				Err: &mockErrorCodeError{code: "SomeErroCode"},
			},
			Expect: oss.UnknownTernary,
		},
		"other error": {
			Err:    fmt.Errorf("some other error"),
			Expect: oss.UnknownTernary,
		},
	}

	r := RetryableErrorCode{Codes: map[string]struct{}{
		"ErrorCode1": {},
		"ErrorCode2": {},
	}}

	for name, c := range cases {
		t.Run(name, func(t *testing.T) {
			if e, a := c.Expect, r.IsErrorRetryable(c.Err); e != a {
				t.Errorf("expect %v, got %v", e, a)
			}
		})
	}
}

func TestCanceledError(t *testing.T) {
	cases := map[string]struct {
		Err    error
		Expect oss.Ternary
	}{
		"canceled error": {
			Err: fmt.Errorf("some error, %w", &oss.RequestCanceledError{
				Err: fmt.Errorf(":("),
			}),
			Expect: oss.FalseTernary,
		},
		"canceled retryable error": {
			Err: fmt.Errorf("some error, %w", &oss.RequestCanceledError{
				Err: mockRetryableError{b: true},
			}),
			Expect: oss.FalseTernary,
		},
		"not canceled error": {
			Err:    fmt.Errorf("some error, %w", mockCanceledError{b: false}),
			Expect: oss.UnknownTernary,
		},
		"retryable error": {
			Err:    fmt.Errorf("some error, %w", mockRetryableError{b: true}),
			Expect: oss.TrueTernary,
		},
	}

	r := IsErrorRetryables{
		NoRetryCanceledError{},
		RetryableError{},
	}
	for name, c := range cases {
		t.Run(name, func(t *testing.T) {
			if e, a := c.Expect, r.IsErrorRetryable(c.Err); e != a {
				t.Errorf("Expect %v retryable, got %v", e, a)
			}
		})
	}
}
