// Code generated by smithy-go-codegen DO NOT EDIT.

package ssoadmin

import (
	"context"
	awsmiddleware "oss-sdk-go/oss/middleware"
	"oss-sdk-go/oss/signer/v4"
	"oss-sdk-go/service/ssoadmin/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Updates the IAM Identity Center identity store attributes that you can use with
// the IAM Identity Center instance for attributes-based access control (ABAC).
// When using an external identity provider as an identity source, you can pass
// attributes through the SAML assertion as an alternative to configuring
// attributes from the IAM Identity Center identity store. If a SAML assertion
// passes any of these attributes, IAM Identity Center replaces the attribute value
// with the value from the IAM Identity Center identity store. For more information
// about ABAC, see Attribute-Based Access Control in the IAM Identity Center User
// Guide.
func (c *Client) UpdateInstanceAccessControlAttributeConfiguration(ctx context.Context, params *UpdateInstanceAccessControlAttributeConfigurationInput, optFns ...func(*Options)) (*UpdateInstanceAccessControlAttributeConfigurationOutput, error) {
	if params == nil {
		params = &UpdateInstanceAccessControlAttributeConfigurationInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateInstanceAccessControlAttributeConfiguration", params, optFns, c.addOperationUpdateInstanceAccessControlAttributeConfigurationMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateInstanceAccessControlAttributeConfigurationOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateInstanceAccessControlAttributeConfigurationInput struct {

	// Updates the attributes for your ABAC configuration.
	//
	// This member is required.
	InstanceAccessControlAttributeConfiguration *types.InstanceAccessControlAttributeConfiguration

	// The ARN of the IAM Identity Center instance under which the operation will be
	// executed.
	//
	// This member is required.
	InstanceArn *string

	noSmithyDocumentSerde
}

type UpdateInstanceAccessControlAttributeConfigurationOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationUpdateInstanceAccessControlAttributeConfigurationMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpUpdateInstanceAccessControlAttributeConfiguration{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpUpdateInstanceAccessControlAttributeConfiguration{}, middleware.After)
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
	if err = addOpUpdateInstanceAccessControlAttributeConfigurationValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateInstanceAccessControlAttributeConfiguration(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opUpdateInstanceAccessControlAttributeConfiguration(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "sso",
		OperationName: "UpdateInstanceAccessControlAttributeConfiguration",
	}
}