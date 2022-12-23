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

// Returns intent information as follows:
//
// * If you specify the nameContains field,
// returns the $LATEST version of all intents that contain the specified string.
//
// *
// If you don't specify the nameContains field, returns information about the
// $LATEST version of all intents.
//
// The operation requires permission for the
// lex:GetIntents action.
func (c *Client) GetIntents(ctx context.Context, params *GetIntentsInput, optFns ...func(*Options)) (*GetIntentsOutput, error) {
	if params == nil {
		params = &GetIntentsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetIntents", params, optFns, c.addOperationGetIntentsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetIntentsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetIntentsInput struct {

	// The maximum number of intents to return in the response. The default is 10.
	MaxResults *int32

	// Substring to match in intent names. An intent will be returned if any part of
	// its name matches the substring. For example, "xyz" matches both "xyzabc" and
	// "abcxyz."
	NameContains *string

	// A pagination token that fetches the next page of intents. If the response to
	// this API call is truncated, Amazon Lex returns a pagination token in the
	// response. To fetch the next page of intents, specify the pagination token in the
	// next request.
	NextToken *string

	noSmithyDocumentSerde
}

type GetIntentsOutput struct {

	// An array of Intent objects. For more information, see PutBot.
	Intents []types.IntentMetadata

	// If the response is truncated, the response includes a pagination token that you
	// can specify in your next request to fetch the next page of intents.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetIntentsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpGetIntents{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpGetIntents{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetIntents(options.Region), middleware.Before); err != nil {
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

// GetIntentsAPIClient is a client that implements the GetIntents operation.
type GetIntentsAPIClient interface {
	GetIntents(context.Context, *GetIntentsInput, ...func(*Options)) (*GetIntentsOutput, error)
}

var _ GetIntentsAPIClient = (*Client)(nil)

// GetIntentsPaginatorOptions is the paginator options for GetIntents
type GetIntentsPaginatorOptions struct {
	// The maximum number of intents to return in the response. The default is 10.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// GetIntentsPaginator is a paginator for GetIntents
type GetIntentsPaginator struct {
	options   GetIntentsPaginatorOptions
	client    GetIntentsAPIClient
	params    *GetIntentsInput
	nextToken *string
	firstPage bool
}

// NewGetIntentsPaginator returns a new GetIntentsPaginator
func NewGetIntentsPaginator(client GetIntentsAPIClient, params *GetIntentsInput, optFns ...func(*GetIntentsPaginatorOptions)) *GetIntentsPaginator {
	if params == nil {
		params = &GetIntentsInput{}
	}

	options := GetIntentsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &GetIntentsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *GetIntentsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next GetIntents page.
func (p *GetIntentsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*GetIntentsOutput, error) {
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

	result, err := p.client.GetIntents(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opGetIntents(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "lex",
		OperationName: "GetIntents",
	}
}