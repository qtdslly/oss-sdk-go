// Code generated by smithy-go-codegen DO NOT EDIT.

package servicequotas

import (
	"context"
	awsmiddleware "oss-sdk-go/oss/middleware"
	"oss-sdk-go/oss/signer/v4"
	"oss-sdk-go/service/servicequotas/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Submits a quota increase request for the specified quota.
func (c *Client) RequestServiceQuotaIncrease(ctx context.Context, params *RequestServiceQuotaIncreaseInput, optFns ...func(*Options)) (*RequestServiceQuotaIncreaseOutput, error) {
	if params == nil {
		params = &RequestServiceQuotaIncreaseInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "RequestServiceQuotaIncrease", params, optFns, c.addOperationRequestServiceQuotaIncreaseMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*RequestServiceQuotaIncreaseOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type RequestServiceQuotaIncreaseInput struct {

	// The new, increased value for the quota.
	//
	// This member is required.
	DesiredValue *float64

	// The quota identifier.
	//
	// This member is required.
	QuotaCode *string

	// The service identifier.
	//
	// This member is required.
	ServiceCode *string

	noSmithyDocumentSerde
}

type RequestServiceQuotaIncreaseOutput struct {

	// Information about the quota increase request.
	RequestedQuota *types.RequestedServiceQuotaChange

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationRequestServiceQuotaIncreaseMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpRequestServiceQuotaIncrease{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpRequestServiceQuotaIncrease{}, middleware.After)
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
	if err = addOpRequestServiceQuotaIncreaseValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opRequestServiceQuotaIncrease(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opRequestServiceQuotaIncrease(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "servicequotas",
		OperationName: "RequestServiceQuotaIncrease",
	}
}