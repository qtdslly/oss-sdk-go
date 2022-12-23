// Code generated by smithy-go-codegen DO NOT EDIT.

package comprehend

import (
	"context"
	"fmt"
	awsmiddleware "oss-sdk-go/oss/middleware"
	"oss-sdk-go/oss/signer/v4"
	"oss-sdk-go/service/comprehend/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Starts an asynchronous PII entity detection job for a collection of documents.
func (c *Client) StartPiiEntitiesDetectionJob(ctx context.Context, params *StartPiiEntitiesDetectionJobInput, optFns ...func(*Options)) (*StartPiiEntitiesDetectionJobOutput, error) {
	if params == nil {
		params = &StartPiiEntitiesDetectionJobInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "StartPiiEntitiesDetectionJob", params, optFns, c.addOperationStartPiiEntitiesDetectionJobMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*StartPiiEntitiesDetectionJobOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type StartPiiEntitiesDetectionJobInput struct {

	// The Amazon Resource Name (ARN) of the AWS Identity and Access Management (IAM)
	// role that grants Amazon Comprehend read access to your input data.
	//
	// This member is required.
	DataAccessRoleArn *string

	// The input properties for a PII entities detection job.
	//
	// This member is required.
	InputDataConfig *types.InputDataConfig

	// The language of the input documents. Currently, English is the only valid
	// language.
	//
	// This member is required.
	LanguageCode types.LanguageCode

	// Specifies whether the output provides the locations (offsets) of PII entities or
	// a file in which PII entities are redacted.
	//
	// This member is required.
	Mode types.PiiEntitiesDetectionMode

	// Provides conﬁguration parameters for the output of PII entity detection jobs.
	//
	// This member is required.
	OutputDataConfig *types.OutputDataConfig

	// A unique identifier for the request. If you don't set the client request token,
	// Amazon Comprehend generates one.
	ClientRequestToken *string

	// The identifier of the job.
	JobName *string

	// Provides configuration parameters for PII entity redaction. This parameter is
	// required if you set the Mode parameter to ONLY_REDACTION. In that case, you must
	// provide a RedactionConfig definition that includes the PiiEntityTypes parameter.
	RedactionConfig *types.RedactionConfig

	// Tags to be associated with the PII entities detection job. A tag is a key-value
	// pair that adds metadata to a resource used by Amazon Comprehend. For example, a
	// tag with "Sales" as the key might be added to a resource to indicate its use by
	// the sales department.
	Tags []types.Tag

	noSmithyDocumentSerde
}

type StartPiiEntitiesDetectionJobOutput struct {

	// The Amazon Resource Name (ARN) of the PII entity detection job. It is a unique,
	// fully qualified identifier for the job. It includes the AWS account, Region, and
	// the job ID. The format of the ARN is as follows:
	// arn::comprehend:::pii-entities-detection-job/ The following is an example job
	// ARN:
	// arn:aws:comprehend:us-west-2:111122223333:pii-entities-detection-job/1234abcd12ab34cd56ef1234567890ab
	JobArn *string

	// The identifier generated for the job.
	JobId *string

	// The status of the job.
	JobStatus types.JobStatus

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationStartPiiEntitiesDetectionJobMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpStartPiiEntitiesDetectionJob{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpStartPiiEntitiesDetectionJob{}, middleware.After)
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
	if err = addIdempotencyToken_opStartPiiEntitiesDetectionJobMiddleware(stack, options); err != nil {
		return err
	}
	if err = addOpStartPiiEntitiesDetectionJobValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opStartPiiEntitiesDetectionJob(options.Region), middleware.Before); err != nil {
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

type idempotencyToken_initializeOpStartPiiEntitiesDetectionJob struct {
	tokenProvider IdempotencyTokenProvider
}

func (*idempotencyToken_initializeOpStartPiiEntitiesDetectionJob) ID() string {
	return "OperationIdempotencyTokenAutoFill"
}

func (m *idempotencyToken_initializeOpStartPiiEntitiesDetectionJob) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	if m.tokenProvider == nil {
		return next.HandleInitialize(ctx, in)
	}

	input, ok := in.Parameters.(*StartPiiEntitiesDetectionJobInput)
	if !ok {
		return out, metadata, fmt.Errorf("expected middleware input to be of type *StartPiiEntitiesDetectionJobInput ")
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
func addIdempotencyToken_opStartPiiEntitiesDetectionJobMiddleware(stack *middleware.Stack, cfg Options) error {
	return stack.Initialize.Add(&idempotencyToken_initializeOpStartPiiEntitiesDetectionJob{tokenProvider: cfg.IdempotencyTokenProvider}, middleware.Before)
}

func newServiceMetadataMiddleware_opStartPiiEntitiesDetectionJob(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "comprehend",
		OperationName: "StartPiiEntitiesDetectionJob",
	}
}