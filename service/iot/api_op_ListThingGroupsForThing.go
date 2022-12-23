// Code generated by smithy-go-codegen DO NOT EDIT.

package iot

import (
	"context"
	"fmt"
	awsmiddleware "oss-sdk-go/oss/middleware"
	"oss-sdk-go/oss/signer/v4"
	"oss-sdk-go/service/iot/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// List the thing groups to which the specified thing belongs. Requires permission
// to access the ListThingGroupsForThing
// (https://docs.aws.amazon.com/service-authorization/latest/reference/list_awsiot.html#awsiot-actions-as-permissions)
// action.
func (c *Client) ListThingGroupsForThing(ctx context.Context, params *ListThingGroupsForThingInput, optFns ...func(*Options)) (*ListThingGroupsForThingOutput, error) {
	if params == nil {
		params = &ListThingGroupsForThingInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListThingGroupsForThing", params, optFns, c.addOperationListThingGroupsForThingMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListThingGroupsForThingOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListThingGroupsForThingInput struct {

	// The thing name.
	//
	// This member is required.
	ThingName *string

	// The maximum number of results to return at one time.
	MaxResults *int32

	// To retrieve the next set of results, the nextToken value from a previous
	// response; otherwise null to receive the first set of results.
	NextToken *string

	noSmithyDocumentSerde
}

type ListThingGroupsForThingOutput struct {

	// The token to use to get the next set of results, or null if there are no
	// additional results.
	NextToken *string

	// The thing groups.
	ThingGroups []types.GroupNameAndArn

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListThingGroupsForThingMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListThingGroupsForThing{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListThingGroupsForThing{}, middleware.After)
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
	if err = addOpListThingGroupsForThingValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListThingGroupsForThing(options.Region), middleware.Before); err != nil {
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

// ListThingGroupsForThingAPIClient is a client that implements the
// ListThingGroupsForThing operation.
type ListThingGroupsForThingAPIClient interface {
	ListThingGroupsForThing(context.Context, *ListThingGroupsForThingInput, ...func(*Options)) (*ListThingGroupsForThingOutput, error)
}

var _ ListThingGroupsForThingAPIClient = (*Client)(nil)

// ListThingGroupsForThingPaginatorOptions is the paginator options for
// ListThingGroupsForThing
type ListThingGroupsForThingPaginatorOptions struct {
	// The maximum number of results to return at one time.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListThingGroupsForThingPaginator is a paginator for ListThingGroupsForThing
type ListThingGroupsForThingPaginator struct {
	options   ListThingGroupsForThingPaginatorOptions
	client    ListThingGroupsForThingAPIClient
	params    *ListThingGroupsForThingInput
	nextToken *string
	firstPage bool
}

// NewListThingGroupsForThingPaginator returns a new
// ListThingGroupsForThingPaginator
func NewListThingGroupsForThingPaginator(client ListThingGroupsForThingAPIClient, params *ListThingGroupsForThingInput, optFns ...func(*ListThingGroupsForThingPaginatorOptions)) *ListThingGroupsForThingPaginator {
	if params == nil {
		params = &ListThingGroupsForThingInput{}
	}

	options := ListThingGroupsForThingPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListThingGroupsForThingPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListThingGroupsForThingPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListThingGroupsForThing page.
func (p *ListThingGroupsForThingPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListThingGroupsForThingOutput, error) {
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

	result, err := p.client.ListThingGroupsForThing(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListThingGroupsForThing(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "execute-api",
		OperationName: "ListThingGroupsForThing",
	}
}
