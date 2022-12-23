// Code generated by smithy-go-codegen DO NOT EDIT.

package organizations

import (
	"context"
	awsmiddleware "oss-sdk-go/oss/middleware"
	"oss-sdk-go/oss/signer/v4"
	"oss-sdk-go/service/organizations/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Retrieves information about the organization that the user's account belongs to.
// This operation can be called from any account in the organization. Even if a
// policy type is shown as available in the organization, you can disable it
// separately at the root level with DisablePolicyType. Use ListRoots to see the
// status of policy types for a specified root.
func (c *Client) DescribeOrganization(ctx context.Context, params *DescribeOrganizationInput, optFns ...func(*Options)) (*DescribeOrganizationOutput, error) {
	if params == nil {
		params = &DescribeOrganizationInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeOrganization", params, optFns, c.addOperationDescribeOrganizationMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeOrganizationOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeOrganizationInput struct {
	noSmithyDocumentSerde
}

type DescribeOrganizationOutput struct {

	// A structure that contains information about the organization. The
	// AvailablePolicyTypes part of the response is deprecated, and you shouldn't use
	// it in your apps. It doesn't include any policy type supported by Organizations
	// other than SCPs. To determine which policy types are enabled in your
	// organization, use the ListRoots operation.
	Organization *types.Organization

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeOrganizationMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDescribeOrganization{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDescribeOrganization{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeOrganization(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDescribeOrganization(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "organizations",
		OperationName: "DescribeOrganization",
	}
}
