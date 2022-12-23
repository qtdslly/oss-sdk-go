// Code generated by smithy-go-codegen DO NOT EDIT.

package codecommit

import (
	"context"
	awsmiddleware "oss-sdk-go/oss/middleware"
	"oss-sdk-go/oss/signer/v4"
	"oss-sdk-go/service/codecommit/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Removes the association between an approval rule template and one or more
// specified repositories.
func (c *Client) BatchDisassociateApprovalRuleTemplateFromRepositories(ctx context.Context, params *BatchDisassociateApprovalRuleTemplateFromRepositoriesInput, optFns ...func(*Options)) (*BatchDisassociateApprovalRuleTemplateFromRepositoriesOutput, error) {
	if params == nil {
		params = &BatchDisassociateApprovalRuleTemplateFromRepositoriesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "BatchDisassociateApprovalRuleTemplateFromRepositories", params, optFns, c.addOperationBatchDisassociateApprovalRuleTemplateFromRepositoriesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*BatchDisassociateApprovalRuleTemplateFromRepositoriesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type BatchDisassociateApprovalRuleTemplateFromRepositoriesInput struct {

	// The name of the template that you want to disassociate from one or more
	// repositories.
	//
	// This member is required.
	ApprovalRuleTemplateName *string

	// The repository names that you want to disassociate from the approval rule
	// template. The length constraint limit is for each string in the array. The array
	// itself can be empty.
	//
	// This member is required.
	RepositoryNames []string

	noSmithyDocumentSerde
}

type BatchDisassociateApprovalRuleTemplateFromRepositoriesOutput struct {

	// A list of repository names that have had their association with the template
	// removed.
	//
	// This member is required.
	DisassociatedRepositoryNames []string

	// A list of any errors that might have occurred while attempting to remove the
	// association between the template and the repositories.
	//
	// This member is required.
	Errors []types.BatchDisassociateApprovalRuleTemplateFromRepositoriesError

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationBatchDisassociateApprovalRuleTemplateFromRepositoriesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpBatchDisassociateApprovalRuleTemplateFromRepositories{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpBatchDisassociateApprovalRuleTemplateFromRepositories{}, middleware.After)
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
	if err = addOpBatchDisassociateApprovalRuleTemplateFromRepositoriesValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opBatchDisassociateApprovalRuleTemplateFromRepositories(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opBatchDisassociateApprovalRuleTemplateFromRepositories(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "codecommit",
		OperationName: "BatchDisassociateApprovalRuleTemplateFromRepositories",
	}
}
