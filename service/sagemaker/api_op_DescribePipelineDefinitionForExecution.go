// Code generated by smithy-go-codegen DO NOT EDIT.

package sagemaker

import (
	"context"
	awsmiddleware "oss-sdk-go/oss/middleware"
	"oss-sdk-go/oss/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Describes the details of an execution's pipeline definition.
func (c *Client) DescribePipelineDefinitionForExecution(ctx context.Context, params *DescribePipelineDefinitionForExecutionInput, optFns ...func(*Options)) (*DescribePipelineDefinitionForExecutionOutput, error) {
	if params == nil {
		params = &DescribePipelineDefinitionForExecutionInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribePipelineDefinitionForExecution", params, optFns, c.addOperationDescribePipelineDefinitionForExecutionMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribePipelineDefinitionForExecutionOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribePipelineDefinitionForExecutionInput struct {

	// The Amazon Resource Name (ARN) of the pipeline execution.
	//
	// This member is required.
	PipelineExecutionArn *string

	noSmithyDocumentSerde
}

type DescribePipelineDefinitionForExecutionOutput struct {

	// The time when the pipeline was created.
	CreationTime *time.Time

	// The JSON pipeline definition.
	PipelineDefinition *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribePipelineDefinitionForExecutionMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDescribePipelineDefinitionForExecution{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDescribePipelineDefinitionForExecution{}, middleware.After)
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
	if err = addOpDescribePipelineDefinitionForExecutionValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribePipelineDefinitionForExecution(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDescribePipelineDefinitionForExecution(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "sagemaker",
		OperationName: "DescribePipelineDefinitionForExecution",
	}
}
