// Code generated by smithy-go-codegen DO NOT EDIT.

package connect

import (
	"context"
	"fmt"
	awsmiddleware "oss-sdk-go/oss/middleware"
	"oss-sdk-go/oss/signer/v4"
	"oss-sdk-go/service/connect/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// This API is in preview release for Amazon Connect and is subject to change.
// Lists agent statuses.
func (c *Client) ListAgentStatuses(ctx context.Context, params *ListAgentStatusesInput, optFns ...func(*Options)) (*ListAgentStatusesOutput, error) {
	if params == nil {
		params = &ListAgentStatusesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListAgentStatuses", params, optFns, c.addOperationListAgentStatusesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListAgentStatusesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListAgentStatusesInput struct {

	// The identifier of the Amazon Connect instance. You can find the instanceId in
	// the ARN of the instance.
	//
	// This member is required.
	InstanceId *string

	// Available agent status types.
	AgentStatusTypes []types.AgentStatusType

	// The maximum number of results to return per page.
	MaxResults *int32

	// The token for the next set of results. Use the value returned in the previous
	// response in the next request to retrieve the next set of results.
	NextToken *string

	noSmithyDocumentSerde
}

type ListAgentStatusesOutput struct {

	// A summary of agent statuses.
	AgentStatusSummaryList []types.AgentStatusSummary

	// If there are additional results, this is the token for the next set of results.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListAgentStatusesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListAgentStatuses{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListAgentStatuses{}, middleware.After)
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
	if err = addOpListAgentStatusesValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListAgentStatuses(options.Region), middleware.Before); err != nil {
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

// ListAgentStatusesAPIClient is a client that implements the ListAgentStatuses
// operation.
type ListAgentStatusesAPIClient interface {
	ListAgentStatuses(context.Context, *ListAgentStatusesInput, ...func(*Options)) (*ListAgentStatusesOutput, error)
}

var _ ListAgentStatusesAPIClient = (*Client)(nil)

// ListAgentStatusesPaginatorOptions is the paginator options for ListAgentStatuses
type ListAgentStatusesPaginatorOptions struct {
	// The maximum number of results to return per page.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListAgentStatusesPaginator is a paginator for ListAgentStatuses
type ListAgentStatusesPaginator struct {
	options   ListAgentStatusesPaginatorOptions
	client    ListAgentStatusesAPIClient
	params    *ListAgentStatusesInput
	nextToken *string
	firstPage bool
}

// NewListAgentStatusesPaginator returns a new ListAgentStatusesPaginator
func NewListAgentStatusesPaginator(client ListAgentStatusesAPIClient, params *ListAgentStatusesInput, optFns ...func(*ListAgentStatusesPaginatorOptions)) *ListAgentStatusesPaginator {
	if params == nil {
		params = &ListAgentStatusesInput{}
	}

	options := ListAgentStatusesPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListAgentStatusesPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListAgentStatusesPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListAgentStatuses page.
func (p *ListAgentStatusesPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListAgentStatusesOutput, error) {
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

	result, err := p.client.ListAgentStatuses(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListAgentStatuses(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "connect",
		OperationName: "ListAgentStatuses",
	}
}
