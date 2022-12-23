// Code generated by smithy-go-codegen DO NOT EDIT.

package iotsitewise

import (
	"context"
	"fmt"
	awsmiddleware "oss-sdk-go/oss/middleware"
	"oss-sdk-go/oss/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Retrieves information about a dashboard.
func (c *Client) DescribeDashboard(ctx context.Context, params *DescribeDashboardInput, optFns ...func(*Options)) (*DescribeDashboardOutput, error) {
	if params == nil {
		params = &DescribeDashboardInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeDashboard", params, optFns, c.addOperationDescribeDashboardMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeDashboardOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeDashboardInput struct {

	// The ID of the dashboard.
	//
	// This member is required.
	DashboardId *string

	noSmithyDocumentSerde
}

type DescribeDashboardOutput struct {

	// The ARN
	// (https://docs.aws.amazon.com/general/latest/gr/aws-arns-and-namespaces.html) of
	// the dashboard, which has the following format.
	// arn:${Partition}:iotsitewise:${Region}:${Account}:dashboard/${DashboardId}
	//
	// This member is required.
	DashboardArn *string

	// The date the dashboard was created, in Unix epoch time.
	//
	// This member is required.
	DashboardCreationDate *time.Time

	// The dashboard's definition JSON literal. For detailed information, see Creating
	// dashboards (CLI)
	// (https://docs.aws.amazon.com/iot-sitewise/latest/userguide/create-dashboards-using-aws-cli.html)
	// in the IoT SiteWise User Guide.
	//
	// This member is required.
	DashboardDefinition *string

	// The ID of the dashboard.
	//
	// This member is required.
	DashboardId *string

	// The date the dashboard was last updated, in Unix epoch time.
	//
	// This member is required.
	DashboardLastUpdateDate *time.Time

	// The name of the dashboard.
	//
	// This member is required.
	DashboardName *string

	// The ID of the project that the dashboard is in.
	//
	// This member is required.
	ProjectId *string

	// The dashboard's description.
	DashboardDescription *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeDashboardMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpDescribeDashboard{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpDescribeDashboard{}, middleware.After)
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
	if err = addEndpointPrefix_opDescribeDashboardMiddleware(stack); err != nil {
		return err
	}
	if err = addOpDescribeDashboardValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeDashboard(options.Region), middleware.Before); err != nil {
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

type endpointPrefix_opDescribeDashboardMiddleware struct {
}

func (*endpointPrefix_opDescribeDashboardMiddleware) ID() string {
	return "EndpointHostPrefix"
}

func (m *endpointPrefix_opDescribeDashboardMiddleware) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
) {
	if smithyhttp.GetHostnameImmutable(ctx) || smithyhttp.IsEndpointHostPrefixDisabled(ctx) {
		return next.HandleSerialize(ctx, in)
	}

	req, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, fmt.Errorf("unknown transport type %T", in.Request)
	}

	req.URL.Host = "monitor." + req.URL.Host

	return next.HandleSerialize(ctx, in)
}
func addEndpointPrefix_opDescribeDashboardMiddleware(stack *middleware.Stack) error {
	return stack.Serialize.Insert(&endpointPrefix_opDescribeDashboardMiddleware{}, `OperationSerializer`, middleware.After)
}

func newServiceMetadataMiddleware_opDescribeDashboard(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "iotsitewise",
		OperationName: "DescribeDashboard",
	}
}
