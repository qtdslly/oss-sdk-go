// Code generated by smithy-go-codegen DO NOT EDIT.

package managedblockchain

import (
	"context"
	"fmt"
	awsmiddleware "oss-sdk-go/oss/middleware"
	"oss-sdk-go/oss/signer/v4"
	"oss-sdk-go/service/managedblockchain/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Returns information about the nodes within a network. Applies to Hyperledger
// Fabric and Ethereum.
func (c *Client) ListNodes(ctx context.Context, params *ListNodesInput, optFns ...func(*Options)) (*ListNodesOutput, error) {
	if params == nil {
		params = &ListNodesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListNodes", params, optFns, c.addOperationListNodesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListNodesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListNodesInput struct {

	// The unique identifier of the network for which to list nodes.
	//
	// This member is required.
	NetworkId *string

	// The maximum number of nodes to list.
	MaxResults *int32

	// The unique identifier of the member who owns the nodes to list. Applies only to
	// Hyperledger Fabric and is required for Hyperledger Fabric.
	MemberId *string

	// The pagination token that indicates the next set of results to retrieve.
	NextToken *string

	// An optional status specifier. If provided, only nodes currently in this status
	// are listed.
	Status types.NodeStatus

	noSmithyDocumentSerde
}

type ListNodesOutput struct {

	// The pagination token that indicates the next set of results to retrieve.
	NextToken *string

	// An array of NodeSummary objects that contain configuration properties for each
	// node.
	Nodes []types.NodeSummary

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListNodesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListNodes{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListNodes{}, middleware.After)
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
	if err = addOpListNodesValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListNodes(options.Region), middleware.Before); err != nil {
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

// ListNodesAPIClient is a client that implements the ListNodes operation.
type ListNodesAPIClient interface {
	ListNodes(context.Context, *ListNodesInput, ...func(*Options)) (*ListNodesOutput, error)
}

var _ ListNodesAPIClient = (*Client)(nil)

// ListNodesPaginatorOptions is the paginator options for ListNodes
type ListNodesPaginatorOptions struct {
	// The maximum number of nodes to list.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListNodesPaginator is a paginator for ListNodes
type ListNodesPaginator struct {
	options   ListNodesPaginatorOptions
	client    ListNodesAPIClient
	params    *ListNodesInput
	nextToken *string
	firstPage bool
}

// NewListNodesPaginator returns a new ListNodesPaginator
func NewListNodesPaginator(client ListNodesAPIClient, params *ListNodesInput, optFns ...func(*ListNodesPaginatorOptions)) *ListNodesPaginator {
	if params == nil {
		params = &ListNodesInput{}
	}

	options := ListNodesPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListNodesPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListNodesPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListNodes page.
func (p *ListNodesPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListNodesOutput, error) {
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

	result, err := p.client.ListNodes(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListNodes(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "managedblockchain",
		OperationName: "ListNodes",
	}
}