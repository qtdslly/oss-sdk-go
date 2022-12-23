// Code generated by smithy-go-codegen DO NOT EDIT.

package waf

import (
	"context"
	awsmiddleware "oss-sdk-go/oss/middleware"
	"oss-sdk-go/oss/signer/v4"
	"oss-sdk-go/service/waf/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// This is AWS WAF Classic documentation. For more information, see AWS WAF Classic
// (https://docs.aws.amazon.com/waf/latest/developerguide/classic-waf-chapter.html)
// in the developer guide. For the latest version of AWS WAF, use the AWS WAFV2 API
// and see the AWS WAF Developer Guide
// (https://docs.aws.amazon.com/waf/latest/developerguide/waf-chapter.html). With
// the latest version, AWS WAF has a single set of endpoints for regional and
// global use. Returns an array of RuleSummary objects.
func (c *Client) ListRules(ctx context.Context, params *ListRulesInput, optFns ...func(*Options)) (*ListRulesOutput, error) {
	if params == nil {
		params = &ListRulesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListRules", params, optFns, c.addOperationListRulesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListRulesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListRulesInput struct {

	// Specifies the number of Rules that you want AWS WAF to return for this request.
	// If you have more Rules than the number that you specify for Limit, the response
	// includes a NextMarker value that you can use to get another batch of Rules.
	Limit int32

	// If you specify a value for Limit and you have more Rules than the value of
	// Limit, AWS WAF returns a NextMarker value in the response that allows you to
	// list another group of Rules. For the second and subsequent ListRules requests,
	// specify the value of NextMarker from the previous response to get information
	// about another batch of Rules.
	NextMarker *string

	noSmithyDocumentSerde
}

type ListRulesOutput struct {

	// If you have more Rules than the number that you specified for Limit in the
	// request, the response includes a NextMarker value. To list more Rules, submit
	// another ListRules request, and specify the NextMarker value from the response in
	// the NextMarker value in the next request.
	NextMarker *string

	// An array of RuleSummary objects.
	Rules []types.RuleSummary

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListRulesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpListRules{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpListRules{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListRules(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opListRules(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "waf",
		OperationName: "ListRules",
	}
}
