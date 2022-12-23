// Code generated by smithy-go-codegen DO NOT EDIT.

package outposts

import (
	"context"
	"fmt"
	awsmiddleware "oss-sdk-go/oss/middleware"
	"oss-sdk-go/oss/signer/v4"
	"oss-sdk-go/service/outposts/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Lists the Outpost orders for your Amazon Web Services account. You can filter
// your request by Outpost to return a more specific list of results.
func (c *Client) ListOrders(ctx context.Context, params *ListOrdersInput, optFns ...func(*Options)) (*ListOrdersOutput, error) {
	if params == nil {
		params = &ListOrdersInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListOrders", params, optFns, c.addOperationListOrdersMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListOrdersOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListOrdersInput struct {

	// The maximum page size.
	MaxResults *int32

	// The pagination token.
	NextToken *string

	// The ID or the Amazon Resource Name (ARN) of the Outpost.
	OutpostIdentifierFilter *string

	noSmithyDocumentSerde
}

type ListOrdersOutput struct {

	// The pagination token.
	NextToken *string

	// Information about the orders.
	Orders []types.OrderSummary

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListOrdersMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListOrders{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListOrders{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListOrders(options.Region), middleware.Before); err != nil {
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

// ListOrdersAPIClient is a client that implements the ListOrders operation.
type ListOrdersAPIClient interface {
	ListOrders(context.Context, *ListOrdersInput, ...func(*Options)) (*ListOrdersOutput, error)
}

var _ ListOrdersAPIClient = (*Client)(nil)

// ListOrdersPaginatorOptions is the paginator options for ListOrders
type ListOrdersPaginatorOptions struct {
	// The maximum page size.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListOrdersPaginator is a paginator for ListOrders
type ListOrdersPaginator struct {
	options   ListOrdersPaginatorOptions
	client    ListOrdersAPIClient
	params    *ListOrdersInput
	nextToken *string
	firstPage bool
}

// NewListOrdersPaginator returns a new ListOrdersPaginator
func NewListOrdersPaginator(client ListOrdersAPIClient, params *ListOrdersInput, optFns ...func(*ListOrdersPaginatorOptions)) *ListOrdersPaginator {
	if params == nil {
		params = &ListOrdersInput{}
	}

	options := ListOrdersPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListOrdersPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListOrdersPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListOrders page.
func (p *ListOrdersPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListOrdersOutput, error) {
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

	result, err := p.client.ListOrders(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListOrders(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "outposts",
		OperationName: "ListOrders",
	}
}
