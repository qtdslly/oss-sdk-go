// Code generated by smithy-go-codegen DO NOT EDIT.

package autoscaling

import (
	"context"
	awsmiddleware "oss-sdk-go/oss/middleware"
	"oss-sdk-go/oss/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Attaches one or more EC2 instances to the specified Auto Scaling group. When you
// attach instances, Amazon EC2 Auto Scaling increases the desired capacity of the
// group by the number of instances being attached. If the number of instances
// being attached plus the desired capacity of the group exceeds the maximum size
// of the group, the operation fails. If there is a Classic Load Balancer attached
// to your Auto Scaling group, the instances are also registered with the load
// balancer. If there are target groups attached to your Auto Scaling group, the
// instances are also registered with the target groups. For more information, see
// Attach EC2 instances to your Auto Scaling group
// (https://docs.aws.amazon.com/autoscaling/ec2/userguide/attach-instance-asg.html)
// in the Amazon EC2 Auto Scaling User Guide.
func (c *Client) AttachInstances(ctx context.Context, params *AttachInstancesInput, optFns ...func(*Options)) (*AttachInstancesOutput, error) {
	if params == nil {
		params = &AttachInstancesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "AttachInstances", params, optFns, c.addOperationAttachInstancesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*AttachInstancesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type AttachInstancesInput struct {

	// The name of the Auto Scaling group.
	//
	// This member is required.
	AutoScalingGroupName *string

	// The IDs of the instances. You can specify up to 20 instances.
	InstanceIds []string

	noSmithyDocumentSerde
}

type AttachInstancesOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationAttachInstancesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsquery_serializeOpAttachInstances{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpAttachInstances{}, middleware.After)
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
	if err = addOpAttachInstancesValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opAttachInstances(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opAttachInstances(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "autoscaling",
		OperationName: "AttachInstances",
	}
}
