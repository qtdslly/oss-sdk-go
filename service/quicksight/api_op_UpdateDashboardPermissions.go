// Code generated by smithy-go-codegen DO NOT EDIT.

package quicksight

import (
	"context"
	awsmiddleware "oss-sdk-go/oss/middleware"
	"oss-sdk-go/oss/signer/v4"
	"oss-sdk-go/service/quicksight/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Updates read and write permissions on a dashboard.
func (c *Client) UpdateDashboardPermissions(ctx context.Context, params *UpdateDashboardPermissionsInput, optFns ...func(*Options)) (*UpdateDashboardPermissionsOutput, error) {
	if params == nil {
		params = &UpdateDashboardPermissionsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateDashboardPermissions", params, optFns, c.addOperationUpdateDashboardPermissionsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateDashboardPermissionsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateDashboardPermissionsInput struct {

	// The ID of the Amazon Web Services account that contains the dashboard whose
	// permissions you're updating.
	//
	// This member is required.
	AwsAccountId *string

	// The ID for the dashboard.
	//
	// This member is required.
	DashboardId *string

	// Grants link permissions to all users in a defined namespace.
	GrantLinkPermissions []types.ResourcePermission

	// The permissions that you want to grant on this resource.
	GrantPermissions []types.ResourcePermission

	// Revokes link permissions from all users in a defined namespace.
	RevokeLinkPermissions []types.ResourcePermission

	// The permissions that you want to revoke from this resource.
	RevokePermissions []types.ResourcePermission

	noSmithyDocumentSerde
}

type UpdateDashboardPermissionsOutput struct {

	// The Amazon Resource Name (ARN) of the dashboard.
	DashboardArn *string

	// The ID for the dashboard.
	DashboardId *string

	// Updates the permissions of a shared link to an Amazon QuickSight dashboard.
	LinkSharingConfiguration *types.LinkSharingConfiguration

	// Information about the permissions on the dashboard.
	Permissions []types.ResourcePermission

	// The Amazon Web Services request ID for this operation.
	RequestId *string

	// The HTTP status of the request.
	Status int32

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationUpdateDashboardPermissionsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpUpdateDashboardPermissions{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpUpdateDashboardPermissions{}, middleware.After)
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
	if err = addOpUpdateDashboardPermissionsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateDashboardPermissions(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opUpdateDashboardPermissions(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "quicksight",
		OperationName: "UpdateDashboardPermissions",
	}
}