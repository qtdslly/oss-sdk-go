// Code generated by smithy-go-codegen DO NOT EDIT.

package wafregional

import (
	"context"
	awsmiddleware "oss-sdk-go/oss/middleware"
	"oss-sdk-go/oss/signer/v4"
	"oss-sdk-go/service/wafregional/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// This is AWS WAF Classic documentation. For more information, see AWS WAF Classic
// (https://docs.aws.amazon.com/waf/latest/developerguide/classic-waf-chapter.html)
// in the developer guide. For the latest version of AWS WAF, use the AWS WAFV2 API
// and see the AWS WAF Developer Guide
// (https://docs.aws.amazon.com/waf/latest/developerguide/waf-chapter.html). With
// the latest version, AWS WAF has a single set of endpoints for regional and
// global use. Inserts or deletes ByteMatchTuple objects (filters) in a
// ByteMatchSet. For each ByteMatchTuple object, you specify the following
// values:
//
// * Whether to insert or delete the object from the array. If you want to
// change a ByteMatchSetUpdate object, you delete the existing object and add a new
// one.
//
// * The part of a web request that you want AWS WAF to inspect, such as a
// query string or the value of the User-Agent header.
//
// * The bytes (typically a
// string that corresponds with ASCII characters) that you want AWS WAF to look
// for. For more information, including how you specify the values for the AWS WAF
// API and the AWS CLI or SDKs, see TargetString in the ByteMatchTuple data
// type.
//
// * Where to look, such as at the beginning or the end of a query
// string.
//
// * Whether to perform any conversions on the request, such as converting
// it to lowercase, before inspecting it for the specified string.
//
// For example,
// you can add a ByteMatchSetUpdate object that matches web requests in which
// User-Agent headers contain the string BadBot. You can then configure AWS WAF to
// block those requests. To create and configure a ByteMatchSet, perform the
// following steps:
//
// * Create a ByteMatchSet. For more information, see
// CreateByteMatchSet.
//
// * Use GetChangeToken to get the change token that you
// provide in the ChangeToken parameter of an UpdateByteMatchSet request.
//
// * Submit
// an UpdateByteMatchSet request to specify the part of the request that you want
// AWS WAF to inspect (for example, the header or the URI) and the value that you
// want AWS WAF to watch for.
//
// For more information about how to use the AWS WAF
// API to allow or block HTTP requests, see the AWS WAF Developer Guide
// (https://docs.aws.amazon.com/waf/latest/developerguide/).
func (c *Client) UpdateByteMatchSet(ctx context.Context, params *UpdateByteMatchSetInput, optFns ...func(*Options)) (*UpdateByteMatchSetOutput, error) {
	if params == nil {
		params = &UpdateByteMatchSetInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateByteMatchSet", params, optFns, c.addOperationUpdateByteMatchSetMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateByteMatchSetOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateByteMatchSetInput struct {

	// The ByteMatchSetId of the ByteMatchSet that you want to update. ByteMatchSetId
	// is returned by CreateByteMatchSet and by ListByteMatchSets.
	//
	// This member is required.
	ByteMatchSetId *string

	// The value returned by the most recent call to GetChangeToken.
	//
	// This member is required.
	ChangeToken *string

	// An array of ByteMatchSetUpdate objects that you want to insert into or delete
	// from a ByteMatchSet. For more information, see the applicable data types:
	//
	// *
	// ByteMatchSetUpdate: Contains Action and ByteMatchTuple
	//
	// * ByteMatchTuple:
	// Contains FieldToMatch, PositionalConstraint, TargetString, and
	// TextTransformation
	//
	// * FieldToMatch: Contains Data and Type
	//
	// This member is required.
	Updates []types.ByteMatchSetUpdate

	noSmithyDocumentSerde
}

type UpdateByteMatchSetOutput struct {

	// The ChangeToken that you used to submit the UpdateByteMatchSet request. You can
	// also use this value to query the status of the request. For more information,
	// see GetChangeTokenStatus.
	ChangeToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationUpdateByteMatchSetMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpUpdateByteMatchSet{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpUpdateByteMatchSet{}, middleware.After)
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
	if err = addOpUpdateByteMatchSetValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateByteMatchSet(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opUpdateByteMatchSet(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "waf-regional",
		OperationName: "UpdateByteMatchSet",
	}
}
