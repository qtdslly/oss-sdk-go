// Code generated by smithy-go-codegen DO NOT EDIT.

package kinesisanalyticsv2

import (
	"context"
	awsmiddleware "oss-sdk-go/oss/middleware"
	"oss-sdk-go/oss/signer/v4"
	"oss-sdk-go/service/kinesisanalyticsv2/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Updates an existing Kinesis Data Analytics application. Using this operation,
// you can update application code, input configuration, and output configuration.
// Kinesis Data Analytics updates the ApplicationVersionId each time you update
// your application. You cannot update the RuntimeEnvironment of an existing
// application. If you need to update an application's RuntimeEnvironment, you must
// delete the application and create it again.
func (c *Client) UpdateApplication(ctx context.Context, params *UpdateApplicationInput, optFns ...func(*Options)) (*UpdateApplicationOutput, error) {
	if params == nil {
		params = &UpdateApplicationInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateApplication", params, optFns, c.addOperationUpdateApplicationMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateApplicationOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateApplicationInput struct {

	// The name of the application to update.
	//
	// This member is required.
	ApplicationName *string

	// Describes application configuration updates.
	ApplicationConfigurationUpdate *types.ApplicationConfigurationUpdate

	// Describes application Amazon CloudWatch logging option updates. You can only
	// update existing CloudWatch logging options with this action. To add a new
	// CloudWatch logging option, use AddApplicationCloudWatchLoggingOption.
	CloudWatchLoggingOptionUpdates []types.CloudWatchLoggingOptionUpdate

	// A value you use to implement strong concurrency for application updates. You
	// must provide the CurrentApplicationVersionId or the ConditionalToken. You get
	// the application's current ConditionalToken using DescribeApplication. For better
	// concurrency support, use the ConditionalToken parameter instead of
	// CurrentApplicationVersionId.
	ConditionalToken *string

	// The current application version ID. You must provide the
	// CurrentApplicationVersionId or the ConditionalToken.You can retrieve the
	// application version ID using DescribeApplication. For better concurrency
	// support, use the ConditionalToken parameter instead of
	// CurrentApplicationVersionId.
	CurrentApplicationVersionId *int64

	// Describes updates to the application's starting parameters.
	RunConfigurationUpdate *types.RunConfigurationUpdate

	// Describes updates to the service execution role.
	ServiceExecutionRoleUpdate *string

	noSmithyDocumentSerde
}

type UpdateApplicationOutput struct {

	// Describes application updates.
	//
	// This member is required.
	ApplicationDetail *types.ApplicationDetail

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationUpdateApplicationMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpUpdateApplication{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpUpdateApplication{}, middleware.After)
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
	if err = addOpUpdateApplicationValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateApplication(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opUpdateApplication(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "kinesisanalytics",
		OperationName: "UpdateApplication",
	}
}