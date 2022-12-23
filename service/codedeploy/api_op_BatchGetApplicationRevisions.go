// Code generated by smithy-go-codegen DO NOT EDIT.

package codedeploy

import (
	"context"
	awsmiddleware "oss-sdk-go/oss/middleware"
	"oss-sdk-go/oss/signer/v4"
	"oss-sdk-go/service/codedeploy/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Gets information about one or more application revisions. The maximum number of
// application revisions that can be returned is 25.
func (c *Client) BatchGetApplicationRevisions(ctx context.Context, params *BatchGetApplicationRevisionsInput, optFns ...func(*Options)) (*BatchGetApplicationRevisionsOutput, error) {
	if params == nil {
		params = &BatchGetApplicationRevisionsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "BatchGetApplicationRevisions", params, optFns, c.addOperationBatchGetApplicationRevisionsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*BatchGetApplicationRevisionsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Represents the input of a BatchGetApplicationRevisions operation.
type BatchGetApplicationRevisionsInput struct {

	// The name of an AWS CodeDeploy application about which to get revision
	// information.
	//
	// This member is required.
	ApplicationName *string

	// An array of RevisionLocation objects that specify information to get about the
	// application revisions, including type and location. The maximum number of
	// RevisionLocation objects you can specify is 25.
	//
	// This member is required.
	Revisions []types.RevisionLocation

	noSmithyDocumentSerde
}

// Represents the output of a BatchGetApplicationRevisions operation.
type BatchGetApplicationRevisionsOutput struct {

	// The name of the application that corresponds to the revisions.
	ApplicationName *string

	// Information about errors that might have occurred during the API call.
	ErrorMessage *string

	// Additional information about the revisions, including the type and location.
	Revisions []types.RevisionInfo

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationBatchGetApplicationRevisionsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpBatchGetApplicationRevisions{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpBatchGetApplicationRevisions{}, middleware.After)
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
	if err = addOpBatchGetApplicationRevisionsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opBatchGetApplicationRevisions(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opBatchGetApplicationRevisions(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "codedeploy",
		OperationName: "BatchGetApplicationRevisions",
	}
}
