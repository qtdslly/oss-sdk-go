// Code generated by smithy-go-codegen DO NOT EDIT.

package appmesh

import (
	"context"
	"fmt"
	awsmiddleware "oss-sdk-go/oss/middleware"
	"oss-sdk-go/oss/signer/v4"
	"oss-sdk-go/service/appmesh/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Returns a list of existing gateway routes that are associated to a virtual
// gateway.
func (c *Client) ListGatewayRoutes(ctx context.Context, params *ListGatewayRoutesInput, optFns ...func(*Options)) (*ListGatewayRoutesOutput, error) {
	if params == nil {
		params = &ListGatewayRoutesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListGatewayRoutes", params, optFns, c.addOperationListGatewayRoutesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListGatewayRoutesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListGatewayRoutesInput struct {

	// The name of the service mesh to list gateway routes in.
	//
	// This member is required.
	MeshName *string

	// The name of the virtual gateway to list gateway routes in.
	//
	// This member is required.
	VirtualGatewayName *string

	// The maximum number of results returned by ListGatewayRoutes in paginated output.
	// When you use this parameter, ListGatewayRoutes returns only limit results in a
	// single page along with a nextToken response element. You can see the remaining
	// results of the initial request by sending another ListGatewayRoutes request with
	// the returned nextToken value. This value can be between 1 and 100. If you don't
	// use this parameter, ListGatewayRoutes returns up to 100 results and a nextToken
	// value if applicable.
	Limit *int32

	// The Amazon Web Services IAM account ID of the service mesh owner. If the account
	// ID is not your own, then it's the ID of the account that shared the mesh with
	// your account. For more information about mesh sharing, see Working with shared
	// meshes (https://docs.aws.amazon.com/app-mesh/latest/userguide/sharing.html).
	MeshOwner *string

	// The nextToken value returned from a previous paginated ListGatewayRoutes request
	// where limit was used and the results exceeded the value of that parameter.
	// Pagination continues from the end of the previous results that returned the
	// nextToken value.
	NextToken *string

	noSmithyDocumentSerde
}

type ListGatewayRoutesOutput struct {

	// The list of existing gateway routes for the specified service mesh and virtual
	// gateway.
	//
	// This member is required.
	GatewayRoutes []types.GatewayRouteRef

	// The nextToken value to include in a future ListGatewayRoutes request. When the
	// results of a ListGatewayRoutes request exceed limit, you can use this value to
	// retrieve the next page of results. This value is null when there are no more
	// results to return.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListGatewayRoutesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListGatewayRoutes{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListGatewayRoutes{}, middleware.After)
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
	if err = addOpListGatewayRoutesValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListGatewayRoutes(options.Region), middleware.Before); err != nil {
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

// ListGatewayRoutesAPIClient is a client that implements the ListGatewayRoutes
// operation.
type ListGatewayRoutesAPIClient interface {
	ListGatewayRoutes(context.Context, *ListGatewayRoutesInput, ...func(*Options)) (*ListGatewayRoutesOutput, error)
}

var _ ListGatewayRoutesAPIClient = (*Client)(nil)

// ListGatewayRoutesPaginatorOptions is the paginator options for ListGatewayRoutes
type ListGatewayRoutesPaginatorOptions struct {
	// The maximum number of results returned by ListGatewayRoutes in paginated output.
	// When you use this parameter, ListGatewayRoutes returns only limit results in a
	// single page along with a nextToken response element. You can see the remaining
	// results of the initial request by sending another ListGatewayRoutes request with
	// the returned nextToken value. This value can be between 1 and 100. If you don't
	// use this parameter, ListGatewayRoutes returns up to 100 results and a nextToken
	// value if applicable.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListGatewayRoutesPaginator is a paginator for ListGatewayRoutes
type ListGatewayRoutesPaginator struct {
	options   ListGatewayRoutesPaginatorOptions
	client    ListGatewayRoutesAPIClient
	params    *ListGatewayRoutesInput
	nextToken *string
	firstPage bool
}

// NewListGatewayRoutesPaginator returns a new ListGatewayRoutesPaginator
func NewListGatewayRoutesPaginator(client ListGatewayRoutesAPIClient, params *ListGatewayRoutesInput, optFns ...func(*ListGatewayRoutesPaginatorOptions)) *ListGatewayRoutesPaginator {
	if params == nil {
		params = &ListGatewayRoutesInput{}
	}

	options := ListGatewayRoutesPaginatorOptions{}
	if params.Limit != nil {
		options.Limit = *params.Limit
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListGatewayRoutesPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListGatewayRoutesPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListGatewayRoutes page.
func (p *ListGatewayRoutesPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListGatewayRoutesOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.NextToken = p.nextToken

	var limit *int32
	if p.options.Limit > 0 {
		limit = &p.options.Limit
	}
	params.Limit = limit

	result, err := p.client.ListGatewayRoutes(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListGatewayRoutes(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "appmesh",
		OperationName: "ListGatewayRoutes",
	}
}
