// Code generated by smithy-go-codegen DO NOT EDIT.

package auditmanager

import (
	"context"
	awsmiddleware "oss-sdk-go/oss/middleware"
	"oss-sdk-go/oss/signer/v4"
	"oss-sdk-go/service/auditmanager/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates a batch of delegations for an assessment in Audit Manager.
func (c *Client) BatchCreateDelegationByAssessment(ctx context.Context, params *BatchCreateDelegationByAssessmentInput, optFns ...func(*Options)) (*BatchCreateDelegationByAssessmentOutput, error) {
	if params == nil {
		params = &BatchCreateDelegationByAssessmentInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "BatchCreateDelegationByAssessment", params, optFns, c.addOperationBatchCreateDelegationByAssessmentMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*BatchCreateDelegationByAssessmentOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type BatchCreateDelegationByAssessmentInput struct {

	// The identifier for the assessment.
	//
	// This member is required.
	AssessmentId *string

	// The API request to batch create delegations in Audit Manager.
	//
	// This member is required.
	CreateDelegationRequests []types.CreateDelegationRequest

	noSmithyDocumentSerde
}

type BatchCreateDelegationByAssessmentOutput struct {

	// The delegations that are associated with the assessment.
	Delegations []types.Delegation

	// A list of errors that the BatchCreateDelegationByAssessment API returned.
	Errors []types.BatchCreateDelegationByAssessmentError

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationBatchCreateDelegationByAssessmentMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpBatchCreateDelegationByAssessment{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpBatchCreateDelegationByAssessment{}, middleware.After)
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
	if err = addOpBatchCreateDelegationByAssessmentValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opBatchCreateDelegationByAssessment(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opBatchCreateDelegationByAssessment(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "auditmanager",
		OperationName: "BatchCreateDelegationByAssessment",
	}
}