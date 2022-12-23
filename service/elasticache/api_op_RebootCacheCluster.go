// Code generated by smithy-go-codegen DO NOT EDIT.

package elasticache

import (
	"context"
	awsmiddleware "oss-sdk-go/oss/middleware"
	"oss-sdk-go/oss/signer/v4"
	"oss-sdk-go/service/elasticache/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Reboots some, or all, of the cache nodes within a provisioned cluster. This
// operation applies any modified cache parameter groups to the cluster. The reboot
// operation takes place as soon as possible, and results in a momentary outage to
// the cluster. During the reboot, the cluster status is set to REBOOTING. The
// reboot causes the contents of the cache (for each cache node being rebooted) to
// be lost. When the reboot is complete, a cluster event is created. Rebooting a
// cluster is currently supported on Memcached and Redis (cluster mode disabled)
// clusters. Rebooting is not supported on Redis (cluster mode enabled) clusters.
// If you make changes to parameters that require a Redis (cluster mode enabled)
// cluster reboot for the changes to be applied, see Rebooting a Cluster
// (http://docs.aws.amazon.com/AmazonElastiCache/latest/red-ug/nodes.rebooting.html)
// for an alternate process.
func (c *Client) RebootCacheCluster(ctx context.Context, params *RebootCacheClusterInput, optFns ...func(*Options)) (*RebootCacheClusterOutput, error) {
	if params == nil {
		params = &RebootCacheClusterInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "RebootCacheCluster", params, optFns, c.addOperationRebootCacheClusterMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*RebootCacheClusterOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Represents the input of a RebootCacheCluster operation.
type RebootCacheClusterInput struct {

	// The cluster identifier. This parameter is stored as a lowercase string.
	//
	// This member is required.
	CacheClusterId *string

	// A list of cache node IDs to reboot. A node ID is a numeric identifier (0001,
	// 0002, etc.). To reboot an entire cluster, specify all of the cache node IDs.
	//
	// This member is required.
	CacheNodeIdsToReboot []string

	noSmithyDocumentSerde
}

type RebootCacheClusterOutput struct {

	// Contains all of the attributes of a specific cluster.
	CacheCluster *types.CacheCluster

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationRebootCacheClusterMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsquery_serializeOpRebootCacheCluster{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpRebootCacheCluster{}, middleware.After)
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
	if err = addOpRebootCacheClusterValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opRebootCacheCluster(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opRebootCacheCluster(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "elasticache",
		OperationName: "RebootCacheCluster",
	}
}
