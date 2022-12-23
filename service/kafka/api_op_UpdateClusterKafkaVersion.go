// Code generated by smithy-go-codegen DO NOT EDIT.

package kafka

import (
	"context"
	awsmiddleware "oss-sdk-go/oss/middleware"
	"oss-sdk-go/oss/signer/v4"
	"oss-sdk-go/service/kafka/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Updates the Apache Kafka version for the cluster.
func (c *Client) UpdateClusterKafkaVersion(ctx context.Context, params *UpdateClusterKafkaVersionInput, optFns ...func(*Options)) (*UpdateClusterKafkaVersionOutput, error) {
	if params == nil {
		params = &UpdateClusterKafkaVersionInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateClusterKafkaVersion", params, optFns, c.addOperationUpdateClusterKafkaVersionMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateClusterKafkaVersionOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateClusterKafkaVersionInput struct {

	// The Amazon Resource Name (ARN) of the cluster to be updated.
	//
	// This member is required.
	ClusterArn *string

	// Current cluster version.
	//
	// This member is required.
	CurrentVersion *string

	// Target Kafka version.
	//
	// This member is required.
	TargetKafkaVersion *string

	// The custom configuration that should be applied on the new version of cluster.
	ConfigurationInfo *types.ConfigurationInfo

	noSmithyDocumentSerde
}

type UpdateClusterKafkaVersionOutput struct {

	// The Amazon Resource Name (ARN) of the cluster.
	ClusterArn *string

	// The Amazon Resource Name (ARN) of the cluster operation.
	ClusterOperationArn *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationUpdateClusterKafkaVersionMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpUpdateClusterKafkaVersion{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpUpdateClusterKafkaVersion{}, middleware.After)
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
	if err = addOpUpdateClusterKafkaVersionValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateClusterKafkaVersion(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opUpdateClusterKafkaVersion(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "kafka",
		OperationName: "UpdateClusterKafkaVersion",
	}
}
