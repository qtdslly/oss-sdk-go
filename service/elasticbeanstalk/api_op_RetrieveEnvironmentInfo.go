// Code generated by smithy-go-codegen DO NOT EDIT.

package elasticbeanstalk

import (
	"context"
	awsmiddleware "oss-sdk-go/oss/middleware"
	"oss-sdk-go/oss/signer/v4"
	"oss-sdk-go/service/elasticbeanstalk/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Retrieves the compiled information from a RequestEnvironmentInfo request.
// Related Topics
//
// * RequestEnvironmentInfo
func (c *Client) RetrieveEnvironmentInfo(ctx context.Context, params *RetrieveEnvironmentInfoInput, optFns ...func(*Options)) (*RetrieveEnvironmentInfoOutput, error) {
	if params == nil {
		params = &RetrieveEnvironmentInfoInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "RetrieveEnvironmentInfo", params, optFns, c.addOperationRetrieveEnvironmentInfoMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*RetrieveEnvironmentInfoOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Request to download logs retrieved with RequestEnvironmentInfo.
type RetrieveEnvironmentInfoInput struct {

	// The type of information to retrieve.
	//
	// This member is required.
	InfoType types.EnvironmentInfoType

	// The ID of the data's environment. If no such environment is found, returns an
	// InvalidParameterValue error. Condition: You must specify either this or an
	// EnvironmentName, or both. If you do not specify either, AWS Elastic Beanstalk
	// returns MissingRequiredParameter error.
	EnvironmentId *string

	// The name of the data's environment. If no such environment is found, returns an
	// InvalidParameterValue error. Condition: You must specify either this or an
	// EnvironmentId, or both. If you do not specify either, AWS Elastic Beanstalk
	// returns MissingRequiredParameter error.
	EnvironmentName *string

	noSmithyDocumentSerde
}

// Result message containing a description of the requested environment info.
type RetrieveEnvironmentInfoOutput struct {

	// The EnvironmentInfoDescription of the environment.
	EnvironmentInfo []types.EnvironmentInfoDescription

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationRetrieveEnvironmentInfoMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsquery_serializeOpRetrieveEnvironmentInfo{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpRetrieveEnvironmentInfo{}, middleware.After)
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
	if err = addOpRetrieveEnvironmentInfoValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opRetrieveEnvironmentInfo(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opRetrieveEnvironmentInfo(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "elasticbeanstalk",
		OperationName: "RetrieveEnvironmentInfo",
	}
}
