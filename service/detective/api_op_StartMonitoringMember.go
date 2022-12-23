// Code generated by smithy-go-codegen DO NOT EDIT.

package detective

import (
	"context"
	awsmiddleware "oss-sdk-go/oss/middleware"
	"oss-sdk-go/oss/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Sends a request to enable data ingest for a member account that has a status of
// ACCEPTED_BUT_DISABLED. For valid member accounts, the status is updated as
// follows.
//
// * If Detective enabled the member account, then the new status is
// ENABLED.
//
// * If Detective cannot enable the member account, the status remains
// ACCEPTED_BUT_DISABLED.
func (c *Client) StartMonitoringMember(ctx context.Context, params *StartMonitoringMemberInput, optFns ...func(*Options)) (*StartMonitoringMemberOutput, error) {
	if params == nil {
		params = &StartMonitoringMemberInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "StartMonitoringMember", params, optFns, c.addOperationStartMonitoringMemberMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*StartMonitoringMemberOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type StartMonitoringMemberInput struct {

	// The account ID of the member account to try to enable. The account must be an
	// invited member account with a status of ACCEPTED_BUT_DISABLED.
	//
	// This member is required.
	AccountId *string

	// The ARN of the behavior graph.
	//
	// This member is required.
	GraphArn *string

	noSmithyDocumentSerde
}

type StartMonitoringMemberOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationStartMonitoringMemberMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpStartMonitoringMember{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpStartMonitoringMember{}, middleware.After)
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
	if err = addOpStartMonitoringMemberValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opStartMonitoringMember(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opStartMonitoringMember(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "detective",
		OperationName: "StartMonitoringMember",
	}
}