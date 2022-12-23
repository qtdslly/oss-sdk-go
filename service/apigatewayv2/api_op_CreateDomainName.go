// Code generated by smithy-go-codegen DO NOT EDIT.

package apigatewayv2

import (
	"context"
	awsmiddleware "oss-sdk-go/oss/middleware"
	"oss-sdk-go/oss/signer/v4"
	"oss-sdk-go/service/apigatewayv2/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates a domain name.
func (c *Client) CreateDomainName(ctx context.Context, params *CreateDomainNameInput, optFns ...func(*Options)) (*CreateDomainNameOutput, error) {
	if params == nil {
		params = &CreateDomainNameInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateDomainName", params, optFns, c.addOperationCreateDomainNameMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateDomainNameOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Creates a new DomainName resource to represent a domain name.
type CreateDomainNameInput struct {

	// The domain name.
	//
	// This member is required.
	DomainName *string

	// The domain name configurations.
	DomainNameConfigurations []types.DomainNameConfiguration

	// The mutual TLS authentication configuration for a custom domain name.
	MutualTlsAuthentication *types.MutualTlsAuthenticationInput

	// The collection of tags associated with a domain name.
	Tags map[string]string

	noSmithyDocumentSerde
}

type CreateDomainNameOutput struct {

	// The API mapping selection expression.
	ApiMappingSelectionExpression *string

	// The name of the DomainName resource.
	DomainName *string

	// The domain name configurations.
	DomainNameConfigurations []types.DomainNameConfiguration

	// The mutual TLS authentication configuration for a custom domain name.
	MutualTlsAuthentication *types.MutualTlsAuthentication

	// The collection of tags associated with a domain name.
	Tags map[string]string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateDomainNameMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpCreateDomainName{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpCreateDomainName{}, middleware.After)
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
	if err = addOpCreateDomainNameValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateDomainName(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateDomainName(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "apigateway",
		OperationName: "CreateDomainName",
	}
}
