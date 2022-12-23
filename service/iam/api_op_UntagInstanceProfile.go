// Code generated by smithy-go-codegen DO NOT EDIT.

package iam

import (
	"context"
	awsmiddleware "oss-sdk-go/oss/middleware"
	"oss-sdk-go/oss/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Removes the specified tags from the IAM instance profile. For more information
// about tagging, see Tagging IAM resources
// (https://docs.aws.amazon.com/IAM/latest/UserGuide/id_tags.html) in the IAM User
// Guide.
func (c *Client) UntagInstanceProfile(ctx context.Context, params *UntagInstanceProfileInput, optFns ...func(*Options)) (*UntagInstanceProfileOutput, error) {
	if params == nil {
		params = &UntagInstanceProfileInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UntagInstanceProfile", params, optFns, c.addOperationUntagInstanceProfileMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UntagInstanceProfileOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UntagInstanceProfileInput struct {

	// The name of the IAM instance profile from which you want to remove tags. This
	// parameter allows (through its regex pattern (http://wikipedia.org/wiki/regex)) a
	// string of characters consisting of upper and lowercase alphanumeric characters
	// with no spaces. You can also include any of the following characters: _+=,.@-
	//
	// This member is required.
	InstanceProfileName *string

	// A list of key names as a simple array of strings. The tags with matching keys
	// are removed from the specified instance profile.
	//
	// This member is required.
	TagKeys []string

	noSmithyDocumentSerde
}

type UntagInstanceProfileOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationUntagInstanceProfileMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsquery_serializeOpUntagInstanceProfile{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpUntagInstanceProfile{}, middleware.After)
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
	if err = addOpUntagInstanceProfileValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUntagInstanceProfile(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opUntagInstanceProfile(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "iam",
		OperationName: "UntagInstanceProfile",
	}
}
