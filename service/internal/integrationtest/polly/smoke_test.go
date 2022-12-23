//go:build integration
// +build integration

package polly

import (
	"context"
	"testing"
	"time"

	"oss-sdk-go/service/internal/integrationtest"
	"oss-sdk-go/service/polly"
)

func TestInteg_00_DescribeVoices(t *testing.T) {
	ctx, cancelFn := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelFn()

	cfg, err := integrationtest.LoadConfigWithDefaultRegion("us-west-2")
	if err != nil {
		t.Fatalf("failed to load config, %v", err)
	}

	client := polly.NewFromConfig(cfg)
	params := &polly.DescribeVoicesInput{}
	_, err = client.DescribeVoices(ctx, params)
	if err != nil {
		t.Errorf("expect no error, got %v", err)
	}
}
