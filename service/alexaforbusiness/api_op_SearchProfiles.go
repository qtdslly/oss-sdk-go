// Code generated by smithy-go-codegen DO NOT EDIT.

package alexaforbusiness

import (
	"context"
	"fmt"
	awsmiddleware "oss-sdk-go/oss/middleware"
	"oss-sdk-go/oss/signer/v4"
	"oss-sdk-go/service/alexaforbusiness/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Searches room profiles and lists the ones that meet a set of filter criteria.
func (c *Client) SearchProfiles(ctx context.Context, params *SearchProfilesInput, optFns ...func(*Options)) (*SearchProfilesOutput, error) {
	if params == nil {
		params = &SearchProfilesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "SearchProfiles", params, optFns, c.addOperationSearchProfilesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*SearchProfilesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type SearchProfilesInput struct {

	// The filters to use to list a specified set of room profiles. Supported filter
	// keys are ProfileName and Address. Required.
	Filters []types.Filter

	// The maximum number of results to include in the response. If more results exist
	// than the specified MaxResults value, a token is included in the response so that
	// the remaining results can be retrieved.
	MaxResults *int32

	// An optional token returned from a prior request. Use this token for pagination
	// of results from this action. If this parameter is specified, the response
	// includes only results beyond the token, up to the value specified by MaxResults.
	NextToken *string

	// The sort order to use in listing the specified set of room profiles. Supported
	// sort keys are ProfileName and Address.
	SortCriteria []types.Sort

	noSmithyDocumentSerde
}

type SearchProfilesOutput struct {

	// The token returned to indicate that there is more data available.
	NextToken *string

	// The profiles that meet the specified set of filter criteria, in sort order.
	Profiles []types.ProfileData

	// The total number of room profiles returned.
	TotalCount *int32

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationSearchProfilesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpSearchProfiles{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpSearchProfiles{}, middleware.After)
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
	if err = addOpSearchProfilesValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opSearchProfiles(options.Region), middleware.Before); err != nil {
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

// SearchProfilesAPIClient is a client that implements the SearchProfiles
// operation.
type SearchProfilesAPIClient interface {
	SearchProfiles(context.Context, *SearchProfilesInput, ...func(*Options)) (*SearchProfilesOutput, error)
}

var _ SearchProfilesAPIClient = (*Client)(nil)

// SearchProfilesPaginatorOptions is the paginator options for SearchProfiles
type SearchProfilesPaginatorOptions struct {
	// The maximum number of results to include in the response. If more results exist
	// than the specified MaxResults value, a token is included in the response so that
	// the remaining results can be retrieved.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// SearchProfilesPaginator is a paginator for SearchProfiles
type SearchProfilesPaginator struct {
	options   SearchProfilesPaginatorOptions
	client    SearchProfilesAPIClient
	params    *SearchProfilesInput
	nextToken *string
	firstPage bool
}

// NewSearchProfilesPaginator returns a new SearchProfilesPaginator
func NewSearchProfilesPaginator(client SearchProfilesAPIClient, params *SearchProfilesInput, optFns ...func(*SearchProfilesPaginatorOptions)) *SearchProfilesPaginator {
	if params == nil {
		params = &SearchProfilesInput{}
	}

	options := SearchProfilesPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &SearchProfilesPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *SearchProfilesPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next SearchProfiles page.
func (p *SearchProfilesPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*SearchProfilesOutput, error) {
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

	result, err := p.client.SearchProfiles(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opSearchProfiles(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "a4b",
		OperationName: "SearchProfiles",
	}
}