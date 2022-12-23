// Code generated by smithy-go-codegen DO NOT EDIT.

package transcribe

import (
	"context"
	awsmiddleware "oss-sdk-go/oss/middleware"
	"oss-sdk-go/oss/signer/v4"
	"oss-sdk-go/service/transcribe/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Provides information about the specified custom language model. This operation
// also shows if the base language model you used to create your custom language
// model has been updated. If Amazon Transcribe has updated the base model, you can
// create a new custom language model using the updated base model. If you tried to
// create a new custom language model and the request wasn't successful, you can
// use DescribeLanguageModel to help identify the reason for this failure. To get a
// list of your custom language models, use the operation.
func (c *Client) DescribeLanguageModel(ctx context.Context, params *DescribeLanguageModelInput, optFns ...func(*Options)) (*DescribeLanguageModelOutput, error) {
	if params == nil {
		params = &DescribeLanguageModelInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeLanguageModel", params, optFns, c.addOperationDescribeLanguageModelMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeLanguageModelOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeLanguageModelInput struct {

	// The name of the custom language model you want information about. Model names
	// are case sensitive.
	//
	// This member is required.
	ModelName *string

	noSmithyDocumentSerde
}

type DescribeLanguageModelOutput struct {

	// Provides information about the specified custom language model. This parameter
	// also shows if the base language model you used to create your custom language
	// model has been updated. If Amazon Transcribe has updated the base model, you can
	// create a new custom language model using the updated base model. If you tried to
	// create a new custom language model and the request wasn't successful, you can
	// use this DescribeLanguageModel to help identify the reason for this failure.
	LanguageModel *types.LanguageModel

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeLanguageModelMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDescribeLanguageModel{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDescribeLanguageModel{}, middleware.After)
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
	if err = addOpDescribeLanguageModelValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeLanguageModel(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDescribeLanguageModel(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "transcribe",
		OperationName: "DescribeLanguageModel",
	}
}