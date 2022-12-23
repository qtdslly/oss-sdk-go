// Code generated by smithy-go-codegen DO NOT EDIT.

package cloudfront

import (
	"context"
	awsmiddleware "oss-sdk-go/oss/middleware"
	"oss-sdk-go/oss/signer/v4"
	"oss-sdk-go/service/cloudfront/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Updates a real-time log configuration. When you update a real-time log
// configuration, all the parameters are updated with the values provided in the
// request. You cannot update some parameters independent of others. To update a
// real-time log configuration:
//
// * Call GetRealtimeLogConfig to get the current
// real-time log configuration.
//
// * Locally modify the parameters in the real-time
// log configuration that you want to update.
//
// * Call this API
// (UpdateRealtimeLogConfig) by providing the entire real-time log configuration,
// including the parameters that you modified and those that you didn’t.
//
// You
// cannot update a real-time log configuration’s Name or ARN.
func (c *Client) UpdateRealtimeLogConfig(ctx context.Context, params *UpdateRealtimeLogConfigInput, optFns ...func(*Options)) (*UpdateRealtimeLogConfigOutput, error) {
	if params == nil {
		params = &UpdateRealtimeLogConfigInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateRealtimeLogConfig", params, optFns, c.addOperationUpdateRealtimeLogConfigMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateRealtimeLogConfigOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateRealtimeLogConfigInput struct {

	// The Amazon Resource Name (ARN) for this real-time log configuration.
	ARN *string

	// Contains information about the Amazon Kinesis data stream where you are sending
	// real-time log data.
	EndPoints []types.EndPoint

	// A list of fields to include in each real-time log record. For more information
	// about fields, see Real-time log configuration fields
	// (https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/real-time-logs.html#understand-real-time-log-config-fields)
	// in the Amazon CloudFront Developer Guide.
	Fields []string

	// The name for this real-time log configuration.
	Name *string

	// The sampling rate for this real-time log configuration. The sampling rate
	// determines the percentage of viewer requests that are represented in the
	// real-time log data. You must provide an integer between 1 and 100, inclusive.
	SamplingRate *int64

	noSmithyDocumentSerde
}

type UpdateRealtimeLogConfigOutput struct {

	// A real-time log configuration.
	RealtimeLogConfig *types.RealtimeLogConfig

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationUpdateRealtimeLogConfigMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestxml_serializeOpUpdateRealtimeLogConfig{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestxml_deserializeOpUpdateRealtimeLogConfig{}, middleware.After)
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
	if err = addOpUpdateRealtimeLogConfigValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateRealtimeLogConfig(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opUpdateRealtimeLogConfig(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "cloudfront",
		OperationName: "UpdateRealtimeLogConfig",
	}
}
