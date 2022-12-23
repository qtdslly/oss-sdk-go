// Code generated by smithy-go-codegen DO NOT EDIT.

package codeguruprofiler

import (
	"context"
	awsmiddleware "oss-sdk-go/oss/middleware"
	"oss-sdk-go/oss/signer/v4"
	"oss-sdk-go/service/codeguruprofiler/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Returns the time series of values for a requested list of frame metrics from a
// time period.
func (c *Client) BatchGetFrameMetricData(ctx context.Context, params *BatchGetFrameMetricDataInput, optFns ...func(*Options)) (*BatchGetFrameMetricDataOutput, error) {
	if params == nil {
		params = &BatchGetFrameMetricDataInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "BatchGetFrameMetricData", params, optFns, c.addOperationBatchGetFrameMetricDataMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*BatchGetFrameMetricDataOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// The structure representing the BatchGetFrameMetricDataRequest.
type BatchGetFrameMetricDataInput struct {

	// The name of the profiling group associated with the the frame metrics used to
	// return the time series values.
	//
	// This member is required.
	ProfilingGroupName *string

	// The end time of the time period for the returned time series values. This is
	// specified using the ISO 8601 format. For example, 2020-06-01T13:15:02.001Z
	// represents 1 millisecond past June 1, 2020 1:15:02 PM UTC.
	EndTime *time.Time

	// The details of the metrics that are used to request a time series of values. The
	// metric includes the name of the frame, the aggregation type to calculate the
	// metric value for the frame, and the thread states to use to get the count for
	// the metric value of the frame.
	FrameMetrics []types.FrameMetric

	// The duration of the frame metrics used to return the time series values. Specify
	// using the ISO 8601 format. The maximum period duration is one day (PT24H or
	// P1D).
	Period *string

	// The start time of the time period for the frame metrics used to return the time
	// series values. This is specified using the ISO 8601 format. For example,
	// 2020-06-01T13:15:02.001Z represents 1 millisecond past June 1, 2020 1:15:02 PM
	// UTC.
	StartTime *time.Time

	// The requested resolution of time steps for the returned time series of values.
	// If the requested target resolution is not available due to data not being
	// retained we provide a best effort result by falling back to the most granular
	// available resolution after the target resolution. There are 3 valid values.
	//
	// *
	// P1D — 1 day
	//
	// * PT1H — 1 hour
	//
	// * PT5M — 5 minutes
	TargetResolution types.AggregationPeriod

	noSmithyDocumentSerde
}

// The structure representing the BatchGetFrameMetricDataResponse.
type BatchGetFrameMetricDataOutput struct {

	// The end time of the time period for the returned time series values. This is
	// specified using the ISO 8601 format. For example, 2020-06-01T13:15:02.001Z
	// represents 1 millisecond past June 1, 2020 1:15:02 PM UTC.
	//
	// This member is required.
	EndTime *time.Time

	// List of instances, or time steps, in the time series. For example, if the period
	// is one day (PT24H)), and the resolution is five minutes (PT5M), then there are
	// 288 endTimes in the list that are each five minutes appart.
	//
	// This member is required.
	EndTimes []types.TimestampStructure

	// Details of the metrics to request a time series of values. The metric includes
	// the name of the frame, the aggregation type to calculate the metric value for
	// the frame, and the thread states to use to get the count for the metric value of
	// the frame.
	//
	// This member is required.
	FrameMetricData []types.FrameMetricDatum

	// Resolution or granularity of the profile data used to generate the time series.
	// This is the value used to jump through time steps in a time series. There are 3
	// valid values.
	//
	// * P1D — 1 day
	//
	// * PT1H — 1 hour
	//
	// * PT5M — 5 minutes
	//
	// This member is required.
	Resolution types.AggregationPeriod

	// The start time of the time period for the returned time series values. This is
	// specified using the ISO 8601 format. For example, 2020-06-01T13:15:02.001Z
	// represents 1 millisecond past June 1, 2020 1:15:02 PM UTC.
	//
	// This member is required.
	StartTime *time.Time

	// List of instances which remained unprocessed. This will create a missing time
	// step in the list of end times.
	//
	// This member is required.
	UnprocessedEndTimes map[string][]types.TimestampStructure

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationBatchGetFrameMetricDataMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpBatchGetFrameMetricData{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpBatchGetFrameMetricData{}, middleware.After)
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
	if err = addOpBatchGetFrameMetricDataValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opBatchGetFrameMetricData(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opBatchGetFrameMetricData(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "codeguru-profiler",
		OperationName: "BatchGetFrameMetricData",
	}
}
