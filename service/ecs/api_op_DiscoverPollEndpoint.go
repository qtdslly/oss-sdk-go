// Code generated by smithy-go-codegen DO NOT EDIT.

package ecs

import (
	"context"
	awsmiddleware "oss-sdk-go/oss/middleware"
	"oss-sdk-go/oss/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// This action is only used by the Amazon ECS agent, and it is not intended for use
// outside of the agent. Returns an endpoint for the Amazon ECS agent to poll for
// updates.
func (c *Client) DiscoverPollEndpoint(ctx context.Context, params *DiscoverPollEndpointInput, optFns ...func(*Options)) (*DiscoverPollEndpointOutput, error) {
	if params == nil {
		params = &DiscoverPollEndpointInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DiscoverPollEndpoint", params, optFns, c.addOperationDiscoverPollEndpointMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DiscoverPollEndpointOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DiscoverPollEndpointInput struct {

	// The short name or full Amazon Resource Name (ARN) of the cluster that the
	// container instance belongs to.
	Cluster *string

	// The container instance ID or full ARN of the container instance. For more
	// information about the ARN format, see Amazon Resource Name (ARN)
	// (https://docs.aws.amazon.com/AmazonECS/latest/developerguide/ecs-account-settings.html#ecs-resource-ids)
	// in the Amazon ECS Developer Guide.
	ContainerInstance *string

	noSmithyDocumentSerde
}

type DiscoverPollEndpointOutput struct {

	// The endpoint for the Amazon ECS agent to poll.
	Endpoint *string

	// The telemetry endpoint for the Amazon ECS agent.
	TelemetryEndpoint *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDiscoverPollEndpointMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDiscoverPollEndpoint{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDiscoverPollEndpoint{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDiscoverPollEndpoint(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDiscoverPollEndpoint(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ecs",
		OperationName: "DiscoverPollEndpoint",
	}
}
