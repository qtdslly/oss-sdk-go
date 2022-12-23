// Code generated by smithy-go-codegen DO NOT EDIT.

package kinesis

import (
	"context"
	awsmiddleware "oss-sdk-go/oss/middleware"
	"oss-sdk-go/oss/signer/v4"
	"oss-sdk-go/service/kinesis/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Provides a summarized description of the specified Kinesis data stream without
// the shard list. The information returned includes the stream name, Amazon
// Resource Name (ARN), status, record retention period, approximate creation time,
// monitoring, encryption details, and open shard count. DescribeStreamSummary has
// a limit of 20 transactions per second per account.
func (c *Client) DescribeStreamSummary(ctx context.Context, params *DescribeStreamSummaryInput, optFns ...func(*Options)) (*DescribeStreamSummaryOutput, error) {
	if params == nil {
		params = &DescribeStreamSummaryInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeStreamSummary", params, optFns, c.addOperationDescribeStreamSummaryMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeStreamSummaryOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeStreamSummaryInput struct {

	// The name of the stream to describe.
	//
	// This member is required.
	StreamName *string

	noSmithyDocumentSerde
}

type DescribeStreamSummaryOutput struct {

	// A StreamDescriptionSummary containing information about the stream.
	//
	// This member is required.
	StreamDescriptionSummary *types.StreamDescriptionSummary

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeStreamSummaryMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDescribeStreamSummary{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDescribeStreamSummary{}, middleware.After)
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
	if err = addOpDescribeStreamSummaryValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeStreamSummary(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDescribeStreamSummary(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "kinesis",
		OperationName: "DescribeStreamSummary",
	}
}
