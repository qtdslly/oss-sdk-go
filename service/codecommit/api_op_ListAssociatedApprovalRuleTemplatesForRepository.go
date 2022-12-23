// Code generated by smithy-go-codegen DO NOT EDIT.

package codecommit

import (
	"context"
	"fmt"
	awsmiddleware "oss-sdk-go/oss/middleware"
	"oss-sdk-go/oss/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Lists all approval rule templates that are associated with a specified
// repository.
func (c *Client) ListAssociatedApprovalRuleTemplatesForRepository(ctx context.Context, params *ListAssociatedApprovalRuleTemplatesForRepositoryInput, optFns ...func(*Options)) (*ListAssociatedApprovalRuleTemplatesForRepositoryOutput, error) {
	if params == nil {
		params = &ListAssociatedApprovalRuleTemplatesForRepositoryInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListAssociatedApprovalRuleTemplatesForRepository", params, optFns, c.addOperationListAssociatedApprovalRuleTemplatesForRepositoryMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListAssociatedApprovalRuleTemplatesForRepositoryOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListAssociatedApprovalRuleTemplatesForRepositoryInput struct {

	// The name of the repository for which you want to list all associated approval
	// rule templates.
	//
	// This member is required.
	RepositoryName *string

	// A non-zero, non-negative integer used to limit the number of returned results.
	MaxResults *int32

	// An enumeration token that, when provided in a request, returns the next batch of
	// the results.
	NextToken *string

	noSmithyDocumentSerde
}

type ListAssociatedApprovalRuleTemplatesForRepositoryOutput struct {

	// The names of all approval rule templates associated with the repository.
	ApprovalRuleTemplateNames []string

	// An enumeration token that allows the operation to batch the next results of the
	// operation.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListAssociatedApprovalRuleTemplatesForRepositoryMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpListAssociatedApprovalRuleTemplatesForRepository{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpListAssociatedApprovalRuleTemplatesForRepository{}, middleware.After)
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
	if err = addOpListAssociatedApprovalRuleTemplatesForRepositoryValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListAssociatedApprovalRuleTemplatesForRepository(options.Region), middleware.Before); err != nil {
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

// ListAssociatedApprovalRuleTemplatesForRepositoryAPIClient is a client that
// implements the ListAssociatedApprovalRuleTemplatesForRepository operation.
type ListAssociatedApprovalRuleTemplatesForRepositoryAPIClient interface {
	ListAssociatedApprovalRuleTemplatesForRepository(context.Context, *ListAssociatedApprovalRuleTemplatesForRepositoryInput, ...func(*Options)) (*ListAssociatedApprovalRuleTemplatesForRepositoryOutput, error)
}

var _ ListAssociatedApprovalRuleTemplatesForRepositoryAPIClient = (*Client)(nil)

// ListAssociatedApprovalRuleTemplatesForRepositoryPaginatorOptions is the
// paginator options for ListAssociatedApprovalRuleTemplatesForRepository
type ListAssociatedApprovalRuleTemplatesForRepositoryPaginatorOptions struct {
	// A non-zero, non-negative integer used to limit the number of returned results.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListAssociatedApprovalRuleTemplatesForRepositoryPaginator is a paginator for
// ListAssociatedApprovalRuleTemplatesForRepository
type ListAssociatedApprovalRuleTemplatesForRepositoryPaginator struct {
	options   ListAssociatedApprovalRuleTemplatesForRepositoryPaginatorOptions
	client    ListAssociatedApprovalRuleTemplatesForRepositoryAPIClient
	params    *ListAssociatedApprovalRuleTemplatesForRepositoryInput
	nextToken *string
	firstPage bool
}

// NewListAssociatedApprovalRuleTemplatesForRepositoryPaginator returns a new
// ListAssociatedApprovalRuleTemplatesForRepositoryPaginator
func NewListAssociatedApprovalRuleTemplatesForRepositoryPaginator(client ListAssociatedApprovalRuleTemplatesForRepositoryAPIClient, params *ListAssociatedApprovalRuleTemplatesForRepositoryInput, optFns ...func(*ListAssociatedApprovalRuleTemplatesForRepositoryPaginatorOptions)) *ListAssociatedApprovalRuleTemplatesForRepositoryPaginator {
	if params == nil {
		params = &ListAssociatedApprovalRuleTemplatesForRepositoryInput{}
	}

	options := ListAssociatedApprovalRuleTemplatesForRepositoryPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListAssociatedApprovalRuleTemplatesForRepositoryPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListAssociatedApprovalRuleTemplatesForRepositoryPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListAssociatedApprovalRuleTemplatesForRepository
// page.
func (p *ListAssociatedApprovalRuleTemplatesForRepositoryPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListAssociatedApprovalRuleTemplatesForRepositoryOutput, error) {
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

	result, err := p.client.ListAssociatedApprovalRuleTemplatesForRepository(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListAssociatedApprovalRuleTemplatesForRepository(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "codecommit",
		OperationName: "ListAssociatedApprovalRuleTemplatesForRepository",
	}
}
