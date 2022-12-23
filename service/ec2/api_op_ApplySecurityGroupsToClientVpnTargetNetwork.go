// Code generated by smithy-go-codegen DO NOT EDIT.

package ec2

import (
	"context"
	awsmiddleware "oss-sdk-go/oss/middleware"
	"oss-sdk-go/oss/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Applies a security group to the association between the target network and the
// Client VPN endpoint. This action replaces the existing security groups with the
// specified security groups.
func (c *Client) ApplySecurityGroupsToClientVpnTargetNetwork(ctx context.Context, params *ApplySecurityGroupsToClientVpnTargetNetworkInput, optFns ...func(*Options)) (*ApplySecurityGroupsToClientVpnTargetNetworkOutput, error) {
	if params == nil {
		params = &ApplySecurityGroupsToClientVpnTargetNetworkInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ApplySecurityGroupsToClientVpnTargetNetwork", params, optFns, c.addOperationApplySecurityGroupsToClientVpnTargetNetworkMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ApplySecurityGroupsToClientVpnTargetNetworkOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ApplySecurityGroupsToClientVpnTargetNetworkInput struct {

	// The ID of the Client VPN endpoint.
	//
	// This member is required.
	ClientVpnEndpointId *string

	// The IDs of the security groups to apply to the associated target network. Up to
	// 5 security groups can be applied to an associated target network.
	//
	// This member is required.
	SecurityGroupIds []string

	// The ID of the VPC in which the associated target network is located.
	//
	// This member is required.
	VpcId *string

	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have the
	// required permissions, the error response is DryRunOperation. Otherwise, it is
	// UnauthorizedOperation.
	DryRun *bool

	noSmithyDocumentSerde
}

type ApplySecurityGroupsToClientVpnTargetNetworkOutput struct {

	// The IDs of the applied security groups.
	SecurityGroupIds []string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationApplySecurityGroupsToClientVpnTargetNetworkMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsEc2query_serializeOpApplySecurityGroupsToClientVpnTargetNetwork{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsEc2query_deserializeOpApplySecurityGroupsToClientVpnTargetNetwork{}, middleware.After)
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
	if err = addOpApplySecurityGroupsToClientVpnTargetNetworkValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opApplySecurityGroupsToClientVpnTargetNetwork(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opApplySecurityGroupsToClientVpnTargetNetwork(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ec2",
		OperationName: "ApplySecurityGroupsToClientVpnTargetNetwork",
	}
}
