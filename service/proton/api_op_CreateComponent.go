// Code generated by smithy-go-codegen DO NOT EDIT.

package proton

import (
	"context"
	awsmiddleware "oss-sdk-go/oss/middleware"
	"oss-sdk-go/oss/signer/v4"
	"oss-sdk-go/service/proton/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Create an Proton component. A component is an infrastructure extension for a
// service instance. For more information about components, see Proton components
// (https://docs.aws.amazon.com/proton/latest/adminguide/ag-components.html) in the
// Proton Administrator Guide.
func (c *Client) CreateComponent(ctx context.Context, params *CreateComponentInput, optFns ...func(*Options)) (*CreateComponentOutput, error) {
	if params == nil {
		params = &CreateComponentInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateComponent", params, optFns, c.addOperationCreateComponentMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateComponentOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateComponentInput struct {

	// A path to a manifest file that lists the Infrastructure as Code (IaC) file,
	// template language, and rendering engine for infrastructure that a custom
	// component provisions.
	//
	// This value conforms to the media type: application/yaml
	//
	// This member is required.
	Manifest *string

	// The customer-provided name of the component.
	//
	// This member is required.
	Name *string

	// A path to the Infrastructure as Code (IaC) file describing infrastructure that a
	// custom component provisions. Components support a single IaC file, even if you
	// use Terraform as your template language.
	//
	// This value conforms to the media type: application/yaml
	//
	// This member is required.
	TemplateFile *string

	// An optional customer-provided description of the component.
	Description *string

	// The name of the Proton environment that you want to associate this component
	// with. You must specify this when you don't specify serviceInstanceName and
	// serviceName.
	EnvironmentName *string

	// The name of the service instance that you want to attach this component to. If
	// you don't specify this, the component isn't attached to any service instance.
	// Specify both serviceInstanceName and serviceName or neither of them.
	ServiceInstanceName *string

	// The name of the service that serviceInstanceName is associated with. If you
	// don't specify this, the component isn't attached to any service instance.
	// Specify both serviceInstanceName and serviceName or neither of them.
	ServiceName *string

	// The service spec that you want the component to use to access service inputs.
	// Set this only when you attach the component to a service instance.
	//
	// This value conforms to the media type: application/yaml
	ServiceSpec *string

	// An optional list of metadata items that you can associate with the Proton
	// component. A tag is a key-value pair. For more information, see Proton resources
	// and tagging in the Proton Administrator Guide
	// (https://docs.aws.amazon.com/proton/latest/adminguide/resources.html) or Proton
	// User Guide (https://docs.aws.amazon.com/proton/latest/userguide/resources.html).
	Tags []types.Tag

	noSmithyDocumentSerde
}

type CreateComponentOutput struct {

	// The detailed data of the created component.
	//
	// This member is required.
	Component *types.Component

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateComponentMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson10_serializeOpCreateComponent{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson10_deserializeOpCreateComponent{}, middleware.After)
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
	if err = addOpCreateComponentValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateComponent(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateComponent(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "proton",
		OperationName: "CreateComponent",
	}
}