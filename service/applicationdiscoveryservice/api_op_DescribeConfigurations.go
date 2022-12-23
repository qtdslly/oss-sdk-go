// Code generated by smithy-go-codegen DO NOT EDIT.

package applicationdiscoveryservice

import (
	"context"
	awsmiddleware "oss-sdk-go/oss/middleware"
	"oss-sdk-go/oss/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Retrieves attributes for a list of configuration item IDs. All of the supplied
// IDs must be for the same asset type from one of the following:
//
// * server
//
// *
// application
//
// * process
//
// * connection
//
// Output fields are specific to the asset
// type specified. For example, the output for a server configuration item includes
// a list of attributes about the server, such as host name, operating system,
// number of network cards, etc. For a complete list of outputs for each asset
// type, see Using the DescribeConfigurations Action
// (https://docs.aws.amazon.com/application-discovery/latest/userguide/discovery-api-queries.html#DescribeConfigurations)
// in the Amazon Web Services Application Discovery Service User Guide.
func (c *Client) DescribeConfigurations(ctx context.Context, params *DescribeConfigurationsInput, optFns ...func(*Options)) (*DescribeConfigurationsOutput, error) {
	if params == nil {
		params = &DescribeConfigurationsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeConfigurations", params, optFns, c.addOperationDescribeConfigurationsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeConfigurationsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeConfigurationsInput struct {

	// One or more configuration IDs.
	//
	// This member is required.
	ConfigurationIds []string

	noSmithyDocumentSerde
}

type DescribeConfigurationsOutput struct {

	// A key in the response map. The value is an array of data.
	Configurations []map[string]string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeConfigurationsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDescribeConfigurations{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDescribeConfigurations{}, middleware.After)
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
	if err = addOpDescribeConfigurationsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeConfigurations(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDescribeConfigurations(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "discovery",
		OperationName: "DescribeConfigurations",
	}
}
