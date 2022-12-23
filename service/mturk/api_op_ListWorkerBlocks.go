// Code generated by smithy-go-codegen DO NOT EDIT.

package mturk

import (
	"context"
	"fmt"
	awsmiddleware "oss-sdk-go/oss/middleware"
	"oss-sdk-go/oss/signer/v4"
	"oss-sdk-go/service/mturk/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// The ListWorkersBlocks operation retrieves a list of Workers who are blocked from
// working on your HITs.
func (c *Client) ListWorkerBlocks(ctx context.Context, params *ListWorkerBlocksInput, optFns ...func(*Options)) (*ListWorkerBlocksOutput, error) {
	if params == nil {
		params = &ListWorkerBlocksInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListWorkerBlocks", params, optFns, c.addOperationListWorkerBlocksMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListWorkerBlocksOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListWorkerBlocksInput struct {
	MaxResults *int32

	// Pagination token
	NextToken *string

	noSmithyDocumentSerde
}

type ListWorkerBlocksOutput struct {

	// If the previous response was incomplete (because there is more data to
	// retrieve), Amazon Mechanical Turk returns a pagination token in the response.
	// You can use this pagination token to retrieve the next set of results.
	NextToken *string

	// The number of assignments on the page in the filtered results list, equivalent
	// to the number of assignments returned by this call.
	NumResults *int32

	// The list of WorkerBlocks, containing the collection of Worker IDs and reasons
	// for blocking.
	WorkerBlocks []types.WorkerBlock

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListWorkerBlocksMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpListWorkerBlocks{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpListWorkerBlocks{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListWorkerBlocks(options.Region), middleware.Before); err != nil {
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

// ListWorkerBlocksAPIClient is a client that implements the ListWorkerBlocks
// operation.
type ListWorkerBlocksAPIClient interface {
	ListWorkerBlocks(context.Context, *ListWorkerBlocksInput, ...func(*Options)) (*ListWorkerBlocksOutput, error)
}

var _ ListWorkerBlocksAPIClient = (*Client)(nil)

// ListWorkerBlocksPaginatorOptions is the paginator options for ListWorkerBlocks
type ListWorkerBlocksPaginatorOptions struct {
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListWorkerBlocksPaginator is a paginator for ListWorkerBlocks
type ListWorkerBlocksPaginator struct {
	options   ListWorkerBlocksPaginatorOptions
	client    ListWorkerBlocksAPIClient
	params    *ListWorkerBlocksInput
	nextToken *string
	firstPage bool
}

// NewListWorkerBlocksPaginator returns a new ListWorkerBlocksPaginator
func NewListWorkerBlocksPaginator(client ListWorkerBlocksAPIClient, params *ListWorkerBlocksInput, optFns ...func(*ListWorkerBlocksPaginatorOptions)) *ListWorkerBlocksPaginator {
	if params == nil {
		params = &ListWorkerBlocksInput{}
	}

	options := ListWorkerBlocksPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListWorkerBlocksPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListWorkerBlocksPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListWorkerBlocks page.
func (p *ListWorkerBlocksPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListWorkerBlocksOutput, error) {
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

	result, err := p.client.ListWorkerBlocks(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListWorkerBlocks(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "mturk-requester",
		OperationName: "ListWorkerBlocks",
	}
}
