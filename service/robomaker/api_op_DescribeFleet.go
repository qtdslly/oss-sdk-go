// Code generated by smithy-go-codegen DO NOT EDIT.

package robomaker

import (
	"context"
	awsmiddleware "oss-sdk-go/oss/middleware"
	"oss-sdk-go/oss/signer/v4"
	"oss-sdk-go/service/robomaker/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Describes a fleet. This API will no longer be supported as of May 2, 2022. Use
// it to remove resources that were created for Deployment Service.
//
// Deprecated: Support for the AWS RoboMaker application deployment feature has
// ended. For additional information, see
// https://docs.aws.amazon.com/robomaker/latest/dg/fleets.html.
func (c *Client) DescribeFleet(ctx context.Context, params *DescribeFleetInput, optFns ...func(*Options)) (*DescribeFleetOutput, error) {
	if params == nil {
		params = &DescribeFleetInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeFleet", params, optFns, c.addOperationDescribeFleetMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeFleetOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeFleetInput struct {

	// The Amazon Resource Name (ARN) of the fleet.
	//
	// This member is required.
	Fleet *string

	noSmithyDocumentSerde
}

type DescribeFleetOutput struct {

	// The Amazon Resource Name (ARN) of the fleet.
	Arn *string

	// The time, in milliseconds since the epoch, when the fleet was created.
	CreatedAt *time.Time

	// The Amazon Resource Name (ARN) of the last deployment job.
	LastDeploymentJob *string

	// The status of the last deployment.
	LastDeploymentStatus types.DeploymentStatus

	// The time of the last deployment.
	LastDeploymentTime *time.Time

	// The name of the fleet.
	Name *string

	// A list of robots.
	Robots []types.Robot

	// The list of all tags added to the specified fleet.
	Tags map[string]string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeFleetMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpDescribeFleet{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpDescribeFleet{}, middleware.After)
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
	if err = addOpDescribeFleetValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeFleet(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDescribeFleet(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "robomaker",
		OperationName: "DescribeFleet",
	}
}
