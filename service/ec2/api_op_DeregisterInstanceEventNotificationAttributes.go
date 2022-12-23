// Code generated by smithy-go-codegen DO NOT EDIT.

package ec2

import (
	"context"
	awsmiddleware "oss-sdk-go/oss/middleware"
	"oss-sdk-go/oss/signer/v4"
	"oss-sdk-go/service/ec2/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Deregisters tag keys to prevent tags that have the specified tag keys from being
// included in scheduled event notifications for resources in the Region.
func (c *Client) DeregisterInstanceEventNotificationAttributes(ctx context.Context, params *DeregisterInstanceEventNotificationAttributesInput, optFns ...func(*Options)) (*DeregisterInstanceEventNotificationAttributesOutput, error) {
	if params == nil {
		params = &DeregisterInstanceEventNotificationAttributesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DeregisterInstanceEventNotificationAttributes", params, optFns, c.addOperationDeregisterInstanceEventNotificationAttributesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DeregisterInstanceEventNotificationAttributesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DeregisterInstanceEventNotificationAttributesInput struct {

	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have the
	// required permissions, the error response is DryRunOperation. Otherwise, it is
	// UnauthorizedOperation.
	DryRun *bool

	// Information about the tag keys to deregister.
	InstanceTagAttribute *types.DeregisterInstanceTagAttributeRequest

	noSmithyDocumentSerde
}

type DeregisterInstanceEventNotificationAttributesOutput struct {

	// The resulting set of tag keys.
	InstanceTagAttribute *types.InstanceTagNotificationAttribute

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDeregisterInstanceEventNotificationAttributesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsEc2query_serializeOpDeregisterInstanceEventNotificationAttributes{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsEc2query_deserializeOpDeregisterInstanceEventNotificationAttributes{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDeregisterInstanceEventNotificationAttributes(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDeregisterInstanceEventNotificationAttributes(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ec2",
		OperationName: "DeregisterInstanceEventNotificationAttributes",
	}
}
