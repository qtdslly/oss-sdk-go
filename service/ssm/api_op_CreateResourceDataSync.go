// Code generated by smithy-go-codegen DO NOT EDIT.

package ssm

import (
	"context"
	awsmiddleware "oss-sdk-go/oss/middleware"
	"oss-sdk-go/oss/signer/v4"
	"oss-sdk-go/service/ssm/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// A resource data sync helps you view data from multiple sources in a single
// location. Amazon Web Services Systems Manager offers two types of resource data
// sync: SyncToDestination and SyncFromSource. You can configure Systems Manager
// Inventory to use the SyncToDestination type to synchronize Inventory data from
// multiple Amazon Web Services Regions to a single Amazon Simple Storage Service
// (Amazon S3) bucket. For more information, see Configuring resource data sync for
// Inventory
// (https://docs.aws.amazon.com/systems-manager/latest/userguide/sysman-inventory-datasync.html)
// in the Amazon Web Services Systems Manager User Guide. You can configure Systems
// Manager Explorer to use the SyncFromSource type to synchronize operational work
// items (OpsItems) and operational data (OpsData) from multiple Amazon Web
// Services Regions to a single Amazon S3 bucket. This type can synchronize
// OpsItems and OpsData from multiple Amazon Web Services accounts and Amazon Web
// Services Regions or EntireOrganization by using Organizations. For more
// information, see Setting up Systems Manager Explorer to display data from
// multiple accounts and Regions
// (https://docs.aws.amazon.com/systems-manager/latest/userguide/Explorer-resource-data-sync.html)
// in the Amazon Web Services Systems Manager User Guide. A resource data sync is
// an asynchronous operation that returns immediately. After a successful initial
// sync is completed, the system continuously syncs data. To check the status of a
// sync, use the ListResourceDataSync. By default, data isn't encrypted in Amazon
// S3. We strongly recommend that you enable encryption in Amazon S3 to ensure
// secure data storage. We also recommend that you secure access to the Amazon S3
// bucket by creating a restrictive bucket policy.
func (c *Client) CreateResourceDataSync(ctx context.Context, params *CreateResourceDataSyncInput, optFns ...func(*Options)) (*CreateResourceDataSyncOutput, error) {
	if params == nil {
		params = &CreateResourceDataSyncInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateResourceDataSync", params, optFns, c.addOperationCreateResourceDataSyncMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateResourceDataSyncOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateResourceDataSyncInput struct {

	// A name for the configuration.
	//
	// This member is required.
	SyncName *string

	// Amazon S3 configuration details for the sync. This parameter is required if the
	// SyncType value is SyncToDestination.
	S3Destination *types.ResourceDataSyncS3Destination

	// Specify information about the data sources to synchronize. This parameter is
	// required if the SyncType value is SyncFromSource.
	SyncSource *types.ResourceDataSyncSource

	// Specify SyncToDestination to create a resource data sync that synchronizes data
	// to an S3 bucket for Inventory. If you specify SyncToDestination, you must
	// provide a value for S3Destination. Specify SyncFromSource to synchronize data
	// from a single account and multiple Regions, or multiple Amazon Web Services
	// accounts and Amazon Web Services Regions, as listed in Organizations for
	// Explorer. If you specify SyncFromSource, you must provide a value for
	// SyncSource. The default value is SyncToDestination.
	SyncType *string

	noSmithyDocumentSerde
}

type CreateResourceDataSyncOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateResourceDataSyncMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpCreateResourceDataSync{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpCreateResourceDataSync{}, middleware.After)
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
	if err = addOpCreateResourceDataSyncValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateResourceDataSync(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateResourceDataSync(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ssm",
		OperationName: "CreateResourceDataSync",
	}
}
