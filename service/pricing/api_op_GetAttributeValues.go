// Code generated by smithy-go-codegen DO NOT EDIT.

package pricing

import (
	"context"
	"fmt"
	awsmiddleware "oss-sdk-go/oss/middleware"
	"oss-sdk-go/oss/signer/v4"
	"oss-sdk-go/service/pricing/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Returns a list of attribute values. Attributes are similar to the details in a
// Price List API offer file. For a list of available attributes, see Offer File
// Definitions
// (https://docs.aws.amazon.com/awsaccountbilling/latest/aboutv2/reading-an-offer.html#pps-defs)
// in the Billing and Cost Management User Guide
// (https://docs.aws.amazon.com/awsaccountbilling/latest/aboutv2/billing-what-is.html).
func (c *Client) GetAttributeValues(ctx context.Context, params *GetAttributeValuesInput, optFns ...func(*Options)) (*GetAttributeValuesOutput, error) {
	if params == nil {
		params = &GetAttributeValuesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetAttributeValues", params, optFns, c.addOperationGetAttributeValuesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetAttributeValuesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetAttributeValuesInput struct {

	// The name of the attribute that you want to retrieve the values for, such as
	// volumeType.
	//
	// This member is required.
	AttributeName *string

	// The service code for the service whose attributes you want to retrieve. For
	// example, if you want the retrieve an EC2 attribute, use AmazonEC2.
	//
	// This member is required.
	ServiceCode *string

	// The maximum number of results to return in response.
	MaxResults *int32

	// The pagination token that indicates the next set of results that you want to
	// retrieve.
	NextToken *string

	noSmithyDocumentSerde
}

type GetAttributeValuesOutput struct {

	// The list of values for an attribute. For example, Throughput Optimized HDD and
	// Provisioned IOPS are two available values for the AmazonEC2volumeType.
	AttributeValues []types.AttributeValue

	// The pagination token that indicates the next set of results to retrieve.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetAttributeValuesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpGetAttributeValues{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpGetAttributeValues{}, middleware.After)
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
	if err = addOpGetAttributeValuesValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetAttributeValues(options.Region), middleware.Before); err != nil {
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

// GetAttributeValuesAPIClient is a client that implements the GetAttributeValues
// operation.
type GetAttributeValuesAPIClient interface {
	GetAttributeValues(context.Context, *GetAttributeValuesInput, ...func(*Options)) (*GetAttributeValuesOutput, error)
}

var _ GetAttributeValuesAPIClient = (*Client)(nil)

// GetAttributeValuesPaginatorOptions is the paginator options for
// GetAttributeValues
type GetAttributeValuesPaginatorOptions struct {
	// The maximum number of results to return in response.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// GetAttributeValuesPaginator is a paginator for GetAttributeValues
type GetAttributeValuesPaginator struct {
	options   GetAttributeValuesPaginatorOptions
	client    GetAttributeValuesAPIClient
	params    *GetAttributeValuesInput
	nextToken *string
	firstPage bool
}

// NewGetAttributeValuesPaginator returns a new GetAttributeValuesPaginator
func NewGetAttributeValuesPaginator(client GetAttributeValuesAPIClient, params *GetAttributeValuesInput, optFns ...func(*GetAttributeValuesPaginatorOptions)) *GetAttributeValuesPaginator {
	if params == nil {
		params = &GetAttributeValuesInput{}
	}

	options := GetAttributeValuesPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &GetAttributeValuesPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *GetAttributeValuesPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next GetAttributeValues page.
func (p *GetAttributeValuesPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*GetAttributeValuesOutput, error) {
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

	result, err := p.client.GetAttributeValues(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opGetAttributeValues(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "pricing",
		OperationName: "GetAttributeValues",
	}
}
