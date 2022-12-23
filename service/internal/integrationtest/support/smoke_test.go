//go:build integration
// +build integration

package support

import (
	"context"
	"errors"
	"testing"
	"time"

	"oss-sdk-go/oss"
	"oss-sdk-go/service/support"
	"github.com/aws/smithy-go"

	"oss-sdk-go/service/internal/integrationtest"
)

func TestInteg_00_DescribeServices(t *testing.T) {
	ctx, cancelFn := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelFn()

	cfg, err := integrationtest.LoadConfigWithDefaultRegion("us-east-1")
	if err != nil {
		t.Fatalf("failed to load config, %v", err)
	}

	client := support.NewFromConfig(cfg)
	params := &support.DescribeServicesInput{}
	_, err = client.DescribeServices(ctx, params)
	if err != nil {
		t.Errorf("expect no error, got %v", err)
	}
}

func TestInteg_01_CreateCase(t *testing.T) {
	ctx, cancelFn := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelFn()

	cfg, err := integrationtest.LoadConfigWithDefaultRegion("us-east-1")
	if err != nil {
		t.Fatalf("failed to load config, %v", err)
	}

	client := support.NewFromConfig(cfg)
	params := &support.CreateCaseInput{
		CategoryCode:      oss.String("category"),
		CommunicationBody: oss.String("communication"),
		ServiceCode:       oss.String("amazon-dynamodb"),
		SeverityCode:      oss.String("low"),
		Subject:           oss.String("subject"),
	}
	_, err = client.CreateCase(ctx, params)
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
