// Code generated by smithy-go-codegen DO NOT EDIT.

package iotwireless

import (
	"context"
	awsmiddleware "oss-sdk-go/oss/middleware"
	"oss-sdk-go/oss/signer/v4"
	"oss-sdk-go/service/iotwireless/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Get the event configuration based on resource types.
func (c *Client) GetEventConfigurationByResourceTypes(ctx context.Context, params *GetEventConfigurationByResourceTypesInput, optFns ...func(*Options)) (*GetEventConfigurationByResourceTypesOutput, error) {
	if params == nil {
		params = &GetEventConfigurationByResourceTypesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetEventConfigurationByResourceTypes", params, optFns, c.addOperationGetEventConfigurationByResourceTypesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetEventConfigurationByResourceTypesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetEventConfigurationByResourceTypesInput struct {
	noSmithyDocumentSerde
}

type GetEventConfigurationByResourceTypesOutput struct {

	// Resource type event configuration for the connection status event.
	ConnectionStatus *types.ConnectionStatusResourceTypeEventConfiguration

	// Resource type event configuration for the device registration state event.
	DeviceRegistrationState *types.DeviceRegistrationStateResourceTypeEventConfiguration

	// Resource type event configuration for the join event.
	Join *types.JoinResourceTypeEventConfiguration

	// Resource type event configuration object for the message delivery status event.
	MessageDeliveryStatus *types.MessageDeliveryStatusResourceTypeEventConfiguration

	// Resource type event configuration for the proximity event.
	Proximity *types.ProximityResourceTypeEventConfiguration

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetEventConfigurationByResourceTypesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpGetEventConfigurationByResourceTypes{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpGetEventConfigurationByResourceTypes{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetEventConfigurationByResourceTypes(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opGetEventConfigurationByResourceTypes(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "iotwireless",
		OperationName: "GetEventConfigurationByResourceTypes",
	}
}