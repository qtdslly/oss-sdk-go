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
	"time"
)

// Lists the Studio Lifecycle Configurations in your Amazon Web Services Account.
func (c *Client) ListStudioLifecycleConfigs(ctx context.Context, params *ListStudioLifecycleConfigsInput, optFns ...func(*Options)) (*ListStudioLifecycleConfigsOutput, error) {
	if params == nil {
		params = &ListStudioLifecycleConfigsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListStudioLifecycleConfigs", params, optFns, c.addOperationListStudioLifecycleConfigsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListStudioLifecycleConfigsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListStudioLifecycleConfigsInput struct {

	// A parameter to search for the App Type to which the Lifecycle Configuration is
	// attached.
	AppTypeEquals types.StudioLifecycleConfigAppType

	// A filter that returns only Lifecycle Configurations created on or after the
	// specified time.
	CreationTimeAfter *time.Time

	// A filter that returns only Lifecycle Configurations created on or before the
	// specified time.
	CreationTimeBefore *time.Time

	// The maximum number of Studio Lifecycle Configurations to return in the response.
	// The default value is 10.
	MaxResults *int32

	// A filter that returns only Lifecycle Configurations modified after the specified
	// time.
	ModifiedTimeAfter *time.Time

	// A filter that returns only Lifecycle Configurations modified before the
	// specified time.
	ModifiedTimeBefore *time.Time

	// A string in the Lifecycle Configuration name. This filter returns only Lifecycle
	// Configurations whose name contains the specified string.
	NameContains *string

	// If the previous call to ListStudioLifecycleConfigs didn't return the full set of
	// Lifecycle Configurations, the call returns a token for getting the next set of
	// Lifecycle Configurations.
	NextToken *string

	// The property used to sort results. The default value is CreationTime.
	SortBy types.StudioLifecycleConfigSortKey

	// The sort order. The default value is Descending.
	SortOrder types.SortOrder

	noSmithyDocumentSerde
}

type ListStudioLifecycleConfigsOutput struct {

	// A token for getting the next set of actions, if there are any.
	NextToken *string

	// A list of Lifecycle Configurations and their properties.
	StudioLifecycleConfigs []types.StudioLifecycleConfigDetails

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListStudioLifecycleConfigsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpListStudioLifecycleConfigs{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpListStudioLifecycleConfigs{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListStudioLifecycleConfigs(options.Region), middleware.Before); err != nil {
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

// ListStudioLifecycleConfigsAPIClient is a client that implements the
// ListStudioLifecycleConfigs operation.
type ListStudioLifecycleConfigsAPIClient interface {
	ListStudioLifecycleConfigs(context.Context, *ListStudioLifecycleConfigsInput, ...func(*Options)) (*ListStudioLifecycleConfigsOutput, error)
}

var _ ListStudioLifecycleConfigsAPIClient = (*Client)(nil)

// ListStudioLifecycleConfigsPaginatorOptions is the paginator options for
// ListStudioLifecycleConfigs
type ListStudioLifecycleConfigsPaginatorOptions struct {
	// The maximum number of Studio Lifecycle Configurations to return in the response.
	// The default value is 10.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListStudioLifecycleConfigsPaginator is a paginator for
// ListStudioLifecycleConfigs
type ListStudioLifecycleConfigsPaginator struct {
	options   ListStudioLifecycleConfigsPaginatorOptions
	client    ListStudioLifecycleConfigsAPIClient
	params    *ListStudioLifecycleConfigsInput
	nextToken *string
	firstPage bool
}

// NewListStudioLifecycleConfigsPaginator returns a new
// ListStudioLifecycleConfigsPaginator
func NewListStudioLifecycleConfigsPaginator(client ListStudioLifecycleConfigsAPIClient, params *ListStudioLifecycleConfigsInput, optFns ...func(*ListStudioLifecycleConfigsPaginatorOptions)) *ListStudioLifecycleConfigsPaginator {
	if params == nil {
		params = &ListStudioLifecycleConfigsInput{}
	}

	options := ListStudioLifecycleConfigsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListStudioLifecycleConfigsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListStudioLifecycleConfigsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListStudioLifecycleConfigs page.
func (p *ListStudioLifecycleConfigsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListStudioLifecycleConfigsOutput, error) {
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

	result, err := p.client.ListStudioLifecycleConfigs(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListStudioLifecycleConfigs(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "sagemaker",
		OperationName: "ListStudioLifecycleConfigs",
	}
}