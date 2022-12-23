// Code generated by smithy-go-codegen DO NOT EDIT.

package applicationcostprofiler

import (
	"context"
	awsmiddleware "oss-sdk-go/oss/middleware"
	"oss-sdk-go/oss/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Deletes the specified report definition in AWS Application Cost Profiler. This
// stops the report from being generated.
func (c *Client) DeleteReportDefinition(ctx context.Context, params *DeleteReportDefinitionInput, optFns ...func(*Options)) (*DeleteReportDefinitionOutput, error) {
	if params == nil {
		params = &DeleteReportDefinitionInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DeleteReportDefinition", params, optFns, c.addOperationDeleteReportDefinitionMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DeleteReportDefinitionOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DeleteReportDefinitionInput struct {

	// Required. ID of the report to delete.
	//
	// This member is required.
	ReportId *string

	noSmithyDocumentSerde
}

type DeleteReportDefinitionOutput struct {

	// ID of the report that was deleted.
	ReportId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDeleteReportDefinitionMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpDeleteReportDefinition{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpDeleteReportDefinition{}, middleware.After)
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
	if err = addOpDeleteReportDefinitionValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDeleteReportDefinition(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDeleteReportDefinition(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "application-cost-profiler",
		OperationName: "DeleteReportDefinition",
	}
}
