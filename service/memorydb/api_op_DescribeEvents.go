// Code generated by smithy-go-codegen DO NOT EDIT.

package memorydb

import (
	"context"
	awsmiddleware "oss-sdk-go/oss/middleware"
	"oss-sdk-go/oss/signer/v4"
	"oss-sdk-go/service/memorydb/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Returns events related to clusters, security groups, and parameter groups. You
// can obtain events specific to a particular cluster, security group, or parameter
// group by providing the name as a parameter. By default, only the events
// occurring within the last hour are returned; however, you can retrieve up to 14
// days' worth of events if necessary.
func (c *Client) DescribeEvents(ctx context.Context, params *DescribeEventsInput, optFns ...func(*Options)) (*DescribeEventsOutput, error) {
	if params == nil {
		params = &DescribeEventsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeEvents", params, optFns, c.addOperationDescribeEventsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeEventsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeEventsInput struct {

	// The number of minutes worth of events to retrieve.
	Duration *int32

	// The end of the time interval for which to retrieve events, specified in ISO 8601
	// format. Example: 2017-03-30T07:03:49.555Z
	EndTime *time.Time

	// The maximum number of records to include in the response. If more records exist
	// than the specified MaxResults value, a token is included in the response so that
	// the remaining results can be retrieved.
	MaxResults *int32

	// An optional argument to pass in case the total number of records exceeds the
	// value of MaxResults. If nextToken is returned, there are more results available.
	// The value of nextToken is a unique pagination token for each page. Make the call
	// again using the returned token to retrieve the next page. Keep all other
	// arguments unchanged.
	NextToken *string

	// The identifier of the event source for which events are returned. If not
	// specified, all sources are included in the response.
	SourceName *string

	// The event source to retrieve events for. If no value is specified, all events
	// are returned.
	SourceType types.SourceType

	// The beginning of the time interval to retrieve events for, specified in ISO 8601
	// format. Example: 2017-03-30T07:03:49.555Z
	StartTime *time.Time

	noSmithyDocumentSerde
}

type DescribeEventsOutput struct {

	// A list of events. Each element in the list contains detailed information about
	// one event.
	Events []types.Event

	// An optional argument to pass in case the total number of records exceeds the
	// value of MaxResults. If nextToken is returned, there are more results available.
	// The value of nextToken is a unique pagination token for each page. Make the call
	// again using the returned token to retrieve the next page. Keep all other
	// arguments unchanged.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeEventsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDescribeEvents{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDescribeEvents{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeEvents(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDescribeEvents(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "memorydb",
		OperationName: "DescribeEvents",
	}
}
