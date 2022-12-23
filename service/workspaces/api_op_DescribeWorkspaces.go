// Code generated by smithy-go-codegen DO NOT EDIT.

package workspaces

import (
	"context"
	"fmt"
	awsmiddleware "oss-sdk-go/oss/middleware"
	"oss-sdk-go/oss/signer/v4"
	"oss-sdk-go/service/workspaces/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Describes the specified WorkSpaces. You can filter the results by using the
// bundle identifier, directory identifier, or owner, but you can specify only one
// filter at a time.
func (c *Client) DescribeWorkspaces(ctx context.Context, params *DescribeWorkspacesInput, optFns ...func(*Options)) (*DescribeWorkspacesOutput, error) {
	if params == nil {
		params = &DescribeWorkspacesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeWorkspaces", params, optFns, c.addOperationDescribeWorkspacesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeWorkspacesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeWorkspacesInput struct {

	// The identifier of the bundle. All WorkSpaces that are created from this bundle
	// are retrieved. You cannot combine this parameter with any other filter.
	BundleId *string

	// The identifier of the directory. In addition, you can optionally specify a
	// specific directory user (see UserName). You cannot combine this parameter with
	// any other filter.
	DirectoryId *string

	// The maximum number of items to return.
	Limit *int32

	// If you received a NextToken from a previous call that was paginated, provide
	// this token to receive the next set of results.
	NextToken *string

	// The name of the directory user. You must specify this parameter with
	// DirectoryId.
	UserName *string

	// The identifiers of the WorkSpaces. You cannot combine this parameter with any
	// other filter. Because the CreateWorkspaces operation is asynchronous, the
	// identifier it returns is not immediately available. If you immediately call
	// DescribeWorkspaces with this identifier, no information is returned.
	WorkspaceIds []string

	noSmithyDocumentSerde
}

type DescribeWorkspacesOutput struct {

	// The token to use to retrieve the next page of results. This value is null when
	// there are no more results to return.
	NextToken *string

	// Information about the WorkSpaces. Because CreateWorkspaces is an asynchronous
	// operation, some of the returned information could be incomplete.
	Workspaces []types.Workspace

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeWorkspacesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDescribeWorkspaces{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDescribeWorkspaces{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeWorkspaces(options.Region), middleware.Before); err != nil {
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

// DescribeWorkspacesAPIClient is a client that implements the DescribeWorkspaces
// operation.
type DescribeWorkspacesAPIClient interface {
	DescribeWorkspaces(context.Context, *DescribeWorkspacesInput, ...func(*Options)) (*DescribeWorkspacesOutput, error)
}

var _ DescribeWorkspacesAPIClient = (*Client)(nil)

// DescribeWorkspacesPaginatorOptions is the paginator options for
// DescribeWorkspaces
type DescribeWorkspacesPaginatorOptions struct {
	// The maximum number of items to return.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// DescribeWorkspacesPaginator is a paginator for DescribeWorkspaces
type DescribeWorkspacesPaginator struct {
	options   DescribeWorkspacesPaginatorOptions
	client    DescribeWorkspacesAPIClient
	params    *DescribeWorkspacesInput
	nextToken *string
	firstPage bool
}

// NewDescribeWorkspacesPaginator returns a new DescribeWorkspacesPaginator
func NewDescribeWorkspacesPaginator(client DescribeWorkspacesAPIClient, params *DescribeWorkspacesInput, optFns ...func(*DescribeWorkspacesPaginatorOptions)) *DescribeWorkspacesPaginator {
	if params == nil {
		params = &DescribeWorkspacesInput{}
	}

	options := DescribeWorkspacesPaginatorOptions{}
	if params.Limit != nil {
		options.Limit = *params.Limit
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &DescribeWorkspacesPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *DescribeWorkspacesPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next DescribeWorkspaces page.
func (p *DescribeWorkspacesPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*DescribeWorkspacesOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.NextToken = p.nextToken

	var limit *int32
	if p.options.Limit > 0 {
		limit = &p.options.Limit
	}
	params.Limit = limit

	result, err := p.client.DescribeWorkspaces(ctx, &params, optFns...)
	if err != nil {
		return nil, err
	}
	p.firstPage = false

	prevToken := p.nextToken
	p.nextToken = result.NextToken

	if p.options.StopOnDuplicateToken &&
		prevToken != nil &&
		p.nextToken != nil &&
		*prevToken == *p.nextToken {
		p.nextToken = nil
	}

	return result, nil
}

func newServiceMetadataMiddleware_opDescribeWorkspaces(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "workspaces",
		OperationName: "DescribeWorkspaces",
	}
}
