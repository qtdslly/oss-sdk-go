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

// Returns bot information as follows:
//
// * If you provide the nameContains field,
// the response includes information for the $LATEST version of all bots whose name
// contains the specified string.
//
// * If you don't specify the nameContains field,
// the operation returns information about the $LATEST version of all of your
// bots.
//
// This operation requires permission for the lex:GetBots action.
func (c *Client) GetBots(ctx context.Context, params *GetBotsInput, optFns ...func(*Options)) (*GetBotsOutput, error) {
	if params == nil {
		params = &GetBotsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetBots", params, optFns, c.addOperationGetBotsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetBotsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetBotsInput struct {

	// The maximum number of bots to return in the response that the request will
	// return. The default is 10.
	MaxResults *int32

	// Substring to match in bot names. A bot will be returned if any part of its name
	// matches the substring. For example, "xyz" matches both "xyzabc" and "abcxyz."
	NameContains *string

	// A pagination token that fetches the next page of bots. If the response to this
	// call is truncated, Amazon Lex returns a pagination token in the response. To
	// fetch the next page of bots, specify the pagination token in the next request.
	NextToken *string

	noSmithyDocumentSerde
}

type GetBotsOutput struct {

	// An array of botMetadata objects, with one entry for each bot.
	Bots []types.BotMetadata

	// If the response is truncated, it includes a pagination token that you can
	// specify in your next request to fetch the next page of bots.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetBotsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpGetBots{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpGetBots{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetBots(options.Region), middleware.Before); err != nil {
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

// GetBotsAPIClient is a client that implements the GetBots operation.
type GetBotsAPIClient interface {
	GetBots(context.Context, *GetBotsInput, ...func(*Options)) (*GetBotsOutput, error)
}

var _ GetBotsAPIClient = (*Client)(nil)

// GetBotsPaginatorOptions is the paginator options for GetBots
type GetBotsPaginatorOptions struct {
	// The maximum number of bots to return in the response that the request will
	// return. The default is 10.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// GetBotsPaginator is a paginator for GetBots
type GetBotsPaginator struct {
	options   GetBotsPaginatorOptions
	client    GetBotsAPIClient
	params    *GetBotsInput
	nextToken *string
	firstPage bool
}

// NewGetBotsPaginator returns a new GetBotsPaginator
func NewGetBotsPaginator(client GetBotsAPIClient, params *GetBotsInput, optFns ...func(*GetBotsPaginatorOptions)) *GetBotsPaginator {
	if params == nil {
		params = &GetBotsInput{}
	}

	options := GetBotsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &GetBotsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *GetBotsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next GetBots page.
func (p *GetBotsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*GetBotsOutput, error) {
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

	result, err := p.client.GetBots(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opGetBots(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "lex",
		OperationName: "GetBots",
	}
}
