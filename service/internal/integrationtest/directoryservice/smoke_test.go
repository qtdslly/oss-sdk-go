//go:build integration
// +build integration

package directoryservice

import (
	"context"
	"errors"
	"testing"
	"time"

	"oss-sdk-go/oss"
	"oss-sdk-go/service/directoryservice"
	"oss-sdk-go/service/directoryservice/types"

	"oss-sdk-go/service/internal/integrationtest"
	"github.com/aws/smithy-go"
)

func TestInteg_00_DescribeDirectories(t *testing.T) {
	ctx, cancelFn := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelFn()

	cfg, err := integrationtest.LoadConfigWithDefaultRegion("us-west-2")
	if err != nil {
		t.Fatalf("failed to load config, %v", err)
	}

	client := directoryservice.NewFromConfig(cfg)
	params := &directoryservice.DescribeDirectoriesInput{}
	_, err = client.DescribeDirectories(ctx, params)
	if err != nil {
		t.Errorf("expect no error, got %v", err)
	}
}

func TestInteg_01_CreateDirectory(t *testing.T) {
	ctx, cancelFn := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelFn()

	cfg, err := integrationtest.LoadConfigWithDefaultRegion("us-west-2")
	if err != nil {
		t.Fatalf("failed to load config, %v", err)
	}

	client := directoryservice.NewFromConfig(cfg)
	params := &directoryservice.CreateDirectoryInput{
		Name:     oss.String(""),
		Password: oss.String(""),
		Size:     types.DirectorySize(""),
	}
	_, err = client.CreateDirectory(ctx, params)
	if err == nil {
		t.Fatalf("expect request to fail")
	}

	var apiErr smithy.APIError
	if !errors.As(err, &apiErr) {
		t.Fatalf("expect error to be API error, was not, %v", err)
	}
	if len(apiErr.ErrorCode()) == 0 {
		t.Errorf("expect non-empty error code")
	}
	if len(apiErr.ErrorMessage()) == 0 {
		t.Errorf("expect non-empty error message")
	}
}
