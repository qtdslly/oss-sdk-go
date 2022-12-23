// Code generated by smithy-go-codegen DO NOT EDIT.

package databasemigrationservice

import (
	"context"
	"fmt"
	awsmiddleware "oss-sdk-go/oss/middleware"
	"oss-sdk-go/oss/signer/v4"
	"oss-sdk-go/service/databasemigrationservice/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Returns information about the replication instance types that can be created in
// the specified region.
func (c *Client) DescribeOrderableReplicationInstances(ctx context.Context, params *DescribeOrderableReplicationInstancesInput, optFns ...func(*Options)) (*DescribeOrderableReplicationInstancesOutput, error) {
	if params == nil {
		params = &DescribeOrderableReplicationInstancesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeOrderableReplicationInstances", params, optFns, c.addOperationDescribeOrderableReplicationInstancesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeOrderableReplicationInstancesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeOrderableReplicationInstancesInput struct {

	// An optional pagination token provided by a previous request. If this parameter
	// is specified, the response includes only records beyond the marker, up to the
	// value specified by MaxRecords.
	Marker *string

	// The maximum number of records to include in the response. If more records exist
	// than the specified MaxRecords value, a pagination token called a marker is
	// included in the response so that the remaining results can be retrieved.
	// Default: 100 Constraints: Minimum 20, maximum 100.
	MaxRecords *int32

	noSmithyDocumentSerde
}

type DescribeOrderableReplicationInstancesOutput struct {

	// An optional pagination token provided by a previous request. If this parameter
	// is specified, the response includes only records beyond the marker, up to the
	// value specified by MaxRecords.
	Marker *string

	// The order-able replication instances available.
	OrderableReplicationInstances []types.OrderableReplicationInstance

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeOrderableReplicationInstancesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDescribeOrderableReplicationInstances{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDescribeOrderableReplicationInstances{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeOrderableReplicationInstances(options.Region), middleware.Before); err != nil {
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

// DescribeOrderableReplicationInstancesAPIClient is a client that implements the
// DescribeOrderableReplicationInstances operation.
type DescribeOrderableReplicationInstancesAPIClient interface {
	DescribeOrderableReplicationInstances(context.Context, *DescribeOrderableReplicationInstancesInput, ...func(*Options)) (*DescribeOrderableReplicationInstancesOutput, error)
}

var _ DescribeOrderableReplicationInstancesAPIClient = (*Client)(nil)

// DescribeOrderableReplicationInstancesPaginatorOptions is the paginator options
// for DescribeOrderableReplicationInstances
type DescribeOrderableReplicationInstancesPaginatorOptions struct {
	// The maximum number of records to include in the response. If more records exist
	// than the specified MaxRecords value, a pagination token called a marker is
	// included in the response so that the remaining results can be retrieved.
	// Default: 100 Constraints: Minimum 20, maximum 100.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// DescribeOrderableReplicationInstancesPaginator is a paginator for
// DescribeOrderableReplicationInstances
type DescribeOrderableReplicationInstancesPaginator struct {
	options   DescribeOrderableReplicationInstancesPaginatorOptions
	client    DescribeOrderableReplicationInstancesAPIClient
	params    *DescribeOrderableReplicationInstancesInput
	nextToken *string
	firstPage bool
}

// NewDescribeOrderableReplicationInstancesPaginator returns a new
// DescribeOrderableReplicationInstancesPaginator
func NewDescribeOrderableReplicationInstancesPaginator(client DescribeOrderableReplicationInstancesAPIClient, params *DescribeOrderableReplicationInstancesInput, optFns ...func(*DescribeOrderableReplicationInstancesPaginatorOptions)) *DescribeOrderableReplicationInstancesPaginator {
	if params == nil {
		params = &DescribeOrderableReplicationInstancesInput{}
	}

	options := DescribeOrderableReplicationInstancesPaginatorOptions{}
	if params.MaxRecords != nil {
		options.Limit = *params.MaxRecords
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &DescribeOrderableReplicationInstancesPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.Marker,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *DescribeOrderableReplicationInstancesPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next DescribeOrderableReplicationInstances page.
func (p *DescribeOrderableReplicationInstancesPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*DescribeOrderableReplicationInstancesOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.Marker = p.nextToken

	var limit *int32
	if p.options.Limit > 0 {
		limit = &p.options.Limit
	}
	params.MaxRecords = limit

	result, err := p.client.DescribeOrderableReplicationInstances(ctx, &params, optFns...)
	if err != nil {
		return nil, err
	}
	p.firstPage = false

	prevToken := p.nextToken
	p.nextToken = result.Marker

	if p.options.StopOnDuplicateToken &&
		prevToken != nil &&
		p.nextToken != nil &&
		*prevToken == *p.nextToken {
		p.nextToken = nil
	}

	return result, nil
}

func newServiceMetadataMiddleware_opDescribeOrderableReplicationInstances(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "dms",
		OperationName: "DescribeOrderableReplicationInstances",
	}
}
