// Code generated by smithy-go-codegen DO NOT EDIT.

package ec2

import (
	"context"
	awsmiddleware "oss-sdk-go/oss/middleware"
	"oss-sdk-go/oss/signer/v4"
	"oss-sdk-go/service/ec2/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// You can modify several parameters of an existing EBS volume, including volume
// size, volume type, and IOPS capacity. If your EBS volume is attached to a
// current-generation EC2 instance type, you might be able to apply these changes
// without stopping the instance or detaching the volume from it. For more
// information about modifying EBS volumes, see Amazon EBS Elastic Volumes
// (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/ebs-modify-volume.html)
// (Linux instances) or Amazon EBS Elastic Volumes
// (https://docs.aws.amazon.com/AWSEC2/latest/WindowsGuide/ebs-modify-volume.html)
// (Windows instances). When you complete a resize operation on your volume, you
// need to extend the volume's file-system size to take advantage of the new
// storage capacity. For more information, see Extend a Linux file system
// (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/ebs-expand-volume.html#recognize-expanded-volume-linux)
// or Extend a Windows file system
// (https://docs.aws.amazon.com/AWSEC2/latest/WindowsGuide/ebs-expand-volume.html#recognize-expanded-volume-windows).
// You can use CloudWatch Events to check the status of a modification to an EBS
// volume. For information about CloudWatch Events, see the Amazon CloudWatch
// Events User Guide (https://docs.aws.amazon.com/AmazonCloudWatch/latest/events/).
// You can also track the status of a modification using
// DescribeVolumesModifications. For information about tracking status changes
// using either method, see Monitor the progress of volume modifications
// (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/monitoring-volume-modifications.html).
// With previous-generation instance types, resizing an EBS volume might require
// detaching and reattaching the volume or stopping and restarting the instance.
// After modifying a volume, you must wait at least six hours and ensure that the
// volume is in the in-use or available state before you can modify the same
// volume. This is sometimes referred to as a cooldown period.
func (c *Client) ModifyVolume(ctx context.Context, params *ModifyVolumeInput, optFns ...func(*Options)) (*ModifyVolumeOutput, error) {
	if params == nil {
		params = &ModifyVolumeInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ModifyVolume", params, optFns, c.addOperationModifyVolumeMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ModifyVolumeOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ModifyVolumeInput struct {

	// The ID of the volume.
	//
	// This member is required.
	VolumeId *string

	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have the
	// required permissions, the error response is DryRunOperation. Otherwise, it is
	// UnauthorizedOperation.
	DryRun *bool

	// The target IOPS rate of the volume. This parameter is valid only for gp3, io1,
	// and io2 volumes. The following are the supported values for each volume type:
	//
	// *
	// gp3: 3,000-16,000 IOPS
	//
	// * io1: 100-64,000 IOPS
	//
	// * io2: 100-64,000 IOPS
	//
	// Default:
	// The existing value is retained if you keep the same volume type. If you change
	// the volume type to io1, io2, or gp3, the default is 3,000.
	Iops *int32

	// Specifies whether to enable Amazon EBS Multi-Attach. If you enable Multi-Attach,
	// you can attach the volume to up to 16  Nitro-based instances
	// (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/instance-types.html#ec2-nitro-instances)
	// in the same Availability Zone. This parameter is supported with io1 and io2
	// volumes only. For more information, see  Amazon EBS Multi-Attach
	// (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/ebs-volumes-multi.html) in
	// the Amazon Elastic Compute Cloud User Guide.
	MultiAttachEnabled *bool

	// The target size of the volume, in GiB. The target volume size must be greater
	// than or equal to the existing size of the volume. The following are the
	// supported volumes sizes for each volume type:
	//
	// * gp2 and gp3: 1-16,384
	//
	// * io1
	// and io2: 4-16,384
	//
	// * st1 and sc1: 125-16,384
	//
	// * standard: 1-1,024
	//
	// Default: The
	// existing size is retained.
	Size *int32

	// The target throughput of the volume, in MiB/s. This parameter is valid only for
	// gp3 volumes. The maximum value is 1,000. Default: The existing value is retained
	// if the source and target volume type is gp3. Otherwise, the default value is
	// 125. Valid Range: Minimum value of 125. Maximum value of 1000.
	Throughput *int32

	// The target EBS volume type of the volume. For more information, see Amazon EBS
	// volume types
	// (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/EBSVolumeTypes.html) in the
	// Amazon Elastic Compute Cloud User Guide. Default: The existing type is retained.
	VolumeType types.VolumeType

	noSmithyDocumentSerde
}

type ModifyVolumeOutput struct {

	// Information about the volume modification.
	VolumeModification *types.VolumeModification

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationModifyVolumeMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsEc2query_serializeOpModifyVolume{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsEc2query_deserializeOpModifyVolume{}, middleware.After)
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
	if err = addOpModifyVolumeValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opModifyVolume(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opModifyVolume(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ec2",
		OperationName: "ModifyVolume",
	}
}