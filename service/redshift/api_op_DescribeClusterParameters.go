// Code generated by smithy-go-codegen DO NOT EDIT.

package redshift

import (
	"context"
	"fmt"
	awsmiddleware "oss-sdk-go/oss/middleware"
	"oss-sdk-go/oss/signer/v4"
	"oss-sdk-go/service/redshift/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Returns a detailed list of parameters contained within the specified Amazon
// Redshift parameter group. For each parameter the response includes information
// such as parameter name, description, data type, value, whether the parameter
// value is modifiable, and so on. You can specify source filter to retrieve
// parameters of only specific type. For example, to retrieve parameters that were
// modified by a user action such as from ModifyClusterParameterGroup, you can
// specify source equal to user. For more information about parameters and
// parameter groups, go to Amazon Redshift Parameter Groups
// (https://docs.aws.amazon.com/redshift/latest/mgmt/working-with-parameter-groups.html)
// in the Amazon Redshift Cluster Management Guide.
func (c *Client) DescribeClusterParameters(ctx context.Context, params *DescribeClusterParametersInput, optFns ...func(*Options)) (*DescribeClusterParametersOutput, error) {
	if params == nil {
		params = &DescribeClusterParametersInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeClusterParameters", params, optFns, c.addOperationDescribeClusterParametersMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeClusterParametersOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeClusterParametersInput struct {

	// The name of a cluster parameter group for which to return details.
	//
	// This member is required.
	ParameterGroupName *string

	// An optional parameter that specifies the starting point to return a set of
	// response records. When the results of a DescribeClusterParameters request exceed
	// the value specified in MaxRecords, Amazon Web Services returns a value in the
	// Marker field of the response. You can retrieve the next set of response records
	// by providing the returned marker value in the Marker parameter and retrying the
	// request.
	Marker *string

	// The maximum number of response records to return in each call. If the number of
	// remaining response records exceeds the specified MaxRecords value, a value is
	// returned in a marker field of the response. You can retrieve the next set of
	// records by retrying the command with the returned marker value. Default: 100
	// Constraints: minimum 20, maximum 100.
	MaxRecords *int32

	// The parameter types to return. Specify user to show parameters that are
	// different form the default. Similarly, specify engine-default to show parameters
	// that are the same as the default parameter group. Default: All parameter types
	// returned. Valid Values: user | engine-default
	Source *string

	noSmithyDocumentSerde
}

// Contains the output from the DescribeClusterParameters action.
type DescribeClusterParametersOutput struct {

	// A value that indicates the starting point for the next set of response records
	// in a subsequent request. If a value is returned in a response, you can retrieve
	// the next set of records by providing this returned marker value in the Marker
	// parameter and retrying the command. If the Marker field is empty, all response
	// records have been retrieved for the request.
	Marker *string

	// A list of Parameter instances. Each instance lists the parameters of one cluster
	// parameter group.
	Parameters []types.Parameter

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeClusterParametersMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsquery_serializeOpDescribeClusterParameters{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpDescribeClusterParameters{}, middleware.After)
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
	if err = addOpDescribeClusterParametersValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeClusterParameters(options.Region), middleware.Before); err != nil {
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

// DescribeClusterParametersAPIClient is a client that implements the
// DescribeClusterParameters operation.
type DescribeClusterParametersAPIClient interface {
	DescribeClusterParameters(context.Context, *DescribeClusterParametersInput, ...func(*Options)) (*DescribeClusterParametersOutput, error)
}

var _ DescribeClusterParametersAPIClient = (*Client)(nil)

// DescribeClusterParametersPaginatorOptions is the paginator options for
// DescribeClusterParameters
type DescribeClusterParametersPaginatorOptions struct {
	// The maximum number of response records to return in each call. If the number of
	// remaining response records exceeds the specified MaxRecords value, a value is
	// returned in a marker field of the response. You can retrieve the next set of
	// records by retrying the command with the returned marker value. Default: 100
	// Constraints: minimum 20, maximum 100.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// DescribeClusterParametersPaginator is a paginator for DescribeClusterParameters
type DescribeClusterParametersPaginator struct {
	options   DescribeClusterParametersPaginatorOptions
	client    DescribeClusterParametersAPIClient
	params    *DescribeClusterParametersInput
	nextToken *string
	firstPage bool
}

// NewDescribeClusterParametersPaginator returns a new
// DescribeClusterParametersPaginator
func NewDescribeClusterParametersPaginator(client DescribeClusterParametersAPIClient, params *DescribeClusterParametersInput, optFns ...func(*DescribeClusterParametersPaginatorOptions)) *DescribeClusterParametersPaginator {
	if params == nil {
		params = &DescribeClusterParametersInput{}
	}

	options := DescribeClusterParametersPaginatorOptions{}
	if params.MaxRecords != nil {
		options.Limit = *params.MaxRecords
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &DescribeClusterParametersPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.Marker,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *DescribeClusterParametersPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next DescribeClusterParameters page.
func (p *DescribeClusterParametersPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*DescribeClusterParametersOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.Marker = p.nextToken

	var limit *int32
	if p.options.Limit > 0 {
		limit = &p.options.Limit
	}
	params.MaxRecords = limit

	result, err := p.client.DescribeClusterParameters(ctx, &params, optFns...)
	if err != nil {
		return nil, err
	}
	p.firstPage = false

	prevToken := p.nextToken
	p.nextToken = result.Marker

	if p.options.StopOnDuplicateToken &&
		prevToken != nil &&
		p.nextToken != nil &&
		*prevToken == *p.nextToken {
		p.nextToken = nil
	}

	return result, nil
}

func newServiceMetadataMiddleware_opDescribeClusterParameters(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "redshift",
		OperationName: "DescribeClusterParameters",
	}
}