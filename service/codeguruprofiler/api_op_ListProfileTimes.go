// Code generated by smithy-go-codegen DO NOT EDIT.

package codeguruprofiler

import (
	"context"
	"fmt"
	awsmiddleware "oss-sdk-go/oss/middleware"
	"oss-sdk-go/oss/signer/v4"
	"oss-sdk-go/service/codeguruprofiler/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Lists the start times of the available aggregated profiles of a profiling group
// for an aggregation period within the specified time range.
func (c *Client) ListProfileTimes(ctx context.Context, params *ListProfileTimesInput, optFns ...func(*Options)) (*ListProfileTimesOutput, error) {
	if params == nil {
		params = &ListProfileTimesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListProfileTimes", params, optFns, c.addOperationListProfileTimesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListProfileTimesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// The structure representing the listProfileTimesRequest.
type ListProfileTimesInput struct {

	// The end time of the time range from which to list the profiles.
	//
	// This member is required.
	EndTime *time.Time

	// The aggregation period. This specifies the period during which an aggregation
	// profile collects posted agent profiles for a profiling group. There are 3 valid
	// values.
	//
	// * P1D — 1 day
	//
	// * PT1H — 1 hour
	//
	// * PT5M — 5 minutes
	//
	// This member is required.
	Period types.AggregationPeriod

	// The name of the profiling group.
	//
	// This member is required.
	ProfilingGroupName *string

	// The start time of the time range from which to list the profiles.
	//
	// This member is required.
	StartTime *time.Time

	// The maximum number of profile time results returned by ListProfileTimes in
	// paginated output. When this parameter is used, ListProfileTimes only returns
	// maxResults results in a single page with a nextToken response element. The
	// remaining results of the initial request can be seen by sending another
	// ListProfileTimes request with the returned nextToken value.
	MaxResults *int32

	// The nextToken value returned from a previous paginated ListProfileTimes request
	// where maxResults was used and the results exceeded the value of that parameter.
	// Pagination continues from the end of the previous results that returned the
	// nextToken value. This token should be treated as an opaque identifier that is
	// only used to retrieve the next items in a list and not for other programmatic
	// purposes.
	NextToken *string

	// The order (ascending or descending by start time of the profile) to use when
	// listing profiles. Defaults to TIMESTAMP_DESCENDING.
	OrderBy types.OrderBy

	noSmithyDocumentSerde
}

// The structure representing the listProfileTimesResponse.
type ListProfileTimesOutput struct {

	// The list of start times of the available profiles for the aggregation period in
	// the specified time range.
	//
	// This member is required.
	ProfileTimes []types.ProfileTime

	// The nextToken value to include in a future ListProfileTimes request. When the
	// results of a ListProfileTimes request exceed maxResults, this value can be used
	// to retrieve the next page of results. This value is null when there are no more
	// results to return.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListProfileTimesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListProfileTimes{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListProfileTimes{}, middleware.After)
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
	if err = addOpListProfileTimesValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListProfileTimes(options.Region), middleware.Before); err != nil {
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

// ListProfileTimesAPIClient is a client that implements the ListProfileTimes
// operation.
type ListProfileTimesAPIClient interface {
	ListProfileTimes(context.Context, *ListProfileTimesInput, ...func(*Options)) (*ListProfileTimesOutput, error)
}

var _ ListProfileTimesAPIClient = (*Client)(nil)

// ListProfileTimesPaginatorOptions is the paginator options for ListProfileTimes
type ListProfileTimesPaginatorOptions struct {
	// The maximum number of profile time results returned by ListProfileTimes in
	// paginated output. When this parameter is used, ListProfileTimes only returns
	// maxResults results in a single page with a nextToken response element. The
	// remaining results of the initial request can be seen by sending another
	// ListProfileTimes request with the returned nextToken value.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListProfileTimesPaginator is a paginator for ListProfileTimes
type ListProfileTimesPaginator struct {
	options   ListProfileTimesPaginatorOptions
	client    ListProfileTimesAPIClient
	params    *ListProfileTimesInput
	nextToken *string
	firstPage bool
}

// NewListProfileTimesPaginator returns a new ListProfileTimesPaginator
func NewListProfileTimesPaginator(client ListProfileTimesAPIClient, params *ListProfileTimesInput, optFns ...func(*ListProfileTimesPaginatorOptions)) *ListProfileTimesPaginator {
	if params == nil {
		params = &ListProfileTimesInput{}
	}

	options := ListProfileTimesPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListProfileTimesPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListProfileTimesPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListProfileTimes page.
func (p *ListProfileTimesPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListProfileTimesOutput, error) {
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

	result, err := p.client.ListProfileTimes(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListProfileTimes(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "codeguru-profiler",
		OperationName: "ListProfileTimes",
	}
}
