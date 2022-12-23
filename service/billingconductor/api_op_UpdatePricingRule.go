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

// Updates an existing pricing rule.
func (c *Client) UpdatePricingRule(ctx context.Context, params *UpdatePricingRuleInput, optFns ...func(*Options)) (*UpdatePricingRuleOutput, error) {
	if params == nil {
		params = &UpdatePricingRuleInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdatePricingRule", params, optFns, c.addOperationUpdatePricingRuleMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdatePricingRuleOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdatePricingRuleInput struct {

	// The Amazon Resource Name (ARN) of the pricing rule to update.
	//
	// This member is required.
	Arn *string

	// The new description for the pricing rule.
	Description *string

	// The new modifier to show pricing plan rates as a percentage.
	ModifierPercentage *float64

	// The new name of the pricing rule. The name must be unique to each pricing rule.
	Name *string

	// The new pricing rule type.
	Type types.PricingRuleType

	noSmithyDocumentSerde
}

type UpdatePricingRuleOutput struct {

	// The Amazon Resource Name (ARN) of the successfully updated pricing rule.
	Arn *string

	// The pricing plans count that this pricing rule is associated with.
	AssociatedPricingPlanCount int64

	// The new description for the pricing rule.
	Description *string

	// The most recent time the pricing rule was modified.
	LastModifiedTime int64

	// The new modifier to show pricing plan rates as a percentage.
	ModifierPercentage *float64

	// The new name of the pricing rule. The name must be unique to each pricing rule.
	Name *string

	// The scope of pricing rule that indicates if it is globally applicable, or is
	// service-specific.
	Scope types.PricingRuleScope

	// If the Scope attribute is set to SERVICE, the attribute indicates which service
	// the PricingRule is applicable for.
	Service *string

	// The new pricing rule type.
	Type types.PricingRuleType

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationUpdatePricingRuleMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpUpdatePricingRule{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpUpdatePricingRule{}, middleware.After)
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
	if err = addOpUpdatePricingRuleValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUpdatePricingRule(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opUpdatePricingRule(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "billingconductor",
		OperationName: "UpdatePricingRule",
	}
}
