package retry

import (
	"context"
	"time"

	"oss-sdk-go/oss"
)

// AddWithErrorCodes returns a Retryer with additional error codes considered
// for determining if the error should be retried.
func AddWithErrorCodes(r oss.Retryer, codes ...string) oss.Retryer {
	retryable := &RetryableErrorCode{
		Codes: map[string]struct{}{},
	}
	for _, c := range codes {
		retryable.Codes[c] = struct{}{}
	}

	return &withIsErrorRetryable{
		RetryerV2: wrapAsRetryerV2(r),
		Retryable: retryable,
	}
}

type withIsErrorRetryable struct {
	oss.RetryerV2
	Retryable IsErrorRetryable
}

func (r *withIsErrorRetryable) IsErrorRetryable(err error) bool {
	if v := r.Retryable.IsErrorRetryable(err); v != oss.UnknownTernary {
		return v.Bool()
	}
	return r.RetryerV2.IsErrorRetryable(err)
}

// AddWithMaxAttempts returns a Retryer with MaxAttempts set to the value
// specified.
func AddWithMaxAttempts(r oss.Retryer, max int) oss.Retryer {
	return &withMaxAttempts{
		RetryerV2: wrapAsRetryerV2(r),
		Max:       max,
	}
}

type withMaxAttempts struct {
	oss.RetryerV2
	Max int
}

func (w *withMaxAttempts) MaxAttempts() int {
	return w.Max
}

// AddWithMaxBackoffDelay returns a retryer wrapping the passed in retryer
// overriding the RetryDelay behavior for a alternate minimum initial backoff
// delay.
func AddWithMaxBackoffDelay(r oss.Retryer, delay time.Duration) oss.Retryer {
	return &withMaxBackoffDelay{
		RetryerV2: wrapAsRetryerV2(r),
		backoff:   NewExponentialJitterBackoff(delay),
	}
}

type withMaxBackoffDelay struct {
	oss.RetryerV2
	backoff *ExponentialJitterBackoff
}

func (r *withMaxBackoffDelay) RetryDelay(attempt int, err error) (time.Duration, error) {
	return r.backoff.BackoffDelay(attempt, err)
}

type wrappedAsRetryerV2 struct {
	oss.Retryer
}

func wrapAsRetryerV2(r oss.Retryer) oss.RetryerV2 {
	v, ok := r.(oss.RetryerV2)
	if !ok {
		v = wrappedAsRetryerV2{Retryer: r}
	}

	return v
}

func (w wrappedAsRetryerV2) GetAttemptToken(context.Context) (func(error) error, error) {
	return w.Retryer.GetInitialToken(), nil
}
