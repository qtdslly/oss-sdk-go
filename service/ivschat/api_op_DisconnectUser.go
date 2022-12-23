// Code generated by smithy-go-codegen DO NOT EDIT.

package ivschat

import (
	"context"
	awsmiddleware "oss-sdk-go/oss/middleware"
	"oss-sdk-go/oss/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Disconnects all connections using a specified user ID from a room. This
// replicates the  DisconnectUser
// (https://docs.aws.amazon.com/ivs/latest/chatmsgapireference/actions-disconnectuser-publish.html)
// WebSocket operation in the Amazon IVS Chat Messaging API.
func (c *Client) DisconnectUser(ctx context.Context, params *DisconnectUserInput, optFns ...func(*Options)) (*DisconnectUserOutput, error) {
	if params == nil {
		params = &DisconnectUserInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DisconnectUser", params, optFns, c.addOperationDisconnectUserMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DisconnectUserOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DisconnectUserInput struct {

	// Identifier of the room from which the user's clients should be disconnected.
	// Currently this must be an ARN.
	//
	// This member is required.
	RoomIdentifier *string

	// ID of the user (connection) to disconnect from the room.
	//
	// This member is required.
	UserId *string

	// Reason for disconnecting the user.
	Reason *string

	noSmithyDocumentSerde
}

type DisconnectUserOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDisconnectUserMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpDisconnectUser{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpDisconnectUser{}, middleware.After)
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
	if err = addOpDisconnectUserValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDisconnectUser(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDisconnectUser(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ivschat",
		OperationName: "DisconnectUser",
	}
}
