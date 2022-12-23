// Code generated by smithy-go-codegen DO NOT EDIT.

package kinesisvideo

import (
	"context"
	awsmiddleware "oss-sdk-go/oss/middleware"
	"oss-sdk-go/oss/signer/v4"
	"oss-sdk-go/service/kinesisvideo/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Updates the existing signaling channel. This is an asynchronous operation and
// takes time to complete. If the MessageTtlSeconds value is updated (either
// increased or reduced), it only applies to new messages sent via this channel
// after it's been updated. Existing messages are still expired as per the previous
// MessageTtlSeconds value.
func (c *Client) UpdateSignalingChannel(ctx context.Context, params *UpdateSignalingChannelInput, optFns ...func(*Options)) (*UpdateSignalingChannelOutput, error) {
	if params == nil {
		params = &UpdateSignalingChannelInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateSignalingChannel", params, optFns, c.addOperationUpdateSignalingChannelMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateSignalingChannelOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateSignalingChannelInput struct {

	// The Amazon Resource Name (ARN) of the signaling channel that you want to update.
	//
	// This member is required.
	ChannelARN *string

	// The current version of the signaling channel that you want to update.
	//
	// This member is required.
	CurrentVersion *string

	// The structure containing the configuration for the SINGLE_MASTER type of the
	// signaling channel that you want to update.
	SingleMasterConfiguration *types.SingleMasterConfiguration

	noSmithyDocumentSerde
}

type UpdateSignalingChannelOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationUpdateSignalingChannelMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpUpdateSignalingChannel{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpUpdateSignalingChannel{}, middleware.After)
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
	if err = addOpUpdateSignalingChannelValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateSignalingChannel(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opUpdateSignalingChannel(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "kinesisvideo",
		OperationName: "UpdateSignalingChannel",
	}
}