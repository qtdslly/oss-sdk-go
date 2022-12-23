// Code generated by smithy-go-codegen DO NOT EDIT.

package fms

import (
	"context"
	awsmiddleware "oss-sdk-go/oss/middleware"
	"oss-sdk-go/oss/signer/v4"
	"oss-sdk-go/service/fms/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Disassociates a Firewall Manager policy administrator from a third-party
// firewall tenant. When you call DisassociateThirdPartyFirewall, the third-party
// firewall vendor deletes all of the firewalls that are associated with the
// account.
func (c *Client) DisassociateThirdPartyFirewall(ctx context.Context, params *DisassociateThirdPartyFirewallInput, optFns ...func(*Options)) (*DisassociateThirdPartyFirewallOutput, error) {
	if params == nil {
		params = &DisassociateThirdPartyFirewallInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DisassociateThirdPartyFirewall", params, optFns, c.addOperationDisassociateThirdPartyFirewallMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DisassociateThirdPartyFirewallOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DisassociateThirdPartyFirewallInput struct {

	// The name of the third-party firewall vendor.
	//
	// This member is required.
	ThirdPartyFirewall types.ThirdPartyFirewall

	noSmithyDocumentSerde
}

type DisassociateThirdPartyFirewallOutput struct {

	// The current status for the disassociation of a Firewall Manager administrators
	// account with a third-party firewall.
	ThirdPartyFirewallStatus types.ThirdPartyFirewallAssociationStatus

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDisassociateThirdPartyFirewallMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDisassociateThirdPartyFirewall{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDisassociateThirdPartyFirewall{}, middleware.After)
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
	if err = addOpDisassociateThirdPartyFirewallValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDisassociateThirdPartyFirewall(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDisassociateThirdPartyFirewall(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "fms",
		OperationName: "DisassociateThirdPartyFirewall",
	}
}