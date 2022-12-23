// Code generated by smithy-go-codegen DO NOT EDIT.

package kafka

import (
	"context"
	awsmiddleware "oss-sdk-go/oss/middleware"
	"oss-sdk-go/oss/signer/v4"
	"oss-sdk-go/service/kafka/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Creates a new MSK configuration.
func (c *Client) CreateConfiguration(ctx context.Context, params *CreateConfigurationInput, optFns ...func(*Options)) (*CreateConfigurationOutput, error) {
	if params == nil {
		params = &CreateConfigurationInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateConfiguration", params, optFns, c.addOperationCreateConfigurationMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateConfigurationOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateConfigurationInput struct {

	// The name of the configuration.
	//
	// This member is required.
	Name *string

	// Contents of the server.properties file. When using the API, you must ensure that
	// the contents of the file are base64 encoded. When using the AWS Management
	// Console, the SDK, or the AWS CLI, the contents of server.properties can be in
	// plaintext.
	//
	// This member is required.
	ServerProperties []byte

	// The description of the configuration.
	Description *string

	// The versions of Apache Kafka with which you can use this MSK configuration.
	KafkaVersions []string

	noSmithyDocumentSerde
}

type CreateConfigurationOutput struct {

	// The Amazon Resource Name (ARN) of the configuration.
	Arn *string

	// The time when the configuration was created.
	CreationTime *time.Time

	// Latest revision of the configuration.
	LatestRevision *types.ConfigurationRevision

	// The name of the configuration.
	Name *string

	// The state of the configuration. The possible states are ACTIVE, DELETING, and
	// DELETE_FAILED.
	State types.ConfigurationState

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateConfigurationMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpCreateConfiguration{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpCreateConfiguration{}, middleware.After)
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
	if err = addOpCreateConfigurationValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateConfiguration(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateConfiguration(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "kafka",
		OperationName: "CreateConfiguration",
	}
}
