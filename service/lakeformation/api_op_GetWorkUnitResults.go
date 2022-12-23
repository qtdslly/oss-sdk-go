// Code generated by smithy-go-codegen DO NOT EDIT.

package lakeformation

import (
	"context"
	"fmt"
	awsmiddleware "oss-sdk-go/oss/middleware"
	"oss-sdk-go/oss/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"io"
)

// Returns the work units resulting from the query. Work units can be executed in
// any order and in parallel.
func (c *Client) GetWorkUnitResults(ctx context.Context, params *GetWorkUnitResultsInput, optFns ...func(*Options)) (*GetWorkUnitResultsOutput, error) {
	if params == nil {
		params = &GetWorkUnitResultsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetWorkUnitResults", params, optFns, c.addOperationGetWorkUnitResultsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetWorkUnitResultsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetWorkUnitResultsInput struct {

	// The ID of the plan query operation for which to get results.
	//
	// This member is required.
	QueryId *string

	// The work unit ID for which to get results. Value generated by enumerating
	// WorkUnitIdMin to WorkUnitIdMax (inclusive) from the WorkUnitRange in the output
	// of GetWorkUnits.
	//
	// This member is required.
	WorkUnitId int64

	// A work token used to query the execution service. Token output from
	// GetWorkUnits.
	//
	// This member is required.
	WorkUnitToken *string

	noSmithyDocumentSerde
}

// A structure for the output.
type GetWorkUnitResultsOutput struct {

	// Rows returned from the GetWorkUnitResults operation as a stream of Apache Arrow
	// v1.0 messages.
	ResultStream io.ReadCloser

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetWorkUnitResultsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpGetWorkUnitResults{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpGetWorkUnitResults{}, middleware.After)
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
	if err = addEndpointPrefix_opGetWorkUnitResultsMiddleware(stack); err != nil {
		return err
	}
	if err = addOpGetWorkUnitResultsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetWorkUnitResults(options.Region), middleware.Before); err != nil {
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

type endpointPrefix_opGetWorkUnitResultsMiddleware struct {
}

func (*endpointPrefix_opGetWorkUnitResultsMiddleware) ID() string {
	return "EndpointHostPrefix"
}

func (m *endpointPrefix_opGetWorkUnitResultsMiddleware) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
) {
	if smithyhttp.GetHostnameImmutable(ctx) || smithyhttp.IsEndpointHostPrefixDisabled(ctx) {
		return next.HandleSerialize(ctx, in)
	}

	req, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, fmt.Errorf("unknown transport type %T", in.Request)
	}

	req.URL.Host = "data-" + req.URL.Host

	return next.HandleSerialize(ctx, in)
}
func addEndpointPrefix_opGetWorkUnitResultsMiddleware(stack *middleware.Stack) error {
	return stack.Serialize.Insert(&endpointPrefix_opGetWorkUnitResultsMiddleware{}, `OperationSerializer`, middleware.After)
}

func newServiceMetadataMiddleware_opGetWorkUnitResults(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "lakeformation",
		OperationName: "GetWorkUnitResults",
	}
}