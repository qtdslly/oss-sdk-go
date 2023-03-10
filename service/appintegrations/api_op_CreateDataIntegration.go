// Code generated by smithy-go-codegen DO NOT EDIT.

package appintegrations

import (
	"context"
	"fmt"
	awsmiddleware "oss-sdk-go/oss/middleware"
	"oss-sdk-go/oss/signer/v4"
	"oss-sdk-go/service/appintegrations/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates and persists a DataIntegration resource. You cannot create a
// DataIntegration association for a DataIntegration that has been previously
// associated. Use a different DataIntegration, or recreate the DataIntegration
// using the CreateDataIntegration API.
func (c *Client) CreateDataIntegration(ctx context.Context, params *CreateDataIntegrationInput, optFns ...func(*Options)) (*CreateDataIntegrationOutput, error) {
	if params == nil {
		params = &CreateDataIntegrationInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateDataIntegration", params, optFns, c.addOperationCreateDataIntegrationMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateDataIntegrationOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateDataIntegrationInput struct {

	// The name of the DataIntegration.
	//
	// This member is required.
	Name *string

	// A unique, case-sensitive identifier that you provide to ensure the idempotency
	// of the request.
	ClientToken *string

	// A description of the DataIntegration.
	Description *string

	// The KMS key for the DataIntegration.
	KmsKey *string

	// The name of the data and how often it should be pulled from the source.
	ScheduleConfig *types.ScheduleConfiguration

	// The URI of the data source.
	SourceURI *string

	// One or more tags.
	Tags map[string]string

	noSmithyDocumentSerde
}

type CreateDataIntegrationOutput struct {

	// The Amazon Resource Name (ARN)
	Arn *string

	// A unique, case-sensitive identifier that you provide to ensure the idempotency
	// of the request.
	ClientToken *string

	// A description of the DataIntegration.
	Description *string

	// A unique identifier.
	Id *string

	// The KMS key for the DataIntegration.
	KmsKey *string

	// The name of the DataIntegration.
	Name *string

	// The name of the data and how often it should be pulled from the source.
	ScheduleConfiguration *types.ScheduleConfiguration

	// The URI of the data source.
	SourceURI *string

	// One or more tags.
	Tags map[string]string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateDataIntegrationMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpCreateDataIntegration{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpCreateDataIntegration{}, middleware.After)
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
	if err = addIdempotencyToken_opCreateDataIntegrationMiddleware(stack, options); err != nil {
		return err
	}
	if err = addOpCreateDataIntegrationValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateDataIntegration(options.Region), middleware.Before); err != nil {
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

type idempotencyToken_initializeOpCreateDataIntegration struct {
	tokenProvider IdempotencyTokenProvider
}

func (*idempotencyToken_initializeOpCreateDataIntegration) ID() string {
	return "OperationIdempotencyTokenAutoFill"
}

func (m *idempotencyToken_initializeOpCreateDataIntegration) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	if m.tokenProvider == nil {
		return next.HandleInitialize(ctx, in)
	}

	input, ok := in.Parameters.(*CreateDataIntegrationInput)
	if !ok {
		return out, metadata, fmt.Errorf("expected middleware input to be of type *CreateDataIntegrationInput ")
	}

	if input.ClientToken == nil {
		t, err := m.tokenProvider.GetIdempotencyToken()
		if err != nil {
			return out, metadata, err
		}
		input.ClientToken = &t
	}
	return next.HandleInitialize(ctx, in)
}
func addIdempotencyToken_opCreateDataIntegrationMiddleware(stack *middleware.Stack, cfg Options) error {
	return stack.Initialize.Add(&idempotencyToken_initializeOpCreateDataIntegration{tokenProvider: cfg.IdempotencyTokenProvider}, middleware.Before)
}

func newServiceMetadataMiddleware_opCreateDataIntegration(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "app-integrations",
		OperationName: "CreateDataIntegration",
	}
}
