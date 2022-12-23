// Code generated by smithy-go-codegen DO NOT EDIT.

package ec2

import (
	"context"
	awsmiddleware "oss-sdk-go/oss/middleware"
	"oss-sdk-go/oss/signer/v4"
	"oss-sdk-go/service/ec2/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Describes the events for the specified EC2 Fleet during the specified time. EC2
// Fleet events are delayed by up to 30 seconds before they can be described. This
// ensures that you can query by the last evaluated time and not miss a recorded
// event. EC2 Fleet events are available for 48 hours. For more information, see
// Monitor fleet events using Amazon EventBridge
// (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/fleet-monitor.html) in the
// Amazon EC2 User Guide.
func (c *Client) DescribeFleetHistory(ctx context.Context, params *DescribeFleetHistoryInput, optFns ...func(*Options)) (*DescribeFleetHistoryOutput, error) {
	if params == nil {
		params = &DescribeFleetHistoryInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeFleetHistory", params, optFns, c.addOperationDescribeFleetHistoryMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeFleetHistoryOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeFleetHistoryInput struct {

	// The ID of the EC2 Fleet.
	//
	// This member is required.
	FleetId *string

	// The start date and time for the events, in UTC format (for example,
	// YYYY-MM-DDTHH:MM:SSZ).
	//
	// This member is required.
	StartTime *time.Time

	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have the
	// required permissions, the error response is DryRunOperation. Otherwise, it is
	// UnauthorizedOperation.
	DryRun *bool

	// The type of events to describe. By default, all events are described.
	EventType types.FleetEventType

	// The maximum number of results to return in a single call. Specify a value
	// between 1 and 1000. The default value is 1000. To retrieve the remaining
	// results, make another call with the returned NextToken value.
	MaxResults *int32

	// The token for the next set of results.
	NextToken *string

	noSmithyDocumentSerde
}

type DescribeFleetHistoryOutput struct {

	// The ID of the EC Fleet.
	FleetId *string

	// Information about the events in the history of the EC2 Fleet.
	HistoryRecords []types.HistoryRecordEntry

	// The last date and time for the events, in UTC format (for example,
	// YYYY-MM-DDTHH:MM:SSZ). All records up to this time were retrieved. If nextToken
	// indicates that there are more results, this value is not present.
	LastEvaluatedTime *time.Time

	// The token for the next set of results.
	NextToken *string

	// The start date and time for the events, in UTC format (for example,
	// YYYY-MM-DDTHH:MM:SSZ).
	StartTime *time.Time

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeFleetHistoryMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsEc2query_serializeOpDescribeFleetHistory{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsEc2query_deserializeOpDescribeFleetHistory{}, middleware.After)
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
	if err = addOpDescribeFleetHistoryValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeFleetHistory(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDescribeFleetHistory(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ec2",
		OperationName: "DescribeFleetHistory",
	}
}
