// Code generated by smithy-go-codegen DO NOT EDIT.

package ivs

import (
	"context"
	"fmt"
	awsmiddleware "oss-sdk-go/oss/middleware"
	"oss-sdk-go/oss/signer/v4"
	"oss-sdk-go/service/ivs/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Gets a summary of current and previous streams for a specified channel in your
// account, in the AWS region where the API request is processed.
func (c *Client) ListStreamSessions(ctx context.Context, params *ListStreamSessionsInput, optFns ...func(*Options)) (*ListStreamSessionsOutput, error) {
	if params == nil {
		params = &ListStreamSessionsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListStreamSessions", params, optFns, c.addOperationListStreamSessionsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListStreamSessionsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListStreamSessionsInput struct {

	// Channel ARN used to filter the list.
	//
	// This member is required.
	ChannelArn *string

	// Maximum number of streams to return. Default: 100.
	MaxResults int32

	// The first stream to retrieve. This is used for pagination; see the nextToken
	// response field.
	NextToken *string

	noSmithyDocumentSerde
}

type ListStreamSessionsOutput struct {

	// List of stream sessions.
	//
	// This member is required.
	StreamSessions []types.StreamSessionSummary

	// If there are more streams than maxResults, use nextToken in the request to get
	// the next set.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListStreamSessionsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListStreamSessions{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListStreamSessions{}, middleware.After)
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
	if err = addOpListStreamSessionsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListStreamSessions(options.Region), middleware.Before); err != nil {
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

// ListStreamSessionsAPIClient is a client that implements the ListStreamSessions
// operation.
type ListStreamSessionsAPIClient interface {
	ListStreamSessions(context.Context, *ListStreamSessionsInput, ...func(*Options)) (*ListStreamSessionsOutput, error)
}

var _ ListStreamSessionsAPIClient = (*Client)(nil)

// ListStreamSessionsPaginatorOptions is the paginator options for
// ListStreamSessions
type ListStreamSessionsPaginatorOptions struct {
	// Maximum number of streams to return. Default: 100.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListStreamSessionsPaginator is a paginator for ListStreamSessions
type ListStreamSessionsPaginator struct {
	options   ListStreamSessionsPaginatorOptions
	client    ListStreamSessionsAPIClient
	params    *ListStreamSessionsInput
	nextToken *string
	firstPage bool
}

// NewListStreamSessionsPaginator returns a new ListStreamSessionsPaginator
func NewListStreamSessionsPaginator(client ListStreamSessionsAPIClient, params *ListStreamSessionsInput, optFns ...func(*ListStreamSessionsPaginatorOptions)) *ListStreamSessionsPaginator {
	if params == nil {
		params = &ListStreamSessionsInput{}
	}

	options := ListStreamSessionsPaginatorOptions{}
	if params.MaxResults != 0 {
		options.Limit = params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListStreamSessionsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListStreamSessionsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListStreamSessions page.
func (p *ListStreamSessionsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListStreamSessionsOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.NextToken = p.nextToken

	params.MaxResults = p.options.Limit

	result, err := p.client.ListStreamSessions(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListStreamSessions(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ivs",
		OperationName: "ListStreamSessions",
	}
}
