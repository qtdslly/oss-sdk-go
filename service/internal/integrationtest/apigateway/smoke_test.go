//go:build integration
// +build integration

package apigateway

import (
	"context"
	"errors"
	"testing"
	"time"

	"oss-sdk-go/oss"
	"oss-sdk-go/service/apigateway"

	"oss-sdk-go/service/internal/integrationtest"
	"github.com/aws/smithy-go"
)

func TestInteg_00_GetDomainNames(t *testing.T) {
	ctx, cancelFn := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelFn()

	cfg, err := integrationtest.LoadConfigWithDefaultRegion("us-west-2")
	if err != nil {
		t.Fatalf("failed to load config, %v", err)
	}

	client := apigateway.NewFromConfig(cfg)
	params := &apigateway.GetDomainNamesInput{}
	_, err = client.GetDomainNames(ctx, params)
	if err != nil {
		t.Errorf("expect no error, got %v", err)
	}
}

func TestInteg_01_CreateUsagePlanKey(t *testing.T) {
	ctx, cancelFn := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelFn()

	cfg, err := integrationtest.LoadConfigWithDefaultRegion("us-west-2")
	if err != nil {
		t.Fatalf("failed to load config, %v", err)
	}

	client := apigateway.NewFromConfig(cfg)
	params := &apigateway.CreateUsagePlanKeyInput{
		KeyId:       oss.String("bar"),
		KeyType:     oss.String("fixx"),
		UsagePlanId: oss.String("foo"),
	}
	_, err = client.CreateUsagePlanKey(ctx, params)
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
