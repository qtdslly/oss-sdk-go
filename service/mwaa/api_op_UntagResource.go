// Code generated by smithy-go-codegen DO NOT EDIT.

package mwaa

import (
	"context"
	"fmt"
	awsmiddleware "oss-sdk-go/oss/middleware"
	"oss-sdk-go/oss/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Removes key-value tag pairs associated to your Amazon Managed Workflows for
// Apache Airflow (MWAA) environment. For example, "Environment": "Staging".
func (c *Client) UntagResource(ctx context.Context, params *UntagResourceInput, optFns ...func(*Options)) (*UntagResourceOutput, error) {
	if params == nil {
		params = &UntagResourceInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UntagResource", params, optFns, c.addOperationUntagResourceMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UntagResourceOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UntagResourceInput struct {

	// The Amazon Resource Name (ARN) of the Amazon MWAA environment. For example,
	// arn:aws:airflow:us-east-1:123456789012:environment/MyMWAAEnvironment.
	//
	// This member is required.
	ResourceArn *string

	// The key-value tag pair you want to remove. For example, "Environment":
	// "Staging".
	//
	// This member is required.
	TagKeys []string

	noSmithyDocumentSerde
}

type UntagResourceOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationUntagResourceMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpUntagResource{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpUntagResource{}, middleware.After)
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
	if err = addEndpointPrefix_opUntagResourceMiddleware(stack); err != nil {
		return err
	}
	if err = addOpUntagResourceValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUntagResource(options.Region), middleware.Before); err != nil {
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

type endpointPrefix_opUntagResourceMiddleware struct {
}

func (*endpointPrefix_opUntagResourceMiddleware) ID() string {
	return "EndpointHostPrefix"
}

func (m *endpointPrefix_opUntagResourceMiddleware) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
) {
	if smithyhttp.GetHostnameImmutable(ctx) || smithyhttp.IsEndpointHostPrefixDisabled(ctx) {
		return next.HandleSerialize(ctx, in)
	}

	req, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, fmt.Errorf("unknown transport type %T", in.Request)
	}

	req.URL.Host = "api." + req.URL.Host

	return next.HandleSerialize(ctx, in)
}
func addEndpointPrefix_opUntagResourceMiddleware(stack *middleware.Stack) error {
	return stack.Serialize.Insert(&endpointPrefix_opUntagResourceMiddleware{}, `OperationSerializer`, middleware.After)
}

func newServiceMetadataMiddleware_opUntagResource(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "airflow",
		OperationName: "UntagResource",
	}
}
