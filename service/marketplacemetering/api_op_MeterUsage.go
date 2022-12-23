// Code generated by smithy-go-codegen DO NOT EDIT.

package marketplacemetering

import (
	"context"
	awsmiddleware "oss-sdk-go/oss/middleware"
	"oss-sdk-go/oss/signer/v4"
	"oss-sdk-go/service/marketplacemetering/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// API to emit metering records. For identical requests, the API is idempotent. It
// simply returns the metering record ID. MeterUsage is authenticated on the
// buyer's AWS account using credentials from the EC2 instance, ECS task, or EKS
// pod. MeterUsage can optionally include multiple usage allocations, to provide
// customers with usage data split into buckets by tags that you define (or allow
// the customer to define). Usage records are expected to be submitted as quickly
// as possible after the event that is being recorded, and are not accepted more
// than 6 hours after the event.
func (c *Client) MeterUsage(ctx context.Context, params *MeterUsageInput, optFns ...func(*Options)) (*MeterUsageOutput, error) {
	if params == nil {
		params = &MeterUsageInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "MeterUsage", params, optFns, c.addOperationMeterUsageMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*MeterUsageOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type MeterUsageInput struct {

	// Product code is used to uniquely identify a product in AWS Marketplace. The
	// product code should be the same as the one used during the publishing of a new
	// product.
	//
	// This member is required.
	ProductCode *string

	// Timestamp, in UTC, for which the usage is being reported. Your application can
	// meter usage for up to one hour in the past. Make sure the timestamp value is not
	// before the start of the software usage.
	//
	// This member is required.
	Timestamp *time.Time

	// It will be one of the fcp dimension name provided during the publishing of the
	// product.
	//
	// This member is required.
	UsageDimension *string

	// Checks whether you have the permissions required for the action, but does not
	// make the request. If you have the permissions, the request returns
	// DryRunOperation; otherwise, it returns UnauthorizedException. Defaults to false
	// if not specified.
	DryRun *bool

	// The set of UsageAllocations to submit. The sum of all UsageAllocation quantities
	// must equal the UsageQuantity of the MeterUsage request, and each UsageAllocation
	// must have a unique set of tags (include no tags).
	UsageAllocations []types.UsageAllocation

	// Consumption value for the hour. Defaults to 0 if not specified.
	UsageQuantity *int32

	noSmithyDocumentSerde
}

type MeterUsageOutput struct {

	// Metering record id.
	MeteringRecordId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationMeterUsageMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpMeterUsage{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpMeterUsage{}, middleware.After)
	if err != nil {
		return err
	}
	if err = addSetLoggerMiddleware(stack, options); err != nil {
		return err
	}
	if err = awsmiddleware.AddClientRequestIDMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddComputeContentLengthMiddleware(stack); err != nil {
		return err
	}
	if err = addResolveEndpointMiddleware(stack, options); err != nil {
		return err
	}
	if err = v4.AddComputePayloadSHA256Middleware(stack); err != nil {
		return err
	}
	if err = addRetryMiddlewares(stack, options); err != nil {
		return err
	}
	if err = addHTTPSignerV4Middleware(stack, options); err != nil {
		return err
	}
	if err = awsmiddleware.AddRawResponseToMetadata(stack); err != nil {
		return err
	}
	if err = awsmiddleware.AddRecordResponseTiming(stack); err != nil {
		return err
	}
	if err = addClientUserAgent(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddErrorCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = addOpMeterUsageValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opMeterUsage(options.Region), middleware.Before); err != nil {
		return err
	}
	if err = addRequestIDRetrieverMiddleware(stack); err != nil {
		return err
	}
	if err = addResponseErrorMiddleware(stack); err != nil {
		return err
	}
	if err = addRequestResponseLogging(stack, options); err != nil {
		return err
	}
	return nil
}

func newServiceMetadataMiddleware_opMeterUsage(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "aws-marketplace",
		OperationName: "MeterUsage",
	}
}
