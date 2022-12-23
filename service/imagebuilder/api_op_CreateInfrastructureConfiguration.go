// Code generated by smithy-go-codegen DO NOT EDIT.

package imagebuilder

import (
	"context"
	"fmt"
	awsmiddleware "oss-sdk-go/oss/middleware"
	"oss-sdk-go/oss/signer/v4"
	"oss-sdk-go/service/imagebuilder/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates a new infrastructure configuration. An infrastructure configuration
// defines the environment in which your image will be built and tested.
func (c *Client) CreateInfrastructureConfiguration(ctx context.Context, params *CreateInfrastructureConfigurationInput, optFns ...func(*Options)) (*CreateInfrastructureConfigurationOutput, error) {
	if params == nil {
		params = &CreateInfrastructureConfigurationInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateInfrastructureConfiguration", params, optFns, c.addOperationCreateInfrastructureConfigurationMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateInfrastructureConfigurationOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateInfrastructureConfigurationInput struct {

	// The idempotency token used to make this request idempotent.
	//
	// This member is required.
	ClientToken *string

	// The instance profile to associate with the instance used to customize your
	// Amazon EC2 AMI.
	//
	// This member is required.
	InstanceProfileName *string

	// The name of the infrastructure configuration.
	//
	// This member is required.
	Name *string

	// The description of the infrastructure configuration.
	Description *string

	// The instance metadata options that you can set for the HTTP requests that
	// pipeline builds use to launch EC2 build and test instances.
	InstanceMetadataOptions *types.InstanceMetadataOptions

	// The instance types of the infrastructure configuration. You can specify one or
	// more instance types to use for this build. The service will pick one of these
	// instance types based on availability.
	InstanceTypes []string

	// The key pair of the infrastructure configuration. You can use this to log on to
	// and debug the instance used to create your image.
	KeyPair *string

	// The logging configuration of the infrastructure configuration.
	Logging *types.Logging

	// The tags attached to the resource created by Image Builder.
	ResourceTags map[string]string

	// The security group IDs to associate with the instance used to customize your
	// Amazon EC2 AMI.
	SecurityGroupIds []string

	// The Amazon Resource Name (ARN) for the SNS topic to which we send image build
	// event notifications. EC2 Image Builder is unable to send notifications to SNS
	// topics that are encrypted using keys from other accounts. The key that is used
	// to encrypt the SNS topic must reside in the account that the Image Builder
	// service runs under.
	SnsTopicArn *string

	// The subnet ID in which to place the instance used to customize your Amazon EC2
	// AMI.
	SubnetId *string

	// The tags of the infrastructure configuration.
	Tags map[string]string

	// The terminate instance on failure setting of the infrastructure configuration.
	// Set to false if you want Image Builder to retain the instance used to configure
	// your AMI if the build or test phase of your workflow fails.
	TerminateInstanceOnFailure *bool

	noSmithyDocumentSerde
}

type CreateInfrastructureConfigurationOutput struct {

	// The idempotency token used to make this request idempotent.
	ClientToken *string

	// The Amazon Resource Name (ARN) of the infrastructure configuration that was
	// created by this request.
	InfrastructureConfigurationArn *string

	// The request ID that uniquely identifies this request.
	RequestId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateInfrastructureConfigurationMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpCreateInfrastructureConfiguration{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpCreateInfrastructureConfiguration{}, middleware.After)
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
	if err = addIdempotencyToken_opCreateInfrastructureConfigurationMiddleware(stack, options); err != nil {
		return err
	}
	if err = addOpCreateInfrastructureConfigurationValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateInfrastructureConfiguration(options.Region), middleware.Before); err != nil {
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

type idempotencyToken_initializeOpCreateInfrastructureConfiguration struct {
	tokenProvider IdempotencyTokenProvider
}

func (*idempotencyToken_initializeOpCreateInfrastructureConfiguration) ID() string {
	return "OperationIdempotencyTokenAutoFill"
}

func (m *idempotencyToken_initializeOpCreateInfrastructureConfiguration) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	if m.tokenProvider == nil {
		return next.HandleInitialize(ctx, in)
	}

	input, ok := in.Parameters.(*CreateInfrastructureConfigurationInput)
	if !ok {
		return out, metadata, fmt.Errorf("expected middleware input to be of type *CreateInfrastructureConfigurationInput ")
	}

	if input.ClientToken == nil {
		t, err := m.tokenProvider.GetIdempotencyToken()
		if err != nil {
			return out, metadata, err
		}
		input.ClientToken = &t
	}
	return next.HandleInitialize(ctx, in)
}
func addIdempotencyToken_opCreateInfrastructureConfigurationMiddleware(stack *middleware.Stack, cfg Options) error {
	return stack.Initialize.Add(&idempotencyToken_initializeOpCreateInfrastructureConfiguration{tokenProvider: cfg.IdempotencyTokenProvider}, middleware.Before)
}

func newServiceMetadataMiddleware_opCreateInfrastructureConfiguration(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "imagebuilder",
		OperationName: "CreateInfrastructureConfiguration",
	}
}
