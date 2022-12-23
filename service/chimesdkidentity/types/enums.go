// Code generated by smithy-go-codegen DO NOT EDIT.

package types

type AllowMessages string

// Enum values for AllowMessages
const (
	AllowMessagesAll  AllowMessages = "ALL"
	AllowMessagesNone AllowMessages = "NONE"
)

// Values returns all known values for AllowMessages. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (AllowMessages) Values() []AllowMessages {
	return []AllowMessages{
		"ALL",
		"NONE",
	}
}

type AppInstanceUserEndpointType string

// Enum values for AppInstanceUserEndpointType
const (
	AppInstanceUserEndpointTypeApns        AppInstanceUserEndpointType = "APNS"
	AppInstanceUserEndpointTypeApnsSandbox AppInstanceUserEndpointType = "APNS_SANDBOX"
	AppInstanceUserEndpointTypeGcm         AppInstanceUserEndpointType = "GCM"
)

// Values returns all known values for AppInstanceUserEndpointType. Note that this
// can be expanded in the future, and so it is only as up to date as the client.
// The ordering of this slice is not guaranteed to be stable across updates.
func (AppInstanceUserEndpointType) Values() []AppInstanceUserEndpointType {
	return []AppInstanceUserEndpointType{
		"APNS",
		"APNS_SANDBOX",
		"GCM",
	}
}

type EndpointStatus string

// Enum values for EndpointStatus
const (
	EndpointStatusActive   EndpointStatus = "ACTIVE"
	EndpointStatusInactive EndpointStatus = "INACTIVE"
)

// Values returns all known values for EndpointStatus. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (EndpointStatus) Values() []EndpointStatus {
	return []EndpointStatus{
		"ACTIVE",
		"INACTIVE",
	}
}

type EndpointStatusReason string

// Enum values for EndpointStatusReason
const (
	EndpointStatusReasonInvalidDeviceToken EndpointStatusReason = "INVALID_DEVICE_TOKEN"
	EndpointStatusReasonInvalidPinpointArn EndpointStatusReason = "INVALID_PINPOINT_ARN"
)

// Values returns all known values for EndpointStatusReason. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (EndpointStatusReason) Values() []EndpointStatusReason {
	return []EndpointStatusReason{
		"INVALID_DEVICE_TOKEN",
		"INVALID_PINPOINT_ARN",
	}
}

type ErrorCode string

// Enum values for ErrorCode
const (
	ErrorCodeBadRequest                           ErrorCode = "BadRequest"
	ErrorCodeConflict                             ErrorCode = "Conflict"
	ErrorCodeForbidden                            ErrorCode = "Forbidden"
	ErrorCodeNotFound                             ErrorCode = "NotFound"
	ErrorCodePreconditionFailed                   ErrorCode = "PreconditionFailed"
	ErrorCodeResourceLimitExceeded                ErrorCode = "ResourceLimitExceeded"
	ErrorCodeServiceFailure                       ErrorCode = "ServiceFailure"
	ErrorCodeAccessDenied                         ErrorCode = "AccessDenied"
	ErrorCodeServiceUnavailable                   ErrorCode = "ServiceUnavailable"
	ErrorCodeThrottled                            ErrorCode = "Throttled"
	ErrorCodeThrottling                           ErrorCode = "Throttling"
	ErrorCodeUnauthorized                         ErrorCode = "Unauthorized"
	ErrorCodeUnprocessable                        ErrorCode = "Unprocessable"
	ErrorCodeVoiceConnectorGroupAssociationsExist ErrorCode = "VoiceConnectorGroupAssociationsExist"
	ErrorCodePhoneNumberAssociationsExist         ErrorCode = "PhoneNumberAssociationsExist"
)

// Values returns all known values for ErrorCode. Note that this can be expanded in
// the future, and so it is only as up to date as the client. The ordering of this
// slice is not guaranteed to be stable across updates.
func (ErrorCode) Values() []ErrorCode {
	return []ErrorCode{
		"BadRequest",
		"Conflict",
		"Forbidden",
		"NotFound",
		"PreconditionFailed",
		"ResourceLimitExceeded",
		"ServiceFailure",
		"AccessDenied",
		"ServiceUnavailable",
		"Throttled",
		"Throttling",
		"Unauthorized",
		"Unprocessable",
		"VoiceConnectorGroupAssociationsExist",
		"PhoneNumberAssociationsExist",
	}
}
