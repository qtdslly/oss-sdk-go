// Code generated by smithy-go-codegen DO NOT EDIT.

package ssm

import (
	"context"
	"fmt"
	awsmiddleware "oss-sdk-go/oss/middleware"
	"oss-sdk-go/oss/signer/v4"
	"oss-sdk-go/service/ssm/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Describes a specific delete inventory operation.
func (c *Client) DescribeInventoryDeletions(ctx context.Context, params *DescribeInventoryDeletionsInput, optFns ...func(*Options)) (*DescribeInventoryDeletionsOutput, error) {
	if params == nil {
		params = &DescribeInventoryDeletionsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeInventoryDeletions", params, optFns, c.addOperationDescribeInventoryDeletionsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeInventoryDeletionsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeInventoryDeletionsInput struct {

	// Specify the delete inventory ID for which you want information. This ID was
	// returned by the DeleteInventory operation.
	DeletionId *string

	// The maximum number of items to return for this call. The call also returns a
	// token that you can specify in a subsequent call to get the next set of results.
	MaxResults *int32

	// A token to start the list. Use this token to get the next set of results.
	NextToken *string

	noSmithyDocumentSerde
}

type DescribeInventoryDeletionsOutput struct {

	// A list of status items for deleted inventory.
	InventoryDeletions []types.InventoryDeletionStatusItem

	// The token for the next set of items to return. Use this token to get the next
	// set of results.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeInventoryDeletionsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDescribeInventoryDeletions{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDescribeInventoryDeletions{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeInventoryDeletions(options.Region), middleware.Before); err != nil {
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

// DescribeInventoryDeletionsAPIClient is a client that implements the
// DescribeInventoryDeletions operation.
type DescribeInventoryDeletionsAPIClient interface {
	DescribeInventoryDeletions(context.Context, *DescribeInventoryDeletionsInput, ...func(*Options)) (*DescribeInventoryDeletionsOutput, error)
}

var _ DescribeInventoryDeletionsAPIClient = (*Client)(nil)

// DescribeInventoryDeletionsPaginatorOptions is the paginator options for
// DescribeInventoryDeletions
type DescribeInventoryDeletionsPaginatorOptions struct {
	// The maximum number of items to return for this call. The call also returns a
	// token that you can specify in a subsequent call to get the next set of results.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// DescribeInventoryDeletionsPaginator is a paginator for
// DescribeInventoryDeletions
type DescribeInventoryDeletionsPaginator struct {
	options   DescribeInventoryDeletionsPaginatorOptions
	client    DescribeInventoryDeletionsAPIClient
	params    *DescribeInventoryDeletionsInput
	nextToken *string
	firstPage bool
}

// NewDescribeInventoryDeletionsPaginator returns a new
// DescribeInventoryDeletionsPaginator
func NewDescribeInventoryDeletionsPaginator(client DescribeInventoryDeletionsAPIClient, params *DescribeInventoryDeletionsInput, optFns ...func(*DescribeInventoryDeletionsPaginatorOptions)) *DescribeInventoryDeletionsPaginator {
	if params == nil {
		params = &DescribeInventoryDeletionsInput{}
	}

	options := DescribeInventoryDeletionsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &DescribeInventoryDeletionsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *DescribeInventoryDeletionsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next DescribeInventoryDeletions page.
func (p *DescribeInventoryDeletionsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*DescribeInventoryDeletionsOutput, error) {
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

	result, err := p.client.DescribeInventoryDeletions(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opDescribeInventoryDeletions(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ssm",
		OperationName: "DescribeInventoryDeletions",
	}
}
