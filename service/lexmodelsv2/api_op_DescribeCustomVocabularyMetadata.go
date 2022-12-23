// Code generated by smithy-go-codegen DO NOT EDIT.

package lexmodelsv2

import (
	"context"
	awsmiddleware "oss-sdk-go/oss/middleware"
	"oss-sdk-go/oss/signer/v4"
	"oss-sdk-go/service/lexmodelsv2/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Provides metadata information about a custom vocabulary.
func (c *Client) DescribeCustomVocabularyMetadata(ctx context.Context, params *DescribeCustomVocabularyMetadataInput, optFns ...func(*Options)) (*DescribeCustomVocabularyMetadataOutput, error) {
	if params == nil {
		params = &DescribeCustomVocabularyMetadataInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeCustomVocabularyMetadata", params, optFns, c.addOperationDescribeCustomVocabularyMetadataMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeCustomVocabularyMetadataOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeCustomVocabularyMetadataInput struct {

	// The unique identifier of the bot that contains the custom vocabulary.
	//
	// This member is required.
	BotId *string

	// The bot version of the bot to return metadata for.
	//
	// This member is required.
	BotVersion *string

	// The locale to return the custom vocabulary information for. The locale must be
	// en_GB.
	//
	// This member is required.
	LocaleId *string

	noSmithyDocumentSerde
}

type DescribeCustomVocabularyMetadataOutput struct {

	// The identifier of the bot that contains the custom vocabulary.
	BotId *string

	// The version of the bot that contains the custom vocabulary to describe.
	BotVersion *string

	// The date and time that the custom vocabulary was created.
	CreationDateTime *time.Time

	// The status of the custom vocabulary. If the status is Ready the custom
	// vocabulary is ready to use.
	CustomVocabularyStatus types.CustomVocabularyStatus

	// The date and time that the custom vocabulary was last updated.
	LastUpdatedDateTime *time.Time

	// The locale that contains the custom vocabulary to describe.
	LocaleId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeCustomVocabularyMetadataMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpDescribeCustomVocabularyMetadata{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpDescribeCustomVocabularyMetadata{}, middleware.After)
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
	if err = addOpDescribeCustomVocabularyMetadataValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeCustomVocabularyMetadata(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDescribeCustomVocabularyMetadata(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "lex",
		OperationName: "DescribeCustomVocabularyMetadata",
	}
}
