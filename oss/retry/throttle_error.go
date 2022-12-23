package retry

import (
	"errors"

	"oss-sdk-go/oss"
)

// IsErrorThrottle provides the interface of an implementation to determine if
// a error response from an operation is a throttling error.
type IsErrorThrottle interface {
	IsErrorThrottle(error) oss.Ternary
}

// IsErrorThrottles is a collection of checks to determine of the error a
// throttle error. Iterates through the checks and returns the state of
// throttle if any check returns something other than unknown.
type IsErrorThrottles []IsErrorThrottle

// IsErrorThrottle returns if the error is a throttle error if any of the
// checks in the list return a value other than unknown.
func (r IsErrorThrottles) IsErrorThrottle(err error) oss.Ternary {
	for _, re := range r {
		if v := re.IsErrorThrottle(err); v != oss.UnknownTernary {
			return v
		}
	}
	return oss.UnknownTernary
}

// IsErrorThrottleFunc wraps a function with the IsErrorThrottle interface.
type IsErrorThrottleFunc func(error) oss.Ternary

// IsErrorThrottle returns if the error is a throttle error.
func (fn IsErrorThrottleFunc) IsErrorThrottle(err error) oss.Ternary {
	return fn(err)
}

// ThrottleErrorCode determines if an attempt should be retried based on the
// API error code.
type ThrottleErrorCode struct {
	Codes map[string]struct{}
}

// IsErrorThrottle return if the error is a throttle error based on the error
// codes. Returns unknown if the error doesn't have a code or it is unknown.
func (r ThrottleErrorCode) IsErrorThrottle(err error) oss.Ternary {
	var v interface{ ErrorCode() string }

	if !errors.As(err, &v) {
		return oss.UnknownTernary
	}

	_, ok := r.Codes[v.ErrorCode()]
	if !ok {
		return oss.UnknownTernary
	}

	return oss.TrueTernary
}
