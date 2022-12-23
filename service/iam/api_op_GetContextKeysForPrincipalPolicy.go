// Code generated by smithy-go-codegen DO NOT EDIT.

package iam

import (
	"context"
	awsmiddleware "oss-sdk-go/oss/middleware"
	"oss-sdk-go/oss/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Gets a list of all of the context keys referenced in all the IAM policies that
// are attached to the specified IAM entity. The entity can be an IAM user, group,
// or role. If you specify a user, then the request also includes all of the
// policies attached to groups that the user is a member of. You can optionally
// include a list of one or more additional policies, specified as strings. If you
// want to include only a list of policies by string, use
// GetContextKeysForCustomPolicy instead. Note: This operation discloses
// information about the permissions granted to other users. If you do not want
// users to see other user's permissions, then consider allowing them to use
// GetContextKeysForCustomPolicy instead. Context keys are variables maintained by
// Amazon Web Services and its services that provide details about the context of
// an API query request. Context keys can be evaluated by testing against a value
// in an IAM policy. Use GetContextKeysForPrincipalPolicy to understand what key
// names and values you must supply when you call SimulatePrincipalPolicy.
func (c *Client) GetContextKeysForPrincipalPolicy(ctx context.Context, params *GetContextKeysForPrincipalPolicyInput, optFns ...func(*Options)) (*GetContextKeysForPrincipalPolicyOutput, error) {
	if params == nil {
		params = &GetContextKeysForPrincipalPolicyInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetContextKeysForPrincipalPolicy", params, optFns, c.addOperationGetContextKeysForPrincipalPolicyMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetContextKeysForPrincipalPolicyOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetContextKeysForPrincipalPolicyInput struct {

	// The ARN of a user, group, or role whose policies contain the context keys that
	// you want listed. If you specify a user, the list includes context keys that are
	// found in all policies that are attached to the user. The list also includes all
	// groups that the user is a member of. If you pick a group or a role, then it
	// includes only those context keys that are found in policies attached to that
	// entity. Note that all parameters are shown in unencoded form here for clarity,
	// but must be URL encoded to be included as a part of a real HTML request. For
	// more information about ARNs, see Amazon Resource Names (ARNs)
	// (https://docs.aws.amazon.com/general/latest/gr/aws-arns-and-namespaces.html) in
	// the Amazon Web Services General Reference.
	//
	// This member is required.
	PolicySourceArn *string

	// An optional list of additional policies for which you want the list of context
	// keys that are referenced. The regex pattern (http://wikipedia.org/wiki/regex)
	// used to validate this parameter is a string of characters consisting of the
	// following:
	//
	// * Any printable ASCII character ranging from the space character
	// (\u0020) through the end of the ASCII character range
	//
	// * The printable
	// characters in the Basic Latin and Latin-1 Supplement character set (through
	// \u00FF)
	//
	// * The special characters tab (\u0009), line feed (\u000A), and carriage
	// return (\u000D)
	PolicyInputList []string

	noSmithyDocumentSerde
}

// Contains the response to a successful GetContextKeysForPrincipalPolicy or
// GetContextKeysForCustomPolicy request.
type GetContextKeysForPrincipalPolicyOutput struct {

	// The list of context keys that are referenced in the input policies.
	ContextKeyNames []string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetContextKeysForPrincipalPolicyMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsquery_serializeOpGetContextKeysForPrincipalPolicy{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpGetContextKeysForPrincipalPolicy{}, middleware.After)
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
	if err = addOpGetContextKeysForPrincipalPolicyValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetContextKeysForPrincipalPolicy(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opGetContextKeysForPrincipalPolicy(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "iam",
		OperationName: "GetContextKeysForPrincipalPolicy",
	}
}
