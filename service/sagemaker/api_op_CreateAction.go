// Code generated by smithy-go-codegen DO NOT EDIT.

package sagemaker

import (
	"context"
	awsmiddleware "oss-sdk-go/oss/middleware"
	"oss-sdk-go/oss/signer/v4"
	"oss-sdk-go/service/sagemaker/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates an action. An action is a lineage tracking entity that represents an
// action or activity. For example, a model deployment or an HPO job. Generally, an
// action involves at least one input or output artifact. For more information, see
// Amazon SageMaker ML Lineage Tracking
// (https://docs.aws.amazon.com/sagemaker/latest/dg/lineage-tracking.html).
func (c *Client) CreateAction(ctx context.Context, params *CreateActionInput, optFns ...func(*Options)) (*CreateActionOutput, error) {
	if params == nil {
		params = &CreateActionInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateAction", params, optFns, c.addOperationCreateActionMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateActionOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateActionInput struct {

	// The name of the action. Must be unique to your account in an Amazon Web Services
	// Region.
	//
	// This member is required.
	ActionName *string

	// The action type.
	//
	// This member is required.
	ActionType *string

	// The source type, ID, and URI.
	//
	// This member is required.
	Source *types.ActionSource

	// The description of the action.
	Description *string

	// Metadata properties of the tracking entity, trial, or trial component.
	MetadataProperties *types.MetadataProperties

	// A list of properties to add to the action.
	Properties map[string]string

	// The status of the action.
	Status types.ActionStatus

	// A list of tags to apply to the action.
	Tags []types.Tag

	noSmithyDocumentSerde
}

type CreateActionOutput struct {

	// The Amazon Resource Name (ARN) of the action.
	ActionArn *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateActionMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpCreateAction{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpCreateAction{}, middleware.After)
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
	if err = addOpCreateActionValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateAction(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateAction(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "sagemaker",
		OperationName: "CreateAction",
	}
}
