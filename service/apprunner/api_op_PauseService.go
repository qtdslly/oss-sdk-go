// Code generated by smithy-go-codegen DO NOT EDIT.

package apprunner

import (
	"context"
	awsmiddleware "oss-sdk-go/oss/middleware"
	"oss-sdk-go/oss/signer/v4"
	"oss-sdk-go/service/apprunner/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Pause an active App Runner service. App Runner reduces compute capacity for the
// service to zero and loses state (for example, ephemeral storage is removed).
// This is an asynchronous operation. On a successful call, you can use the
// returned OperationId and the ListOperations call to track the operation's
// progress.
func (c *Client) PauseService(ctx context.Context, params *PauseServiceInput, optFns ...func(*Options)) (*PauseServiceOutput, error) {
	if params == nil {
		params = &PauseServiceInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "PauseService", params, optFns, c.addOperationPauseServiceMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*PauseServiceOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type PauseServiceInput struct {

	// The Amazon Resource Name (ARN) of the App Runner service that you want to pause.
	//
	// This member is required.
	ServiceArn *string

	noSmithyDocumentSerde
}

type PauseServiceOutput struct {

	// A description of the App Runner service that this request just paused.
	//
	// This member is required.
	Service *types.Service

	// The unique ID of the asynchronous operation that this request started. You can
	// use it combined with the ListOperations call to track the operation's progress.
	OperationId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationPauseServiceMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson10_serializeOpPauseService{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson10_deserializeOpPauseService{}, middleware.After)
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
	if err = addOpPauseServiceValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opPauseService(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opPauseService(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "apprunner",
		OperationName: "PauseService",
	}
}
