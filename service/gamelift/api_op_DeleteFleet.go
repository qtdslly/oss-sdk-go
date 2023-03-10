// Code generated by smithy-go-codegen DO NOT EDIT.

package gamelift

import (
	"context"
	awsmiddleware "oss-sdk-go/oss/middleware"
	"oss-sdk-go/oss/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Deletes all resources and information related a fleet. Any current fleet
// instances, including those in remote locations, are shut down. You don't need to
// call DeleteFleetLocations separately. If the fleet being deleted has a VPC
// peering connection, you first need to get a valid authorization (good for 24
// hours) by calling CreateVpcPeeringAuthorization. You do not need to explicitly
// delete the VPC peering connection--this is done as part of the delete fleet
// process. To delete a fleet, specify the fleet ID to be terminated. During the
// deletion process the fleet status is changed to DELETING. When completed, the
// status switches to TERMINATED and the fleet event FLEET_DELETED is sent. Learn
// more Setting up GameLift Fleets
// (https://docs.aws.amazon.com/gamelift/latest/developerguide/fleets-intro.html)
// Related actions CreateFleetLocations | UpdateFleetAttributes |
// UpdateFleetCapacity | UpdateFleetPortSettings | UpdateRuntimeConfiguration |
// StopFleetActions | StartFleetActions | PutScalingPolicy | DeleteFleet |
// DeleteFleetLocations | DeleteScalingPolicy | All APIs by task
// (https://docs.aws.amazon.com/gamelift/latest/developerguide/reference-awssdk.html#reference-awssdk-resources-fleets)
func (c *Client) DeleteFleet(ctx context.Context, params *DeleteFleetInput, optFns ...func(*Options)) (*DeleteFleetOutput, error) {
	if params == nil {
		params = &DeleteFleetInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DeleteFleet", params, optFns, c.addOperationDeleteFleetMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DeleteFleetOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Represents the input for a request operation.
type DeleteFleetInput struct {

	// A unique identifier for the fleet to be deleted. You can use either the fleet ID
	// or ARN value.
	//
	// This member is required.
	FleetId *string

	noSmithyDocumentSerde
}

type DeleteFleetOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDeleteFleetMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDeleteFleet{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDeleteFleet{}, middleware.After)
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
	if err = addOpDeleteFleetValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDeleteFleet(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDeleteFleet(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "gamelift",
		OperationName: "DeleteFleet",
	}
}
