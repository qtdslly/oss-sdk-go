// Code generated by smithy-go-codegen DO NOT EDIT.

package drs

import (
	"context"
	awsmiddleware "oss-sdk-go/oss/middleware"
	"oss-sdk-go/oss/signer/v4"
	"oss-sdk-go/service/drs/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Disconnects a specific Source Server from Elastic Disaster Recovery. Data
// replication is stopped immediately. All AWS resources created by Elastic
// Disaster Recovery for enabling the replication of the Source Server will be
// terminated / deleted within 90 minutes. You cannot disconnect a Source Server if
// it has a Recovery Instance. If the agent on the Source Server has not been
// prevented from communicating with the Elastic Disaster Recovery service, then it
// will receive a command to uninstall itself (within approximately 10 minutes).
// The following properties of the SourceServer will be changed immediately:
// dataReplicationInfo.dataReplicationState will be set to DISCONNECTED; The
// totalStorageBytes property for each of dataReplicationInfo.replicatedDisks will
// be set to zero; dataReplicationInfo.lagDuration and
// dataReplicationInfo.lagDuration will be nullified.
func (c *Client) DisconnectSourceServer(ctx context.Context, params *DisconnectSourceServerInput, optFns ...func(*Options)) (*DisconnectSourceServerOutput, error) {
	if params == nil {
		params = &DisconnectSourceServerInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DisconnectSourceServer", params, optFns, c.addOperationDisconnectSourceServerMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DisconnectSourceServerOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DisconnectSourceServerInput struct {

	// The ID of the Source Server to disconnect.
	//
	// This member is required.
	SourceServerID *string

	noSmithyDocumentSerde
}

type DisconnectSourceServerOutput struct {

	// The ARN of the Source Server.
	Arn *string

	// The Data Replication Info of the Source Server.
	DataReplicationInfo *types.DataReplicationInfo

	// The status of the last recovery launch of this Source Server.
	LastLaunchResult types.LastLaunchResult

	// The lifecycle information of this Source Server.
	LifeCycle *types.LifeCycle

	// The ID of the Recovery Instance associated with this Source Server.
	RecoveryInstanceId *string

	// The source properties of the Source Server.
	SourceProperties *types.SourceProperties

	// The ID of the Source Server.
	SourceServerID *string

	// The staging area of the source server.
	StagingArea *types.StagingArea

	// The tags associated with the Source Server.
	Tags map[string]string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDisconnectSourceServerMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpDisconnectSourceServer{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpDisconnectSourceServer{}, middleware.After)
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
	if err = addOpDisconnectSourceServerValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDisconnectSourceServer(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDisconnectSourceServer(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "drs",
		OperationName: "DisconnectSourceServer",
	}
}