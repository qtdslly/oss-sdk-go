// Code generated by smithy-go-codegen DO NOT EDIT.

package ec2

import (
	"context"
	awsmiddleware "oss-sdk-go/oss/middleware"
	"oss-sdk-go/oss/signer/v4"
	"oss-sdk-go/service/ec2/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Registers members (network interfaces) with the transit gateway multicast group.
// A member is a network interface associated with a supported EC2 instance that
// receives multicast traffic. For information about supported instances, see
// Multicast Consideration
// (https://docs.aws.amazon.com/vpc/latest/tgw/transit-gateway-limits.html#multicast-limits)
// in Amazon VPC Transit Gateways. After you add the members, use
// SearchTransitGatewayMulticastGroups
// (https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_SearchTransitGatewayMulticastGroups.html)
// to verify that the members were added to the transit gateway multicast group.
func (c *Client) RegisterTransitGatewayMulticastGroupMembers(ctx context.Context, params *RegisterTransitGatewayMulticastGroupMembersInput, optFns ...func(*Options)) (*RegisterTransitGatewayMulticastGroupMembersOutput, error) {
	if params == nil {
		params = &RegisterTransitGatewayMulticastGroupMembersInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "RegisterTransitGatewayMulticastGroupMembers", params, optFns, c.addOperationRegisterTransitGatewayMulticastGroupMembersMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*RegisterTransitGatewayMulticastGroupMembersOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type RegisterTransitGatewayMulticastGroupMembersInput struct {

	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have the
	// required permissions, the error response is DryRunOperation. Otherwise, it is
	// UnauthorizedOperation.
	DryRun *bool

	// The IP address assigned to the transit gateway multicast group.
	GroupIpAddress *string

	// The group members' network interface IDs to register with the transit gateway
	// multicast group.
	NetworkInterfaceIds []string

	// The ID of the transit gateway multicast domain.
	TransitGatewayMulticastDomainId *string

	noSmithyDocumentSerde
}

type RegisterTransitGatewayMulticastGroupMembersOutput struct {

	// Information about the registered transit gateway multicast group members.
	RegisteredMulticastGroupMembers *types.TransitGatewayMulticastRegisteredGroupMembers

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationRegisterTransitGatewayMulticastGroupMembersMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsEc2query_serializeOpRegisterTransitGatewayMulticastGroupMembers{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsEc2query_deserializeOpRegisterTransitGatewayMulticastGroupMembers{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opRegisterTransitGatewayMulticastGroupMembers(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opRegisterTransitGatewayMulticastGroupMembers(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ec2",
		OperationName: "RegisterTransitGatewayMulticastGroupMembers",
	}
}
