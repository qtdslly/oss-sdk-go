// Code generated by smithy-go-codegen DO NOT EDIT.

package kendra

import (
	"context"
	awsmiddleware "oss-sdk-go/oss/middleware"
	"oss-sdk-go/oss/signer/v4"
	"oss-sdk-go/service/kendra/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Defines the specific permissions of users or groups in your Amazon Web Services
// SSO identity source with access to your Amazon Kendra experience. You can create
// an Amazon Kendra experience such as a search application. For more information
// on creating a search application experience, see Building a search experience
// with no code
// (https://docs.aws.amazon.com/kendra/latest/dg/deploying-search-experience-no-code.html).
func (c *Client) AssociatePersonasToEntities(ctx context.Context, params *AssociatePersonasToEntitiesInput, optFns ...func(*Options)) (*AssociatePersonasToEntitiesOutput, error) {
	if params == nil {
		params = &AssociatePersonasToEntitiesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "AssociatePersonasToEntities", params, optFns, c.addOperationAssociatePersonasToEntitiesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*AssociatePersonasToEntitiesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type AssociatePersonasToEntitiesInput struct {

	// The identifier of your Amazon Kendra experience.
	//
	// This member is required.
	Id *string

	// The identifier of the index for your Amazon Kendra experience.
	//
	// This member is required.
	IndexId *string

	// The personas that define the specific permissions of users or groups in your
	// Amazon Web Services SSO identity source. The available personas or access roles
	// are Owner and Viewer. For more information on these personas, see Providing
	// access to your search page
	// (https://docs.aws.amazon.com/kendra/latest/dg/deploying-search-experience-no-code.html#access-search-experience).
	//
	// This member is required.
	Personas []types.EntityPersonaConfiguration

	noSmithyDocumentSerde
}

type AssociatePersonasToEntitiesOutput struct {

	// Lists the users or groups in your Amazon Web Services SSO identity source that
	// failed to properly configure with your Amazon Kendra experience.
	FailedEntityList []types.FailedEntity

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationAssociatePersonasToEntitiesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpAssociatePersonasToEntities{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpAssociatePersonasToEntities{}, middleware.After)
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
	if err = addOpAssociatePersonasToEntitiesValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opAssociatePersonasToEntities(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opAssociatePersonasToEntities(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "kendra",
		OperationName: "AssociatePersonasToEntities",
	}
}
