// Code generated by smithy-go-codegen DO NOT EDIT.

package ses

import (
	"context"
	awsmiddleware "oss-sdk-go/oss/middleware"
	"oss-sdk-go/oss/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Enables or disables email sending for messages sent using a specific
// configuration set in a given AWS Region. You can use this operation in
// conjunction with Amazon CloudWatch alarms to temporarily pause email sending for
// a configuration set when the reputation metrics for that configuration set (such
// as your bounce on complaint rate) exceed certain thresholds. You can execute
// this operation no more than once per second.
func (c *Client) UpdateConfigurationSetSendingEnabled(ctx context.Context, params *UpdateConfigurationSetSendingEnabledInput, optFns ...func(*Options)) (*UpdateConfigurationSetSendingEnabledOutput, error) {
	if params == nil {
		params = &UpdateConfigurationSetSendingEnabledInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateConfigurationSetSendingEnabled", params, optFns, c.addOperationUpdateConfigurationSetSendingEnabledMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateConfigurationSetSendingEnabledOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Represents a request to enable or disable the email sending capabilities for a
// specific configuration set.
type UpdateConfigurationSetSendingEnabledInput struct {

	// The name of the configuration set that you want to update.
	//
	// This member is required.
	ConfigurationSetName *string

	// Describes whether email sending is enabled or disabled for the configuration
	// set.
	//
	// This member is required.
	Enabled bool

	noSmithyDocumentSerde
}

type UpdateConfigurationSetSendingEnabledOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationUpdateConfigurationSetSendingEnabledMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsquery_serializeOpUpdateConfigurationSetSendingEnabled{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpUpdateConfigurationSetSendingEnabled{}, middleware.After)
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
	if err = addOpUpdateConfigurationSetSendingEnabledValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateConfigurationSetSendingEnabled(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opUpdateConfigurationSetSendingEnabled(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ses",
		OperationName: "UpdateConfigurationSetSendingEnabled",
	}
}
