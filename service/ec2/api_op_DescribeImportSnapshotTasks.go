// Code generated by smithy-go-codegen DO NOT EDIT.

package ec2

import (
	"context"
	"fmt"
	awsmiddleware "oss-sdk-go/oss/middleware"
	"oss-sdk-go/oss/signer/v4"
	"oss-sdk-go/service/ec2/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Describes your import snapshot tasks.
func (c *Client) DescribeImportSnapshotTasks(ctx context.Context, params *DescribeImportSnapshotTasksInput, optFns ...func(*Options)) (*DescribeImportSnapshotTasksOutput, error) {
	if params == nil {
		params = &DescribeImportSnapshotTasksInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeImportSnapshotTasks", params, optFns, c.addOperationDescribeImportSnapshotTasksMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeImportSnapshotTasksOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeImportSnapshotTasksInput struct {

	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have the
	// required permissions, the error response is DryRunOperation. Otherwise, it is
	// UnauthorizedOperation.
	DryRun *bool

	// The filters.
	Filters []types.Filter

	// A list of import snapshot task IDs.
	ImportTaskIds []string

	// The maximum number of results to return in a single call. To retrieve the
	// remaining results, make another call with the returned NextToken value.
	MaxResults *int32

	// A token that indicates the next page of results.
	NextToken *string

	noSmithyDocumentSerde
}

type DescribeImportSnapshotTasksOutput struct {

	// A list of zero or more import snapshot tasks that are currently active or were
	// completed or canceled in the previous 7 days.
	ImportSnapshotTasks []types.ImportSnapshotTask

	// The token to use to get the next page of results. This value is null when there
	// are no more results to return.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeImportSnapshotTasksMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsEc2query_serializeOpDescribeImportSnapshotTasks{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsEc2query_deserializeOpDescribeImportSnapshotTasks{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeImportSnapshotTasks(options.Region), middleware.Before); err != nil {
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

// DescribeImportSnapshotTasksAPIClient is a client that implements the
// DescribeImportSnapshotTasks operation.
type DescribeImportSnapshotTasksAPIClient interface {
	DescribeImportSnapshotTasks(context.Context, *DescribeImportSnapshotTasksInput, ...func(*Options)) (*DescribeImportSnapshotTasksOutput, error)
}

var _ DescribeImportSnapshotTasksAPIClient = (*Client)(nil)

// DescribeImportSnapshotTasksPaginatorOptions is the paginator options for
// DescribeImportSnapshotTasks
type DescribeImportSnapshotTasksPaginatorOptions struct {
	// The maximum number of results to return in a single call. To retrieve the
	// remaining results, make another call with the returned NextToken value.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// DescribeImportSnapshotTasksPaginator is a paginator for
// DescribeImportSnapshotTasks
type DescribeImportSnapshotTasksPaginator struct {
	options   DescribeImportSnapshotTasksPaginatorOptions
	client    DescribeImportSnapshotTasksAPIClient
	params    *DescribeImportSnapshotTasksInput
	nextToken *string
	firstPage bool
}

// NewDescribeImportSnapshotTasksPaginator returns a new
// DescribeImportSnapshotTasksPaginator
func NewDescribeImportSnapshotTasksPaginator(client DescribeImportSnapshotTasksAPIClient, params *DescribeImportSnapshotTasksInput, optFns ...func(*DescribeImportSnapshotTasksPaginatorOptions)) *DescribeImportSnapshotTasksPaginator {
	if params == nil {
		params = &DescribeImportSnapshotTasksInput{}
	}

	options := DescribeImportSnapshotTasksPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &DescribeImportSnapshotTasksPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *DescribeImportSnapshotTasksPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next DescribeImportSnapshotTasks page.
func (p *DescribeImportSnapshotTasksPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*DescribeImportSnapshotTasksOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.NextToken = p.nextToken

	var limit *int32
	if p.options.Limit > 0 {
		limit = &p.options.Limit
	}
	params.MaxResults = limit

	result, err := p.client.DescribeImportSnapshotTasks(ctx, &params, optFns...)
	if err != nil {
		return nil, err
	}
	p.firstPage = false

	prevToken := p.nextToken
	p.nextToken = result.NextToken

	if p.options.StopOnDuplicateToken &&
		prevToken != nil &&
		p.nextToken != nil &&
		*prevToken == *p.nextToken {
		p.nextToken = nil
	}

	return result, nil
}

func newServiceMetadataMiddleware_opDescribeImportSnapshotTasks(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ec2",
		OperationName: "DescribeImportSnapshotTasks",
	}
}
