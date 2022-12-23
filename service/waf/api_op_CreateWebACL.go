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
// global use. Creates a WebACL, which contains the Rules that identify the
// CloudFront web requests that you want to allow, block, or count. AWS WAF
// evaluates Rules in order based on the value of Priority for each Rule. You also
// specify a default action, either ALLOW or BLOCK. If a web request doesn't match
// any of the Rules in a WebACL, AWS WAF responds to the request with the default
// action. To create and configure a WebACL, perform the following steps:
//
// * Create
// and update the ByteMatchSet objects and other predicates that you want to
// include in Rules. For more information, see CreateByteMatchSet,
// UpdateByteMatchSet, CreateIPSet, UpdateIPSet, CreateSqlInjectionMatchSet, and
// UpdateSqlInjectionMatchSet.
//
// * Create and update the Rules that you want to
// include in the WebACL. For more information, see CreateRule and UpdateRule.
//
// *
// Use GetChangeToken to get the change token that you provide in the ChangeToken
// parameter of a CreateWebACL request.
//
// * Submit a CreateWebACL request.
//
// * Use
// GetChangeToken to get the change token that you provide in the ChangeToken
// parameter of an UpdateWebACL request.
//
// * Submit an UpdateWebACL request to
// specify the Rules that you want to include in the WebACL, to specify the default
// action, and to associate the WebACL with a CloudFront distribution.
//
// For more
// information about how to use the AWS WAF API, see the AWS WAF Developer Guide
// (https://docs.aws.amazon.com/waf/latest/developerguide/).
func (c *Client) CreateWebACL(ctx context.Context, params *CreateWebACLInput, optFns ...func(*Options)) (*CreateWebACLOutput, error) {
	if params == nil {
		params = &CreateWebACLInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateWebACL", params, optFns, c.addOperationCreateWebACLMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateWebACLOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateWebACLInput struct {

	// The value returned by the most recent call to GetChangeToken.
	//
	// This member is required.
	ChangeToken *string

	// The action that you want AWS WAF to take when a request doesn't match the
	// criteria specified in any of the Rule objects that are associated with the
	// WebACL.
	//
	// This member is required.
	DefaultAction *types.WafAction

	// A friendly name or description for the metrics for this WebACL.The name can
	// contain only alphanumeric characters (A-Z, a-z, 0-9), with maximum length 128
	// and minimum length one. It can't contain whitespace or metric names reserved for
	// AWS WAF, including "All" and "Default_Action." You can't change MetricName after
	// you create the WebACL.
	//
	// This member is required.
	MetricName *string

	// A friendly name or description of the WebACL. You can't change Name after you
	// create the WebACL.
	//
	// This member is required.
	Name *string

	//
	Tags []types.Tag

	noSmithyDocumentSerde
}

type CreateWebACLOutput struct {

	// The ChangeToken that you used to submit the CreateWebACL request. You can also
	// use this value to query the status of the request. For more information, see
	// GetChangeTokenStatus.
	ChangeToken *string

	// The WebACL returned in the CreateWebACL response.
	WebACL *types.WebACL

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateWebACLMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpCreateWebACL{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpCreateWebACL{}, middleware.After)
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
	if err = addOpCreateWebACLValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateWebACL(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateWebACL(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "waf",
		OperationName: "CreateWebACL",
	}
}
