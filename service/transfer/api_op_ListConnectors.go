// Code generated by smithy-go-codegen DO NOT EDIT.

package transfer

import (
	"context"
	"fmt"
	awsmiddleware "oss-sdk-go/oss/middleware"
	"oss-sdk-go/oss/signer/v4"
	"oss-sdk-go/service/transfer/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Lists the connectors for the specified Region.
func (c *Client) ListConnectors(ctx context.Context, params *ListConnectorsInput, optFns ...func(*Options)) (*ListConnectorsOutput, error) {
	if params == nil {
		params = &ListConnectorsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListConnectors", params, optFns, c.addOperationListConnectorsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListConnectorsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListConnectorsInput struct {

	// The maximum number of connectors to return.
	MaxResults *int32

	// When you can get additional results from the ListConnectors call, a NextToken
	// parameter is returned in the output. You can then pass in a subsequent command
	// to the NextToken parameter to continue listing additional connectors.
	NextToken *string

	noSmithyDocumentSerde
}

type ListConnectorsOutput struct {

	// Returns an array, where each item contains the details of a connector.
	//
	// This member is required.
	Connectors []types.ListedConnector

	// Returns a token that you can use to call ListConnectors again and receive
	// additional results, if there are any.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListConnectorsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpListConnectors{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpListConnectors{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListConnectors(options.Region), middleware.Before); err != nil {
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

// ListConnectorsAPIClient is a client that implements the ListConnectors
// operation.
type ListConnectorsAPIClient interface {
	ListConnectors(context.Context, *ListConnectorsInput, ...func(*Options)) (*ListConnectorsOutput, error)
}

var _ ListConnectorsAPIClient = (*Client)(nil)

// ListConnectorsPaginatorOptions is the paginator options for ListConnectors
type ListConnectorsPaginatorOptions struct {
	// The maximum number of connectors to return.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListConnectorsPaginator is a paginator for ListConnectors
type ListConnectorsPaginator struct {
	options   ListConnectorsPaginatorOptions
	client    ListConnectorsAPIClient
	params    *ListConnectorsInput
	nextToken *string
	firstPage bool
}

// NewListConnectorsPaginator returns a new ListConnectorsPaginator
func NewListConnectorsPaginator(client ListConnectorsAPIClient, params *ListConnectorsInput, optFns ...func(*ListConnectorsPaginatorOptions)) *ListConnectorsPaginator {
	if params == nil {
		params = &ListConnectorsInput{}
	}

	options := ListConnectorsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListConnectorsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListConnectorsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListConnectors page.
func (p *ListConnectorsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListConnectorsOutput, error) {
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

	result, err := p.client.ListConnectors(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListConnectors(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "transfer",
		OperationName: "ListConnectors",
	}
}
