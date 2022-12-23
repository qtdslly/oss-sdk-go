// Code generated by smithy-go-codegen DO NOT EDIT.

package location

import (
	"context"
	"fmt"
	awsmiddleware "oss-sdk-go/oss/middleware"
	"oss-sdk-go/oss/signer/v4"
	"oss-sdk-go/service/location/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Lists the latest device positions for requested devices.
func (c *Client) BatchGetDevicePosition(ctx context.Context, params *BatchGetDevicePositionInput, optFns ...func(*Options)) (*BatchGetDevicePositionOutput, error) {
	if params == nil {
		params = &BatchGetDevicePositionInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "BatchGetDevicePosition", params, optFns, c.addOperationBatchGetDevicePositionMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*BatchGetDevicePositionOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type BatchGetDevicePositionInput struct {

	// Devices whose position you want to retrieve.
	//
	// * For example, for two devices:
	// device-ids=DeviceId1&device-ids=DeviceId2
	//
	// This member is required.
	DeviceIds []string

	// The tracker resource retrieving the device position.
	//
	// This member is required.
	TrackerName *string

	noSmithyDocumentSerde
}

type BatchGetDevicePositionOutput struct {

	// Contains device position details such as the device ID, position, and timestamps
	// for when the position was received and sampled.
	//
	// This member is required.
	DevicePositions []types.DevicePosition

	// Contains error details for each device that failed to send its position to the
	// tracker resource.
	//
	// This member is required.
	Errors []types.BatchGetDevicePositionError

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationBatchGetDevicePositionMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpBatchGetDevicePosition{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpBatchGetDevicePosition{}, middleware.After)
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
	if err = addEndpointPrefix_opBatchGetDevicePositionMiddleware(stack); err != nil {
		return err
	}
	if err = addOpBatchGetDevicePositionValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opBatchGetDevicePosition(options.Region), middleware.Before); err != nil {
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

type endpointPrefix_opBatchGetDevicePositionMiddleware struct {
}

func (*endpointPrefix_opBatchGetDevicePositionMiddleware) ID() string {
	return "EndpointHostPrefix"
}

func (m *endpointPrefix_opBatchGetDevicePositionMiddleware) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
) {
	if smithyhttp.GetHostnameImmutable(ctx) || smithyhttp.IsEndpointHostPrefixDisabled(ctx) {
		return next.HandleSerialize(ctx, in)
	}

	req, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, fmt.Errorf("unknown transport type %T", in.Request)
	}

	req.URL.Host = "tracking." + req.URL.Host

	return next.HandleSerialize(ctx, in)
}
func addEndpointPrefix_opBatchGetDevicePositionMiddleware(stack *middleware.Stack) error {
	return stack.Serialize.Insert(&endpointPrefix_opBatchGetDevicePositionMiddleware{}, `OperationSerializer`, middleware.After)
}

func newServiceMetadataMiddleware_opBatchGetDevicePosition(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "geo",
		OperationName: "BatchGetDevicePosition",
	}
}
