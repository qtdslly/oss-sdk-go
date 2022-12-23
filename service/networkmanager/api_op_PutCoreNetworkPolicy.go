// Code generated by smithy-go-codegen DO NOT EDIT.

package networkmanager

import (
	"context"
	"fmt"
	awsmiddleware "oss-sdk-go/oss/middleware"
	"oss-sdk-go/oss/signer/v4"
	"oss-sdk-go/service/networkmanager/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates a new, immutable version of a core network policy. A subsequent change
// set is created showing the differences between the LIVE policy and the submitted
// policy.
func (c *Client) PutCoreNetworkPolicy(ctx context.Context, params *PutCoreNetworkPolicyInput, optFns ...func(*Options)) (*PutCoreNetworkPolicyOutput, error) {
	if params == nil {
		params = &PutCoreNetworkPolicyInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "PutCoreNetworkPolicy", params, optFns, c.addOperationPutCoreNetworkPolicyMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*PutCoreNetworkPolicyOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type PutCoreNetworkPolicyInput struct {

	// The ID of a core network.
	//
	// This member is required.
	CoreNetworkId *string

	// The policy document.
	//
	// This value conforms to the media type: application/json
	//
	// This member is required.
	PolicyDocument *string

	// The client token associated with the request.
	ClientToken *string

	// a core network policy description.
	Description *string

	// The ID of a core network policy.
	LatestVersionId *int32

	noSmithyDocumentSerde
}

type PutCoreNetworkPolicyOutput struct {

	// Describes the changed core network policy.
	CoreNetworkPolicy *types.CoreNetworkPolicy

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationPutCoreNetworkPolicyMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpPutCoreNetworkPolicy{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpPutCoreNetworkPolicy{}, middleware.After)
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
	if err = addIdempotencyToken_opPutCoreNetworkPolicyMiddleware(stack, options); err != nil {
		return err
	}
	if err = addOpPutCoreNetworkPolicyValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opPutCoreNetworkPolicy(options.Region), middleware.Before); err != nil {
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

type idempotencyToken_initializeOpPutCoreNetworkPolicy struct {
	tokenProvider IdempotencyTokenProvider
}

func (*idempotencyToken_initializeOpPutCoreNetworkPolicy) ID() string {
	return "OperationIdempotencyTokenAutoFill"
}

func (m *idempotencyToken_initializeOpPutCoreNetworkPolicy) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	if m.tokenProvider == nil {
		return next.HandleInitialize(ctx, in)
	}

	input, ok := in.Parameters.(*PutCoreNetworkPolicyInput)
	if !ok {
		return out, metadata, fmt.Errorf("expected middleware input to be of type *PutCoreNetworkPolicyInput ")
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
func addIdempotencyToken_opPutCoreNetworkPolicyMiddleware(stack *middleware.Stack, cfg Options) error {
	return stack.Initialize.Add(&idempotencyToken_initializeOpPutCoreNetworkPolicy{tokenProvider: cfg.IdempotencyTokenProvider}, middleware.Before)
}

func newServiceMetadataMiddleware_opPutCoreNetworkPolicy(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "networkmanager",
		OperationName: "PutCoreNetworkPolicy",
	}
}