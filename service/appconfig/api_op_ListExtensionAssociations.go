// Code generated by smithy-go-codegen DO NOT EDIT.

package appconfig

import (
	"context"
	"fmt"
	awsmiddleware "oss-sdk-go/oss/middleware"
	"oss-sdk-go/oss/signer/v4"
	"oss-sdk-go/service/appconfig/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Lists all AppConfig extension associations in the account. For more information
// about extensions and associations, see Working with AppConfig extensions
// (https://docs.aws.amazon.com/appconfig/latest/userguide/working-with-appconfig-extensions.html)
// in the AppConfig User Guide.
func (c *Client) ListExtensionAssociations(ctx context.Context, params *ListExtensionAssociationsInput, optFns ...func(*Options)) (*ListExtensionAssociationsOutput, error) {
	if params == nil {
		params = &ListExtensionAssociationsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListExtensionAssociations", params, optFns, c.addOperationListExtensionAssociationsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListExtensionAssociationsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListExtensionAssociationsInput struct {

	// The name, the ID, or the Amazon Resource Name (ARN) of the extension.
	ExtensionIdentifier *string

	// The version number for the extension defined in the association.
	ExtensionVersionNumber *int32

	// The maximum number of items to return for this call. The call also returns a
	// token that you can specify in a subsequent call to get the next set of results.
	MaxResults *int32

	// A token to start the list. Use this token to get the next set of results or pass
	// null to get the first set of results.
	NextToken *string

	// The ARN of an application, configuration profile, or environment.
	ResourceIdentifier *string

	noSmithyDocumentSerde
}

type ListExtensionAssociationsOutput struct {

	// The list of extension associations. Each item represents an extension
	// association to an application, environment, or configuration profile.
	Items []types.ExtensionAssociationSummary

	// The token for the next set of items to return. Use this token to get the next
	// set of results.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListExtensionAssociationsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListExtensionAssociations{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListExtensionAssociations{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListExtensionAssociations(options.Region), middleware.Before); err != nil {
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

// ListExtensionAssociationsAPIClient is a client that implements the
// ListExtensionAssociations operation.
type ListExtensionAssociationsAPIClient interface {
	ListExtensionAssociations(context.Context, *ListExtensionAssociationsInput, ...func(*Options)) (*ListExtensionAssociationsOutput, error)
}

var _ ListExtensionAssociationsAPIClient = (*Client)(nil)

// ListExtensionAssociationsPaginatorOptions is the paginator options for
// ListExtensionAssociations
type ListExtensionAssociationsPaginatorOptions struct {
	// The maximum number of items to return for this call. The call also returns a
	// token that you can specify in a subsequent call to get the next set of results.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListExtensionAssociationsPaginator is a paginator for ListExtensionAssociations
type ListExtensionAssociationsPaginator struct {
	options   ListExtensionAssociationsPaginatorOptions
	client    ListExtensionAssociationsAPIClient
	params    *ListExtensionAssociationsInput
	nextToken *string
	firstPage bool
}

// NewListExtensionAssociationsPaginator returns a new
// ListExtensionAssociationsPaginator
func NewListExtensionAssociationsPaginator(client ListExtensionAssociationsAPIClient, params *ListExtensionAssociationsInput, optFns ...func(*ListExtensionAssociationsPaginatorOptions)) *ListExtensionAssociationsPaginator {
	if params == nil {
		params = &ListExtensionAssociationsInput{}
	}

	options := ListExtensionAssociationsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListExtensionAssociationsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListExtensionAssociationsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListExtensionAssociations page.
func (p *ListExtensionAssociationsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListExtensionAssociationsOutput, error) {
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

	result, err := p.client.ListExtensionAssociations(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListExtensionAssociations(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "appconfig",
		OperationName: "ListExtensionAssociations",
	}
}
