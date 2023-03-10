// Code generated by smithy-go-codegen DO NOT EDIT.

package opensearch

import (
	"context"
	awsmiddleware "oss-sdk-go/oss/middleware"
	"oss-sdk-go/oss/signer/v4"
	"oss-sdk-go/service/opensearch/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Updates a package for use with Amazon OpenSearch Service domains.
func (c *Client) UpdatePackage(ctx context.Context, params *UpdatePackageInput, optFns ...func(*Options)) (*UpdatePackageOutput, error) {
	if params == nil {
		params = &UpdatePackageInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdatePackage", params, optFns, c.addOperationUpdatePackageMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdatePackageOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Container for request parameters to the UpdatePackage operation.
type UpdatePackageInput struct {

	// The unique identifier for the package.
	//
	// This member is required.
	PackageID *string

	// The Amazon S3 location for importing the package specified as S3BucketName and
	// S3Key
	//
	// This member is required.
	PackageSource *types.PackageSource

	// A commit message for the new version which is shown as part of
	// GetPackageVersionHistoryResponse.
	CommitMessage *string

	// A new description of the package.
	PackageDescription *string

	noSmithyDocumentSerde
}

// Container for the response returned by the UpdatePackage operation.
type UpdatePackageOutput struct {

	// Information about the package.
	PackageDetails *types.PackageDetails

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationUpdatePackageMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpUpdatePackage{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpUpdatePackage{}, middleware.After)
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
	if err = addOpUpdatePackageValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUpdatePackage(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opUpdatePackage(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "es",
		OperationName: "UpdatePackage",
	}
}
