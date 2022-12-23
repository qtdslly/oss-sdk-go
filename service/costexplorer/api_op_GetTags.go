// Code generated by smithy-go-codegen DO NOT EDIT.

package costexplorer

import (
	"context"
	awsmiddleware "oss-sdk-go/oss/middleware"
	"oss-sdk-go/oss/signer/v4"
	"oss-sdk-go/service/costexplorer/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Queries for available tag keys and tag values for a specified period. You can
// search the tag values for an arbitrary string.
func (c *Client) GetTags(ctx context.Context, params *GetTagsInput, optFns ...func(*Options)) (*GetTagsOutput, error) {
	if params == nil {
		params = &GetTagsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetTags", params, optFns, c.addOperationGetTagsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetTagsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetTagsInput struct {

	// The start and end dates for retrieving the dimension values. The start date is
	// inclusive, but the end date is exclusive. For example, if start is 2017-01-01
	// and end is 2017-05-01, then the cost and usage data is retrieved from 2017-01-01
	// up to and including 2017-04-30 but not including 2017-05-01.
	//
	// This member is required.
	TimePeriod *types.DateInterval

	// Use Expression to filter by cost or by usage. There are two patterns:
	//
	// * Simple
	// dimension values - You can set the dimension name and values for the filters
	// that you plan to use. For example, you can filter for REGION==us-east-1 OR
	// REGION==us-west-1. For GetRightsizingRecommendation, the Region is a full name
	// (for example, REGION==US East (N. Virginia). The Expression example is as
	// follows: { "Dimensions": { "Key": "REGION", "Values": [ "us-east-1", “us-west-1”
	// ] } } The list of dimension values are OR'd together to retrieve cost or usage
	// data. You can create Expression and DimensionValues objects using either with*
	// methods or set* methods in multiple lines.
	//
	// * Compound dimension values with
	// logical operations - You can use multiple Expression types and the logical
	// operators AND/OR/NOT to create a list of one or more Expression objects. By
	// doing this, you can filter on more advanced options. For example, you can filter
	// on ((REGION == us-east-1 OR REGION == us-west-1) OR (TAG.Type == Type1)) AND
	// (USAGE_TYPE != DataTransfer). The Expression for that is as follows: { "And": [
	// {"Or": [ {"Dimensions": { "Key": "REGION", "Values": [ "us-east-1", "us-west-1"
	// ] }}, {"Tags": { "Key": "TagName", "Values": ["Value1"] } } ]}, {"Not":
	// {"Dimensions": { "Key": "USAGE_TYPE", "Values": ["DataTransfer"] }}} ] }
	// Because each Expression can have only one operator, the service returns an error
	// if more than one is specified. The following example shows an Expression object
	// that creates an error.  { "And": [ ... ], "DimensionValues": { "Dimension":
	// "USAGE_TYPE", "Values": [ "DataTransfer" ] } }
	//
	// For the
	// GetRightsizingRecommendation action, a combination of OR and NOT isn't
	// supported. OR isn't supported between different dimensions, or dimensions and
	// tags. NOT operators aren't supported. Dimensions are also limited to
	// LINKED_ACCOUNT, REGION, or RIGHTSIZING_TYPE. For the
	// GetReservationPurchaseRecommendation action, only NOT is supported. AND and OR
	// aren't supported. Dimensions are limited to LINKED_ACCOUNT.
	Filter *types.Expression

	// This field is only used when SortBy is provided in the request. The maximum
	// number of objects that are returned for this request. If MaxResults isn't
	// specified with SortBy, the request returns 1000 results as the default value for
	// this parameter. For GetTags, MaxResults has an upper quota of 1000.
	MaxResults int32

	// The token to retrieve the next set of results. Amazon Web Services provides the
	// token when the response from a previous call has more results than the maximum
	// page size.
	NextPageToken *string

	// The value that you want to search for.
	SearchString *string

	// The value that you want to sort the data by. The key represents cost and usage
	// metrics. The following values are supported:
	//
	// * BlendedCost
	//
	// * UnblendedCost
	//
	// *
	// AmortizedCost
	//
	// * NetAmortizedCost
	//
	// * NetUnblendedCost
	//
	// * UsageQuantity
	//
	// *
	// NormalizedUsageAmount
	//
	// The supported values for SortOrder are ASCENDING and
	// DESCENDING. When you use SortBy, NextPageToken and SearchString aren't
	// supported.
	SortBy []types.SortDefinition

	// The key of the tag that you want to return values for.
	TagKey *string

	noSmithyDocumentSerde
}

type GetTagsOutput struct {

	// The number of query results that Amazon Web Services returns at a time.
	//
	// This member is required.
	ReturnSize *int32

	// The tags that match your request.
	//
	// This member is required.
	Tags []string

	// The total number of query results.
	//
	// This member is required.
	TotalSize *int32

	// The token for the next set of retrievable results. Amazon Web Services provides
	// the token when the response from a previous call has more results than the
	// maximum page size.
	NextPageToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetTagsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpGetTags{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpGetTags{}, middleware.After)
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
	if err = addOpGetTagsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetTags(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opGetTags(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ce",
		OperationName: "GetTags",
	}
}
