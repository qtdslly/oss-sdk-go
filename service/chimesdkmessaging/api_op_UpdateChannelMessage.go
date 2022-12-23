// Code generated by smithy-go-codegen DO NOT EDIT.

package chimesdkmessaging

import (
	"context"
	awsmiddleware "oss-sdk-go/oss/middleware"
	"oss-sdk-go/oss/signer/v4"
	"oss-sdk-go/service/chimesdkmessaging/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Updates the content of a message. The x-amz-chime-bearer request header is
// mandatory. Use the AppInstanceUserArn of the user that makes the API call as the
// value in the header.
func (c *Client) UpdateChannelMessage(ctx context.Context, params *UpdateChannelMessageInput, optFns ...func(*Options)) (*UpdateChannelMessageOutput, error) {
	if params == nil {
		params = &UpdateChannelMessageInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateChannelMessage", params, optFns, c.addOperationUpdateChannelMessageMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateChannelMessageOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateChannelMessageInput struct {

	// The ARN of the channel.
	//
	// This member is required.
	ChannelArn *string

	// The AppInstanceUserArn of the user that makes the API call.
	//
	// This member is required.
	ChimeBearer *string

	// The ID string of the message being updated.
	//
	// This member is required.
	MessageId *string

	// The content of the message being updated.
	Content *string

	// The metadata of the message being updated.
	Metadata *string

	// The ID of the SubChannel in the request. Only required when updating messages in
	// a SubChannel that the user belongs to.
	SubChannelId *string

	noSmithyDocumentSerde
}

type UpdateChannelMessageOutput struct {

	// The ARN of the channel.
	ChannelArn *string

	// The ID string of the message being updated.
	MessageId *string

	// The status of the message update.
	Status *types.ChannelMessageStatusStructure

	// The ID of the SubChannel in the response.
	SubChannelId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationUpdateChannelMessageMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpUpdateChannelMessage{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpUpdateChannelMessage{}, middleware.After)
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
	if err = addOpUpdateChannelMessageValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateChannelMessage(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opUpdateChannelMessage(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "chime",
		OperationName: "UpdateChannelMessage",
	}
}