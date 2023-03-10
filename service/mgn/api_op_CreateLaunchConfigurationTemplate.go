// Code generated by smithy-go-codegen DO NOT EDIT.

package mgn

import (
	"context"
	awsmiddleware "oss-sdk-go/oss/middleware"
	"oss-sdk-go/oss/signer/v4"
	"oss-sdk-go/service/mgn/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates a new ReplicationConfigurationTemplate.
func (c *Client) CreateLaunchConfigurationTemplate(ctx context.Context, params *CreateLaunchConfigurationTemplateInput, optFns ...func(*Options)) (*CreateLaunchConfigurationTemplateOutput, error) {
	if params == nil {
		params = &CreateLaunchConfigurationTemplateInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateLaunchConfigurationTemplate", params, optFns, c.addOperationCreateLaunchConfigurationTemplateMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateLaunchConfigurationTemplateOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateLaunchConfigurationTemplateInput struct {

	// Request to associate the default Application Migration Service Security group
	// with the Replication Settings template.
	PostLaunchActions *types.PostLaunchActions

	// Request to associate the default Application Migration Service Security group
	// with the Replication Settings template.
	Tags map[string]string

	noSmithyDocumentSerde
}

type CreateLaunchConfigurationTemplateOutput struct {

	// Copy Private IP during Launch Configuration.
	//
	// This member is required.
	LaunchConfigurationTemplateID *string

	// Copy Private IP during Launch Configuration.
	Arn *string

	// Copy Private IP during Launch Configuration.
	PostLaunchActions *types.PostLaunchActions

	// Copy Private IP during Launch Configuration.
	Tags map[string]string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateLaunchConfigurationTemplateMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpCreateLaunchConfigurationTemplate{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpCreateLaunchConfigurationTemplate{}, middleware.After)
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
	if err = addOpCreateLaunchConfigurationTemplateValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateLaunchConfigurationTemplate(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateLaunchConfigurationTemplate(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "mgn",
		OperationName: "CreateLaunchConfigurationTemplate",
	}
}
