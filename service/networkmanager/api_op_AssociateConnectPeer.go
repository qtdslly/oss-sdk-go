// Code generated by smithy-go-codegen DO NOT EDIT.

package networkmanager

import (
	"context"
	awsmiddleware "oss-sdk-go/oss/middleware"
	"oss-sdk-go/oss/signer/v4"
	"oss-sdk-go/service/networkmanager/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Associates a core network Connect peer with a device and optionally, with a
// link. If you specify a link, it must be associated with the specified device.
// You can only associate core network Connect peers that have been created on a
// core network Connect attachment on a core network.
func (c *Client) AssociateConnectPeer(ctx context.Context, params *AssociateConnectPeerInput, optFns ...func(*Options)) (*AssociateConnectPeerOutput, error) {
	if params == nil {
		params = &AssociateConnectPeerInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "AssociateConnectPeer", params, optFns, c.addOperationAssociateConnectPeerMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*AssociateConnectPeerOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type AssociateConnectPeerInput struct {

	// The ID of the Connect peer.
	//
	// This member is required.
	ConnectPeerId *string

	// The ID of the device.
	//
	// This member is required.
	DeviceId *string

	// The ID of your global network.
	//
	// This member is required.
	GlobalNetworkId *string

	// The ID of the link.
	LinkId *string

	noSmithyDocumentSerde
}

type AssociateConnectPeerOutput struct {

	// The response to the Connect peer request.
	ConnectPeerAssociation *types.ConnectPeerAssociation

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationAssociateConnectPeerMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpAssociateConnectPeer{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpAssociateConnectPeer{}, middleware.After)
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
	if err = addOpAssociateConnectPeerValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opAssociateConnectPeer(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opAssociateConnectPeer(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "networkmanager",
		OperationName: "AssociateConnectPeer",
	}
}
