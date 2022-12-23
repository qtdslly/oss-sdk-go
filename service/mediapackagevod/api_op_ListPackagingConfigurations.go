// Code generated by smithy-go-codegen DO NOT EDIT.

package mediapackagevod

import (
	"context"
	"fmt"
	awsmiddleware "oss-sdk-go/oss/middleware"
	"oss-sdk-go/oss/signer/v4"
	"oss-sdk-go/service/mediapackagevod/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Returns a collection of MediaPackage VOD PackagingConfiguration resources.
func (c *Client) ListPackagingConfigurations(ctx context.Context, params *ListPackagingConfigurationsInput, optFns ...func(*Options)) (*ListPackagingConfigurationsOutput, error) {
	if params == nil {
		params = &ListPackagingConfigurationsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListPackagingConfigurations", params, optFns, c.addOperationListPackagingConfigurationsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListPackagingConfigurationsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListPackagingConfigurationsInput struct {

	// Upper bound on number of records to return.
	MaxResults int32

	// A token used to resume pagination from the end of a previous request.
	NextToken *string

	// Returns MediaPackage VOD PackagingConfigurations associated with the specified
	// PackagingGroup.
	PackagingGroupId *string

	noSmithyDocumentSerde
}

type ListPackagingConfigurationsOutput struct {

	// A token that can be used to resume pagination from the end of the collection.
	NextToken *string

	// A list of MediaPackage VOD PackagingConfiguration resources.
	PackagingConfigurations []types.PackagingConfiguration

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListPackagingConfigurationsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListPackagingConfigurations{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListPackagingConfigurations{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListPackagingConfigurations(options.Region), middleware.Before); err != nil {
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

// ListPackagingConfigurationsAPIClient is a client that implements the
// ListPackagingConfigurations operation.
type ListPackagingConfigurationsAPIClient interface {
	ListPackagingConfigurations(context.Context, *ListPackagingConfigurationsInput, ...func(*Options)) (*ListPackagingConfigurationsOutput, error)
}

var _ ListPackagingConfigurationsAPIClient = (*Client)(nil)

// ListPackagingConfigurationsPaginatorOptions is the paginator options for
// ListPackagingConfigurations
type ListPackagingConfigurationsPaginatorOptions struct {
	// Upper bound on number of records to return.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListPackagingConfigurationsPaginator is a paginator for
// ListPackagingConfigurations
type ListPackagingConfigurationsPaginator struct {
	options   ListPackagingConfigurationsPaginatorOptions
	client    ListPackagingConfigurationsAPIClient
	params    *ListPackagingConfigurationsInput
	nextToken *string
	firstPage bool
}

// NewListPackagingConfigurationsPaginator returns a new
// ListPackagingConfigurationsPaginator
func NewListPackagingConfigurationsPaginator(client ListPackagingConfigurationsAPIClient, params *ListPackagingConfigurationsInput, optFns ...func(*ListPackagingConfigurationsPaginatorOptions)) *ListPackagingConfigurationsPaginator {
	if params == nil {
		params = &ListPackagingConfigurationsInput{}
	}

	options := ListPackagingConfigurationsPaginatorOptions{}
	if params.MaxResults != 0 {
		options.Limit = params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListPackagingConfigurationsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListPackagingConfigurationsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListPackagingConfigurations page.
func (p *ListPackagingConfigurationsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListPackagingConfigurationsOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.NextToken = p.nextToken

	params.MaxResults = p.options.Limit

	result, err := p.client.ListPackagingConfigurations(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListPackagingConfigurations(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "mediapackage-vod",
		OperationName: "ListPackagingConfigurations",
	}
}
