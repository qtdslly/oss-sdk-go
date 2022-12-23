// Code generated by smithy-go-codegen DO NOT EDIT.

package detective

import (
	"context"
	awsmiddleware "oss-sdk-go/oss/middleware"
	"oss-sdk-go/oss/signer/v4"
	"oss-sdk-go/service/detective/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Gets data source package information for the behavior graph.
func (c *Client) BatchGetGraphMemberDatasources(ctx context.Context, params *BatchGetGraphMemberDatasourcesInput, optFns ...func(*Options)) (*BatchGetGraphMemberDatasourcesOutput, error) {
	if params == nil {
		params = &BatchGetGraphMemberDatasourcesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "BatchGetGraphMemberDatasources", params, optFns, c.addOperationBatchGetGraphMemberDatasourcesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*BatchGetGraphMemberDatasourcesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type BatchGetGraphMemberDatasourcesInput struct {

	// The list of Amazon Web Services accounts to get data source package information
	// on.
	//
	// This member is required.
	AccountIds []string

	// The ARN of the behavior graph.
	//
	// This member is required.
	GraphArn *string

	noSmithyDocumentSerde
}

type BatchGetGraphMemberDatasourcesOutput struct {

	// Details on the status of data source packages for members of the behavior graph.
	MemberDatasources []types.MembershipDatasources

	// Accounts that data source package information could not be retrieved for.
	UnprocessedAccounts []types.UnprocessedAccount

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationBatchGetGraphMemberDatasourcesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpBatchGetGraphMemberDatasources{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpBatchGetGraphMemberDatasources{}, middleware.After)
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
	if err = addOpBatchGetGraphMemberDatasourcesValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opBatchGetGraphMemberDatasources(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opBatchGetGraphMemberDatasources(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "detective",
		OperationName: "BatchGetGraphMemberDatasources",
	}
}
