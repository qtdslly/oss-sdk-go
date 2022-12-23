// Code generated by smithy-go-codegen DO NOT EDIT.

package translate

import (
	"context"
	"fmt"
	awsmiddleware "oss-sdk-go/oss/middleware"
	"oss-sdk-go/oss/signer/v4"
	"oss-sdk-go/service/translate/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Provides a list of your parallel data resources in Amazon Translate.
func (c *Client) ListParallelData(ctx context.Context, params *ListParallelDataInput, optFns ...func(*Options)) (*ListParallelDataOutput, error) {
	if params == nil {
		params = &ListParallelDataInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListParallelData", params, optFns, c.addOperationListParallelDataMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListParallelDataOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListParallelDataInput struct {

	// The maximum number of parallel data resources returned for each request.
	MaxResults *int32

	// A string that specifies the next page of results to return in a paginated
	// response.
	NextToken *string

	noSmithyDocumentSerde
}

type ListParallelDataOutput struct {

	// The string to use in a subsequent request to get the next page of results in a
	// paginated response. This value is null if there are no additional pages.
	NextToken *string

	// The properties of the parallel data resources returned by this request.
	ParallelDataPropertiesList []types.ParallelDataProperties

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListParallelDataMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpListParallelData{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpListParallelData{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListParallelData(options.Region), middleware.Before); err != nil {
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

// ListParallelDataAPIClient is a client that implements the ListParallelData
// operation.
type ListParallelDataAPIClient interface {
	ListParallelData(context.Context, *ListParallelDataInput, ...func(*Options)) (*ListParallelDataOutput, error)
}

var _ ListParallelDataAPIClient = (*Client)(nil)

// ListParallelDataPaginatorOptions is the paginator options for ListParallelData
type ListParallelDataPaginatorOptions struct {
	// The maximum number of parallel data resources returned for each request.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListParallelDataPaginator is a paginator for ListParallelData
type ListParallelDataPaginator struct {
	options   ListParallelDataPaginatorOptions
	client    ListParallelDataAPIClient
	params    *ListParallelDataInput
	nextToken *string
	firstPage bool
}

// NewListParallelDataPaginator returns a new ListParallelDataPaginator
func NewListParallelDataPaginator(client ListParallelDataAPIClient, params *ListParallelDataInput, optFns ...func(*ListParallelDataPaginatorOptions)) *ListParallelDataPaginator {
	if params == nil {
		params = &ListParallelDataInput{}
	}

	options := ListParallelDataPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListParallelDataPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListParallelDataPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListParallelData page.
func (p *ListParallelDataPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListParallelDataOutput, error) {
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

	result, err := p.client.ListParallelData(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListParallelData(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "translate",
		OperationName: "ListParallelData",
	}
}
