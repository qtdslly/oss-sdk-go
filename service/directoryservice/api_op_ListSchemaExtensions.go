// Code generated by smithy-go-codegen DO NOT EDIT.

package directoryservice

import (
	"context"
	awsmiddleware "oss-sdk-go/oss/middleware"
	"oss-sdk-go/oss/signer/v4"
	"oss-sdk-go/service/directoryservice/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Lists all schema extensions applied to a Microsoft AD Directory.
func (c *Client) ListSchemaExtensions(ctx context.Context, params *ListSchemaExtensionsInput, optFns ...func(*Options)) (*ListSchemaExtensionsOutput, error) {
	if params == nil {
		params = &ListSchemaExtensionsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListSchemaExtensions", params, optFns, c.addOperationListSchemaExtensionsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListSchemaExtensionsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListSchemaExtensionsInput struct {

	// The identifier of the directory from which to retrieve the schema extension
	// information.
	//
	// This member is required.
	DirectoryId *string

	// The maximum number of items to return.
	Limit *int32

	// The ListSchemaExtensions.NextToken value from a previous call to
	// ListSchemaExtensions. Pass null if this is the first call.
	NextToken *string

	noSmithyDocumentSerde
}

type ListSchemaExtensionsOutput struct {

	// If not null, more results are available. Pass this value for the NextToken
	// parameter in a subsequent call to ListSchemaExtensions to retrieve the next set
	// of items.
	NextToken *string

	// Information about the schema extensions applied to the directory.
	SchemaExtensionsInfo []types.SchemaExtensionInfo

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListSchemaExtensionsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpListSchemaExtensions{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpListSchemaExtensions{}, middleware.After)
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
	if err = addOpListSchemaExtensionsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListSchemaExtensions(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opListSchemaExtensions(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ds",
		OperationName: "ListSchemaExtensions",
	}
}
