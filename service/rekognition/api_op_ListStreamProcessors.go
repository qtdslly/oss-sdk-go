// Code generated by smithy-go-codegen DO NOT EDIT.

package rekognition

import (
	"context"
	"fmt"
	awsmiddleware "oss-sdk-go/oss/middleware"
	"oss-sdk-go/oss/signer/v4"
	"oss-sdk-go/service/rekognition/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Gets a list of stream processors that you have created with
// CreateStreamProcessor.
func (c *Client) ListStreamProcessors(ctx context.Context, params *ListStreamProcessorsInput, optFns ...func(*Options)) (*ListStreamProcessorsOutput, error) {
	if params == nil {
		params = &ListStreamProcessorsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListStreamProcessors", params, optFns, c.addOperationListStreamProcessorsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListStreamProcessorsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListStreamProcessorsInput struct {

	// Maximum number of stream processors you want Amazon Rekognition Video to return
	// in the response. The default is 1000.
	MaxResults *int32

	// If the previous response was incomplete (because there are more stream
	// processors to retrieve), Amazon Rekognition Video returns a pagination token in
	// the response. You can use this pagination token to retrieve the next set of
	// stream processors.
	NextToken *string

	noSmithyDocumentSerde
}

type ListStreamProcessorsOutput struct {

	// If the response is truncated, Amazon Rekognition Video returns this token that
	// you can use in the subsequent request to retrieve the next set of stream
	// processors.
	NextToken *string

	// List of stream processors that you have created.
	StreamProcessors []types.StreamProcessor

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListStreamProcessorsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpListStreamProcessors{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpListStreamProcessors{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListStreamProcessors(options.Region), middleware.Before); err != nil {
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

// ListStreamProcessorsAPIClient is a client that implements the
// ListStreamProcessors operation.
type ListStreamProcessorsAPIClient interface {
	ListStreamProcessors(context.Context, *ListStreamProcessorsInput, ...func(*Options)) (*ListStreamProcessorsOutput, error)
}

var _ ListStreamProcessorsAPIClient = (*Client)(nil)

// ListStreamProcessorsPaginatorOptions is the paginator options for
// ListStreamProcessors
type ListStreamProcessorsPaginatorOptions struct {
	// Maximum number of stream processors you want Amazon Rekognition Video to return
	// in the response. The default is 1000.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListStreamProcessorsPaginator is a paginator for ListStreamProcessors
type ListStreamProcessorsPaginator struct {
	options   ListStreamProcessorsPaginatorOptions
	client    ListStreamProcessorsAPIClient
	params    *ListStreamProcessorsInput
	nextToken *string
	firstPage bool
}

// NewListStreamProcessorsPaginator returns a new ListStreamProcessorsPaginator
func NewListStreamProcessorsPaginator(client ListStreamProcessorsAPIClient, params *ListStreamProcessorsInput, optFns ...func(*ListStreamProcessorsPaginatorOptions)) *ListStreamProcessorsPaginator {
	if params == nil {
		params = &ListStreamProcessorsInput{}
	}

	options := ListStreamProcessorsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListStreamProcessorsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListStreamProcessorsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListStreamProcessors page.
func (p *ListStreamProcessorsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListStreamProcessorsOutput, error) {
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

	result, err := p.client.ListStreamProcessors(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListStreamProcessors(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "rekognition",
		OperationName: "ListStreamProcessors",
	}
}
