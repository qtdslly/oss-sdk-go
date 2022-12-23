// Code generated by smithy-go-codegen DO NOT EDIT.

package worklink

import (
	"context"
	awsmiddleware "oss-sdk-go/oss/middleware"
	"oss-sdk-go/oss/signer/v4"
	"oss-sdk-go/service/worklink/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Updates the identity provider configuration for the fleet.
//
// Deprecated: Amazon WorkLink is no longer supported. This will be removed in a
// future version of the SDK.
func (c *Client) UpdateIdentityProviderConfiguration(ctx context.Context, params *UpdateIdentityProviderConfigurationInput, optFns ...func(*Options)) (*UpdateIdentityProviderConfigurationOutput, error) {
	if params == nil {
		params = &UpdateIdentityProviderConfigurationInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateIdentityProviderConfiguration", params, optFns, c.addOperationUpdateIdentityProviderConfigurationMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateIdentityProviderConfigurationOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateIdentityProviderConfigurationInput struct {

	// The ARN of the fleet.
	//
	// This member is required.
	FleetArn *string

	// The type of identity provider.
	//
	// This member is required.
	IdentityProviderType types.IdentityProviderType

	// The SAML metadata document provided by the customer’s identity provider. The
	// existing IdentityProviderSamlMetadata is unset if null is passed.
	IdentityProviderSamlMetadata *string

	noSmithyDocumentSerde
}

type UpdateIdentityProviderConfigurationOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationUpdateIdentityProviderConfigurationMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpUpdateIdentityProviderConfiguration{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpUpdateIdentityProviderConfiguration{}, middleware.After)
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
	if err = addOpUpdateIdentityProviderConfigurationValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateIdentityProviderConfiguration(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opUpdateIdentityProviderConfiguration(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "worklink",
		OperationName: "UpdateIdentityProviderConfiguration",
	}
}
