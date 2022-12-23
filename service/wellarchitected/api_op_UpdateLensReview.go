// Code generated by smithy-go-codegen DO NOT EDIT.

package wellarchitected

import (
	"context"
	awsmiddleware "oss-sdk-go/oss/middleware"
	"oss-sdk-go/oss/signer/v4"
	"oss-sdk-go/service/wellarchitected/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Update lens review.
func (c *Client) UpdateLensReview(ctx context.Context, params *UpdateLensReviewInput, optFns ...func(*Options)) (*UpdateLensReviewOutput, error) {
	if params == nil {
		params = &UpdateLensReviewInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateLensReview", params, optFns, c.addOperationUpdateLensReviewMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateLensReviewOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Input for update lens review.
type UpdateLensReviewInput struct {

	// The alias of the lens. For Amazon Web Services official lenses, this is either
	// the lens alias, such as serverless, or the lens ARN, such as
	// arn:aws:wellarchitected:us-west-2::lens/serverless. For custom lenses, this is
	// the lens ARN, such as
	// arn:aws:wellarchitected:us-east-1:123456789012:lens/my-lens. Each lens is
	// identified by its LensSummary$LensAlias.
	//
	// This member is required.
	LensAlias *string

	// The ID assigned to the workload. This ID is unique within an Amazon Web Services
	// Region.
	//
	// This member is required.
	WorkloadId *string

	// The notes associated with the workload.
	LensNotes *string

	// List of pillar notes of a lens review in a workload.
	PillarNotes map[string]string

	noSmithyDocumentSerde
}

// Output of a update lens review call.
type UpdateLensReviewOutput struct {

	// A lens review of a question.
	LensReview *types.LensReview

	// The ID assigned to the workload. This ID is unique within an Amazon Web Services
	// Region.
	WorkloadId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationUpdateLensReviewMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpUpdateLensReview{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpUpdateLensReview{}, middleware.After)
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
	if err = addOpUpdateLensReviewValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateLensReview(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opUpdateLensReview(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "wellarchitected",
		OperationName: "UpdateLensReview",
	}
}
