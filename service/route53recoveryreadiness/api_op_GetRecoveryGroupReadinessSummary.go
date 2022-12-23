// Code generated by smithy-go-codegen DO NOT EDIT.

package route53recoveryreadiness

import (
	"context"
	"fmt"
	awsmiddleware "oss-sdk-go/oss/middleware"
	"oss-sdk-go/oss/signer/v4"
	"oss-sdk-go/service/route53recoveryreadiness/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Displays a summary of information about a recovery group's readiness status.
// Includes the readiness checks for resources in the recovery group and the
// readiness status of each one.
func (c *Client) GetRecoveryGroupReadinessSummary(ctx context.Context, params *GetRecoveryGroupReadinessSummaryInput, optFns ...func(*Options)) (*GetRecoveryGroupReadinessSummaryOutput, error) {
	if params == nil {
		params = &GetRecoveryGroupReadinessSummaryInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetRecoveryGroupReadinessSummary", params, optFns, c.addOperationGetRecoveryGroupReadinessSummaryMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetRecoveryGroupReadinessSummaryOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetRecoveryGroupReadinessSummaryInput struct {

	// The name of a recovery group.
	//
	// This member is required.
	RecoveryGroupName *string

	// The number of objects that you want to return with this call.
	MaxResults int32

	// The token that identifies which batch of results you want to see.
	NextToken *string

	noSmithyDocumentSerde
}

type GetRecoveryGroupReadinessSummaryOutput struct {

	// The token that identifies which batch of results you want to see.
	NextToken *string

	// The readiness status at a recovery group level.
	Readiness types.Readiness

	// Summaries of the readiness checks for the recovery group.
	ReadinessChecks []types.ReadinessCheckSummary

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetRecoveryGroupReadinessSummaryMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpGetRecoveryGroupReadinessSummary{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpGetRecoveryGroupReadinessSummary{}, middleware.After)
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
	if err = addOpGetRecoveryGroupReadinessSummaryValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetRecoveryGroupReadinessSummary(options.Region), middleware.Before); err != nil {
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

// GetRecoveryGroupReadinessSummaryAPIClient is a client that implements the
// GetRecoveryGroupReadinessSummary operation.
type GetRecoveryGroupReadinessSummaryAPIClient interface {
	GetRecoveryGroupReadinessSummary(context.Context, *GetRecoveryGroupReadinessSummaryInput, ...func(*Options)) (*GetRecoveryGroupReadinessSummaryOutput, error)
}

var _ GetRecoveryGroupReadinessSummaryAPIClient = (*Client)(nil)

// GetRecoveryGroupReadinessSummaryPaginatorOptions is the paginator options for
// GetRecoveryGroupReadinessSummary
type GetRecoveryGroupReadinessSummaryPaginatorOptions struct {
	// The number of objects that you want to return with this call.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// GetRecoveryGroupReadinessSummaryPaginator is a paginator for
// GetRecoveryGroupReadinessSummary
type GetRecoveryGroupReadinessSummaryPaginator struct {
	options   GetRecoveryGroupReadinessSummaryPaginatorOptions
	client    GetRecoveryGroupReadinessSummaryAPIClient
	params    *GetRecoveryGroupReadinessSummaryInput
	nextToken *string
	firstPage bool
}

// NewGetRecoveryGroupReadinessSummaryPaginator returns a new
// GetRecoveryGroupReadinessSummaryPaginator
func NewGetRecoveryGroupReadinessSummaryPaginator(client GetRecoveryGroupReadinessSummaryAPIClient, params *GetRecoveryGroupReadinessSummaryInput, optFns ...func(*GetRecoveryGroupReadinessSummaryPaginatorOptions)) *GetRecoveryGroupReadinessSummaryPaginator {
	if params == nil {
		params = &GetRecoveryGroupReadinessSummaryInput{}
	}

	options := GetRecoveryGroupReadinessSummaryPaginatorOptions{}
	if params.MaxResults != 0 {
		options.Limit = params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &GetRecoveryGroupReadinessSummaryPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *GetRecoveryGroupReadinessSummaryPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next GetRecoveryGroupReadinessSummary page.
func (p *GetRecoveryGroupReadinessSummaryPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*GetRecoveryGroupReadinessSummaryOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.NextToken = p.nextToken

	params.MaxResults = p.options.Limit

	result, err := p.client.GetRecoveryGroupReadinessSummary(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opGetRecoveryGroupReadinessSummary(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "route53-recovery-readiness",
		OperationName: "GetRecoveryGroupReadinessSummary",
	}
}
