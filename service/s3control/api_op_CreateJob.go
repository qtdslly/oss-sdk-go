// Code generated by smithy-go-codegen DO NOT EDIT.

package s3control

import (
	"context"
	"fmt"
	awsmiddleware "oss-sdk-go/oss/middleware"
	"oss-sdk-go/oss/signer/v4"
	s3controlcust "oss-sdk-go/service/s3control/internal/customizations"
	"oss-sdk-go/service/s3control/types"
	smithy "github.com/aws/smithy-go"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"strings"
)

// You can use S3 Batch Operations to perform large-scale batch actions on Amazon
// S3 objects. Batch Operations can run a single action on lists of Amazon S3
// objects that you specify. For more information, see S3 Batch Operations
// (https://docs.aws.amazon.com/AmazonS3/latest/userguide/batch-ops.html) in the
// Amazon S3 User Guide. This action creates a S3 Batch Operations job. Related
// actions include:
//
// * DescribeJob
// (https://docs.aws.amazon.com/AmazonS3/latest/API/API_control_DescribeJob.html)
//
// *
// ListJobs
// (https://docs.aws.amazon.com/AmazonS3/latest/API/API_control_ListJobs.html)
//
// *
// UpdateJobPriority
// (https://docs.aws.amazon.com/AmazonS3/latest/API/API_control_UpdateJobPriority.html)
//
// *
// UpdateJobStatus
// (https://docs.aws.amazon.com/AmazonS3/latest/API/API_control_UpdateJobStatus.html)
//
// *
// JobOperation
// (https://docs.aws.amazon.com/AmazonS3/latest/API/API_control_JobOperation.html)
func (c *Client) CreateJob(ctx context.Context, params *CreateJobInput, optFns ...func(*Options)) (*CreateJobOutput, error) {
	if params == nil {
		params = &CreateJobInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateJob", params, optFns, c.addOperationCreateJobMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateJobOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateJobInput struct {

	// The Amazon Web Services account ID that creates the job.
	//
	// This member is required.
	AccountId *string

	// An idempotency token to ensure that you don't accidentally submit the same
	// request twice. You can use any string up to the maximum length.
	//
	// This member is required.
	ClientRequestToken *string

	// The action that you want this job to perform on every object listed in the
	// manifest. For more information about the available actions, see Operations
	// (https://docs.aws.amazon.com/AmazonS3/latest/dev/batch-ops-actions.html) in the
	// Amazon S3 User Guide.
	//
	// This member is required.
	Operation *types.JobOperation

	// The numerical priority for this job. Higher numbers indicate higher priority.
	//
	// This member is required.
	Priority *int32

	// Configuration parameters for the optional job-completion report.
	//
	// This member is required.
	Report *types.JobReport

	// The Amazon Resource Name (ARN) for the Identity and Access Management (IAM) role
	// that Batch Operations will use to run this job's action on every object in the
	// manifest.
	//
	// This member is required.
	RoleArn *string

	// Indicates whether confirmation is required before Amazon S3 runs the job.
	// Confirmation is only required for jobs created through the Amazon S3 console.
	ConfirmationRequired *bool

	// A description for this job. You can use any string within the permitted length.
	// Descriptions don't need to be unique and can be used for multiple jobs.
	Description *string

	// Configuration parameters for the manifest.
	Manifest *types.JobManifest

	// The attribute container for the ManifestGenerator details. Jobs must be created
	// with either a manifest file or a ManifestGenerator, but not both.
	ManifestGenerator types.JobManifestGenerator

	// A set of tags to associate with the S3 Batch Operations job. This is an optional
	// parameter.
	Tags []types.S3Tag

	noSmithyDocumentSerde
}

type CreateJobOutput struct {

	// The ID for this job. Amazon S3 generates this ID automatically and returns it
	// after a successful Create Job request.
	JobId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateJobMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestxml_serializeOpCreateJob{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestxml_deserializeOpCreateJob{}, middleware.After)
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
	if err = addEndpointPrefix_opCreateJobMiddleware(stack); err != nil {
		return err
	}
	if err = addIdempotencyToken_opCreateJobMiddleware(stack, options); err != nil {
		return err
	}
	if err = addOpCreateJobValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateJob(options.Region), middleware.Before); err != nil {
		return err
	}
	if err = addMetadataRetrieverMiddleware(stack); err != nil {
		return err
	}
	if err = addCreateJobUpdateEndpoint(stack, options); err != nil {
		return err
	}
	if err = addResponseErrorMiddleware(stack); err != nil {
		return err
	}
	if err = v4.AddContentSHA256HeaderMiddleware(stack); err != nil {
		return err
	}
	if err = addRequestResponseLogging(stack, options); err != nil {
		return err
	}
	return nil
}

