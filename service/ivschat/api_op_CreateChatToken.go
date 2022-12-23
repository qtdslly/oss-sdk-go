// Code generated by smithy-go-codegen DO NOT EDIT.

package ivschat

import (
	"context"
	awsmiddleware "oss-sdk-go/oss/middleware"
	"oss-sdk-go/oss/signer/v4"
	"oss-sdk-go/service/ivschat/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Creates an encrypted token that is used to establish an individual WebSocket
// connection to a room. The token is valid for one minute, and a connection
// (session) established with the token is valid for the specified duration.
// Encryption keys are owned by Amazon IVS Chat and never used directly by your
// application.
func (c *Client) CreateChatToken(ctx context.Context, params *CreateChatTokenInput, optFns ...func(*Options)) (*CreateChatTokenOutput, error) {
	if params == nil {
		params = &CreateChatTokenInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateChatToken", params, optFns, c.addOperationCreateChatTokenMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateChatTokenOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateChatTokenInput struct {

	// Identifier of the room that the client is trying to access. Currently this must
	// be an ARN.
	//
	// This member is required.
	RoomIdentifier *string

	// Application-provided ID that uniquely identifies the user associated with this
	// token. This can be any UTF-8 encoded text.
	//
	// This member is required.
	UserId *string

	// Application-provided attributes to encode into the token and attach to a chat
	// session. Map keys and values can contain UTF-8 encoded text. The maximum length
	// of this field is 1 KB total.
	Attributes map[string]string

	// Set of capabilities that the user is allowed to perform in the room. Default:
	// None (the capability to view messages is implicitly included in all requests).
	Capabilities []types.ChatTokenCapability

	// Session duration (in minutes), after which the session expires. Default: 60 (1
	// hour).
	SessionDurationInMinutes int32

	noSmithyDocumentSerde
}

type CreateChatTokenOutput struct {

	// Time after which an end user's session is no longer valid. This is an ISO 8601
	// timestamp; note that this is returned as a string.
	SessionExpirationTime *time.Time

	// The issued client token, encrypted.
	Token *string

	// Time after which the token is no longer valid and cannot be used to connect to a
	// room. This is an ISO 8601 timestamp; note that this is returned as a string.
	TokenExpirationTime *time.Time

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateChatTokenMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpCreateChatToken{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpCreateChatToken{}, middleware.After)
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
	if err = addOpCreateChatTokenValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateChatToken(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateChatToken(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ivschat",
		OperationName: "CreateChatToken",
	}
}
