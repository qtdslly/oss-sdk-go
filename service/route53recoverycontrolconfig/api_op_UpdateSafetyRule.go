// Code generated by smithy-go-codegen DO NOT EDIT.

package route53recoverycontrolconfig

import (
	"context"
	awsmiddleware "oss-sdk-go/oss/middleware"
	"oss-sdk-go/oss/signer/v4"
	"oss-sdk-go/service/route53recoverycontrolconfig/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Update a safety rule (an assertion rule or gating rule). You can only update the
// name and the waiting period for a safety rule. To make other updates, delete the
// safety rule and create a new one.
func (c *Client) UpdateSafetyRule(ctx context.Context, params *UpdateSafetyRuleInput, optFns ...func(*Options)) (*UpdateSafetyRuleOutput, error) {
	if params == nil {
		params = &UpdateSafetyRuleInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateSafetyRule", params, optFns, c.addOperationUpdateSafetyRuleMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateSafetyRuleOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// A rule that you add to Application Recovery Controller to ensure that recovery
// actions don't accidentally impair your application's availability.
type UpdateSafetyRuleInput struct {

	// The assertion rule to update.
	AssertionRuleUpdate *types.AssertionRuleUpdate

	// The gating rule to update.
	GatingRuleUpdate *types.GatingRuleUpdate

	noSmithyDocumentSerde
}

type UpdateSafetyRuleOutput struct {

	// The assertion rule updated.
	AssertionRule *types.AssertionRule

	// The gating rule updated.
	GatingRule *types.GatingRule

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationUpdateSafetyRuleMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpUpdateSafetyRule{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpUpdateSafetyRule{}, middleware.After)
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
	if err = addOpUpdateSafetyRuleValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateSafetyRule(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opUpdateSafetyRule(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "route53-recovery-control-config",
		OperationName: "UpdateSafetyRule",
	}
}