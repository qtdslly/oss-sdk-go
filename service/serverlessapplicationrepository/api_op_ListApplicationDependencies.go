// Code generated by smithy-go-codegen DO NOT EDIT.

package serverlessapplicationrepository

import (
	"context"
	"fmt"
	awsmiddleware "oss-sdk-go/oss/middleware"
	"oss-sdk-go/oss/signer/v4"
	"oss-sdk-go/service/serverlessapplicationrepository/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Retrieves the list of applications nested in the containing application.
func (c *Client) ListApplicationDependencies(ctx context.Context, params *ListApplicationDependenciesInput, optFns ...func(*Options)) (*ListApplicationDependenciesOutput, error) {
	if params == nil {
		params = &ListApplicationDependenciesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListApplicationDependencies", params, optFns, c.addOperationListApplicationDependenciesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListApplicationDependenciesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListApplicationDependenciesInput struct {

	// The Amazon Resource Name (ARN) of the application.
	//
	// This member is required.
	ApplicationId *string

	// The total number of items to return.
	MaxItems int32

	// A token to specify where to start paginating.
	NextToken *string

	// The semantic version of the application to get.
	SemanticVersion *string

	noSmithyDocumentSerde
}

type ListApplicationDependenciesOutput struct {

	// An array of application summaries nested in the application.
	Dependencies []types.ApplicationDependencySummary

	// The token to request the next page of results.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListApplicationDependenciesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListApplicationDependencies{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListApplicationDependencies{}, middleware.After)
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
	if err = addOpListApplicationDependenciesValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListApplicationDependencies(options.Region), middleware.Before); err != nil {
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

// ListApplicationDependenciesAPIClient is a client that implements the
// ListApplicationDependencies operation.
type ListApplicationDependenciesAPIClient interface {
	ListApplicationDependencies(context.Context, *ListApplicationDependenciesInput, ...func(*Options)) (*ListApplicationDependenciesOutput, error)
}

var _ ListApplicationDependenciesAPIClient = (*Client)(nil)

// ListApplicationDependenciesPaginatorOptions is the paginator options for
// ListApplicationDependencies
type ListApplicationDependenciesPaginatorOptions struct {
	// The total number of items to return.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListApplicationDependenciesPaginator is a paginator for
// ListApplicationDependencies
type ListApplicationDependenciesPaginator struct {
	options   ListApplicationDependenciesPaginatorOptions
	client    ListApplicationDependenciesAPIClient
	params    *ListApplicationDependenciesInput
	nextToken *string
	firstPage bool
}

// NewListApplicationDependenciesPaginator returns a new
// ListApplicationDependenciesPaginator
func NewListApplicationDependenciesPaginator(client ListApplicationDependenciesAPIClient, params *ListApplicationDependenciesInput, optFns ...func(*ListApplicationDependenciesPaginatorOptions)) *ListApplicationDependenciesPaginator {
	if params == nil {
		params = &ListApplicationDependenciesInput{}
	}

	options := ListApplicationDependenciesPaginatorOptions{}
	if params.MaxItems != 0 {
		options.Limit = params.MaxItems
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListApplicationDependenciesPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListApplicationDependenciesPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListApplicationDependencies page.
func (p *ListApplicationDependenciesPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListApplicationDependenciesOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.NextToken = p.nextToken

	params.MaxItems = p.options.Limit

	result, err := p.client.ListApplicationDependencies(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListApplicationDependencies(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "serverlessrepo",
		OperationName: "ListApplicationDependencies",
	}
}
