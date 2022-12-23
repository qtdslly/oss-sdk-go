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

// Attaches the specified customer managed policy to the specified PermissionSet.
func (c *Client) AttachCustomerManagedPolicyReferenceToPermissionSet(ctx context.Context, params *AttachCustomerManagedPolicyReferenceToPermissionSetInput, optFns ...func(*Options)) (*AttachCustomerManagedPolicyReferenceToPermissionSetOutput, error) {
	if params == nil {
		params = &AttachCustomerManagedPolicyReferenceToPermissionSetInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "AttachCustomerManagedPolicyReferenceToPermissionSet", params, optFns, c.addOperationAttachCustomerManagedPolicyReferenceToPermissionSetMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*AttachCustomerManagedPolicyReferenceToPermissionSetOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type AttachCustomerManagedPolicyReferenceToPermissionSetInput struct {

	// Specifies the name and path of a customer managed policy. You must have an IAM
	// policy that matches the name and path in each AWS account where you want to
	// deploy your permission set.
	//
	// This member is required.
	CustomerManagedPolicyReference *types.CustomerManagedPolicyReference

	// The ARN of the IAM Identity Center instance under which the operation will be
	// executed.
	//
	// This member is required.
	InstanceArn *string

	// The ARN of the PermissionSet.
	//
	// This member is required.
	PermissionSetArn *string

	noSmithyDocumentSerde
}

type AttachCustomerManagedPolicyReferenceToPermissionSetOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationAttachCustomerManagedPolicyReferenceToPermissionSetMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpAttachCustomerManagedPolicyReferenceToPermissionSet{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpAttachCustomerManagedPolicyReferenceToPermissionSet{}, middleware.After)
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
	if err = addOpAttachCustomerManagedPolicyReferenceToPermissionSetValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opAttachCustomerManagedPolicyReferenceToPermissionSet(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opAttachCustomerManagedPolicyReferenceToPermissionSet(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "sso",
		OperationName: "AttachCustomerManagedPolicyReferenceToPermissionSet",
	}
}
