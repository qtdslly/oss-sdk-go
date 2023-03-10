package retry_test

import (
	"context"
	"fmt"
	"log"
	"time"

	"oss-sdk-go/oss"
	"oss-sdk-go/oss/retry"
	config "oss-sdk-go/oss/retry/internal/mock"
	s3 "oss-sdk-go/oss/retry/internal/mock"
	types "oss-sdk-go/oss/retry/internal/mock"
)

func Example_overrideForAllClients() {
	cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithRetryer(func() oss.Retryer {
		// Generally you will always want to return new instance of a Retryer. This will avoid a global rate limit
		// bucket being shared between across all service clients.
		return retry.AddWithMaxBackoffDelay(retry.NewStandard(), time.Second*5)
	}))
	if err != nil {
		log.Fatal(err)
		return
	}

	client := s3.NewFromConfig(cfg)
	_ = client
}

func Example_overrideForSpecificClient() {
	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		log.Fatal(err)
		return
	}

	client := s3.NewFromConfig(cfg, func(options *s3.Options) {
		options.Retryer = retry.AddWithMaxBackoffDelay(options.Retryer, time.Second*5)
	})
	_ = client
}

func Example_overrideSpecificOperation() {
	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		log.Fatal(err)
		return
	}

	client := s3.NewFromConfig(cfg)

	// Wrap the default client retryer with an additional retryable error code for this specific
	// operation invocation
	_, err = client.GetObject(context.Background(), &s3.GetObjectInput{
		Bucket: oss.String("my-bucket"),
		Key:    oss.String("my-key"),
	}, func(options *types.Options) {
		options.Retryer = retry.AddWithErrorCodes(options.Retryer, (*types.NoSuchBucketException)(nil).ErrorCode())
	})
	if err != nil {
		log.Fatal(err)
		return
	}
}

func ExampleAddWithErrorCodes() {
	// Wrap a standard retyer and add the types.NoSuchBucketException Amazon S3 error code as retryable
	custom := retry.AddWithErrorCodes(retry.NewStandard(), (*types.NoSuchBucketException)(nil).ErrorCode())

	fmt.Println(custom.IsErrorRetryable(&types.NoSuchBucketException{}))
	// Output: true
}

func ExampleAddWithMaxAttempts() {
	// Wrap a standard retyer and set the max attempts to 5
	custom := retry.AddWithMaxAttempts(retry.NewStandard(), 5)

	fmt.Println(custom.MaxAttempts())
	// Output: 5
}

func ExampleAddWithMaxBackoffDelay() {
	// Wrap a standard retyer and add tthe NoSuchBucket API error code to the list of retryables
	custom := retry.AddWithMaxBackoffDelay(retry.NewStandard(), time.Second*5)
	_ = custom
}
