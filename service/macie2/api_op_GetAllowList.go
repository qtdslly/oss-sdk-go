// Code generated by smithy-go-codegen DO NOT EDIT.

package macie2

import (
	"context"
	awsmiddleware "oss-sdk-go/oss/middleware"
	"oss-sdk-go/oss/signer/v4"
	"oss-sdk-go/service/macie2/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Retrieves the settings and status of an allow list.
func (c *Client) GetAllowList(ctx context.Context, params *GetAllowListInput, optFns ...func(*Options)) (*GetAllowListOutput, error) {
	if params == nil {
		params = &GetAllowListInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetAllowList", params, optFns, c.addOperationGetAllowListMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetAllowListOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetAllowListInput struct {

	// The unique identifier for the Amazon Macie resource that the request applies to.
	//
	// This member is required.
	Id *string

	noSmithyDocumentSerde
}

type GetAllowListOutput struct {

	// The Amazon Resource Name (ARN) of the allow list.
	Arn *string

	// The date and time, in UTC and extended ISO 8601 format, when the allow list was
	// created in Amazon Macie.
	CreatedAt *time.Time

	// The criteria that specify the text or text pattern to ignore. The criteria can
	// be the location and name of an S3 object that lists specific text to ignore
	// (s3WordsList), or a regular expression (regex) that defines a text pattern to
	// ignore.
	Criteria *types.AllowListCriteria

	// The custom description of the allow list.
	Description *string

	// The unique identifier for the allow list.
	Id *string

	// The custom name of the allow list.
	Name *string

	// The current status of the allow list, which indicates whether Amazon Macie can
	// access and use the list's criteria.
	Status *types.AllowListStatus

	// A map of key-value pairs that specifies which tags (keys and values) are
	// associated with the allow list.
	Tags map[string]string

	// The date and time, in UTC and extended ISO 8601 format, when the allow list's
	// settings were most recently changed in Amazon Macie.
	UpdatedAt *time.Time

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetAllowListMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpGetAllowList{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpGetAllowList{}, middleware.After)
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
	if err = addOpGetAllowListValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetAllowList(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opGetAllowList(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "macie2",
		OperationName: "GetAllowList",
	}
}
