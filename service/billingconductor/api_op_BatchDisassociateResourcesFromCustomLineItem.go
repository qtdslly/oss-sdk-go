// Code generated by smithy-go-codegen DO NOT EDIT.

package billingconductor

import (
	"context"
	awsmiddleware "oss-sdk-go/oss/middleware"
	"oss-sdk-go/oss/signer/v4"
	"oss-sdk-go/service/billingconductor/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Disassociates a batch of resources from a percentage custom line item.
func (c *Client) BatchDisassociateResourcesFromCustomLineItem(ctx context.Context, params *BatchDisassociateResourcesFromCustomLineItemInput, optFns ...func(*Options)) (*BatchDisassociateResourcesFromCustomLineItemOutput, error) {
	if params == nil {
		params = &BatchDisassociateResourcesFromCustomLineItemInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "BatchDisassociateResourcesFromCustomLineItem", params, optFns, c.addOperationBatchDisassociateResourcesFromCustomLineItemMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*BatchDisassociateResourcesFromCustomLineItemOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type BatchDisassociateResourcesFromCustomLineItemInput struct {

	// A list containing the ARNs of resources to be disassociated.
	//
	// This member is required.
	ResourceArns []string

	// A percentage custom line item ARN to disassociate the resources from.
	//
	// This member is required.
	TargetArn *string

	// The billing period range in which the custom line item request will be applied.
	BillingPeriodRange *types.CustomLineItemBillingPeriodRange

	noSmithyDocumentSerde
}

type BatchDisassociateResourcesFromCustomLineItemOutput struct {

	// A list of DisassociateResourceResponseElement for each resource that failed
	// disassociation from a percentage custom line item.
	FailedDisassociatedResources []types.DisassociateResourceResponseElement

	// A list of DisassociateResourceResponseElement for each resource that's been
	// disassociated from a percentage custom line item successfully.
	SuccessfullyDisassociatedResources []types.DisassociateResourceResponseElement

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationBatchDisassociateResourcesFromCustomLineItemMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpBatchDisassociateResourcesFromCustomLineItem{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpBatchDisassociateResourcesFromCustomLineItem{}, middleware.After)
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
	if err = addOpBatchDisassociateResourcesFromCustomLineItemValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opBatchDisassociateResourcesFromCustomLineItem(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opBatchDisassociateResourcesFromCustomLineItem(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "billingconductor",
		OperationName: "BatchDisassociateResourcesFromCustomLineItem",
	}
}
