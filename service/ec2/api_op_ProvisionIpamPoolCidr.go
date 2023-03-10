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

// Provision a CIDR to an IPAM pool. You can use this action to provision new CIDRs
// to a top-level pool or to transfer a CIDR from a top-level pool to a pool within
// it. For more information, see Provision CIDRs to pools
// (https://docs.aws.amazon.com/vpc/latest/ipam/prov-cidr-ipam.html) in the Amazon
// VPC IPAM User Guide.
func (c *Client) ProvisionIpamPoolCidr(ctx context.Context, params *ProvisionIpamPoolCidrInput, optFns ...func(*Options)) (*ProvisionIpamPoolCidrOutput, error) {
	if params == nil {
		params = &ProvisionIpamPoolCidrInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ProvisionIpamPoolCidr", params, optFns, c.addOperationProvisionIpamPoolCidrMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ProvisionIpamPoolCidrOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ProvisionIpamPoolCidrInput struct {

	// The ID of the IPAM pool to which you want to assign a CIDR.
	//
	// This member is required.
	IpamPoolId *string

	// The CIDR you want to assign to the IPAM pool.
	Cidr *string

	// A signed document that proves that you are authorized to bring a specified IP
	// address range to Amazon using BYOIP. This option applies to public pools only.
	CidrAuthorizationContext *types.IpamCidrAuthorizationContext

	// A check for whether you have the required permissions for the action without
	// actually making the request and provides an error response. If you have the
	// required permissions, the error response is DryRunOperation. Otherwise, it is
	// UnauthorizedOperation.
	DryRun *bool

	noSmithyDocumentSerde
}

type ProvisionIpamPoolCidrOutput struct {

	// Information about the provisioned CIDR.
	IpamPoolCidr *types.IpamPoolCidr

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationProvisionIpamPoolCidrMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsEc2query_serializeOpProvisionIpamPoolCidr{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsEc2query_deserializeOpProvisionIpamPoolCidr{}, middleware.After)
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
	if err = addOpProvisionIpamPoolCidrValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opProvisionIpamPoolCidr(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opProvisionIpamPoolCidr(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ec2",
		OperationName: "ProvisionIpamPoolCidr",
	}
}
