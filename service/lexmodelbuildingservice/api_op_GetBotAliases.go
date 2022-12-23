// Code generated by smithy-go-codegen DO NOT EDIT.

package lexmodelbuildingservice

import (
	"context"
	"fmt"
	awsmiddleware "oss-sdk-go/oss/middleware"
	"oss-sdk-go/oss/signer/v4"
	"oss-sdk-go/service/lexmodelbuildingservice/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Returns a list of aliases for a specified Amazon Lex bot. This operation
// requires permissions for the lex:GetBotAliases action.
func (c *Client) GetBotAliases(ctx context.Context, params *GetBotAliasesInput, optFns ...func(*Options)) (*GetBotAliasesOutput, error) {
	if params == nil {
		params = &GetBotAliasesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetBotAliases", params, optFns, c.addOperationGetBotAliasesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetBotAliasesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetBotAliasesInput struct {

	// The name of the bot.
	//
	// This member is required.
	BotName *string

	// The maximum number of aliases to return in the response. The default is 50. .
	MaxResults *int32

	// Substring to match in bot alias names. An alias will be returned if any part of
	// its name matches the substring. For example, "xyz" matches both "xyzabc" and
	// "abcxyz."
	NameContains *string

	// A pagination token for fetching the next page of aliases. If the response to
	// this call is truncated, Amazon Lex returns a pagination token in the response.
	// To fetch the next page of aliases, specify the pagination token in the next
	// request.
	NextToken *string

	noSmithyDocumentSerde
}

type GetBotAliasesOutput struct {

	// An array of BotAliasMetadata objects, each describing a bot alias.
	BotAliases []types.BotAliasMetadata

	// A pagination token for fetching next page of aliases. If the response to this
	// call is truncated, Amazon Lex returns a pagination token in the response. To
	// fetch the next page of aliases, specify the pagination token in the next
	// request.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetBotAliasesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpGetBotAliases{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpGetBotAliases{}, middleware.After)
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
	if err = addOpGetBotAliasesValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetBotAliases(options.Region), middleware.Before); err != nil {
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

// GetBotAliasesAPIClient is a client that implements the GetBotAliases operation.
type GetBotAliasesAPIClient interface {
	GetBotAliases(context.Context, *GetBotAliasesInput, ...func(*Options)) (*GetBotAliasesOutput, error)
}

var _ GetBotAliasesAPIClient = (*Client)(nil)

// GetBotAliasesPaginatorOptions is the paginator options for GetBotAliases
type GetBotAliasesPaginatorOptions struct {
	// The maximum number of aliases to return in the response. The default is 50. .
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// GetBotAliasesPaginator is a paginator for GetBotAliases
type GetBotAliasesPaginator struct {
	options   GetBotAliasesPaginatorOptions
	client    GetBotAliasesAPIClient
	params    *GetBotAliasesInput
	nextToken *string
	firstPage bool
}

// NewGetBotAliasesPaginator returns a new GetBotAliasesPaginator
func NewGetBotAliasesPaginator(client GetBotAliasesAPIClient, params *GetBotAliasesInput, optFns ...func(*GetBotAliasesPaginatorOptions)) *GetBotAliasesPaginator {
	if params == nil {
		params = &GetBotAliasesInput{}
	}

	options := GetBotAliasesPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &GetBotAliasesPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *GetBotAliasesPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next GetBotAliases page.
func (p *GetBotAliasesPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*GetBotAliasesOutput, error) {
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

	result, err := p.client.GetBotAliases(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opGetBotAliases(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "lex",
		OperationName: "GetBotAliases",
	}
}
