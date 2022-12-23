// Code generated by smithy-go-codegen DO NOT EDIT.

package chime

import (
	"context"
	awsmiddleware "oss-sdk-go/oss/middleware"
	"oss-sdk-go/oss/signer/v4"
	"oss-sdk-go/service/chime/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates a SIP media application.
func (c *Client) CreateSipMediaApplication(ctx context.Context, params *CreateSipMediaApplicationInput, optFns ...func(*Options)) (*CreateSipMediaApplicationOutput, error) {
	if params == nil {
		params = &CreateSipMediaApplicationInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateSipMediaApplication", params, optFns, c.addOperationCreateSipMediaApplicationMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateSipMediaApplicationOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateSipMediaApplicationInput struct {

	// The AWS Region assigned to the SIP media application.
	//
	// This member is required.
	AwsRegion *string

	// List of endpoints (Lambda Amazon Resource Names) specified for the SIP media
	// application. Currently, only one endpoint is supported.
	//
	// This member is required.
	Endpoints []types.SipMediaApplicationEndpoint

	// The SIP media application name.
	//
	// This member is required.
	Name *string

	noSmithyDocumentSerde
}

type CreateSipMediaApplicationOutput struct {

	// The SIP media application details.
	SipMediaApplication *types.SipMediaApplication

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateSipMediaApplicationMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpCreateSipMediaApplication{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpCreateSipMediaApplication{}, middleware.After)
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
	if err = addOpCreateSipMediaApplicationValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateSipMediaApplication(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateSipMediaApplication(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "chime",
		OperationName: "CreateSipMediaApplication",
	}
}