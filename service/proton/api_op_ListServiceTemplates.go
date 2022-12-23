// Code generated by smithy-go-codegen DO NOT EDIT.

package proton

import (
	"context"
	"fmt"
	awsmiddleware "oss-sdk-go/oss/middleware"
	"oss-sdk-go/oss/signer/v4"
	"oss-sdk-go/service/proton/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// List service templates with detail data.
func (c *Client) ListServiceTemplates(ctx context.Context, params *ListServiceTemplatesInput, optFns ...func(*Options)) (*ListServiceTemplatesOutput, error) {
	if params == nil {
		params = &ListServiceTemplatesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListServiceTemplates", params, optFns, c.addOperationListServiceTemplatesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListServiceTemplatesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListServiceTemplatesInput struct {

	// The maximum number of service templates to list.
	MaxResults *int32

	// A token that indicates the location of the next service template in the array of
	// service templates, after the list of service templates previously requested.
	NextToken *string

	noSmithyDocumentSerde
}

type ListServiceTemplatesOutput struct {

	// An array of service templates with detail data.
	//
	// This member is required.
	Templates []types.ServiceTemplateSummary

	// A token that indicates the location of the next service template in the array of
	// service templates, after the current requested list of service templates.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListServiceTemplatesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson10_serializeOpListServiceTemplates{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson10_deserializeOpListServiceTemplates{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListServiceTemplates(options.Region), middleware.Before); err != nil {
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

// ListServiceTemplatesAPIClient is a client that implements the
// ListServiceTemplates operation.
type ListServiceTemplatesAPIClient interface {
	ListServiceTemplates(context.Context, *ListServiceTemplatesInput, ...func(*Options)) (*ListServiceTemplatesOutput, error)
}

var _ ListServiceTemplatesAPIClient = (*Client)(nil)

// ListServiceTemplatesPaginatorOptions is the paginator options for
// ListServiceTemplates
type ListServiceTemplatesPaginatorOptions struct {
	// The maximum number of service templates to list.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListServiceTemplatesPaginator is a paginator for ListServiceTemplates
type ListServiceTemplatesPaginator struct {
	options   ListServiceTemplatesPaginatorOptions
	client    ListServiceTemplatesAPIClient
	params    *ListServiceTemplatesInput
	nextToken *string
	firstPage bool
}

// NewListServiceTemplatesPaginator returns a new ListServiceTemplatesPaginator
func NewListServiceTemplatesPaginator(client ListServiceTemplatesAPIClient, params *ListServiceTemplatesInput, optFns ...func(*ListServiceTemplatesPaginatorOptions)) *ListServiceTemplatesPaginator {
	if params == nil {
		params = &ListServiceTemplatesInput{}
	}

	options := ListServiceTemplatesPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListServiceTemplatesPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListServiceTemplatesPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListServiceTemplates page.
func (p *ListServiceTemplatesPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListServiceTemplatesOutput, error) {
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

	result, err := p.client.ListServiceTemplates(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListServiceTemplates(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "proton",
		OperationName: "ListServiceTemplates",
	}
}
