// Code generated by smithy-go-codegen DO NOT EDIT.

package iot

import (
	"context"
	"fmt"
	awsmiddleware "oss-sdk-go/oss/middleware"
	"oss-sdk-go/oss/signer/v4"
	"oss-sdk-go/service/iot/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Gets a list of domain configurations for the user. This list is sorted
// alphabetically by domain configuration name. Requires permission to access the
// ListDomainConfigurations
// (https://docs.aws.amazon.com/service-authorization/latest/reference/list_awsiot.html#awsiot-actions-as-permissions)
// action.
func (c *Client) ListDomainConfigurations(ctx context.Context, params *ListDomainConfigurationsInput, optFns ...func(*Options)) (*ListDomainConfigurationsOutput, error) {
	if params == nil {
		params = &ListDomainConfigurationsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListDomainConfigurations", params, optFns, c.addOperationListDomainConfigurationsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListDomainConfigurationsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListDomainConfigurationsInput struct {

	// The marker for the next set of results.
	Marker *string

	// The result page size.
	PageSize *int32

	// The type of service delivered by the endpoint.
	ServiceType types.ServiceType

	noSmithyDocumentSerde
}

type ListDomainConfigurationsOutput struct {

	// A list of objects that contain summary information about the user's domain
	// configurations.
	DomainConfigurations []types.DomainConfigurationSummary

	// The marker for the next set of results.
	NextMarker *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListDomainConfigurationsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListDomainConfigurations{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListDomainConfigurations{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListDomainConfigurations(options.Region), middleware.Before); err != nil {
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

// ListDomainConfigurationsAPIClient is a client that implements the
// ListDomainConfigurations operation.
type ListDomainConfigurationsAPIClient interface {
	ListDomainConfigurations(context.Context, *ListDomainConfigurationsInput, ...func(*Options)) (*ListDomainConfigurationsOutput, error)
}

var _ ListDomainConfigurationsAPIClient = (*Client)(nil)

// ListDomainConfigurationsPaginatorOptions is the paginator options for
// ListDomainConfigurations
type ListDomainConfigurationsPaginatorOptions struct {
	// The result page size.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListDomainConfigurationsPaginator is a paginator for ListDomainConfigurations
type ListDomainConfigurationsPaginator struct {
	options   ListDomainConfigurationsPaginatorOptions
	client    ListDomainConfigurationsAPIClient
	params    *ListDomainConfigurationsInput
	nextToken *string
	firstPage bool
}

// NewListDomainConfigurationsPaginator returns a new
// ListDomainConfigurationsPaginator
func NewListDomainConfigurationsPaginator(client ListDomainConfigurationsAPIClient, params *ListDomainConfigurationsInput, optFns ...func(*ListDomainConfigurationsPaginatorOptions)) *ListDomainConfigurationsPaginator {
	if params == nil {
		params = &ListDomainConfigurationsInput{}
	}

	options := ListDomainConfigurationsPaginatorOptions{}
	if params.PageSize != nil {
		options.Limit = *params.PageSize
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListDomainConfigurationsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.Marker,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListDomainConfigurationsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListDomainConfigurations page.
func (p *ListDomainConfigurationsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListDomainConfigurationsOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.Marker = p.nextToken

	var limit *int32
	if p.options.Limit > 0 {
		limit = &p.options.Limit
	}
	params.PageSize = limit

	result, err := p.client.ListDomainConfigurations(ctx, &params, optFns...)
	if err != nil {
		return nil, err
	}
	p.firstPage = false

	prevToken := p.nextToken
	p.nextToken = result.NextMarker

	if p.options.StopOnDuplicateToken &&
		prevToken != nil &&
		p.nextToken != nil &&
		*prevToken == *p.nextToken {
		p.nextToken = nil
	}

	return result, nil
}

func newServiceMetadataMiddleware_opListDomainConfigurations(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "execute-api",
		OperationName: "ListDomainConfigurations",
	}
}
