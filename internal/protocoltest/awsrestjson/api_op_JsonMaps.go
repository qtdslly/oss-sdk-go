// Code generated by smithy-go-codegen DO NOT EDIT.

package awsrestjson

import (
	"context"
	awsmiddleware "oss-sdk-go/oss/middleware"
	"oss-sdk-go/internal/protocoltest/awsrestjson/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// The example tests basic map serialization.
func (c *Client) JsonMaps(ctx context.Context, params *JsonMapsInput, optFns ...func(*Options)) (*JsonMapsOutput, error) {
	if params == nil {
		params = &JsonMapsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "JsonMaps", params, optFns, c.addOperationJsonMapsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*JsonMapsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type JsonMapsInput struct {
	DenseBooleanMap map[string]bool

	DenseNumberMap map[string]int32

	DenseSetMap map[string][]string

	DenseStringMap map[string]string

	DenseStructMap map[string]types.GreetingStruct

	SparseBooleanMap map[string]*bool

	SparseNumberMap map[string]*int32

	SparseSetMap map[string][]string

	SparseStringMap map[string]*string

	SparseStructMap map[string]*types.GreetingStruct

	noSmithyDocumentSerde
}

type JsonMapsOutput struct {
	DenseBooleanMap map[string]bool

	DenseNumberMap map[string]int32

	DenseSetMap map[string][]string

	DenseStringMap map[string]string

	DenseStructMap map[string]types.GreetingStruct

	SparseBooleanMap map[string]*bool

	SparseNumberMap map[string]*int32

	SparseSetMap map[string][]string

	SparseStringMap map[string]*string

	SparseStructMap map[string]*types.GreetingStruct

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationJsonMapsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpJsonMaps{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpJsonMaps{}, middleware.After)
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
	if err = addRetryMiddlewares(stack, options); err != nil {
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opJsonMaps(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opJsonMaps(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "JsonMaps",
	}
}
