// Code generated by smithy-go-codegen DO NOT EDIT.

package inspector

import (
	"context"
	"fmt"
	awsmiddleware "oss-sdk-go/oss/middleware"
	"oss-sdk-go/oss/signer/v4"
	"oss-sdk-go/service/inspector/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Lists all the event subscriptions for the assessment template that is specified
// by the ARN of the assessment template. For more information, see
// SubscribeToEvent and UnsubscribeFromEvent.
func (c *Client) ListEventSubscriptions(ctx context.Context, params *ListEventSubscriptionsInput, optFns ...func(*Options)) (*ListEventSubscriptionsOutput, error) {
	if params == nil {
		params = &ListEventSubscriptionsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListEventSubscriptions", params, optFns, c.addOperationListEventSubscriptionsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListEventSubscriptionsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListEventSubscriptionsInput struct {

	// You can use this parameter to indicate the maximum number of items you want in
	// the response. The default value is 10. The maximum value is 500.
	MaxResults *int32

	// You can use this parameter when paginating results. Set the value of this
	// parameter to null on your first call to the ListEventSubscriptions action.
	// Subsequent calls to the action fill nextToken in the request with the value of
	// NextToken from the previous response to continue listing data.
	NextToken *string

	// The ARN of the assessment template for which you want to list the existing event
	// subscriptions.
	ResourceArn *string

	noSmithyDocumentSerde
}

type ListEventSubscriptionsOutput struct {

	// Details of the returned event subscriptions.
	//
	// This member is required.
	Subscriptions []types.Subscription

	// When a response is generated, if there is more data to be listed, this parameter
	// is present in the response and contains the value to use for the nextToken
	// parameter in a subsequent pagination request. If there is no more data to be
	// listed, this parameter is set to null.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListEventSubscriptionsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpListEventSubscriptions{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpListEventSubscriptions{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListEventSubscriptions(options.Region), middleware.Before); err != nil {
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

// ListEventSubscriptionsAPIClient is a client that implements the
// ListEventSubscriptions operation.
type ListEventSubscriptionsAPIClient interface {
	ListEventSubscriptions(context.Context, *ListEventSubscriptionsInput, ...func(*Options)) (*ListEventSubscriptionsOutput, error)
}

var _ ListEventSubscriptionsAPIClient = (*Client)(nil)

// ListEventSubscriptionsPaginatorOptions is the paginator options for
// ListEventSubscriptions
type ListEventSubscriptionsPaginatorOptions struct {
	// You can use this parameter to indicate the maximum number of items you want in
	// the response. The default value is 10. The maximum value is 500.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListEventSubscriptionsPaginator is a paginator for ListEventSubscriptions
type ListEventSubscriptionsPaginator struct {
	options   ListEventSubscriptionsPaginatorOptions
	client    ListEventSubscriptionsAPIClient
	params    *ListEventSubscriptionsInput
	nextToken *string
	firstPage bool
}

// NewListEventSubscriptionsPaginator returns a new ListEventSubscriptionsPaginator
func NewListEventSubscriptionsPaginator(client ListEventSubscriptionsAPIClient, params *ListEventSubscriptionsInput, optFns ...func(*ListEventSubscriptionsPaginatorOptions)) *ListEventSubscriptionsPaginator {
	if params == nil {
		params = &ListEventSubscriptionsInput{}
	}

	options := ListEventSubscriptionsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListEventSubscriptionsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListEventSubscriptionsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListEventSubscriptions page.
func (p *ListEventSubscriptionsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListEventSubscriptionsOutput, error) {
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

	result, err := p.client.ListEventSubscriptions(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListEventSubscriptions(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "inspector",
		OperationName: "ListEventSubscriptions",
	}
}
