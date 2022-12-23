// Code generated by smithy-go-codegen DO NOT EDIT.

package sagemaker

import (
	"context"
	"fmt"
	awsmiddleware "oss-sdk-go/oss/middleware"
	"oss-sdk-go/oss/signer/v4"
	"oss-sdk-go/service/sagemaker/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// List the candidates created for the job.
func (c *Client) ListCandidatesForAutoMLJob(ctx context.Context, params *ListCandidatesForAutoMLJobInput, optFns ...func(*Options)) (*ListCandidatesForAutoMLJobOutput, error) {
	if params == nil {
		params = &ListCandidatesForAutoMLJobInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListCandidatesForAutoMLJob", params, optFns, c.addOperationListCandidatesForAutoMLJobMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListCandidatesForAutoMLJobOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListCandidatesForAutoMLJobInput struct {

	// List the candidates created for the job by providing the job's name.
	//
	// This member is required.
	AutoMLJobName *string

	// List the candidates for the job and filter by candidate name.
	CandidateNameEquals *string

	// List the job's candidates up to a specified limit.
	MaxResults *int32

	// If the previous response was truncated, you receive this token. Use it in your
	// next request to receive the next set of results.
	NextToken *string

	// The parameter by which to sort the results. The default is Descending.
	SortBy types.CandidateSortBy

	// The sort order for the results. The default is Ascending.
	SortOrder types.AutoMLSortOrder

	// List the candidates for the job and filter by status.
	StatusEquals types.CandidateStatus

	noSmithyDocumentSerde
}

type ListCandidatesForAutoMLJobOutput struct {

	// Summaries about the AutoMLCandidates.
	//
	// This member is required.
	Candidates []types.AutoMLCandidate

	// If the previous response was truncated, you receive this token. Use it in your
	// next request to receive the next set of results.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListCandidatesForAutoMLJobMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpListCandidatesForAutoMLJob{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpListCandidatesForAutoMLJob{}, middleware.After)
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
	if err = addOpListCandidatesForAutoMLJobValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListCandidatesForAutoMLJob(options.Region), middleware.Before); err != nil {
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

// ListCandidatesForAutoMLJobAPIClient is a client that implements the
// ListCandidatesForAutoMLJob operation.
type ListCandidatesForAutoMLJobAPIClient interface {
	ListCandidatesForAutoMLJob(context.Context, *ListCandidatesForAutoMLJobInput, ...func(*Options)) (*ListCandidatesForAutoMLJobOutput, error)
}

var _ ListCandidatesForAutoMLJobAPIClient = (*Client)(nil)

// ListCandidatesForAutoMLJobPaginatorOptions is the paginator options for
// ListCandidatesForAutoMLJob
type ListCandidatesForAutoMLJobPaginatorOptions struct {
	// List the job's candidates up to a specified limit.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListCandidatesForAutoMLJobPaginator is a paginator for
// ListCandidatesForAutoMLJob
type ListCandidatesForAutoMLJobPaginator struct {
	options   ListCandidatesForAutoMLJobPaginatorOptions
	client    ListCandidatesForAutoMLJobAPIClient
	params    *ListCandidatesForAutoMLJobInput
	nextToken *string
	firstPage bool
}

// NewListCandidatesForAutoMLJobPaginator returns a new
// ListCandidatesForAutoMLJobPaginator
func NewListCandidatesForAutoMLJobPaginator(client ListCandidatesForAutoMLJobAPIClient, params *ListCandidatesForAutoMLJobInput, optFns ...func(*ListCandidatesForAutoMLJobPaginatorOptions)) *ListCandidatesForAutoMLJobPaginator {
	if params == nil {
		params = &ListCandidatesForAutoMLJobInput{}
	}

	options := ListCandidatesForAutoMLJobPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListCandidatesForAutoMLJobPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListCandidatesForAutoMLJobPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListCandidatesForAutoMLJob page.
func (p *ListCandidatesForAutoMLJobPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListCandidatesForAutoMLJobOutput, error) {
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

	result, err := p.client.ListCandidatesForAutoMLJob(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListCandidatesForAutoMLJob(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "sagemaker",
		OperationName: "ListCandidatesForAutoMLJob",
	}
}
