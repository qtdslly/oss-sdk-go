// Code generated by smithy-go-codegen DO NOT EDIT.

package finspacedata

import (
	"context"
	awsmiddleware "oss-sdk-go/oss/middleware"
	"oss-sdk-go/oss/signer/v4"
	"oss-sdk-go/service/finspacedata/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Lists all the permission groups that are associated with a specific user
// account.
func (c *Client) ListPermissionGroupsByUser(ctx context.Context, params *ListPermissionGroupsByUserInput, optFns ...func(*Options)) (*ListPermissionGroupsByUserOutput, error) {
	if params == nil {
		params = &ListPermissionGroupsByUserInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListPermissionGroupsByUser", params, optFns, c.addOperationListPermissionGroupsByUserMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListPermissionGroupsByUserOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListPermissionGroupsByUserInput struct {

	// The maximum number of results per page.
	//
	// This member is required.
	MaxResults int32

	// The unique identifier for the user.
	//
	// This member is required.
	UserId *string

	// A token that indicates where a results page should begin.
	NextToken *string

	noSmithyDocumentSerde
}

type ListPermissionGroupsByUserOutput struct {

	// A token that indicates where a results page should begin.
	NextToken *string

	// A list of returned permission groups.
	PermissionGroups []types.PermissionGroupByUser

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListPermissionGroupsByUserMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListPermissionGroupsByUser{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListPermissionGroupsByUser{}, middleware.After)
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
	if err = addRestJsonContentTypeCustomization(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddErrorCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = addOpListPermissionGroupsByUserValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListPermissionGroupsByUser(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opListPermissionGroupsByUser(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "finspace-api",
		OperationName: "ListPermissionGroupsByUser",
	}
}
