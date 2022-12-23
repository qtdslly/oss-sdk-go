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

// Lists Amazon GuardDuty usage statistics over the last 30 days for the specified
// detector ID. For newly enabled detectors or data sources, the cost returned will
// include only the usage so far under 30 days. This may differ from the cost
// metrics in the console, which project usage over 30 days to provide a monthly
// cost estimate. For more information, see Understanding How Usage Costs are
// Calculated
// (https://docs.aws.amazon.com/guardduty/latest/ug/monitoring_costs.html#usage-calculations).
func (c *Client) GetUsageStatistics(ctx context.Context, params *GetUsageStatisticsInput, optFns ...func(*Options)) (*GetUsageStatisticsOutput, error) {
	if params == nil {
		params = &GetUsageStatisticsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetUsageStatistics", params, optFns, c.addOperationGetUsageStatisticsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetUsageStatisticsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetUsageStatisticsInput struct {

	// The ID of the detector that specifies the GuardDuty service whose usage
	// statistics you want to retrieve.
	//
	// This member is required.
	DetectorId *string

	// Represents the criteria used for querying usage.
	//
	// This member is required.
	UsageCriteria *types.UsageCriteria

	// The type of usage statistics to retrieve.
	//
	// This member is required.
	UsageStatisticType types.UsageStatisticType

	// The maximum number of results to return in the response.
	MaxResults int32

	// A token to use for paginating results that are returned in the response. Set the
	// value of this parameter to null for the first request to a list action. For
	// subsequent calls, use the NextToken value returned from the previous request to
	// continue listing results after the first page.
	NextToken *string

	// The currency unit you would like to view your usage statistics in. Current valid
	// values are USD.
	Unit *string

	noSmithyDocumentSerde
}

type GetUsageStatisticsOutput struct {

	// The pagination parameter to be used on the next list operation to retrieve more
	// items.
	NextToken *string

	// The usage statistics object. If a UsageStatisticType was provided, the objects
	// representing other types will be null.
	UsageStatistics *types.UsageStatistics

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetUsageStatisticsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpGetUsageStatistics{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpGetUsageStatistics{}, middleware.After)
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
	if err = addOpGetUsageStatisticsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetUsageStatistics(options.Region), middleware.Before); err != nil {
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

// GetUsageStatisticsAPIClient is a client that implements the GetUsageStatistics
// operation.
type GetUsageStatisticsAPIClient interface {
	GetUsageStatistics(context.Context, *GetUsageStatisticsInput, ...func(*Options)) (*GetUsageStatisticsOutput, error)
}

var _ GetUsageStatisticsAPIClient = (*Client)(nil)

// GetUsageStatisticsPaginatorOptions is the paginator options for
// GetUsageStatistics
type GetUsageStatisticsPaginatorOptions struct {
	// The maximum number of results to return in the response.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// GetUsageStatisticsPaginator is a paginator for GetUsageStatistics
type GetUsageStatisticsPaginator struct {
	options   GetUsageStatisticsPaginatorOptions
	client    GetUsageStatisticsAPIClient
	params    *GetUsageStatisticsInput
	nextToken *string
	firstPage bool
}

// NewGetUsageStatisticsPaginator returns a new GetUsageStatisticsPaginator
func NewGetUsageStatisticsPaginator(client GetUsageStatisticsAPIClient, params *GetUsageStatisticsInput, optFns ...func(*GetUsageStatisticsPaginatorOptions)) *GetUsageStatisticsPaginator {
	if params == nil {
		params = &GetUsageStatisticsInput{}
	}

	options := GetUsageStatisticsPaginatorOptions{}
	if params.MaxResults != 0 {
		options.Limit = params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &GetUsageStatisticsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *GetUsageStatisticsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next GetUsageStatistics page.
func (p *GetUsageStatisticsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*GetUsageStatisticsOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.NextToken = p.nextToken

	params.MaxResults = p.options.Limit

	result, err := p.client.GetUsageStatistics(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opGetUsageStatistics(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "guardduty",
		OperationName: "GetUsageStatistics",
	}
}
