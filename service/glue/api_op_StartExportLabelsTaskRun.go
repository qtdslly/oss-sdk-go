// Code generated by smithy-go-codegen DO NOT EDIT.

package glue

import (
	"context"
	awsmiddleware "oss-sdk-go/oss/middleware"
	"oss-sdk-go/oss/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Begins an asynchronous task to export all labeled data for a particular
// transform. This task is the only label-related API call that is not part of the
// typical active learning workflow. You typically use StartExportLabelsTaskRun
// when you want to work with all of your existing labels at the same time, such as
// when you want to remove or change labels that were previously submitted as
// truth. This API operation accepts the TransformId whose labels you want to
// export and an Amazon Simple Storage Service (Amazon S3) path to export the
// labels to. The operation returns a TaskRunId. You can check on the status of
// your task run by calling the GetMLTaskRun API.
func (c *Client) StartExportLabelsTaskRun(ctx context.Context, params *StartExportLabelsTaskRunInput, optFns ...func(*Options)) (*StartExportLabelsTaskRunOutput, error) {
	if params == nil {
		params = &StartExportLabelsTaskRunInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "StartExportLabelsTaskRun", params, optFns, c.addOperationStartExportLabelsTaskRunMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*StartExportLabelsTaskRunOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type StartExportLabelsTaskRunInput struct {

	// The Amazon S3 path where you export the labels.
	//
	// This member is required.
	OutputS3Path *string

	// The unique identifier of the machine learning transform.
	//
	// This member is required.
	TransformId *string

	noSmithyDocumentSerde
}

type StartExportLabelsTaskRunOutput struct {

	// The unique identifier for the task run.
	TaskRunId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationStartExportLabelsTaskRunMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpStartExportLabelsTaskRun{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpStartExportLabelsTaskRun{}, middleware.After)
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
	if err = addOpStartExportLabelsTaskRunValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opStartExportLabelsTaskRun(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opStartExportLabelsTaskRun(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "glue",
		OperationName: "StartExportLabelsTaskRun",
	}
}
