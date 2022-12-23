// Code generated by smithy-go-codegen DO NOT EDIT.

package redshiftserverless

import (
	"context"
	awsmiddleware "oss-sdk-go/oss/middleware"
	"oss-sdk-go/oss/signer/v4"
	"oss-sdk-go/service/redshiftserverless/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates a namespace in Amazon Redshift Serverless.
func (c *Client) CreateNamespace(ctx context.Context, params *CreateNamespaceInput, optFns ...func(*Options)) (*CreateNamespaceOutput, error) {
	if params == nil {
		params = &CreateNamespaceInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateNamespace", params, optFns, c.addOperationCreateNamespaceMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateNamespaceOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateNamespaceInput struct {

	// The name of the namespace.
	//
	// This member is required.
	NamespaceName *string

	// The password of the administrator for the first database created in the
	// namespace.
	AdminUserPassword *string

	// The username of the administrator for the first database created in the
	// namespace.
	AdminUsername *string

	// The name of the first database created in the namespace.
	DbName *string

	// The Amazon Resource Name (ARN) of the IAM role to set as a default in the
	// namespace.
	DefaultIamRoleArn *string

	// A list of IAM roles to associate with the namespace.
	IamRoles []string

	// The ID of the Amazon Web Services Key Management Service key used to encrypt
	// your data.
	KmsKeyId *string

	// The types of logs the namespace can export. Available export types are userlog,
	// connectionlog, and useractivitylog.
	LogExports []types.LogExport

	// A list of tag instances.
	Tags []types.Tag

	noSmithyDocumentSerde
}

type CreateNamespaceOutput struct {

	// The created namespace object.
	Namespace *types.Namespace

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateNamespaceMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpCreateNamespace{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpCreateNamespace{}, middleware.After)
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
	if err = addOpCreateNamespaceValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateNamespace(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateNamespace(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "redshift-serverless",
		OperationName: "CreateNamespace",
	}
}
