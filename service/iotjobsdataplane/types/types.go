// Code generated by smithy-go-codegen DO NOT EDIT.

package types

import (
	smithydocument "github.com/aws/smithy-go/document"
)

// Contains data about a job execution.
type JobExecution struct {

	// The estimated number of seconds that remain before the job execution status will
	// be changed to TIMED_OUT.
	ApproximateSecondsBeforeTimedOut *int64

	// A number that identifies a particular job execution on a particular device. It
	// can be used later in commands that return or update job execution information.
	ExecutionNumber *int64

	// The content of the job document.
	JobDocument *string

	// The unique identifier you assigned to this job when it was created.
	JobId *string

	// The time, in milliseconds since the epoch, when the job execution was last
	// updated.
	LastUpdatedAt int64

	// The time, in milliseconds since the epoch, when the job execution was enqueued.
	QueuedAt int64

	// The time, in milliseconds since the epoch, when the job execution was started.
	StartedAt *int64

	// The status of the job execution. Can be one of: "QUEUED", "IN_PROGRESS",
	// "FAILED", "SUCCESS", "CANCELED", "REJECTED", or "REMOVED".
	Status JobExecutionStatus

	// A collection of name/value pairs that describe the status of the job execution.
	StatusDetails map[string]string

	// The name of the thing that is executing the job.
	ThingName *string

	// The version of the job execution. Job execution versions are incremented each
	// time they are updated by a device.
	VersionNumber int64

	noSmithyDocumentSerde
}

// Contains data about the state of a job execution.
type JobExecutionState struct {

	// The status of the job execution. Can be one of: "QUEUED", "IN_PROGRESS",
	// "FAILED", "SUCCESS", "CANCELED", "REJECTED", or "REMOVED".
	Status JobExecutionStatus

	// A collection of name/value pairs that describe the status of the job execution.
	StatusDetails map[string]string

	// The version of the job execution. Job execution versions are incremented each
	// time they are updated by a device.
	VersionNumber int64

	noSmithyDocumentSerde
}

// Contains a subset of information about a job execution.
type JobExecutionSummary struct {

	// A number that identifies a particular job execution on a particular device.
	ExecutionNumber *int64

	// The unique identifier you assigned to this job when it was created.
	JobId *string

	// The time, in milliseconds since the epoch, when the job execution was last
	// updated.
	LastUpdatedAt int64

	// The time, in milliseconds since the epoch, when the job execution was enqueued.
	QueuedAt int64

	// The time, in milliseconds since the epoch, when the job execution started.
	StartedAt *int64

	// The version of the job execution. Job execution versions are incremented each
	// time AWS IoT Jobs receives an update from a device.
	VersionNumber int64

	noSmithyDocumentSerde
}

type noSmithyDocumentSerde = smithydocument.NoSerde
