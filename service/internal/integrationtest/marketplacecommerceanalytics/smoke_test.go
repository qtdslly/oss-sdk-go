//go:build integration
// +build integration

package marketplacecommerceanalytics

import (
	"context"
	"errors"
	"testing"
	"time"

	"oss-sdk-go/oss"
	"oss-sdk-go/service/internal/integrationtest"
	"oss-sdk-go/service/marketplacecommerceanalytics"
	"oss-sdk-go/service/marketplacecommerceanalytics/types"
	"github.com/aws/smithy-go"
	smithytime "github.com/aws/smithy-go/time"
)

func TestInteg_00_GenerateDataSet(t *testing.T) {
	ctx, cancelFn := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelFn()

	cfg, err := integrationtest.LoadConfigWithDefaultRegion("us-east-1")
	if err != nil {
		t.Fatalf("failed to load config, %v", err)
	}

	client := marketplacecommerceanalytics.NewFromConfig(cfg)
	params := &marketplacecommerceanalytics.GenerateDataSetInput{
		DataSetPublicationDate:  oss.Time(smithytime.ParseEpochSeconds(0.000000)),
		DataSetType:             types.DataSetTypeDailyBusinessFees,
		DestinationS3BucketName: oss.String("fake-bucket"),
		RoleNameArn:             oss.String("fake-arn"),
		SnsTopicArn:             oss.String("fake-arn"),
	}
	_, err = client.GenerateDataSet(ctx, params)
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
