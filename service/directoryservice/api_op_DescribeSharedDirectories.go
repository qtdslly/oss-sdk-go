// Code generated by smithy-go-codegen DO NOT EDIT.

package directoryservice

import (
	"context"
	awsmiddleware "oss-sdk-go/oss/middleware"
	"oss-sdk-go/oss/signer/v4"
	"oss-sdk-go/service/directoryservice/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Returns the shared directories in your account.
func (c *Client) DescribeSharedDirectories(ctx context.Context, params *DescribeSharedDirectoriesInput, optFns ...func(*Options)) (*DescribeSharedDirectoriesOutput, error) {
	if params == nil {
		params = &DescribeSharedDirectoriesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeSharedDirectories", params, optFns, c.addOperationDescribeSharedDirectoriesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeSharedDirectoriesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeSharedDirectoriesInput struct {

	// Returns the identifier of the directory in the directory owner account.
	//
	// This member is required.
	OwnerDirectoryId *string

	// The number of shared directories to return in the response object.
	Limit *int32

	// The DescribeSharedDirectoriesResult.NextToken value from a previous call to
	// DescribeSharedDirectories. Pass null if this is the first call.
	NextToken *string

	// A list of identifiers of all shared directories in your account.
	SharedDirectoryIds []string

	noSmithyDocumentSerde
}

type DescribeSharedDirectoriesOutput struct {

	// If not null, token that indicates that more results are available. Pass this
	// value for the NextToken parameter in a subsequent call to
	// DescribeSharedDirectories to retrieve the next set of items.
	NextToken *string

	// A list of all shared directories in your account.
	SharedDirectories []types.SharedDirectory

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeSharedDirectoriesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDescribeSharedDirectories{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDescribeSharedDirectories{}, middleware.After)
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
	if err = addOpDescribeSharedDirectoriesValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeSharedDirectories(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDescribeSharedDirectories(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ds",
		OperationName: "DescribeSharedDirectories",
	}
}
