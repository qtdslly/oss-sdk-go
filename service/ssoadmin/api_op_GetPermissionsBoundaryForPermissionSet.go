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

// Obtains the permissions boundary for a specified PermissionSet.
func (c *Client) GetPermissionsBoundaryForPermissionSet(ctx context.Context, params *GetPermissionsBoundaryForPermissionSetInput, optFns ...func(*Options)) (*GetPermissionsBoundaryForPermissionSetOutput, error) {
	if params == nil {
		params = &GetPermissionsBoundaryForPermissionSetInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetPermissionsBoundaryForPermissionSet", params, optFns, c.addOperationGetPermissionsBoundaryForPermissionSetMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetPermissionsBoundaryForPermissionSetOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetPermissionsBoundaryForPermissionSetInput struct {

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

type GetPermissionsBoundaryForPermissionSetOutput struct {

	// The permissions boundary attached to the specified permission set.
	PermissionsBoundary *types.PermissionsBoundary

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetPermissionsBoundaryForPermissionSetMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpGetPermissionsBoundaryForPermissionSet{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpGetPermissionsBoundaryForPermissionSet{}, middleware.After)
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
	if err = addOpGetPermissionsBoundaryForPermissionSetValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetPermissionsBoundaryForPermissionSet(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opGetPermissionsBoundaryForPermissionSet(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "sso",
		OperationName: "GetPermissionsBoundaryForPermissionSet",
	}
}