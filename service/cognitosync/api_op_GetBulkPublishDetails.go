// Code generated by smithy-go-codegen DO NOT EDIT.

package cognitosync

import (
	"context"
	awsmiddleware "oss-sdk-go/oss/middleware"
	"oss-sdk-go/oss/signer/v4"
	"oss-sdk-go/service/cognitosync/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Get the status of the last BulkPublish operation for an identity pool.This API
// can only be called with developer credentials. You cannot call this API with the
// temporary user credentials provided by Cognito Identity.
func (c *Client) GetBulkPublishDetails(ctx context.Context, params *GetBulkPublishDetailsInput, optFns ...func(*Options)) (*GetBulkPublishDetailsOutput, error) {
	if params == nil {
		params = &GetBulkPublishDetailsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetBulkPublishDetails", params, optFns, c.addOperationGetBulkPublishDetailsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetBulkPublishDetailsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// The input for the GetBulkPublishDetails operation.
type GetBulkPublishDetailsInput struct {

	// A name-spaced GUID (for example, us-east-1:23EC4050-6AEA-7089-A2DD-08002EXAMPLE)
	// created by Amazon Cognito. GUID generation is unique within a region.
	//
	// This member is required.
	IdentityPoolId *string

	noSmithyDocumentSerde
}

// The output for the GetBulkPublishDetails operation.
type GetBulkPublishDetailsOutput struct {

	// If BulkPublishStatus is SUCCEEDED, the time the last bulk publish operation
	// completed.
	BulkPublishCompleteTime *time.Time

	// The date/time at which the last bulk publish was initiated.
	BulkPublishStartTime *time.Time

	// Status of the last bulk publish operation, valid values are: NOT_STARTED - No
	// bulk publish has been requested for this identity pool IN_PROGRESS - Data is
	// being published to the configured stream SUCCEEDED - All data for the identity
	// pool has been published to the configured stream FAILED - Some portion of the
	// data has failed to publish, check FailureMessage for the cause.
	BulkPublishStatus types.BulkPublishStatus

	// If BulkPublishStatus is FAILED this field will contain the error message that
	// caused the bulk publish to fail.
	FailureMessage *string

	// A name-spaced GUID (for example, us-east-1:23EC4050-6AEA-7089-A2DD-08002EXAMPLE)
	// created by Amazon Cognito. GUID generation is unique within a region.
	IdentityPoolId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetBulkPublishDetailsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpGetBulkPublishDetails{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpGetBulkPublishDetails{}, middleware.After)
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
	if err = addOpGetBulkPublishDetailsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetBulkPublishDetails(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opGetBulkPublishDetails(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "cognito-sync",
		OperationName: "GetBulkPublishDetails",
	}
}
