// Code generated by smithy-go-codegen DO NOT EDIT.

package cognitoidentityprovider

import (
	"context"
	awsmiddleware "oss-sdk-go/oss/middleware"
	"oss-sdk-go/oss/signer/v4"
	"oss-sdk-go/service/cognitoidentityprovider/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// This action is no longer supported. You can use it to configure only SMS MFA.
// You can't use it to configure time-based one-time password (TOTP) software token
// MFA. To configure either type of MFA, use AdminSetUserMFAPreference
// (https://docs.aws.amazon.com/cognito-user-identity-pools/latest/APIReference/API_AdminSetUserMFAPreference.html)
// instead.
func (c *Client) AdminSetUserSettings(ctx context.Context, params *AdminSetUserSettingsInput, optFns ...func(*Options)) (*AdminSetUserSettingsOutput, error) {
	if params == nil {
		params = &AdminSetUserSettingsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "AdminSetUserSettings", params, optFns, c.addOperationAdminSetUserSettingsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*AdminSetUserSettingsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// You can use this parameter to set an MFA configuration that uses the SMS
// delivery medium.
type AdminSetUserSettingsInput struct {

	// You can use this parameter only to set an SMS configuration that uses SMS for
	// delivery.
	//
	// This member is required.
	MFAOptions []types.MFAOptionType

	// The ID of the user pool that contains the user whose options you're setting.
	//
	// This member is required.
	UserPoolId *string

	// The user name of the user whose options you're setting.
	//
	// This member is required.
	Username *string

	noSmithyDocumentSerde
}

// Represents the response from the server to set user settings as an
// administrator.
type AdminSetUserSettingsOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationAdminSetUserSettingsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpAdminSetUserSettings{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpAdminSetUserSettings{}, middleware.After)
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
	if err = addOpAdminSetUserSettingsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opAdminSetUserSettings(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opAdminSetUserSettings(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "cognito-idp",
		OperationName: "AdminSetUserSettings",
	}
}
