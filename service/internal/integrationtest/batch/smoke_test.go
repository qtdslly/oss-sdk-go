//go:build integration
// +build integration

package batch

import (
	"context"
	"testing"
	"time"

	"oss-sdk-go/service/batch"

	"oss-sdk-go/service/internal/integrationtest"
)

func TestInteg_00_DescribeComputeEnvironments(t *testing.T) {
	ctx, cancelFn := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelFn()

	cfg, err := integrationtest.LoadConfigWithDefaultRegion("us-west-2")
	if err != nil {
		t.Fatalf("failed to load config, %v", err)
	}

	client := batch.NewFromConfig(cfg)
	params := &batch.DescribeComputeEnvironmentsInput{}
	_, err = client.DescribeComputeEnvironments(ctx, params)
	if err != nil {
		t.Errorf("expect no error, got %v", err)
	}
}
