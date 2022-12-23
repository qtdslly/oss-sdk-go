// Code generated by smithy-go-codegen DO NOT EDIT.

package workmail

import (
	"context"
	awsmiddleware "oss-sdk-go/oss/middleware"
	"oss-sdk-go/oss/signer/v4"
	"oss-sdk-go/service/workmail/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Performs a test on an availability provider to ensure that access is allowed.
// For EWS, it verifies the provided credentials can be used to successfully log
// in. For Lambda, it verifies that the Lambda function can be invoked and that the
// resource access policy was configured to deny anonymous access. An anonymous
// invocation is one done without providing either a SourceArn or SourceAccount
// header. The request must contain either one provider definition (EwsProvider or
// LambdaProvider) or the DomainName parameter. If the DomainName parameter is
// provided, the configuration stored under the DomainName will be tested.
func (c *Client) TestAvailabilityConfiguration(ctx context.Context, params *TestAvailabilityConfigurationInput, optFns ...func(*Options)) (*TestAvailabilityConfigurationOutput, error) {
	if params == nil {
		params = &TestAvailabilityConfigurationInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "TestAvailabilityConfiguration", params, optFns, c.addOperationTestAvailabilityConfigurationMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*TestAvailabilityConfigurationOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type TestAvailabilityConfigurationInput struct {

	// The Amazon WorkMail organization where the availability provider will be tested.
	//
	// This member is required.
	OrganizationId *string

	// The domain to which the provider applies. If this field is provided, a stored
	// availability provider associated to this domain name will be tested.
	DomainName *string

	// Describes an EWS based availability provider. This is only used as input to the
	// service.
	EwsProvider *types.EwsAvailabilityProvider

	// Describes a Lambda based availability provider.
	LambdaProvider *types.LambdaAvailabilityProvider

	noSmithyDocumentSerde
}

type TestAvailabilityConfigurationOutput struct {

	// String containing the reason for a failed test if TestPassed is false.
	FailureReason *string

	// Boolean indicating whether the test passed or failed.
	TestPassed bool

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationTestAvailabilityConfigurationMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpTestAvailabilityConfiguration{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpTestAvailabilityConfiguration{}, middleware.After)
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
	if err = addOpTestAvailabilityConfigurationValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opTestAvailabilityConfiguration(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opTestAvailabilityConfiguration(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "workmail",
		OperationName: "TestAvailabilityConfiguration",
	}
}
