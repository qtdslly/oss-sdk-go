// Code generated by smithy-go-codegen DO NOT EDIT.

package fsx

import (
	"context"
	"fmt"
	awsmiddleware "oss-sdk-go/oss/middleware"
	"oss-sdk-go/oss/signer/v4"
	"oss-sdk-go/service/fsx/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates a backup of an existing Amazon FSx for Windows File Server file system,
// Amazon FSx for Lustre file system, Amazon FSx for NetApp ONTAP volume, or Amazon
// FSx for OpenZFS file system. We recommend creating regular backups so that you
// can restore a file system or volume from a backup if an issue arises with the
// original file system or volume. For Amazon FSx for Lustre file systems, you can
// create a backup only for file systems that have the following configuration:
//
// *
// A Persistent deployment type
//
// * Are not linked to a data repository
//
// For more
// information about backups, see the following:
//
// * For Amazon FSx for Lustre, see
// Working with FSx for Lustre backups
// (https://docs.aws.amazon.com/fsx/latest/LustreGuide/using-backups-fsx.html).
//
// *
// For Amazon FSx for Windows, see Working with FSx for Windows backups
// (https://docs.aws.amazon.com/fsx/latest/WindowsGuide/using-backups.html).
//
// * For
// Amazon FSx for NetApp ONTAP, see Working with FSx for NetApp ONTAP backups
// (https://docs.aws.amazon.com/fsx/latest/ONTAPGuide/using-backups.html).
//
// * For
// Amazon FSx for OpenZFS, see Working with FSx for OpenZFS backups
// (https://docs.aws.amazon.com/fsx/latest/OpenZFSGuide/using-backups.html).
//
// If a
// backup with the specified client request token exists and the parameters match,
// this operation returns the description of the existing backup. If a backup with
// the specified client request token exists and the parameters don't match, this
// operation returns IncompatibleParameterError. If a backup with the specified
// client request token doesn't exist, CreateBackup does the following:
//
// * Creates
// a new Amazon FSx backup with an assigned ID, and an initial lifecycle state of
// CREATING.
//
// * Returns the description of the backup.
//
// By using the idempotent
// operation, you can retry a CreateBackup operation without the risk of creating
// an extra backup. This approach can be useful when an initial call fails in a way
// that makes it unclear whether a backup was created. If you use the same client
// request token and the initial call created a backup, the operation returns a
// successful result because all the parameters are the same. The CreateBackup
// operation returns while the backup's lifecycle state is still CREATING. You can
// check the backup creation status by calling the DescribeBackups
// (https://docs.aws.amazon.com/fsx/latest/APIReference/API_DescribeBackups.html)
// operation, which returns the backup state along with other information.
func (c *Client) CreateBackup(ctx context.Context, params *CreateBackupInput, optFns ...func(*Options)) (*CreateBackupOutput, error) {
	if params == nil {
		params = &CreateBackupInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateBackup", params, optFns, c.addOperationCreateBackupMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateBackupOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// The request object for the CreateBackup operation.
type CreateBackupInput struct {

	// (Optional) A string of up to 64 ASCII characters that Amazon FSx uses to ensure
	// idempotent creation. This string is automatically filled on your behalf when you
	// use the Command Line Interface (CLI) or an Amazon Web Services SDK.
	ClientRequestToken *string

	// The ID of the file system to back up.
	FileSystemId *string

	// (Optional) The tags to apply to the backup at backup creation. The key value of
	// the Name tag appears in the console as the backup name. If you have set
	// CopyTagsToBackups to true, and you specify one or more tags using the
	// CreateBackup operation, no existing file system tags are copied from the file
	// system to the backup.
	Tags []types.Tag

	// (Optional) The ID of the FSx for ONTAP volume to back up.
	VolumeId *string

	noSmithyDocumentSerde
}

// The response object for the CreateBackup operation.
type CreateBackupOutput struct {

	// A description of the backup.
	Backup *types.Backup

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateBackupMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpCreateBackup{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpCreateBackup{}, middleware.After)
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
	if err = addIdempotencyToken_opCreateBackupMiddleware(stack, options); err != nil {
		return err
	}
	if err = addOpCreateBackupValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateBackup(options.Region), middleware.Before); err != nil {
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

type idempotencyToken_initializeOpCreateBackup struct {
	tokenProvider IdempotencyTokenProvider
}

func (*idempotencyToken_initializeOpCreateBackup) ID() string {
	return "OperationIdempotencyTokenAutoFill"
}

func (m *idempotencyToken_initializeOpCreateBackup) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	if m.tokenProvider == nil {
		return next.HandleInitialize(ctx, in)
	}

	input, ok := in.Parameters.(*CreateBackupInput)
	if !ok {
		return out, metadata, fmt.Errorf("expected middleware input to be of type *CreateBackupInput ")
	}

	if input.ClientRequestToken == nil {
		t, err := m.tokenProvider.GetIdempotencyToken()
		if err != nil {
			return out, metadata, err
		}
		input.ClientRequestToken = &t
	}
	return next.HandleInitialize(ctx, in)
}
func addIdempotencyToken_opCreateBackupMiddleware(stack *middleware.Stack, cfg Options) error {
	return stack.Initialize.Add(&idempotencyToken_initializeOpCreateBackup{tokenProvider: cfg.IdempotencyTokenProvider}, middleware.Before)
}

func newServiceMetadataMiddleware_opCreateBackup(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "fsx",
		OperationName: "CreateBackup",
	}
}
