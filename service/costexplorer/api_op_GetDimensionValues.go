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

// Retrieves all available filter values for a specified filter over a period of
// time. You can search the dimension values for an arbitrary string.
func (c *Client) GetDimensionValues(ctx context.Context, params *GetDimensionValuesInput, optFns ...func(*Options)) (*GetDimensionValuesOutput, error) {
	if params == nil {
		params = &GetDimensionValuesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetDimensionValues", params, optFns, c.addOperationGetDimensionValuesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetDimensionValuesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetDimensionValuesInput struct {

	// The name of the dimension. Each Dimension is available for a different Context.
	// For more information, see Context.
	//
	// This member is required.
	Dimension types.Dimension

	// The start date and end date for retrieving the dimension values. The start date
	// is inclusive, but the end date is exclusive. For example, if start is 2017-01-01
	// and end is 2017-05-01, then the cost and usage data is retrieved from 2017-01-01
	// up to and including 2017-04-30 but not including 2017-05-01.
	//
	// This member is required.
	TimePeriod *types.DateInterval

	// The context for the call to GetDimensionValues. This can be RESERVATIONS or
	// COST_AND_USAGE. The default value is COST_AND_USAGE. If the context is set to
	// RESERVATIONS, the resulting dimension values can be used in the
	// GetReservationUtilization operation. If the context is set to COST_AND_USAGE,
	// the resulting dimension values can be used in the GetCostAndUsage operation. If
	// you set the context to COST_AND_USAGE, you can use the following dimensions for
	// searching:
	//
	// * AZ - The Availability Zone. An example is us-east-1a.
	//
	// *
	// BILLING_ENTITY - The Amazon Web Services seller that your account is with.
	// Possible values are the following: - Amazon Web Services(Amazon Web Services):
	// The entity that sells Amazon Web Services. - AISPL (Amazon Internet Services
	// Pvt. Ltd.): The local Indian entity that's an acting reseller for Amazon Web
	// Services in India. - Amazon Web Services Marketplace: The entity that supports
	// the sale of solutions that are built on Amazon Web Services by third-party
	// software providers.
	//
	// * CACHE_ENGINE - The Amazon ElastiCache operating system.
	// Examples are Windows or Linux.
	//
	// * DEPLOYMENT_OPTION - The scope of Amazon
	// Relational Database Service deployments. Valid values are SingleAZ and
	// MultiAZ.
	//
	// * DATABASE_ENGINE - The Amazon Relational Database Service database.
	// Examples are Aurora or MySQL.
	//
	// * INSTANCE_TYPE - The type of Amazon EC2
	// instance. An example is m4.xlarge.
	//
	// * INSTANCE_TYPE_FAMILY - A family of
	// instance types optimized to fit different use cases. Examples are Compute
	// Optimized (for example, C4, C5, C6g, and C7g), Memory Optimization (for example,
	// R4, R5n, R5b, and R6g).
	//
	// * INVOICING_ENTITY - The name of the entity that issues
	// the Amazon Web Services invoice.
	//
	// * LEGAL_ENTITY_NAME - The name of the
	// organization that sells you Amazon Web Services services, such as Amazon Web
	// Services.
	//
	// * LINKED_ACCOUNT - The description in the attribute map that includes
	// the full name of the member account. The value field contains the Amazon Web
	// Services ID of the member account.
	//
	// * OPERATING_SYSTEM - The operating system.
	// Examples are Windows or Linux.
	//
	// * OPERATION - The action performed. Examples
	// include RunInstance and CreateBucket.
	//
	// * PLATFORM - The Amazon EC2 operating
	// system. Examples are Windows or Linux.
	//
	// * PURCHASE_TYPE - The reservation type
	// of the purchase that this usage is related to. Examples include On-Demand
	// Instances and Standard Reserved Instances.
	//
	// * RESERVATION_ID - The unique
	// identifier for an Amazon Web Services Reservation Instance.
	//
	// * SAVINGS_PLAN_ARN
	// - The unique identifier for your Savings Plans.
	//
	// * SAVINGS_PLANS_TYPE - Type of
	// Savings Plans (EC2 Instance or Compute).
	//
	// * SERVICE - The Amazon Web Services
	// service such as Amazon DynamoDB.
	//
	// * TENANCY - The tenancy of a resource.
	// Examples are shared or dedicated.
	//
	// * USAGE_TYPE - The type of usage. An example
	// is DataTransfer-In-Bytes. The response for the GetDimensionValues operation
	// includes a unit attribute. Examples include GB and Hrs.
	//
	// * USAGE_TYPE_GROUP -
	// The grouping of common usage types. An example is Amazon EC2: CloudWatch –
	// Alarms. The response for this operation includes a unit attribute.
	//
	// * REGION -
	// The Amazon Web Services Region.
	//
	// * RECORD_TYPE - The different types of charges
	// such as Reserved Instance (RI) fees, usage costs, tax refunds, and credits.
	//
	// *
	// RESOURCE_ID - The unique identifier of the resource. ResourceId is an opt-in
	// feature only available for last 14 days for EC2-Compute Service.
	//
	// If you set the
	// context to RESERVATIONS, you can use the following dimensions for searching:
	//
	// *
	// AZ - The Availability Zone. An example is us-east-1a.
	//
	// * CACHE_ENGINE - The
	// Amazon ElastiCache operating system. Examples are Windows or Linux.
	//
	// *
	// DEPLOYMENT_OPTION - The scope of Amazon Relational Database Service deployments.
	// Valid values are SingleAZ and MultiAZ.
	//
	// * INSTANCE_TYPE - The type of Amazon EC2
	// instance. An example is m4.xlarge.
	//
	// * LINKED_ACCOUNT - The description in the
	// attribute map that includes the full name of the member account. The value field
	// contains the Amazon Web Services ID of the member account.
	//
	// * PLATFORM - The
	// Amazon EC2 operating system. Examples are Windows or Linux.
	//
	// * REGION - The
	// Amazon Web Services Region.
	//
	// * SCOPE (Utilization only) - The scope of a
	// Reserved Instance (RI). Values are regional or a single Availability Zone.
	//
	// *
	// TAG (Coverage only) - The tags that are associated with a Reserved Instance
	// (RI).
	//
	// * TENANCY - The tenancy of a resource. Examples are shared or
	// dedicated.
	//
	// If you set the context to SAVINGS_PLANS, you can use the following
	// dimensions for searching:
	//
	// * SAVINGS_PLANS_TYPE - Type of Savings Plans (EC2
	// Instance or Compute)
	//
	// * PAYMENT_OPTION - The payment option for the given
	// Savings Plans (for example, All Upfront)
	//
	// * REGION - The Amazon Web Services
	// Region.
	//
	// * INSTANCE_TYPE_FAMILY - The family of instances (For example, m5)
	//
	// *
	// LINKED_ACCOUNT - The description in the attribute map that includes the full
	// name of the member account. The value field contains the Amazon Web Services ID
	// of the member account.
	//
	// * SAVINGS_PLAN_ARN - The unique identifier for your
	// Savings Plans.
	Context types.Context

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
	// this parameter. For GetDimensionValues, MaxResults has an upper limit of 1000.
	MaxResults int32

	// The token to retrieve the next set of results. Amazon Web Services provides the
	// token when the response from a previous call has more results than the maximum
	// page size.
	NextPageToken *string

	// The value that you want to search the filter values for.
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
	// The supported values for the SortOrder key are ASCENDING
	// or DESCENDING. When you specify a SortBy paramater, the context must be
	// COST_AND_USAGE. Further, when using SortBy, NextPageToken and SearchString
	// aren't supported.
	SortBy []types.SortDefinition

	noSmithyDocumentSerde
}

type GetDimensionValuesOutput struct {

	// The filters that you used to filter your request. Some dimensions are available
	// only for a specific context. If you set the context to COST_AND_USAGE, you can
	// use the following dimensions for searching:
	//
	// * AZ - The Availability Zone. An
	// example is us-east-1a.
	//
	// * DATABASE_ENGINE - The Amazon Relational Database
	// Service database. Examples are Aurora or MySQL.
	//
	// * INSTANCE_TYPE - The type of
	// Amazon EC2 instance. An example is m4.xlarge.
	//
	// * LEGAL_ENTITY_NAME - The name of
	// the organization that sells you Amazon Web Services services, such as Amazon Web
	// Services.
	//
	// * LINKED_ACCOUNT - The description in the attribute map that includes
	// the full name of the member account. The value field contains the Amazon Web
	// Services ID of the member account.
	//
	// * OPERATING_SYSTEM - The operating system.
	// Examples are Windows or Linux.
	//
	// * OPERATION - The action performed. Examples
	// include RunInstance and CreateBucket.
	//
	// * PLATFORM - The Amazon EC2 operating
	// system. Examples are Windows or Linux.
	//
	// * PURCHASE_TYPE - The reservation type
	// of the purchase to which this usage is related. Examples include On-Demand
	// Instances and Standard Reserved Instances.
	//
	// * SERVICE - The Amazon Web Services
	// service such as Amazon DynamoDB.
	//
	// * USAGE_TYPE - The type of usage. An example
	// is DataTransfer-In-Bytes. The response for the GetDimensionValues operation
	// includes a unit attribute. Examples include GB and Hrs.
	//
	// * USAGE_TYPE_GROUP -
	// The grouping of common usage types. An example is Amazon EC2: CloudWatch –
	// Alarms. The response for this operation includes a unit attribute.
	//
	// *
	// RECORD_TYPE - The different types of charges such as RI fees, usage costs, tax
	// refunds, and credits.
	//
	// * RESOURCE_ID - The unique identifier of the resource.
	// ResourceId is an opt-in feature only available for last 14 days for EC2-Compute
	// Service.
	//
	// If you set the context to RESERVATIONS, you can use the following
	// dimensions for searching:
	//
	// * AZ - The Availability Zone. An example is
	// us-east-1a.
	//
	// * CACHE_ENGINE - The Amazon ElastiCache operating system. Examples
	// are Windows or Linux.
	//
	// * DEPLOYMENT_OPTION - The scope of Amazon Relational
	// Database Service deployments. Valid values are SingleAZ and MultiAZ.
	//
	// *
	// INSTANCE_TYPE - The type of Amazon EC2 instance. An example is m4.xlarge.
	//
	// *
	// LINKED_ACCOUNT - The description in the attribute map that includes the full
	// name of the member account. The value field contains the Amazon Web Services ID
	// of the member account.
	//
	// * PLATFORM - The Amazon EC2 operating system. Examples
	// are Windows or Linux.
	//
	// * REGION - The Amazon Web Services Region.
	//
	// * SCOPE
	// (Utilization only) - The scope of a Reserved Instance (RI). Values are regional
	// or a single Availability Zone.
	//
	// * TAG (Coverage only) - The tags that are
	// associated with a Reserved Instance (RI).
	//
	// * TENANCY - The tenancy of a
	// resource. Examples are shared or dedicated.
	//
	// If you set the context to
	// SAVINGS_PLANS, you can use the following dimensions for searching:
	//
	// *
	// SAVINGS_PLANS_TYPE - Type of Savings Plans (EC2 Instance or Compute)
	//
	// *
	// PAYMENT_OPTION - Payment option for the given Savings Plans (for example, All
	// Upfront)
	//
	// * REGION - The Amazon Web Services Region.
	//
	// * INSTANCE_TYPE_FAMILY -
	// The family of instances (For example, m5)
	//
	// * LINKED_ACCOUNT - The description in
	// the attribute map that includes the full name of the member account. The value
	// field contains the Amazon Web Services ID of the member account.
	//
	// *
	// SAVINGS_PLAN_ARN - The unique identifier for your Savings Plan
	//
	// This member is required.
	DimensionValues []types.DimensionValuesWithAttributes

	// The number of results that Amazon Web Services returned at one time.
	//
	// This member is required.
	ReturnSize *int32

	// The total number of search results.
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

func (c *Client) addOperationGetDimensionValuesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpGetDimensionValues{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpGetDimensionValues{}, middleware.After)
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
	if err = addOpGetDimensionValuesValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetDimensionValues(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opGetDimensionValues(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ce",
		OperationName: "GetDimensionValues",
	}
}
