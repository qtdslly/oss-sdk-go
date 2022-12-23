// Code generated by smithy-go-codegen DO NOT EDIT.

package iotsecuretunneling

import (
	"context"
	awsmiddleware "oss-sdk-go/oss/middleware"
	"oss-sdk-go/oss/signer/v4"
	"oss-sdk-go/service/iotsecuretunneling/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Revokes the current client access token (CAT) and returns new CAT for clients to
// use when reconnecting to secure tunneling to access the same tunnel. Requires
// permission to access the RotateTunnelAccessToken
// (https://docs.aws.amazon.com/service-authorization/latest/reference/list_awsiot.html#awsiot-actions-as-permissions)
// action. Rotating the CAT doesn't extend the tunnel duration. For example, say
// the tunnel duration is 12 hours and the tunnel has already been open for 4
// hours. When you rotate the access tokens, the new tokens that are generated can
// only be used for the remaining 8 hours.
func (c *Client) RotateTunnelAccessToken(ctx context.Context, params *RotateTunnelAccessTokenInput, optFns ...func(*Options)) (*RotateTunnelAccessTokenOutput, error) {
	if params == nil {
		params = &RotateTunnelAccessTokenInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "RotateTunnelAccessToken", params, optFns, c.addOperationRotateTunnelAccessTokenMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*RotateTunnelAccessTokenOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type RotateTunnelAccessTokenInput struct {

	// The mode of the client that will use the client token, which can be either the
	// source or destination, or both source and destination.
	//
	// This member is required.
	ClientMode types.ClientMode

	// The tunnel for which you want to rotate the access tokens.
	//
	// This member is required.
	TunnelId *string

	// The destination configuration.
	DestinationConfig *types.DestinationConfig

	noSmithyDocumentSerde
}

type RotateTunnelAccessTokenOutput struct {

	// The client access token that the destination local proxy uses to connect to IoT
	// Secure Tunneling.
	DestinationAccessToken *string

	// The client access token that the source local proxy uses to connect to IoT
	// Secure Tunneling.
	SourceAccessToken *string

	// The Amazon Resource Name for the tunnel.
	TunnelArn *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationRotateTunnelAccessTokenMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpRotateTunnelAccessToken{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpRotateTunnelAccessToken{}, middleware.After)
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
	if err = addOpRotateTunnelAccessTokenValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opRotateTunnelAccessToken(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opRotateTunnelAccessToken(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "IoTSecuredTunneling",
		OperationName: "RotateTunnelAccessToken",
	}
}
