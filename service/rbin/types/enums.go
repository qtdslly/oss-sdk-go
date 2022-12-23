// Code generated by smithy-go-codegen DO NOT EDIT.

package types

type ResourceNotFoundExceptionReason string

// Enum values for ResourceNotFoundExceptionReason
const (
	ResourceNotFoundExceptionReasonRuleNotFound ResourceNotFoundExceptionReason = "RULE_NOT_FOUND"
)

// Values returns all known values for ResourceNotFoundExceptionReason. Note that
// this can be expanded in the future, and so it is only as up to date as the
// client. The ordering of this slice is not guaranteed to be stable across
// updates.
func (ResourceNotFoundExceptionReason) Values() []ResourceNotFoundExceptionReason {
	return []ResourceNotFoundExceptionReason{
		"RULE_NOT_FOUND",
	}
}

type ResourceType string

// Enum values for ResourceType
const (
	ResourceTypeEbsSnapshot ResourceType = "EBS_SNAPSHOT"
	ResourceTypeEc2Image    ResourceType = "EC2_IMAGE"
)

// Values returns all known values for ResourceType. Note that this can be expanded
// in the future, and so it is only as up to date as the client. The ordering of
// this slice is not guaranteed to be stable across updates.
func (ResourceType) Values() []ResourceType {
	return []ResourceType{
		"EBS_SNAPSHOT",
		"EC2_IMAGE",
	}
}

type RetentionPeriodUnit string

// Enum values for RetentionPeriodUnit
const (
	RetentionPeriodUnitDays RetentionPeriodUnit = "DAYS"
)

// Values returns all known values for RetentionPeriodUnit. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (RetentionPeriodUnit) Values() []RetentionPeriodUnit {
	return []RetentionPeriodUnit{
		"DAYS",
	}
}

type RuleStatus string

// Enum values for RuleStatus
const (
	RuleStatusPending   RuleStatus = "pending"
	RuleStatusAvailable RuleStatus = "available"
)

// Values returns all known values for RuleStatus. Note that this can be expanded
// in the future, and so it is only as up to date as the client. The ordering of
// this slice is not guaranteed to be stable across updates.
func (RuleStatus) Values() []RuleStatus {
	return []RuleStatus{
		"pending",
		"available",
	}
}

type ServiceQuotaExceededExceptionReason string

// Enum values for ServiceQuotaExceededExceptionReason
const (
	ServiceQuotaExceededExceptionReasonServiceQuotaExceeded ServiceQuotaExceededExceptionReason = "SERVICE_QUOTA_EXCEEDED"
)

// Values returns all known values for ServiceQuotaExceededExceptionReason. Note
// that this can be expanded in the future, and so it is only as up to date as the
// client. The ordering of this slice is not guaranteed to be stable across
// updates.
func (ServiceQuotaExceededExceptionReason) Values() []ServiceQuotaExceededExceptionReason {
	return []ServiceQuotaExceededExceptionReason{
		"SERVICE_QUOTA_EXCEEDED",
	}
}

type ValidationExceptionReason string

// Enum values for ValidationExceptionReason
const (
	ValidationExceptionReasonInvalidPageToken      ValidationExceptionReason = "INVALID_PAGE_TOKEN"
	ValidationExceptionReasonInvalidParameterValue ValidationExceptionReason = "INVALID_PARAMETER_VALUE"
)

// Values returns all known values for ValidationExceptionReason. Note that this
// can be expanded in the future, and so it is only as up to date as the client.
// The ordering of this slice is not guaranteed to be stable across updates.
func (ValidationExceptionReason) Values() []ValidationExceptionReason {
	return []ValidationExceptionReason{
		"INVALID_PAGE_TOKEN",
		"INVALID_PARAMETER_VALUE",
	}
}
