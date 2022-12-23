// Code generated by smithy-go-codegen DO NOT EDIT.

package sagemaker

import (
	"context"
	awsmiddleware "oss-sdk-go/oss/middleware"
	"oss-sdk-go/oss/signer/v4"
	"oss-sdk-go/service/sagemaker/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Describes the details of a pipeline.
func (c *Client) DescribePipeline(ctx context.Context, params *DescribePipelineInput, optFns ...func(*Options)) (*DescribePipelineOutput, error) {
	if params == nil {
		params = &DescribePipelineInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribePipeline", params, optFns, c.addOperationDescribePipelineMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribePipelineOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribePipelineInput struct {

	// The name of the pipeline to describe.
	//
	// This member is required.
	PipelineName *string

	noSmithyDocumentSerde
}

type DescribePipelineOutput struct {

	// Information about the user who created or modified an experiment, trial, trial
	// component, lineage group, or project.
	CreatedBy *types.UserContext

	// The time when the pipeline was created.
	CreationTime *time.Time

	// Information about the user who created or modified an experiment, trial, trial
	// component, lineage group, or project.
	LastModifiedBy *types.UserContext

	// The time when the pipeline was last modified.
	LastModifiedTime *time.Time

	// The time when the pipeline was last run.
	LastRunTime *time.Time

	// Lists the parallelism configuration applied to the pipeline.
	ParallelismConfiguration *types.ParallelismConfiguration

	// The Amazon Resource Name (ARN) of the pipeline.
	PipelineArn *string

	// The JSON pipeline definition.
	PipelineDefinition *string

	// The description of the pipeline.
	PipelineDescription *string

	// The display name of the pipeline.
	PipelineDisplayName *string

	// The name of the pipeline.
	PipelineName *string

	// The status of the pipeline execution.
	PipelineStatus types.PipelineStatus

	// The Amazon Resource Name (ARN) that the pipeline uses to execute.
	RoleArn *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribePipelineMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDescribePipeline{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDescribePipeline{}, middleware.After)
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
	if err = addOpDescribePipelineValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribePipeline(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDescribePipeline(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "sagemaker",
		OperationName: "DescribePipeline",
	}
}
