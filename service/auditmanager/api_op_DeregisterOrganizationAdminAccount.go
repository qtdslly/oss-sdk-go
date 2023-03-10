// Code generated by smithy-go-codegen DO NOT EDIT.

package auditmanager

import (
	"context"
	awsmiddleware "oss-sdk-go/oss/middleware"
	"oss-sdk-go/oss/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Removes the specified Amazon Web Services account as a delegated administrator
// for Audit Manager. When you remove a delegated administrator from your Audit
// Manager settings, you continue to have access to the evidence that you
// previously collected under that account. This is also the case when you
// deregister a delegated administrator from Organizations. However, Audit Manager
// will stop collecting and attaching evidence to that delegated administrator
// account moving forward. When you deregister a delegated administrator account
// for Audit Manager, the data for that account isn’t deleted. If you want to
// delete resource data for a delegated administrator account, you must perform
// that task separately before you deregister the account. Either, you can do this
// in the Audit Manager console. Or, you can use one of the delete API operations
// that are provided by Audit Manager. To delete your Audit Manager resource data,
// see the following instructions:
//
// * DeleteAssessment
// (https://docs.aws.amazon.com/audit-manager/latest/APIReference/API_DeleteAssessment.html)
// (see also: Deleting an assessment
// (https://docs.aws.amazon.com/audit-manager/latest/userguide/delete-assessment.html)
// in the Audit Manager User Guide)
//
// * DeleteAssessmentFramework
// (https://docs.aws.amazon.com/audit-manager/latest/APIReference/API_DeleteAssessmentFramework.html)
// (see also: Deleting a custom framework
// (https://docs.aws.amazon.com/audit-manager/latest/userguide/delete-custom-framework.html)
// in the Audit Manager User Guide)
//
// * DeleteAssessmentFrameworkShare
// (https://docs.aws.amazon.com/audit-manager/latest/APIReference/API_DeleteAssessmentFrameworkShare.html)
// (see also: Deleting a share request
// (https://docs.aws.amazon.com/audit-manager/latest/userguide/deleting-shared-framework-requests.html)
// in the Audit Manager User Guide)
//
// * DeleteAssessmentReport
// (https://docs.aws.amazon.com/audit-manager/latest/APIReference/API_DeleteAssessmentReport.html)
// (see also: Deleting an assessment report
// (https://docs.aws.amazon.com/audit-manager/latest/userguide/generate-assessment-report.html#delete-assessment-report-steps)
// in the Audit Manager User Guide)
//
// * DeleteControl
// (https://docs.aws.amazon.com/audit-manager/latest/APIReference/API_DeleteControl.html)
// (see also: Deleting a custom control
// (https://docs.aws.amazon.com/audit-manager/latest/userguide/delete-controls.html)
// in the Audit Manager User Guide)
//
// At this time, Audit Manager doesn't provide an
// option to delete evidence. All available delete operations are listed above.
func (c *Client) DeregisterOrganizationAdminAccount(ctx context.Context, params *DeregisterOrganizationAdminAccountInput, optFns ...func(*Options)) (*DeregisterOrganizationAdminAccountOutput, error) {
	if params == nil {
		params = &DeregisterOrganizationAdminAccountInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DeregisterOrganizationAdminAccount", params, optFns, c.addOperationDeregisterOrganizationAdminAccountMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DeregisterOrganizationAdminAccountOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DeregisterOrganizationAdminAccountInput struct {

	// The identifier for the administrator account.
	AdminAccountId *string

	noSmithyDocumentSerde
}

type DeregisterOrganizationAdminAccountOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDeregisterOrganizationAdminAccountMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpDeregisterOrganizationAdminAccount{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpDeregisterOrganizationAdminAccount{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDeregisterOrganizationAdminAccount(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDeregisterOrganizationAdminAccount(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "auditmanager",
		OperationName: "DeregisterOrganizationAdminAccount",
	}
}
