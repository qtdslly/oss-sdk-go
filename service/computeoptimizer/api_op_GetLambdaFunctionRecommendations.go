// Code generated by smithy-go-codegen DO NOT EDIT.

package computeoptimizer

import (
	"context"
	awsmiddleware "oss-sdk-go/oss/middleware"
	"oss-sdk-go/oss/signer/v4"
	"oss-sdk-go/service/computeoptimizer/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Returns Lambda function recommendations. Compute Optimizer generates
// recommendations for functions that meet a specific set of requirements. For more
// information, see the Supported resources and requirements
// (https://docs.aws.amazon.com/compute-optimizer/latest/ug/requirements.html) in
// the Compute Optimizer User Guide.
func (c *Client) GetLambdaFunctionRecommendations(ctx context.Context, params *GetLambdaFunctionRecommendationsInput, optFns ...func(*Options)) (*GetLambdaFunctionRecommendationsOutput, error) {
	if params == nil {
		params = &GetLambdaFunctionRecommendationsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetLambdaFunctionRecommendations", params, optFns, c.addOperationGetLambdaFunctionRecommendationsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetLambdaFunctionRecommendationsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetLambdaFunctionRecommendationsInput struct {

	// The ID of the Amazon Web Services account for which to return function
	// recommendations. If your account is the management account of an organization,
	// use this parameter to specify the member account for which you want to return
	// function recommendations. Only one account ID can be specified per request.
	AccountIds []string

	// An array of objects to specify a filter that returns a more specific list of
	// function recommendations.
	Filters []types.LambdaFunctionRecommendationFilter

	// The Amazon Resource Name (ARN) of the functions for which to return
	// recommendations. You can specify a qualified or unqualified ARN. If you specify
	// an unqualified ARN without a function version suffix, Compute Optimizer will
	// return recommendations for the latest ($LATEST) version of the function. If you
	// specify a qualified ARN with a version suffix, Compute Optimizer will return
	// recommendations for the specified function version. For more information about
	// using function versions, see Using versions
	// (https://docs.aws.amazon.com/lambda/latest/dg/configuration-versions.html#versioning-versions-using)
	// in the Lambda Developer Guide.
	FunctionArns []string

	// The maximum number of function recommendations to return with a single request.
	// To retrieve the remaining results, make another request with the returned
	// nextToken value.
	MaxResults *int32

	// The token to advance to the next page of function recommendations.
	NextToken *string

	noSmithyDocumentSerde
}

type GetLambdaFunctionRecommendationsOutput struct {

	// An array of objects that describe function recommendations.
	LambdaFunctionRecommendations []types.LambdaFunctionRecommendation

	// The token to use to advance to the next page of function recommendations. This
	// value is null when there are no more pages of function recommendations to
	// return.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetLambdaFunctionRecommendationsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson10_serializeOpGetLambdaFunctionRecommendations{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson10_deserializeOpGetLambdaFunctionRecommendations{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetLambdaFunctionRecommendations(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opGetLambdaFunctionRecommendations(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "compute-optimizer",
		OperationName: "GetLambdaFunctionRecommendations",
	}
}
