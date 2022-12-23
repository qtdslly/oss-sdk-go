// Code generated by smithy-go-codegen DO NOT EDIT.

package gamesparks

import (
	"context"
	"fmt"
	awsmiddleware "oss-sdk-go/oss/middleware"
	"oss-sdk-go/oss/signer/v4"
	"oss-sdk-go/service/gamesparks/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Gets a paginated list of code generation jobs for a snapshot.
func (c *Client) ListGeneratedCodeJobs(ctx context.Context, params *ListGeneratedCodeJobsInput, optFns ...func(*Options)) (*ListGeneratedCodeJobsOutput, error) {
	if params == nil {
		params = &ListGeneratedCodeJobsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListGeneratedCodeJobs", params, optFns, c.addOperationListGeneratedCodeJobsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListGeneratedCodeJobsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListGeneratedCodeJobsInput struct {

	// The name of the game.
	//
	// This member is required.
	GameName *string

	// The identifier of the snapshot.
	//
	// This member is required.
	SnapshotId *string

	// The maximum number of results to return. Use this parameter with NextToken to
	// get results as a set of sequential pages.
	MaxResults *int32

	// The token that indicates the start of the next sequential page of results. Use
	// the token that is returned with a previous call to this operation. To start at
	// the beginning of the result set, do not specify a value.
	NextToken *string

	noSmithyDocumentSerde
}

type ListGeneratedCodeJobsOutput struct {

	// The list of generated code jobs.
	GeneratedCodeJobs []types.GeneratedCodeJobDetails

	// The token that indicates the start of the next sequential page of results. Use
	// this value when making the next call to this operation to continue where the
	// last one finished.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListGeneratedCodeJobsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListGeneratedCodeJobs{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListGeneratedCodeJobs{}, middleware.After)
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
	if err = addOpListGeneratedCodeJobsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListGeneratedCodeJobs(options.Region), middleware.Before); err != nil {
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

// ListGeneratedCodeJobsAPIClient is a client that implements the
// ListGeneratedCodeJobs operation.
type ListGeneratedCodeJobsAPIClient interface {
	ListGeneratedCodeJobs(context.Context, *ListGeneratedCodeJobsInput, ...func(*Options)) (*ListGeneratedCodeJobsOutput, error)
}

var _ ListGeneratedCodeJobsAPIClient = (*Client)(nil)

// ListGeneratedCodeJobsPaginatorOptions is the paginator options for
// ListGeneratedCodeJobs
type ListGeneratedCodeJobsPaginatorOptions struct {
	// The maximum number of results to return. Use this parameter with NextToken to
	// get results as a set of sequential pages.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListGeneratedCodeJobsPaginator is a paginator for ListGeneratedCodeJobs
type ListGeneratedCodeJobsPaginator struct {
	options   ListGeneratedCodeJobsPaginatorOptions
	client    ListGeneratedCodeJobsAPIClient
	params    *ListGeneratedCodeJobsInput
	nextToken *string
	firstPage bool
}

// NewListGeneratedCodeJobsPaginator returns a new ListGeneratedCodeJobsPaginator
func NewListGeneratedCodeJobsPaginator(client ListGeneratedCodeJobsAPIClient, params *ListGeneratedCodeJobsInput, optFns ...func(*ListGeneratedCodeJobsPaginatorOptions)) *ListGeneratedCodeJobsPaginator {
	if params == nil {
		params = &ListGeneratedCodeJobsInput{}
	}

	options := ListGeneratedCodeJobsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListGeneratedCodeJobsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListGeneratedCodeJobsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListGeneratedCodeJobs page.
func (p *ListGeneratedCodeJobsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListGeneratedCodeJobsOutput, error) {
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

	result, err := p.client.ListGeneratedCodeJobs(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListGeneratedCodeJobs(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "gamesparks",
		OperationName: "ListGeneratedCodeJobs",
	}
}
