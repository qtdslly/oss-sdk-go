// Code generated by smithy-go-codegen DO NOT EDIT.

package cognitoidentityprovider

import (
	"context"
	awsmiddleware "oss-sdk-go/oss/middleware"
	"oss-sdk-go/oss/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Removes the specified user from the specified group. Calling this action
// requires developer credentials.
func (c *Client) AdminRemoveUserFromGroup(ctx context.Context, params *AdminRemoveUserFromGroupInput, optFns ...func(*Options)) (*AdminRemoveUserFromGroupOutput, error) {
	if params == nil {
		params = &AdminRemoveUserFromGroupInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "AdminRemoveUserFromGroup", params, optFns, c.addOperationAdminRemoveUserFromGroupMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*AdminRemoveUserFromGroupOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type AdminRemoveUserFromGroupInput struct {

	// The group name.
	//
	// This member is required.
	GroupName *string

	// The user pool ID for the user pool.
	//
	// This member is required.
	UserPoolId *string

	// The username for the user.
	//
	// This member is required.
	Username *string

	noSmithyDocumentSerde
}

type AdminRemoveUserFromGroupOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationAdminRemoveUserFromGroupMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpAdminRemoveUserFromGroup{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpAdminRemoveUserFromGroup{}, middleware.After)
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
	if err = addOpAdminRemoveUserFromGroupValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opAdminRemoveUserFromGroup(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opAdminRemoveUserFromGroup(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "cognito-idp",
		OperationName: "AdminRemoveUserFromGroup",
	}
}
