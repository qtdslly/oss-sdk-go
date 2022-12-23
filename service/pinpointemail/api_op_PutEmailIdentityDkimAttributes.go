// Code generated by smithy-go-codegen DO NOT EDIT.

package pinpointemail

import (
	"context"
	awsmiddleware "oss-sdk-go/oss/middleware"
	"oss-sdk-go/oss/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Used to enable or disable DKIM authentication for an email identity.
func (c *Client) PutEmailIdentityDkimAttributes(ctx context.Context, params *PutEmailIdentityDkimAttributesInput, optFns ...func(*Options)) (*PutEmailIdentityDkimAttributesOutput, error) {
	if params == nil {
		params = &PutEmailIdentityDkimAttributesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "PutEmailIdentityDkimAttributes", params, optFns, c.addOperationPutEmailIdentityDkimAttributesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*PutEmailIdentityDkimAttributesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// A request to enable or disable DKIM signing of email that you send from an email
// identity.
type PutEmailIdentityDkimAttributesInput struct {

	// The email identity that you want to change the DKIM settings for.
	//
	// This member is required.
	EmailIdentity *string

	// Sets the DKIM signing configuration for the identity. When you set this value
	// true, then the messages that Amazon Pinpoint sends from the identity are
	// DKIM-signed. When you set this value to false, then the messages that Amazon
	// Pinpoint sends from the identity aren't DKIM-signed.
	SigningEnabled bool

	noSmithyDocumentSerde
}

// An HTTP 200 response if the request succeeds, or an error message if the request
// fails.
type PutEmailIdentityDkimAttributesOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationPutEmailIdentityDkimAttributesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpPutEmailIdentityDkimAttributes{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpPutEmailIdentityDkimAttributes{}, middleware.After)
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
	if err = addOpPutEmailIdentityDkimAttributesValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opPutEmailIdentityDkimAttributes(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opPutEmailIdentityDkimAttributes(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ses",
		OperationName: "PutEmailIdentityDkimAttributes",
	}
}
