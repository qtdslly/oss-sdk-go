// Code generated by smithy-go-codegen DO NOT EDIT.

package account

import (
	"context"
	awsmiddleware "oss-sdk-go/oss/middleware"
	"oss-sdk-go/oss/signer/v4"
	"oss-sdk-go/service/account/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Modifies the specified alternate contact attached to an Amazon Web Services
// account. For complete details about how to use the alternate contact operations,
// see Access or updating the alternate contacts
// (https://docs.aws.amazon.com/accounts/latest/reference/manage-acct-update-contact.html).
// Before you can update the alternate contact information for an Amazon Web
// Services account that is managed by Organizations, you must first enable
// integration between Amazon Web Services Account Management and Organizations.
// For more information, see Enabling trusted access for Amazon Web Services
// Account Management
// (https://docs.aws.amazon.com/accounts/latest/reference/using-orgs-trusted-access.html).
func (c *Client) PutAlternateContact(ctx context.Context, params *PutAlternateContactInput, optFns ...func(*Options)) (*PutAlternateContactOutput, error) {
	if params == nil {
		params = &PutAlternateContactInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "PutAlternateContact", params, optFns, c.addOperationPutAlternateContactMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*PutAlternateContactOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type PutAlternateContactInput struct {

	// Specifies which alternate contact you want to create or update.
	//
	// This member is required.
	AlternateContactType types.AlternateContactType

	// Specifies an email address for the alternate contact.
	//
	// This member is required.
	EmailAddress *string

	// Specifies a name for the alternate contact.
	//
	// This member is required.
	Name *string

	// Specifies a phone number for the alternate contact.
	//
	// This member is required.
	PhoneNumber *string

	// Specifies a title for the alternate contact.
	//
	// This member is required.
	Title *string

	// Specifies the 12 digit account ID number of the Amazon Web Services account that
	// you want to access or modify with this operation. If you do not specify this
	// parameter, it defaults to the Amazon Web Services account of the identity used
	// to call the operation. To use this parameter, the caller must be an identity in
	// the organization's management account
	// (https://docs.aws.amazon.com/organizations/latest/userguide/orgs_getting-started_concepts.html#account)
	// or a delegated administrator account, and the specified account ID must be a
	// member account in the same organization. The organization must have all features
	// enabled
	// (https://docs.aws.amazon.com/organizations/latest/userguide/orgs_manage_org_support-all-features.html),
	// and the organization must have trusted access
	// (https://docs.aws.amazon.com/organizations/latest/userguide/using-orgs-trusted-access.html)
	// enabled for the Account Management service, and optionally a delegated admin
	// (https://docs.aws.amazon.com/organizations/latest/userguide/using-orgs-delegated-admin.html)
	// account assigned. The management account can't specify its own AccountId; it
	// must call the operation in standalone context by not including the AccountId
	// parameter. To call this operation on an account that is not a member of an
	// organization, then don't specify this parameter, and call the operation using an
	// identity belonging to the account whose contacts you wish to retrieve or modify.
	AccountId *string

	noSmithyDocumentSerde
}

type PutAlternateContactOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationPutAlternateContactMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpPutAlternateContact{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpPutAlternateContact{}, middleware.After)
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
	if err = addOpPutAlternateContactValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opPutAlternateContact(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opPutAlternateContact(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "account",
		OperationName: "PutAlternateContact",
	}
}
