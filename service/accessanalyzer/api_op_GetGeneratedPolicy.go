// Code generated by smithy-go-codegen DO NOT EDIT.

package accessanalyzer

import (
	"context"
	awsmiddleware "oss-sdk-go/oss/middleware"
	"oss-sdk-go/oss/signer/v4"
	"oss-sdk-go/service/accessanalyzer/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Retrieves the policy that was generated using StartPolicyGeneration.
func (c *Client) GetGeneratedPolicy(ctx context.Context, params *GetGeneratedPolicyInput, optFns ...func(*Options)) (*GetGeneratedPolicyOutput, error) {
	if params == nil {
		params = &GetGeneratedPolicyInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetGeneratedPolicy", params, optFns, c.addOperationGetGeneratedPolicyMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetGeneratedPolicyOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetGeneratedPolicyInput struct {

	// The JobId that is returned by the StartPolicyGeneration operation. The JobId can
	// be used with GetGeneratedPolicy to retrieve the generated policies or used with
	// CancelPolicyGeneration to cancel the policy generation request.
	//
	// This member is required.
	JobId *string

	// The level of detail that you want to generate. You can specify whether to
	// generate policies with placeholders for resource ARNs for actions that support
	// resource level granularity in policies. For example, in the resource section of
	// a policy, you can receive a placeholder such as
	// "Resource":"arn:aws:s3:::${BucketName}" instead of "*".
	IncludeResourcePlaceholders *bool

	// The level of detail that you want to generate. You can specify whether to
	// generate service-level policies. IAM Access Analyzer uses
	// iam:servicelastaccessed to identify services that have been used recently to
	// create this service-level template.
	IncludeServiceLevelTemplate *bool

	noSmithyDocumentSerde
}

type GetGeneratedPolicyOutput struct {

	// A GeneratedPolicyResult object that contains the generated policies and
	// associated details.
	//
	// This member is required.
	GeneratedPolicyResult *types.GeneratedPolicyResult

	// A GeneratedPolicyDetails object that contains details about the generated
	// policy.
	//
	// This member is required.
	JobDetails *types.JobDetails

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetGeneratedPolicyMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpGetGeneratedPolicy{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpGetGeneratedPolicy{}, middleware.After)
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
	if err = addOpGetGeneratedPolicyValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetGeneratedPolicy(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opGetGeneratedPolicy(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "access-analyzer",
		OperationName: "GetGeneratedPolicy",
	}
}
