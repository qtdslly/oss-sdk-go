// Code generated by smithy-go-codegen DO NOT EDIT.

package customerprofiles

import (
	"context"
	awsmiddleware "oss-sdk-go/oss/middleware"
	"oss-sdk-go/oss/signer/v4"
	"oss-sdk-go/service/customerprofiles/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Returns information about a specific domain.
func (c *Client) GetDomain(ctx context.Context, params *GetDomainInput, optFns ...func(*Options)) (*GetDomainOutput, error) {
	if params == nil {
		params = &GetDomainInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetDomain", params, optFns, c.addOperationGetDomainMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetDomainOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetDomainInput struct {

	// The unique name of the domain.
	//
	// This member is required.
	DomainName *string

	noSmithyDocumentSerde
}

type GetDomainOutput struct {

	// The timestamp of when the domain was created.
	//
	// This member is required.
	CreatedAt *time.Time

	// The unique name of the domain.
	//
	// This member is required.
	DomainName *string

	// The timestamp of when the domain was most recently edited.
	//
	// This member is required.
	LastUpdatedAt *time.Time

	// The URL of the SQS dead letter queue, which is used for reporting errors
	// associated with ingesting data from third party applications.
	DeadLetterQueueUrl *string

	// The default encryption key, which is an AWS managed key, is used when no
	// specific type of encryption key is specified. It is used to encrypt all data
	// before it is placed in permanent or semi-permanent storage.
	DefaultEncryptionKey *string

	// The default number of days until the data within the domain expires.
	DefaultExpirationDays *int32

	// The process of matching duplicate profiles. If Matching = true, Amazon Connect
	// Customer Profiles starts a weekly batch process called Identity Resolution Job.
	// If you do not specify a date and time for Identity Resolution Job to run, by
	// default it runs every Saturday at 12AM UTC to detect duplicate profiles in your
	// domains. After the Identity Resolution Job completes, use the GetMatches
	// (https://docs.aws.amazon.com/customerprofiles/latest/APIReference/API_GetMatches.html)
	// API to return and review the results. Or, if you have configured ExportingConfig
	// in the MatchingRequest, you can download the results from S3.
	Matching *types.MatchingResponse

	// Usage-specific statistics about the domain.
	Stats *types.DomainStats

	// The tags used to organize, track, or control access for this resource.
	Tags map[string]string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetDomainMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpGetDomain{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpGetDomain{}, middleware.After)
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
	if err = addOpGetDomainValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetDomain(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opGetDomain(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "profile",
		OperationName: "GetDomain",
	}
}