type endpointPrefix_opCreateJobMiddleware struct {
}

func (*endpointPrefix_opCreateJobMiddleware) ID() string {
	return "EndpointHostPrefix"
}

func (m *endpointPrefix_opCreateJobMiddleware) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
) {
	if smithyhttp.GetHostnameImmutable(ctx) || smithyhttp.IsEndpointHostPrefixDisabled(ctx) {
		return next.HandleSerialize(ctx, in)
	}

	req, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, fmt.Errorf("unknown transport type %T", in.Request)
	}

	input, ok := in.Parameters.(*CreateJobInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input type %T", in.Parameters)
	}

	var prefix strings.Builder
	if input.AccountId == nil {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("AccountId forms part of the endpoint host and so may not be nil")}
	} else if !smithyhttp.ValidHostLabel(*input.AccountId) {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("AccountId forms part of the endpoint host and so must match \"[a-zA-Z0-9-]{1,63}\", but was \"%s\"", *input.AccountId)}
	} else {
		prefix.WriteString(*input.AccountId)
	}
	prefix.WriteString(".")
	req.URL.Host = prefix.String() + req.URL.Host

	return next.HandleSerialize(ctx, in)
}
func addEndpointPrefix_opCreateJobMiddleware(stack *middleware.Stack) error {
	return stack.Serialize.Insert(&endpointPrefix_opCreateJobMiddleware{}, `OperationSerializer`, middleware.After)
}

type idempotencyToken_initializeOpCreateJob struct {
	tokenProvider IdempotencyTokenProvider
}

func (*idempotencyToken_initializeOpCreateJob) ID() string {
	return "OperationIdempotencyTokenAutoFill"
}

func (m *idempotencyToken_initializeOpCreateJob) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	if m.tokenProvider == nil {
		return next.HandleInitialize(ctx, in)
	}

	input, ok := in.Parameters.(*CreateJobInput)
	if !ok {
		return out, metadata, fmt.Errorf("expected middleware input to be of type *CreateJobInput ")
	}

	if input.ClientRequestToken == nil {
		t, err := m.tokenProvider.GetIdempotencyToken()
		if err != nil {
			return out, metadata, err
		}
		input.ClientRequestToken = &t
	}
	return next.HandleInitialize(ctx, in)
}
func addIdempotencyToken_opCreateJobMiddleware(stack *middleware.Stack, cfg Options) error {
	return stack.Initialize.Add(&idempotencyToken_initializeOpCreateJob{tokenProvider: cfg.IdempotencyTokenProvider}, middleware.Before)
}

func newServiceMetadataMiddleware_opCreateJob(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "s3",
		OperationName: "CreateJob",
	}
}

func copyCreateJobInputForUpdateEndpoint(params interface{}) (interface{}, error) {
	input, ok := params.(*CreateJobInput)
	if !ok {
		return nil, fmt.Errorf("expect *CreateJobInput type, got %T", params)
	}
	cpy := *input
	return &cpy, nil
}
func backFillCreateJobAccountID(input interface{}, v string) error {
	in := input.(*CreateJobInput)
	if in.AccountId != nil {
		if !strings.EqualFold(*in.AccountId, v) {
			return fmt.Errorf("error backfilling account id")
		}
		return nil
	}
	in.AccountId = &v
	return nil
}
func addCreateJobUpdateEndpoint(stack *middleware.Stack, options Options) error {
	return s3controlcust.UpdateEndpoint(stack, s3controlcust.UpdateEndpointOptions{
		Accessor: s3controlcust.UpdateEndpointParameterAccessor{GetARNInput: nopGetARNAccessor,
			BackfillAccountID: nopBackfillAccountIDAccessor,
			GetOutpostIDInput: nopGetOutpostIDFromInput,
			UpdateARNField:    nopSetARNAccessor,
			CopyInput:         copyCreateJobInputForUpdateEndpoint,
		},
		EndpointResolver:        options.EndpointResolver,
		EndpointResolverOptions: options.EndpointOptions,
		UseARNRegion:            options.UseARNRegion,
	})
}