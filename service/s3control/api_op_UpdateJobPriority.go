// Code generated by smithy-go-codegen DO NOT EDIT.

package s3control

import (
	"context"
	"fmt"
	awsmiddleware "oss-sdk-go/oss/middleware"
	"oss-sdk-go/oss/signer/v4"
	s3controlcust "oss-sdk-go/service/s3control/internal/customizations"
	smithy "github.com/aws/smithy-go"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"strings"
)

// Updates an existing S3 Batch Operations job's priority. For more information,
// see S3 Batch Operations
// (https://docs.aws.amazon.com/AmazonS3/latest/userguide/batch-ops.html) in the
// Amazon S3 User Guide. Related actions include:
//
// * CreateJob
// (https://docs.aws.amazon.com/AmazonS3/latest/API/API_control_CreateJob.html)
//
// *
// ListJobs
// (https://docs.aws.amazon.com/AmazonS3/latest/API/API_control_ListJobs.html)
//
// *
// DescribeJob
// (https://docs.aws.amazon.com/AmazonS3/latest/API/API_control_DescribeJob.html)
//
// *
// UpdateJobStatus
// (https://docs.aws.amazon.com/AmazonS3/latest/API/API_control_UpdateJobStatus.html)
func (c *Client) UpdateJobPriority(ctx context.Context, params *UpdateJobPriorityInput, optFns ...func(*Options)) (*UpdateJobPriorityOutput, error) {
	if params == nil {
		params = &UpdateJobPriorityInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateJobPriority", params, optFns, c.addOperationUpdateJobPriorityMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateJobPriorityOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateJobPriorityInput struct {

	// The Amazon Web Services account ID associated with the S3 Batch Operations job.
	//
	// This member is required.
	AccountId *string

	// The ID for the job whose priority you want to update.
	//
	// This member is required.
	JobId *string

	// The priority you want to assign to this job.
	//
	// This member is required.
	Priority int32

	noSmithyDocumentSerde
}

type UpdateJobPriorityOutput struct {

	// The ID for the job whose priority Amazon S3 updated.
	//
	// This member is required.
	JobId *string

	// The new priority assigned to the specified job.
	//
	// This member is required.
	Priority int32

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationUpdateJobPriorityMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestxml_serializeOpUpdateJobPriority{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestxml_deserializeOpUpdateJobPriority{}, middleware.After)
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
	if err = addEndpointPrefix_opUpdateJobPriorityMiddleware(stack); err != nil {
		return err
	}
	if err = addOpUpdateJobPriorityValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateJobPriority(options.Region), middleware.Before); err != nil {
		return err
	}
	if err = addMetadataRetrieverMiddleware(stack); err != nil {
		return err
	}
	if err = addUpdateJobPriorityUpdateEndpoint(stack, options); err != nil {
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

type endpointPrefix_opUpdateJobPriorityMiddleware struct {
}

func (*endpointPrefix_opUpdateJobPriorityMiddleware) ID() string {
	return "EndpointHostPrefix"
}

func (m *endpointPrefix_opUpdateJobPriorityMiddleware) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
) {
	if smithyhttp.GetHostnameImmutable(ctx) || smithyhttp.IsEndpointHostPrefixDisabled(ctx) {
		return next.HandleSerialize(ctx, in)
	}

	req, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, fmt.Errorf("unknown transport type %T", in.Request)
	}

	input, ok := in.Parameters.(*UpdateJobPriorityInput)
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
func addEndpointPrefix_opUpdateJobPriorityMiddleware(stack *middleware.Stack) error {
	return stack.Serialize.Insert(&endpointPrefix_opUpdateJobPriorityMiddleware{}, `OperationSerializer`, middleware.After)
}

func newServiceMetadataMiddleware_opUpdateJobPriority(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "s3",
		OperationName: "UpdateJobPriority",
	}
}

func copyUpdateJobPriorityInputForUpdateEndpoint(params interface{}) (interface{}, error) {
	input, ok := params.(*UpdateJobPriorityInput)
	if !ok {
		return nil, fmt.Errorf("expect *UpdateJobPriorityInput type, got %T", params)
	}
	cpy := *input
	return &cpy, nil
}
func backFillUpdateJobPriorityAccountID(input interface{}, v string) error {
	in := input.(*UpdateJobPriorityInput)
	if in.AccountId != nil {
		if !strings.EqualFold(*in.AccountId, v) {
			return fmt.Errorf("error backfilling account id")
		}
		return nil
	}
	in.AccountId = &v
	return nil
}
func addUpdateJobPriorityUpdateEndpoint(stack *middleware.Stack, options Options) error {
	return s3controlcust.UpdateEndpoint(stack, s3controlcust.UpdateEndpointOptions{
		Accessor: s3controlcust.UpdateEndpointParameterAccessor{GetARNInput: nopGetARNAccessor,
			BackfillAccountID: nopBackfillAccountIDAccessor,
			GetOutpostIDInput: nopGetOutpostIDFromInput,
			UpdateARNField:    nopSetARNAccessor,
			CopyInput:         copyUpdateJobPriorityInputForUpdateEndpoint,
		},
		EndpointResolver:        options.EndpointResolver,
		EndpointResolverOptions: options.EndpointOptions,
		UseARNRegion:            options.UseARNRegion,
	})
}