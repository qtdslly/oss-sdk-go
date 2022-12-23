// Code generated by smithy-go-codegen DO NOT EDIT.

package ebs

import (
	"context"
	"fmt"
	awsmiddleware "oss-sdk-go/oss/middleware"
	"oss-sdk-go/oss/signer/v4"
	"oss-sdk-go/service/ebs/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Creates a new Amazon EBS snapshot. The new snapshot enters the pending state
// after the request completes. After creating the snapshot, use  PutSnapshotBlock
// (https://docs.aws.amazon.com/ebs/latest/APIReference/API_PutSnapshotBlock.html)
// to write blocks of data to the snapshot.
func (c *Client) StartSnapshot(ctx context.Context, params *StartSnapshotInput, optFns ...func(*Options)) (*StartSnapshotOutput, error) {
	if params == nil {
		params = &StartSnapshotInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "StartSnapshot", params, optFns, c.addOperationStartSnapshotMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*StartSnapshotOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type StartSnapshotInput struct {

	// The size of the volume, in GiB. The maximum size is 65536 GiB (64 TiB).
	//
	// This member is required.
	VolumeSize *int64

	// A unique, case-sensitive identifier that you provide to ensure the idempotency
	// of the request. Idempotency ensures that an API request completes only once.
	// With an idempotent request, if the original request completes successfully. The
	// subsequent retries with the same client token return the result from the
	// original successful request and they have no additional effect. If you do not
	// specify a client token, one is automatically generated by the Amazon Web
	// Services SDK. For more information, see  Idempotency for StartSnapshot API
	// (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/ebs-direct-api-idempotency.html)
	// in the Amazon Elastic Compute Cloud User Guide.
	ClientToken *string

	// A description for the snapshot.
	Description *string

	// Indicates whether to encrypt the snapshot. You can't specify Encrypted and
	// ParentSnapshotId in the same request. If you specify both parameters, the
	// request fails with ValidationException. The encryption status of the snapshot
	// depends on the values that you specify for Encrypted, KmsKeyArn, and
	// ParentSnapshotId, and whether your Amazon Web Services account is enabled for
	// encryption by default
	// (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/EBSEncryption.html#encryption-by-default).
	// For more information, see  Using encryption
	// (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/ebsapis-using-encryption.html)
	// in the Amazon Elastic Compute Cloud User Guide. To create an encrypted snapshot,
	// you must have permission to use the KMS key. For more information, see
	// Permissions to use Key Management Service keys
	// (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/ebsapi-permissions.html#ebsapi-kms-permissions)
	// in the Amazon Elastic Compute Cloud User Guide.
	Encrypted *bool

	// The Amazon Resource Name (ARN) of the Key Management Service (KMS) key to be
	// used to encrypt the snapshot. The encryption status of the snapshot depends on
	// the values that you specify for Encrypted, KmsKeyArn, and ParentSnapshotId, and
	// whether your Amazon Web Services account is enabled for  encryption by default
	// (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/EBSEncryption.html#encryption-by-default).
	// For more information, see  Using encryption
	// (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/ebsapis-using-encryption.html)
	// in the Amazon Elastic Compute Cloud User Guide. To create an encrypted snapshot,
	// you must have permission to use the KMS key. For more information, see
	// Permissions to use Key Management Service keys
	// (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/ebsapi-permissions.html#ebsapi-kms-permissions)
	// in the Amazon Elastic Compute Cloud User Guide.
	KmsKeyArn *string

	// The ID of the parent snapshot. If there is no parent snapshot, or if you are
	// creating the first snapshot for an on-premises volume, omit this parameter. You
	// can't specify ParentSnapshotId and Encrypted in the same request. If you specify
	// both parameters, the request fails with ValidationException. The encryption
	// status of the snapshot depends on the values that you specify for Encrypted,
	// KmsKeyArn, and ParentSnapshotId, and whether your Amazon Web Services account is
	// enabled for  encryption by default
	// (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/EBSEncryption.html#encryption-by-default).
	// For more information, see  Using encryption
	// (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/ebsapis-using-encryption.html)
	// in the Amazon Elastic Compute Cloud User Guide. If you specify an encrypted
	// parent snapshot, you must have permission to use the KMS key that was used to
	// encrypt the parent snapshot. For more information, see  Permissions to use Key
	// Management Service keys
	// (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/ebsapi-permissions.html#ebsapi-kms-permissions)
	// in the Amazon Elastic Compute Cloud User Guide.
	ParentSnapshotId *string

	// The tags to apply to the snapshot.
	Tags []types.Tag

	// The amount of time (in minutes) after which the snapshot is automatically
	// cancelled if:
	//
	// * No blocks are written to the snapshot.
	//
	// * The snapshot is not
	// completed after writing the last block of data.
	//
	// If no value is specified, the
	// timeout defaults to 60 minutes.
	Timeout *int32

	noSmithyDocumentSerde
}

type StartSnapshotOutput struct {

	// The size of the blocks in the snapshot, in bytes.
	BlockSize *int32

	// The description of the snapshot.
	Description *string

	// The Amazon Resource Name (ARN) of the Key Management Service (KMS) key used to
	// encrypt the snapshot.
	KmsKeyArn *string

	// The Amazon Web Services account ID of the snapshot owner.
	OwnerId *string

	// The ID of the parent snapshot.
	ParentSnapshotId *string

	// The ID of the snapshot.
	SnapshotId *string

	// The timestamp when the snapshot was created.
	StartTime *time.Time

	// The status of the snapshot.
	Status types.Status

	// The tags applied to the snapshot. You can specify up to 50 tags per snapshot.
	// For more information, see  Tagging your Amazon EC2 resources
	// (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/Using_Tags.html) in the
	// Amazon Elastic Compute Cloud User Guide.
	Tags []types.Tag

	// The size of the volume, in GiB.
	VolumeSize *int64

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationStartSnapshotMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpStartSnapshot{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpStartSnapshot{}, middleware.After)
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
	if err = addIdempotencyToken_opStartSnapshotMiddleware(stack, options); err != nil {
		return err
	}
	if err = addOpStartSnapshotValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opStartSnapshot(options.Region), middleware.Before); err != nil {
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

type idempotencyToken_initializeOpStartSnapshot struct {
	tokenProvider IdempotencyTokenProvider
}

func (*idempotencyToken_initializeOpStartSnapshot) ID() string {
	return "OperationIdempotencyTokenAutoFill"
}

func (m *idempotencyToken_initializeOpStartSnapshot) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	if m.tokenProvider == nil {
		return next.HandleInitialize(ctx, in)
	}

	input, ok := in.Parameters.(*StartSnapshotInput)
	if !ok {
		return out, metadata, fmt.Errorf("expected middleware input to be of type *StartSnapshotInput ")
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
func addIdempotencyToken_opStartSnapshotMiddleware(stack *middleware.Stack, cfg Options) error {
	return stack.Initialize.Add(&idempotencyToken_initializeOpStartSnapshot{tokenProvider: cfg.IdempotencyTokenProvider}, middleware.Before)
}

func newServiceMetadataMiddleware_opStartSnapshot(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ebs",
		OperationName: "StartSnapshot",
	}
}