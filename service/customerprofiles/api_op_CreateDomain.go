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

// Creates a domain, which is a container for all customer data, such as customer
// profile attributes, object types, profile keys, and encryption keys. You can
// create multiple domains, and each domain can have multiple third-party
// integrations. Each Amazon Connect instance can be associated with only one
// domain. Multiple Amazon Connect instances can be associated with one domain. Use
// this API or UpdateDomain
// (https://docs.aws.amazon.com/customerprofiles/latest/APIReference/API_UpdateDomain.html)
// to enable identity resolution
// (https://docs.aws.amazon.com/customerprofiles/latest/APIReference/API_GetMatches.html):
// set Matching to true. To prevent cross-service impersonation when you call this
// API, see Cross-service confused deputy prevention
// (https://docs.aws.amazon.com/connect/latest/adminguide/cross-service-confused-deputy-prevention.html)
// for sample policies that you should apply.
func (c *Client) CreateDomain(ctx context.Context, params *CreateDomainInput, optFns ...func(*Options)) (*CreateDomainOutput, error) {
	if params == nil {
		params = &CreateDomainInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateDomain", params, optFns, c.addOperationCreateDomainMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateDomainOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateDomainInput struct {

	// The default number of days until the data within the domain expires.
	//
	// This member is required.
	DefaultExpirationDays *int32

	// The unique name of the domain.
	//
	// This member is required.
	DomainName *string

	// The URL of the SQS dead letter queue, which is used for reporting errors
	// associated with ingesting data from third party applications. You must set up a
	// policy on the DeadLetterQueue for the SendMessage operation to enable Amazon
	// Connect Customer Profiles to send messages to the DeadLetterQueue.
	DeadLetterQueueUrl *string

	// The default encryption key, which is an AWS managed key, is used when no
	// specific type of encryption key is specified. It is used to encrypt all data
	// before it is placed in permanent or semi-permanent storage.
	DefaultEncryptionKey *string

	// The process of matching duplicate profiles. If Matching = true, Amazon Connect
	// Customer Profiles starts a weekly batch process called Identity Resolution Job.
	// If you do not specify a date and time for Identity Resolution Job to run, by
	// default it runs every Saturday at 12AM UTC to detect duplicate profiles in your
	// domains. After the Identity Resolution Job completes, use the GetMatches
	// (https://docs.aws.amazon.com/customerprofiles/latest/APIReference/API_GetMatches.html)
	// API to return and review the results. Or, if you have configured ExportingConfig
	// in the MatchingRequest, you can download the results from S3.
	Matching *types.MatchingRequest

	// The tags used to organize, track, or control access for this resource.
	Tags map[string]string

	noSmithyDocumentSerde
}

type CreateDomainOutput struct {

	// The timestamp of when the domain was created.
	//
	// This member is required.
	CreatedAt *time.Time

	// The default number of days until the data within the domain expires.
	//
	// This member is required.
	DefaultExpirationDays *int32

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

	// The process of matching duplicate profiles. If Matching = true, Amazon Connect
	// Customer Profiles starts a weekly batch process called Identity Resolution Job.
	// If you do not specify a date and time for Identity Resolution Job to run, by
	// default it runs every Saturday at 12AM UTC to detect duplicate profiles in your
	// domains. After the Identity Resolution Job completes, use the GetMatches
	// (https://docs.aws.amazon.com/customerprofiles/latest/APIReference/API_GetMatches.html)
	// API to return and review the results. Or, if you have configured ExportingConfig
	// in the MatchingRequest, you can download the results from S3.
	Matching *types.MatchingResponse

	// The tags used to organize, track, or control access for this resource.
	Tags map[string]string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateDomainMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpCreateDomain{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpCreateDomain{}, middleware.After)
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
	if err = addOpCreateDomainValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateDomain(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateDomain(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "profile",
		OperationName: "CreateDomain",
	}
}