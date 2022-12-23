// Code generated by smithy-go-codegen DO NOT EDIT.

package personalize

import (
	"context"
	awsmiddleware "oss-sdk-go/oss/middleware"
	"oss-sdk-go/oss/signer/v4"
	"oss-sdk-go/service/personalize/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Describes the given campaign, including its status. A campaign can be in one of
// the following states:
//
// * CREATE PENDING > CREATE IN_PROGRESS > ACTIVE -or-
// CREATE FAILED
//
// * DELETE PENDING > DELETE IN_PROGRESS
//
// When the status is CREATE
// FAILED, the response includes the failureReason key, which describes why. For
// more information on campaigns, see CreateCampaign
// (https://docs.aws.amazon.com/personalize/latest/dg/API_CreateCampaign.html).
func (c *Client) DescribeCampaign(ctx context.Context, params *DescribeCampaignInput, optFns ...func(*Options)) (*DescribeCampaignOutput, error) {
	if params == nil {
		params = &DescribeCampaignInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeCampaign", params, optFns, c.addOperationDescribeCampaignMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeCampaignOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeCampaignInput struct {

	// The Amazon Resource Name (ARN) of the campaign.
	//
	// This member is required.
	CampaignArn *string

	noSmithyDocumentSerde
}

type DescribeCampaignOutput struct {

	// The properties of the campaign.
	Campaign *types.Campaign

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeCampaignMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDescribeCampaign{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDescribeCampaign{}, middleware.After)
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
	if err = addOpDescribeCampaignValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeCampaign(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDescribeCampaign(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "personalize",
		OperationName: "DescribeCampaign",
	}
}
