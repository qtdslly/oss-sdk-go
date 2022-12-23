//go:build integration
// +build integration

package timestreamwrite

import (
	"context"
	"oss-sdk-go/oss"
	"oss-sdk-go/service/internal/integrationtest"
	tw "oss-sdk-go/service/timestreamwrite"
	"testing"
	"time"
)

func TestInteg_00_ListDatabases(t *testing.T) {
	ctx, cancelFn := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancelFn()

	cfg, err := integrationtest.LoadConfigWithDefaultRegion("us-east-1")
	if err != nil {
		t.Fatalf("failed to load config, %v", err)
	}

	// Create an Amazon timestreamwrite service client
	client := tw.NewFromConfig(cfg)

	// ListDatabases
	_, err = client.ListDatabases(ctx, &tw.ListDatabasesInput{
		MaxResults: oss.Int32(1),
	})
	if err != nil {
		t.Fatal(err)
	}
}
