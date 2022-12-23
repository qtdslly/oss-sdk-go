// Code generated by smithy-go-codegen DO NOT EDIT.

package route53resolver

import (
	"context"
	awsmiddleware "oss-sdk-go/oss/middleware"
	"oss-sdk-go/oss/signer/v4"
	"oss-sdk-go/service/route53resolver/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Retrieves the behavior configuration of Route 53 Resolver behavior for a single
// VPC from Amazon Virtual Private Cloud.
func (c *Client) GetResolverConfig(ctx context.Context, params *GetResolverConfigInput, optFns ...func(*Options)) (*GetResolverConfigOutput, error) {
	if params == nil {
		params = &GetResolverConfigInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetResolverConfig", params, optFns, c.addOperationGetResolverConfigMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetResolverConfigOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetResolverConfigInput struct {

	// Resource ID of the Amazon VPC that you want to get information about.
	//
	// This member is required.
	ResourceId *string

	noSmithyDocumentSerde
}

type GetResolverConfigOutput struct {

	// Information about the behavior configuration of Route 53 Resolver behavior for
	// the VPC you specified in the GetResolverConfig request.
	ResolverConfig *types.ResolverConfig

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetResolverConfigMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpGetResolverConfig{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpGetResolverConfig{}, middleware.After)
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
	if err = addOpGetResolverConfigValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetResolverConfig(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opGetResolverConfig(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "route53resolver",
		OperationName: "GetResolverConfig",
	}
}
