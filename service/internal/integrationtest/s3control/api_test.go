//go:build integration
// +build integration

package s3control

import (
	"context"
	"errors"
	"testing"

	"oss-sdk-go/oss"
	"oss-sdk-go/service/s3control"
	"oss-sdk-go/service/s3control/types"
)

func TestInteg_PublicAccessBlock(t *testing.T) {
	ctx := context.Background()
	_, err := svc.GetPublicAccessBlock(ctx, &s3control.GetPublicAccessBlockInput{
		AccountId: oss.String(accountID),
	})
	if err != nil {
		// Ignore NoSuchPublicAccessBlockConfiguration, but fail on any other error.
		var e *types.NoSuchPublicAccessBlockConfiguration
		if !errors.As(err, &e) {
			t.Fatalf("expect no error for GetPublicAccessBlock, got %v", err)
		}
	}
}
