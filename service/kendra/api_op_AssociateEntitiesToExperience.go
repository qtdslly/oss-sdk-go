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

// Grants users or groups in your Amazon Web Services SSO identity source access to
// your Amazon Kendra experience. You can create an Amazon Kendra experience such
// as a search application. For more information on creating a search application
// experience, see Building a search experience with no code
// (https://docs.aws.amazon.com/kendra/latest/dg/deploying-search-experience-no-code.html).
func (c *Client) AssociateEntitiesToExperience(ctx context.Context, params *AssociateEntitiesToExperienceInput, optFns ...func(*Options)) (*AssociateEntitiesToExperienceOutput, error) {
	if params == nil {
		params = &AssociateEntitiesToExperienceInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "AssociateEntitiesToExperience", params, optFns, c.addOperationAssociateEntitiesToExperienceMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*AssociateEntitiesToExperienceOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type AssociateEntitiesToExperienceInput struct {

	// Lists users or groups in your Amazon Web Services SSO identity source.
	//
	// This member is required.
	EntityList []types.EntityConfiguration

	// The identifier of your Amazon Kendra experience.
	//
	// This member is required.
	Id *string

	// The identifier of the index for your Amazon Kendra experience.
	//
	// This member is required.
	IndexId *string

	noSmithyDocumentSerde
}

type AssociateEntitiesToExperienceOutput struct {

	// Lists the users or groups in your Amazon Web Services SSO identity source that
	// failed to properly configure with your Amazon Kendra experience.
	FailedEntityList []types.FailedEntity

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationAssociateEntitiesToExperienceMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpAssociateEntitiesToExperience{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpAssociateEntitiesToExperience{}, middleware.After)
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
	if err = addOpAssociateEntitiesToExperienceValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opAssociateEntitiesToExperience(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opAssociateEntitiesToExperience(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "kendra",
		OperationName: "AssociateEntitiesToExperience",
	}
}
