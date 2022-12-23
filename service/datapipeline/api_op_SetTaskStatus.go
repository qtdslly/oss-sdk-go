// Code generated by smithy-go-codegen DO NOT EDIT.

package datapipeline

import (
	"context"
	awsmiddleware "oss-sdk-go/oss/middleware"
	"oss-sdk-go/oss/signer/v4"
	"oss-sdk-go/service/datapipeline/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Task runners call SetTaskStatus to notify AWS Data Pipeline that a task is
// completed and provide information about the final status. A task runner makes
// this call regardless of whether the task was sucessful. A task runner does not
// need to call SetTaskStatus for tasks that are canceled by the web service during
// a call to ReportTaskProgress. POST / HTTP/1.1 Content-Type:
// application/x-amz-json-1.1 X-Amz-Target: DataPipeline.SetTaskStatus
// Content-Length: 847 Host: datapipeline.us-east-1.amazonaws.com X-Amz-Date: Mon,
// 12 Nov 2012 17:49:52 GMT Authorization: AuthParams {"taskId":
// "aaGgHT4LuH0T0Y0oLrJRjas5qH0d8cDPADxqq3tn+zCWGELkCdV2JprLreXm1oxeP5EFZHFLJ69kjSsLYE0iYHYBYVGBrB+E/pYq7ANEEeGJFnSBMRiXZVA+8UJ3OzcInvXeinqBmBaKwii7hnnKb/AXjXiNTXyxgydX1KAyg1AxkwBYG4cfPYMZbuEbQJFJvv5C/2+GVXz1w94nKYTeUeepwUOFOuRLS6JVtZoYwpF56E+Yfk1IcGpFOvCZ01B4Bkuu7x3J+MD/j6kJgZLAgbCJQtI3eiW3kdGmX0p0I2BdY1ZsX6b4UiSvM3OMj6NEHJCJL4E0ZfitnhCoe24Kvjo6C2hFbZq+ei/HPgSXBQMSagkr4vS9c0ChzxH2+LNYvec6bY4kymkaZI1dvOzmpa0FcnGf5AjSK4GpsViZ/ujz6zxFv81qBXzjF0/4M1775rjV1VUdyKaixiA/sJiACNezqZqETidp8d24BDPRhGsj6pBCrnelqGFrk/gXEXUsJ+xwMifRC8UVwiKekpAvHUywVk7Ku4jH/n3i2VoLRP6FXwpUbelu34iiZ9czpXyLtyPKwxa87dlrnRVURwkcVjOt2Mcrcaqe+cbWHvNRhyrPkkdfSF3ac8/wfgVbXvLEB2k9mKc67aD9rvdc1PKX09Tk8BKklsMTpZ3TRCd4NzQlJKigMe8Jat9+1tKj4Ole5ZzW6uyTu2s2iFjEV8KXu4MaiRJyNKCdKeGhhZWY37Qk4NBK4Ppgu+C6Y41dpfOh288SLDEVx0/UySlqOEdhba7c6BiPp5r3hKj3mk9lFy5OYp1aoGLeeFmjXveTnPdf2gkWqXXg7AUbJ7jEs1F0lKZQg4szep2gcKyAJXgvXLfJJHcha8Lfb/Ee7wYmyOcAaRpDBoFNSbtoVXar46teIrpho+ZDvynUXvU0grHWGOk=:wn3SgymHZM99bEXAMPLE",
// "taskStatus": "FINISHED"} x-amzn-RequestId: 8c8deb53-0788-11e2-af9c-6bc7a6be6qr8
// Content-Type: application/x-amz-json-1.1 Content-Length: 0 Date: Mon, 12 Nov
// 2012 17:50:53 GMT {}
func (c *Client) SetTaskStatus(ctx context.Context, params *SetTaskStatusInput, optFns ...func(*Options)) (*SetTaskStatusOutput, error) {
	if params == nil {
		params = &SetTaskStatusInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "SetTaskStatus", params, optFns, c.addOperationSetTaskStatusMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*SetTaskStatusOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Contains the parameters for SetTaskStatus.
type SetTaskStatusInput struct {

	// The ID of the task assigned to the task runner. This value is provided in the
	// response for PollForTask.
	//
	// This member is required.
	TaskId *string

	// If FINISHED, the task successfully completed. If FAILED, the task ended
	// unsuccessfully. Preconditions use false.
	//
	// This member is required.
	TaskStatus types.TaskStatus

	// If an error occurred during the task, this value specifies the error code. This
	// value is set on the physical attempt object. It is used to display error
	// information to the user. It should not start with string "Service_" which is
	// reserved by the system.
	ErrorId *string

	// If an error occurred during the task, this value specifies a text description of
	// the error. This value is set on the physical attempt object. It is used to
	// display error information to the user. The web service does not parse this
	// value.
	ErrorMessage *string

	// If an error occurred during the task, this value specifies the stack trace
	// associated with the error. This value is set on the physical attempt object. It
	// is used to display error information to the user. The web service does not parse
	// this value.
	ErrorStackTrace *string

	noSmithyDocumentSerde
}

// Contains the output of SetTaskStatus.
type SetTaskStatusOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationSetTaskStatusMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpSetTaskStatus{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpSetTaskStatus{}, middleware.After)
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
	if err = addOpSetTaskStatusValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opSetTaskStatus(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opSetTaskStatus(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "datapipeline",
		OperationName: "SetTaskStatus",
	}
}
