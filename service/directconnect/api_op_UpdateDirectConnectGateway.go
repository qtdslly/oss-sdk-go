// Code generated by smithy-go-codegen DO NOT EDIT.

package directconnect

import (
	"context"
	awsmiddleware "oss-sdk-go/oss/middleware"
	"oss-sdk-go/oss/signer/v4"
	"oss-sdk-go/service/directconnect/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Updates the name of a current Direct Connect gateway.
func (c *Client) UpdateDirectConnectGateway(ctx context.Context, params *UpdateDirectConnectGatewayInput, optFns ...func(*Options)) (*UpdateDirectConnectGatewayOutput, error) {
	if params == nil {
		params = &UpdateDirectConnectGatewayInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateDirectConnectGateway", params, optFns, c.addOperationUpdateDirectConnectGatewayMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateDirectConnectGatewayOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateDirectConnectGatewayInput struct {

	// The ID of the Direct Connect gateway to update.
	//
	// This member is required.
	DirectConnectGatewayId *string

	// The new name for the Direct Connect gateway.
	//
	// This member is required.
	NewDirectConnectGatewayName *string

	noSmithyDocumentSerde
}

type UpdateDirectConnectGatewayOutput struct {

	// Information about a Direct Connect gateway, which enables you to connect virtual
	// interfaces and virtual private gateway or transit gateways.
	DirectConnectGateway *types.DirectConnectGateway

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationUpdateDirectConnectGatewayMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpUpdateDirectConnectGateway{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpUpdateDirectConnectGateway{}, middleware.After)
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
	if err = addOpUpdateDirectConnectGatewayValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateDirectConnectGateway(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opUpdateDirectConnectGateway(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "directconnect",
		OperationName: "UpdateDirectConnectGateway",
	}
}