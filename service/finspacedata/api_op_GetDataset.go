// Code generated by smithy-go-codegen DO NOT EDIT.

package finspacedata

import (
	"context"
	awsmiddleware "oss-sdk-go/oss/middleware"
	"oss-sdk-go/oss/signer/v4"
	"oss-sdk-go/service/finspacedata/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Returns information about a Dataset.
func (c *Client) GetDataset(ctx context.Context, params *GetDatasetInput, optFns ...func(*Options)) (*GetDatasetOutput, error) {
	if params == nil {
		params = &GetDatasetInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetDataset", params, optFns, c.addOperationGetDatasetMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetDatasetOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Request for the GetDataset operation.
type GetDatasetInput struct {

	// The unique identifier for a Dataset.
	//
	// This member is required.
	DatasetId *string

	noSmithyDocumentSerde
}

// Response for the GetDataset operation
type GetDatasetOutput struct {

	// The unique resource identifier for a Dataset.
	Alias *string

	// The timestamp at which the Dataset was created in FinSpace. The value is
	// determined as epoch time in milliseconds. For example, the value for Monday,
	// November 1, 2021 12:00:00 PM UTC is specified as 1635768000000.
	CreateTime int64

	// The ARN identifier of the Dataset.
	DatasetArn *string

	// A description of the Dataset.
	DatasetDescription *string

	// The unique identifier for a Dataset.
	DatasetId *string

	// Display title for a Dataset.
	DatasetTitle *string

	// The format in which Dataset data is structured.
	//
	// * TABULAR – Data is structured
	// in a tabular format.
	//
	// * NON_TABULAR – Data is structured in a non-tabular
	// format.
	Kind types.DatasetKind

	// The last time that the Dataset was modified. The value is determined as epoch
	// time in milliseconds. For example, the value for Monday, November 1, 2021
	// 12:00:00 PM UTC is specified as 1635768000000.
	LastModifiedTime int64

	// Definition for a schema on a tabular Dataset.
	SchemaDefinition *types.SchemaUnion

	// Status of the Dataset creation.
	//
	// * PENDING – Dataset is pending creation.
	//
	// *
	// FAILED – Dataset creation has failed.
	//
	// * SUCCESS – Dataset creation has
	// succeeded.
	//
	// * RUNNING – Dataset creation is running.
	Status types.DatasetStatus

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetDatasetMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpGetDataset{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpGetDataset{}, middleware.After)
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
	if err = addRestJsonContentTypeCustomization(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddErrorCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = addOpGetDatasetValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetDataset(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opGetDataset(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "finspace-api",
		OperationName: "GetDataset",
	}
}
