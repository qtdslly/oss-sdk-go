// Code generated by smithy-go-codegen DO NOT EDIT.

package transfer

import (
	"context"
	"fmt"
	awsmiddleware "oss-sdk-go/oss/middleware"
	"oss-sdk-go/oss/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Lists the security policies that are attached to your file transfer
// protocol-enabled servers.
func (c *Client) ListSecurityPolicies(ctx context.Context, params *ListSecurityPoliciesInput, optFns ...func(*Options)) (*ListSecurityPoliciesOutput, error) {
	if params == nil {
		params = &ListSecurityPoliciesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListSecurityPolicies", params, optFns, c.addOperationListSecurityPoliciesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListSecurityPoliciesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListSecurityPoliciesInput struct {

	// Specifies the number of security policies to return as a response to the
	// ListSecurityPolicies query.
	MaxResults *int32

	// When additional results are obtained from the ListSecurityPolicies command, a
	// NextToken parameter is returned in the output. You can then pass the NextToken
	// parameter in a subsequent command to continue listing additional security
	// policies.
	NextToken *string

	noSmithyDocumentSerde
}

type ListSecurityPoliciesOutput struct {

	// An array of security policies that were listed.
	//
	// This member is required.
	SecurityPolicyNames []string

	// When you can get additional results from the ListSecurityPolicies operation, a
	// NextToken parameter is returned in the output. In a following command, you can
	// pass in the NextToken parameter to continue listing security policies.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListSecurityPoliciesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpListSecurityPolicies{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpListSecurityPolicies{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListSecurityPolicies(options.Region), middleware.Before); err != nil {
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

// ListSecurityPoliciesAPIClient is a client that implements the
// ListSecurityPolicies operation.
type ListSecurityPoliciesAPIClient interface {
	ListSecurityPolicies(context.Context, *ListSecurityPoliciesInput, ...func(*Options)) (*ListSecurityPoliciesOutput, error)
}

var _ ListSecurityPoliciesAPIClient = (*Client)(nil)

// ListSecurityPoliciesPaginatorOptions is the paginator options for
// ListSecurityPolicies
type ListSecurityPoliciesPaginatorOptions struct {
	// Specifies the number of security policies to return as a response to the
	// ListSecurityPolicies query.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListSecurityPoliciesPaginator is a paginator for ListSecurityPolicies
type ListSecurityPoliciesPaginator struct {
	options   ListSecurityPoliciesPaginatorOptions
	client    ListSecurityPoliciesAPIClient
	params    *ListSecurityPoliciesInput
	nextToken *string
	firstPage bool
}

// NewListSecurityPoliciesPaginator returns a new ListSecurityPoliciesPaginator
func NewListSecurityPoliciesPaginator(client ListSecurityPoliciesAPIClient, params *ListSecurityPoliciesInput, optFns ...func(*ListSecurityPoliciesPaginatorOptions)) *ListSecurityPoliciesPaginator {
	if params == nil {
		params = &ListSecurityPoliciesInput{}
	}

	options := ListSecurityPoliciesPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListSecurityPoliciesPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListSecurityPoliciesPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListSecurityPolicies page.
func (p *ListSecurityPoliciesPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListSecurityPoliciesOutput, error) {
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

	result, err := p.client.ListSecurityPolicies(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListSecurityPolicies(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "transfer",
		OperationName: "ListSecurityPolicies",
	}
}
