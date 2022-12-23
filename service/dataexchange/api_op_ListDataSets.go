// Code generated by smithy-go-codegen DO NOT EDIT.

package dataexchange

import (
	"context"
	"fmt"
	awsmiddleware "oss-sdk-go/oss/middleware"
	"oss-sdk-go/oss/signer/v4"
	"oss-sdk-go/service/dataexchange/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// This operation lists your data sets. When listing by origin OWNED, results are
// sorted by CreatedAt in descending order. When listing by origin ENTITLED, there
// is no order and the maxResults parameter is ignored.
func (c *Client) ListDataSets(ctx context.Context, params *ListDataSetsInput, optFns ...func(*Options)) (*ListDataSetsOutput, error) {
	if params == nil {
		params = &ListDataSetsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListDataSets", params, optFns, c.addOperationListDataSetsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListDataSetsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListDataSetsInput struct {

	// The maximum number of results returned by a single call.
	MaxResults int32

	// The token value retrieved from a previous call to access the next page of
	// results.
	NextToken *string

	// A property that defines the data set as OWNED by the account (for providers) or
	// ENTITLED to the account (for subscribers).
	Origin *string

	noSmithyDocumentSerde
}

type ListDataSetsOutput struct {

	// The data set objects listed by the request.
	DataSets []types.DataSetEntry

	// The token value retrieved from a previous call to access the next page of
	// results.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListDataSetsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListDataSets{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListDataSets{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListDataSets(options.Region), middleware.Before); err != nil {
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

// ListDataSetsAPIClient is a client that implements the ListDataSets operation.
type ListDataSetsAPIClient interface {
	ListDataSets(context.Context, *ListDataSetsInput, ...func(*Options)) (*ListDataSetsOutput, error)
}

var _ ListDataSetsAPIClient = (*Client)(nil)

// ListDataSetsPaginatorOptions is the paginator options for ListDataSets
type ListDataSetsPaginatorOptions struct {
	// The maximum number of results returned by a single call.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListDataSetsPaginator is a paginator for ListDataSets
type ListDataSetsPaginator struct {
	options   ListDataSetsPaginatorOptions
	client    ListDataSetsAPIClient
	params    *ListDataSetsInput
	nextToken *string
	firstPage bool
}

// NewListDataSetsPaginator returns a new ListDataSetsPaginator
func NewListDataSetsPaginator(client ListDataSetsAPIClient, params *ListDataSetsInput, optFns ...func(*ListDataSetsPaginatorOptions)) *ListDataSetsPaginator {
	if params == nil {
		params = &ListDataSetsInput{}
	}

	options := ListDataSetsPaginatorOptions{}
	if params.MaxResults != 0 {
		options.Limit = params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListDataSetsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListDataSetsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListDataSets page.
func (p *ListDataSetsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListDataSetsOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.NextToken = p.nextToken

	params.MaxResults = p.options.Limit

	result, err := p.client.ListDataSets(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListDataSets(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "dataexchange",
		OperationName: "ListDataSets",
	}
}
