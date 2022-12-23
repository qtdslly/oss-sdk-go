//go:build integration
// +build integration

package manager_test

import (
	"context"
	"testing"

	"oss-sdk-go/oss"
	"oss-sdk-go/feature/s3/manager"
	"oss-sdk-go/service/s3"
)

func TestInteg_GetBucketRegion(t *testing.T) {
	expectRegion := integConfig.Region

	region, err := manager.GetBucketRegion(context.Background(), s3.NewFromConfig(integConfig), oss.ToString(bucketName))
	if err != nil {
		t.Fatalf("expect no error, got %v", err)
	}

	if e, a := expectRegion, region; e != a {
		t.Errorf("expect %s bucket region, got %s", e, a)
	}
}
