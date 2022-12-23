// Code generated by smithy-go-codegen DO NOT EDIT.

package alexaforbusiness

import (
	"context"
	awsmiddleware "oss-sdk-go/oss/middleware"
	"oss-sdk-go/oss/signer/v4"
	"oss-sdk-go/service/alexaforbusiness/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Updates an existing conference provider's settings.
func (c *Client) UpdateConferenceProvider(ctx context.Context, params *UpdateConferenceProviderInput, optFns ...func(*Options)) (*UpdateConferenceProviderOutput, error) {
	if params == nil {
		params = &UpdateConferenceProviderInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateConferenceProvider", params, optFns, c.addOperationUpdateConferenceProviderMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateConferenceProviderOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateConferenceProviderInput struct {

	// The ARN of the conference provider.
	//
	// This member is required.
	ConferenceProviderArn *string

	// The type of the conference provider.
	//
	// This member is required.
	ConferenceProviderType types.ConferenceProviderType

	// The meeting settings for the conference provider.
	//
	// This member is required.
	MeetingSetting *types.MeetingSetting

	// The IP endpoint and protocol for calling.
	IPDialIn *types.IPDialIn

	// The information for PSTN conferencing.
	PSTNDialIn *types.PSTNDialIn

	noSmithyDocumentSerde
}

type UpdateConferenceProviderOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationUpdateConferenceProviderMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpUpdateConferenceProvider{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpUpdateConferenceProvider{}, middleware.After)
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
	if err = addOpUpdateConferenceProviderValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateConferenceProvider(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opUpdateConferenceProvider(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "a4b",
		OperationName: "UpdateConferenceProvider",
	}
}
