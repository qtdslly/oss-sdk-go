// Code generated by smithy-go-codegen DO NOT EDIT.

package guardduty

import (
	"context"
	"fmt"
	awsmiddleware "oss-sdk-go/oss/middleware"
	"oss-sdk-go/oss/signer/v4"
	"oss-sdk-go/service/guardduty/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Lists the accounts configured as GuardDuty delegated administrators.
func (c *Client) ListOrganizationAdminAccounts(ctx context.Context, params *ListOrganizationAdminAccountsInput, optFns ...func(*Options)) (*ListOrganizationAdminAccountsOutput, error) {
	if params == nil {
		params = &ListOrganizationAdminAccountsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListOrganizationAdminAccounts", params, optFns, c.addOperationListOrganizationAdminAccountsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListOrganizationAdminAccountsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListOrganizationAdminAccountsInput struct {

	// The maximum number of results to return in the response.
	MaxResults int32

	// A token to use for paginating results that are returned in the response. Set the
	// value of this parameter to null for the first request to a list action. For
	// subsequent calls, use the NextToken value returned from the previous request to
	// continue listing results after the first page.
	NextToken *string

	noSmithyDocumentSerde
}

type ListOrganizationAdminAccountsOutput struct {

	// A list of accounts configured as GuardDuty delegated administrators.
	AdminAccounts []types.AdminAccount

	// The pagination parameter to be used on the next list operation to retrieve more
	// items.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListOrganizationAdminAccountsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListOrganizationAdminAccounts{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListOrganizationAdminAccounts{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListOrganizationAdminAccounts(options.Region), middleware.Before); err != nil {
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

// ListOrganizationAdminAccountsAPIClient is a client that implements the
// ListOrganizationAdminAccounts operation.
type ListOrganizationAdminAccountsAPIClient interface {
	ListOrganizationAdminAccounts(context.Context, *ListOrganizationAdminAccountsInput, ...func(*Options)) (*ListOrganizationAdminAccountsOutput, error)
}

var _ ListOrganizationAdminAccountsAPIClient = (*Client)(nil)

// ListOrganizationAdminAccountsPaginatorOptions is the paginator options for
// ListOrganizationAdminAccounts
type ListOrganizationAdminAccountsPaginatorOptions struct {
	// The maximum number of results to return in the response.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListOrganizationAdminAccountsPaginator is a paginator for
// ListOrganizationAdminAccounts
type ListOrganizationAdminAccountsPaginator struct {
	options   ListOrganizationAdminAccountsPaginatorOptions
	client    ListOrganizationAdminAccountsAPIClient
	params    *ListOrganizationAdminAccountsInput
	nextToken *string
	firstPage bool
}

// NewListOrganizationAdminAccountsPaginator returns a new
// ListOrganizationAdminAccountsPaginator
func NewListOrganizationAdminAccountsPaginator(client ListOrganizationAdminAccountsAPIClient, params *ListOrganizationAdminAccountsInput, optFns ...func(*ListOrganizationAdminAccountsPaginatorOptions)) *ListOrganizationAdminAccountsPaginator {
	if params == nil {
		params = &ListOrganizationAdminAccountsInput{}
	}

	options := ListOrganizationAdminAccountsPaginatorOptions{}
	if params.MaxResults != 0 {
		options.Limit = params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListOrganizationAdminAccountsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListOrganizationAdminAccountsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListOrganizationAdminAccounts page.
func (p *ListOrganizationAdminAccountsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListOrganizationAdminAccountsOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.NextToken = p.nextToken

	params.MaxResults = p.options.Limit

	result, err := p.client.ListOrganizationAdminAccounts(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListOrganizationAdminAccounts(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "guardduty",
		OperationName: "ListOrganizationAdminAccounts",
	}
}
