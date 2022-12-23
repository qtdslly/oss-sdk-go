// Code generated by smithy-go-codegen DO NOT EDIT.

package braket

import (
	"context"
	awsmiddleware "oss-sdk-go/oss/middleware"
	"oss-sdk-go/oss/signer/v4"
	"oss-sdk-go/service/braket/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Retrieves the specified Amazon Braket job.
func (c *Client) GetJob(ctx context.Context, params *GetJobInput, optFns ...func(*Options)) (*GetJobOutput, error) {
	if params == nil {
		params = &GetJobInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetJob", params, optFns, c.addOperationGetJobMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetJobOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetJobInput struct {

	// The ARN of the job to retrieve.
	//
	// This member is required.
	JobArn *string

	noSmithyDocumentSerde
}

type GetJobOutput struct {

	// Definition of the Amazon Braket job created. Specifies the container image the
	// job uses, information about the Python scripts used for entry and training, and
	// the user-defined metrics used to evaluation the job.
	//
	// This member is required.
	AlgorithmSpecification *types.AlgorithmSpecification

	// The date and time that the Amazon Braket job was created.
	//
	// This member is required.
	CreatedAt *time.Time

	// The resource instances to use while running the hybrid job on Amazon Braket.
	//
	// This member is required.
	InstanceConfig *types.InstanceConfig

	// The ARN of the Amazon Braket job.
	//
	// This member is required.
	JobArn *string

	// The name of the Amazon Braket job.
	//
	// This member is required.
	JobName *string

	// The path to the S3 location where job artifacts are stored and the encryption
	// key used to store them there.
	//
	// This member is required.
	OutputDataConfig *types.JobOutputDataConfig

	// The Amazon Resource Name (ARN) of an IAM role that Amazon Braket can assume to
	// perform tasks on behalf of a user. It can access user resources, run an Amazon
	// Braket job container on behalf of user, and output resources to the s3 buckets
	// of a user.
	//
	// This member is required.
	RoleArn *string

	// The status of the Amazon Braket job.
	//
	// This member is required.
	Status types.JobPrimaryStatus

	// The billable time the Amazon Braket job used to complete.
	BillableDuration *int32

	// Information about the output locations for job checkpoint data.
	CheckpointConfig *types.JobCheckpointConfig

	// The quantum processing unit (QPU) or simulator used to run the Amazon Braket
	// job.
	DeviceConfig *types.DeviceConfig

	// The date and time that the Amazon Braket job ended.
	EndedAt *time.Time

	// Details about the type and time events occurred related to the Amazon Braket
	// job.
	Events []types.JobEventDetails

	// A description of the reason why an Amazon Braket job failed, if it failed.
	FailureReason *string

	// Algorithm-specific parameters used by an Amazon Braket job that influence the
	// quality of the traiing job. The values are set with a string of JSON key:value
	// pairs, where the key is the name of the hyperparameter and the value is the
	// value of th hyperparameter.
	HyperParameters map[string]string

	// A list of parameters that specify the name and type of input data and where it
	// is located.
	InputDataConfig []types.InputFileConfig

	// The date and time that the Amazon Braket job was started.
	StartedAt *time.Time

	// The user-defined criteria that specifies when to stop a job running.
	StoppingCondition *types.JobStoppingCondition

	// A tag object that consists of a key and an optional value, used to manage
	// metadata for Amazon Braket resources.
	Tags map[string]string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetJobMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpGetJob{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpGetJob{}, middleware.After)
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
	if err = addOpGetJobValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetJob(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opGetJob(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "braket",
		OperationName: "GetJob",
	}
}