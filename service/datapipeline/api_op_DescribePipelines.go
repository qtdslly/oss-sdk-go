// Code generated by smithy-go-codegen DO NOT EDIT.

package datapipeline

import (
	"context"
	awsmiddleware "oss-sdk-go/oss/middleware"
	"oss-sdk-go/oss/signer/v4"
	"oss-sdk-go/service/datapipeline/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Retrieves metadata about one or more pipelines. The information retrieved
// includes the name of the pipeline, the pipeline identifier, its current state,
// and the user account that owns the pipeline. Using account credentials, you can
// retrieve metadata about pipelines that you or your IAM users have created. If
// you are using an IAM user account, you can retrieve metadata about only those
// pipelines for which you have read permissions. To retrieve the full pipeline
// definition instead of metadata about the pipeline, call GetPipelineDefinition.
// POST / HTTP/1.1 Content-Type: application/x-amz-json-1.1 X-Amz-Target:
// DataPipeline.DescribePipelines Content-Length: 70 Host:
// datapipeline.us-east-1.amazonaws.com X-Amz-Date: Mon, 12 Nov 2012 17:49:52 GMT
// Authorization: AuthParams {"pipelineIds": ["df-08785951KAKJEXAMPLE"] }
// x-amzn-RequestId: 02870eb7-0736-11e2-af6f-6bc7a6be60d9 Content-Type:
// application/x-amz-json-1.1 Content-Length: 767 Date: Mon, 12 Nov 2012 17:50:53
// GMT {"pipelineDescriptionList": [ {"description": "This is my first pipeline",
// "fields": [ {"key": "@pipelineState", "stringValue": "SCHEDULED"}, {"key":
// "description", "stringValue": "This is my first pipeline"}, {"key": "name",
// "stringValue": "myPipeline"}, {"key": "@creationTime", "stringValue":
// "2012-12-13T01:24:06"}, {"key": "@id", "stringValue": "df-0937003356ZJEXAMPLE"},
// {"key": "@sphere", "stringValue": "PIPELINE"}, {"key": "@version",
// "stringValue": "1"}, {"key": "@userId", "stringValue": "924374875933"}, {"key":
// "@accountId", "stringValue": "924374875933"}, {"key": "uniqueId", "stringValue":
// "1234567890"} ], "name": "myPipeline", "pipelineId": "df-0937003356ZJEXAMPLE"} ]
// }
func (c *Client) DescribePipelines(ctx context.Context, params *DescribePipelinesInput, optFns ...func(*Options)) (*DescribePipelinesOutput, error) {
	if params == nil {
		params = &DescribePipelinesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribePipelines", params, optFns, c.addOperationDescribePipelinesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribePipelinesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Contains the parameters for DescribePipelines.
type DescribePipelinesInput struct {

	// The IDs of the pipelines to describe. You can pass as many as 25 identifiers in
	// a single call. To obtain pipeline IDs, call ListPipelines.
	//
	// This member is required.
	PipelineIds []string

	noSmithyDocumentSerde
}

// Contains the output of DescribePipelines.
type DescribePipelinesOutput struct {

	// An array of descriptions for the specified pipelines.
	//
	// This member is required.
	PipelineDescriptionList []types.PipelineDescription

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribePipelinesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDescribePipelines{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDescribePipelines{}, middleware.After)
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
	if err = addOpDescribePipelinesValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribePipelines(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDescribePipelines(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "datapipeline",
		OperationName: "DescribePipelines",
	}
}
