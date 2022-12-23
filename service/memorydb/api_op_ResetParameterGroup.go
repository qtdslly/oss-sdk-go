// Code generated by smithy-go-codegen DO NOT EDIT.

package memorydb

import (
	"context"
	awsmiddleware "oss-sdk-go/oss/middleware"
	"oss-sdk-go/oss/signer/v4"
	"oss-sdk-go/service/memorydb/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Modifies the parameters of a parameter group to the engine or system default
// value. You can reset specific parameters by submitting a list of parameter
// names. To reset the entire parameter group, specify the AllParameters and
// ParameterGroupName parameters.
func (c *Client) ResetParameterGroup(ctx context.Context, params *ResetParameterGroupInput, optFns ...func(*Options)) (*ResetParameterGroupOutput, error) {
	if params == nil {
		params = &ResetParameterGroupInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ResetParameterGroup", params, optFns, c.addOperationResetParameterGroupMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ResetParameterGroupOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ResetParameterGroupInput struct {

	// The name of the parameter group to reset.
	//
	// This member is required.
	ParameterGroupName *string

	// If true, all parameters in the parameter group are reset to their default
	// values. If false, only the parameters listed by ParameterNames are reset to
	// their default values.
	AllParameters bool

	// An array of parameter names to reset to their default values. If AllParameters
	// is true, do not use ParameterNames. If AllParameters is false, you must specify
	// the name of at least one parameter to reset.
	ParameterNames []string

	noSmithyDocumentSerde
}

type ResetParameterGroupOutput struct {

	// The parameter group being reset.
	ParameterGroup *types.ParameterGroup

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationResetParameterGroupMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpResetParameterGroup{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpResetParameterGroup{}, middleware.After)
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
	if err = addOpResetParameterGroupValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opResetParameterGroup(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opResetParameterGroup(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "memorydb",
		OperationName: "ResetParameterGroup",
	}
}
