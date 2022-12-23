// Code generated by smithy-go-codegen DO NOT EDIT.

package gamelift

import (
	"context"
	"fmt"
	awsmiddleware "oss-sdk-go/oss/middleware"
	"oss-sdk-go/oss/signer/v4"
	"oss-sdk-go/service/gamelift/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Retrieves all active game sessions that match a set of search criteria and sorts
// them into a specified order. This operation is not designed to be continually
// called to track game session status. This practice can cause you to exceed your
// API limit, which results in errors. Instead, you must configure configure an
// Amazon Simple Notification Service (SNS) topic to receive notifications from
// FlexMatch or queues. Continuously polling game session status with
// DescribeGameSessions should only be used for games in development with low game
// session usage. When searching for game sessions, you specify exactly where you
// want to search and provide a search filter expression, a sort expression, or
// both. A search request can search only one fleet, but it can search all of a
// fleet's locations. This operation can be used in the following ways:
//
// * To
// search all game sessions that are currently running on all locations in a fleet,
// provide a fleet or alias ID. This approach returns game sessions in the fleet's
// home Region and all remote locations that fit the search criteria.
//
// * To search
// all game sessions that are currently running on a specific fleet location,
// provide a fleet or alias ID and a location name. For location, you can specify a
// fleet's home Region or any remote location.
//
// Use the pagination parameters to
// retrieve results as a set of sequential pages. If successful, a GameSession
// object is returned for each game session that matches the request. Search finds
// game sessions that are in ACTIVE status only. To retrieve information on game
// sessions in other statuses, use DescribeGameSessions. You can search or sort by
// the following game session attributes:
//
// * gameSessionId -- A unique identifier
// for the game session. You can use either a GameSessionId or GameSessionArn
// value.
//
// * gameSessionName -- Name assigned to a game session. This value is set
// when requesting a new game session with CreateGameSession or updating with
// UpdateGameSession. Game session names do not need to be unique to a game
// session.
//
// * gameSessionProperties -- Custom data defined in a game session's
// GameProperty parameter. GameProperty values are stored as key:value pairs; the
// filter expression must indicate the key and a string to search the data values
// for. For example, to search for game sessions with custom data containing the
// key:value pair "gameMode:brawl", specify the following:
// gameSessionProperties.gameMode = "brawl". All custom data values are searched as
// strings.
//
// * maximumSessions -- Maximum number of player sessions allowed for a
// game session. This value is set when requesting a new game session with
// CreateGameSession or updating with UpdateGameSession.
//
// * creationTimeMillis --
// Value indicating when a game session was created. It is expressed in Unix time
// as milliseconds.
//
// * playerSessionCount -- Number of players currently connected
// to a game session. This value changes rapidly as players join the session or
// drop out.
//
// * hasAvailablePlayerSessions -- Boolean value indicating whether a
// game session has reached its maximum number of players. It is highly recommended
// that all search requests include this filter attribute to optimize search
// performance and return only sessions that players can join.
//
// Returned values for
// playerSessionCount and hasAvailablePlayerSessions change quickly as players join
// sessions and others drop out. Results should be considered a snapshot in time.
// Be sure to refresh search results often, and handle sessions that fill up before
// a player can join. Related actions CreateGameSession | DescribeGameSessions |
// DescribeGameSessionDetails | SearchGameSessions | UpdateGameSession |
// GetGameSessionLogUrl | StartGameSessionPlacement | DescribeGameSessionPlacement
// | StopGameSessionPlacement | All APIs by task
// (https://docs.aws.amazon.com/gamelift/latest/developerguide/reference-awssdk.html#reference-awssdk-resources-fleets)
func (c *Client) SearchGameSessions(ctx context.Context, params *SearchGameSessionsInput, optFns ...func(*Options)) (*SearchGameSessionsOutput, error) {
	if params == nil {
		params = &SearchGameSessionsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "SearchGameSessions", params, optFns, c.addOperationSearchGameSessionsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*SearchGameSessionsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Represents the input for a request operation.
type SearchGameSessionsInput struct {

	// A unique identifier for the alias associated with the fleet to search for active
	// game sessions. You can use either the alias ID or ARN value. Each request must
	// reference either a fleet ID or alias ID, but not both.
	AliasId *string

	// String containing the search criteria for the session search. If no filter
	// expression is included, the request returns results for all game sessions in the
	// fleet that are in ACTIVE status. A filter expression can contain one or multiple
	// conditions. Each condition consists of the following:
	//
	// * Operand -- Name of a
	// game session attribute. Valid values are gameSessionName, gameSessionId,
	// gameSessionProperties, maximumSessions, creationTimeMillis, playerSessionCount,
	// hasAvailablePlayerSessions.
	//
	// * Comparator -- Valid comparators are: =, <>, <, >,
	// <=, >=.
	//
	// * Value -- Value to be searched for. Values may be numbers, boolean
	// values (true/false) or strings depending on the operand. String values are case
	// sensitive and must be enclosed in single quotes. Special characters must be
	// escaped. Boolean and string values can only be used with the comparators = and
	// <>. For example, the following filter expression searches on gameSessionName:
	// "FilterExpression": "gameSessionName = 'Matt\\'s Awesome Game 1'".
	//
	// To chain
	// multiple conditions in a single expression, use the logical keywords AND, OR,
	// and NOT and parentheses as needed. For example: x AND y AND NOT z, NOT (x OR y).
	// Session search evaluates conditions from left to right using the following
	// precedence rules:
	//
	// * =, <>, <, >, <=, >=
	//
	// * Parentheses
	//
	// * NOT
	//
	// * AND
	//
	// * OR
	//
	// For
	// example, this filter expression retrieves game sessions hosting at least ten
	// players that have an open player slot: "maximumSessions>=10 AND
	// hasAvailablePlayerSessions=true".
	FilterExpression *string

	// A unique identifier for the fleet to search for active game sessions. You can
	// use either the fleet ID or ARN value. Each request must reference either a fleet
	// ID or alias ID, but not both.
	FleetId *string

	// The maximum number of results to return. Use this parameter with NextToken to
	// get results as a set of sequential pages. The maximum number of results returned
	// is 20, even if this value is not set or is set higher than 20.
	Limit *int32

	// A fleet location to search for game sessions. You can specify a fleet's home
	// Region or a remote location. Use the Amazon Web Services Region code format,
	// such as us-west-2.
	Location *string

	// A token that indicates the start of the next sequential page of results. Use the
	// token that is returned with a previous call to this operation. To start at the
	// beginning of the result set, do not specify a value.
	NextToken *string

	// Instructions on how to sort the search results. If no sort expression is
	// included, the request returns results in random order. A sort expression
	// consists of the following elements:
	//
	// * Operand -- Name of a game session
	// attribute. Valid values are gameSessionName, gameSessionId,
	// gameSessionProperties, maximumSessions, creationTimeMillis, playerSessionCount,
	// hasAvailablePlayerSessions.
	//
	// * Order -- Valid sort orders are ASC (ascending)
	// and DESC (descending).
	//
	// For example, this sort expression returns the oldest
	// active sessions first: "SortExpression": "creationTimeMillis ASC". Results with
	// a null value for the sort operand are returned at the end of the list.
	SortExpression *string

	noSmithyDocumentSerde
}

// Represents the returned data in response to a request operation.
type SearchGameSessionsOutput struct {

	// A collection of objects containing game session properties for each session that
	// matches the request.
	GameSessions []types.GameSession

	// A token that indicates where to resume retrieving results on the next call to
	// this operation. If no token is returned, these results represent the end of the
	// list.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationSearchGameSessionsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpSearchGameSessions{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpSearchGameSessions{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opSearchGameSessions(options.Region), middleware.Before); err != nil {
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

// SearchGameSessionsAPIClient is a client that implements the SearchGameSessions
// operation.
type SearchGameSessionsAPIClient interface {
	SearchGameSessions(context.Context, *SearchGameSessionsInput, ...func(*Options)) (*SearchGameSessionsOutput, error)
}

var _ SearchGameSessionsAPIClient = (*Client)(nil)

// SearchGameSessionsPaginatorOptions is the paginator options for
// SearchGameSessions
type SearchGameSessionsPaginatorOptions struct {
	// The maximum number of results to return. Use this parameter with NextToken to
	// get results as a set of sequential pages. The maximum number of results returned
	// is 20, even if this value is not set or is set higher than 20.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// SearchGameSessionsPaginator is a paginator for SearchGameSessions
type SearchGameSessionsPaginator struct {
	options   SearchGameSessionsPaginatorOptions
	client    SearchGameSessionsAPIClient
	params    *SearchGameSessionsInput
	nextToken *string
	firstPage bool
}

// NewSearchGameSessionsPaginator returns a new SearchGameSessionsPaginator
func NewSearchGameSessionsPaginator(client SearchGameSessionsAPIClient, params *SearchGameSessionsInput, optFns ...func(*SearchGameSessionsPaginatorOptions)) *SearchGameSessionsPaginator {
	if params == nil {
		params = &SearchGameSessionsInput{}
	}

	options := SearchGameSessionsPaginatorOptions{}
	if params.Limit != nil {
		options.Limit = *params.Limit
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &SearchGameSessionsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *SearchGameSessionsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next SearchGameSessions page.
func (p *SearchGameSessionsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*SearchGameSessionsOutput, error) {
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

	result, err := p.client.SearchGameSessions(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opSearchGameSessions(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "gamelift",
		OperationName: "SearchGameSessions",
	}
}
