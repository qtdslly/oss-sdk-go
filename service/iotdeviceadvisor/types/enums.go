// Code generated by smithy-go-codegen DO NOT EDIT.

package types

type Protocol string

// Enum values for Protocol
const (
	ProtocolMqttV311 Protocol = "MqttV3_1_1"
	ProtocolMqttV5   Protocol = "MqttV5"
)

// Values returns all known values for Protocol. Note that this can be expanded in
// the future, and so it is only as up to date as the client. The ordering of this
// slice is not guaranteed to be stable across updates.
func (Protocol) Values() []Protocol {
	return []Protocol{
		"MqttV3_1_1",
		"MqttV5",
	}
}

type Status string

// Enum values for Status
const (
	StatusPass             Status = "PASS"
	StatusFail             Status = "FAIL"
	StatusCanceled         Status = "CANCELED"
	StatusPending          Status = "PENDING"
	StatusRunning          Status = "RUNNING"
	StatusStopping         Status = "STOPPING"
	StatusStopped          Status = "STOPPED"
	StatusPassWithWarnings Status = "PASS_WITH_WARNINGS"
	StatusError            Status = "ERROR"
)

// Values returns all known values for Status. Note that this can be expanded in
// the future, and so it is only as up to date as the client. The ordering of this
// slice is not guaranteed to be stable across updates.
func (Status) Values() []Status {
	return []Status{
		"PASS",
		"FAIL",
		"CANCELED",
		"PENDING",
		"RUNNING",
		"STOPPING",
		"STOPPED",
		"PASS_WITH_WARNINGS",
		"ERROR",
	}
}

type SuiteRunStatus string

// Enum values for SuiteRunStatus
const (
	SuiteRunStatusPass             SuiteRunStatus = "PASS"
	SuiteRunStatusFail             SuiteRunStatus = "FAIL"
	SuiteRunStatusCanceled         SuiteRunStatus = "CANCELED"
	SuiteRunStatusPending          SuiteRunStatus = "PENDING"
	SuiteRunStatusRunning          SuiteRunStatus = "RUNNING"
	SuiteRunStatusStopping         SuiteRunStatus = "STOPPING"
	SuiteRunStatusStopped          SuiteRunStatus = "STOPPED"
	SuiteRunStatusPassWithWarnings SuiteRunStatus = "PASS_WITH_WARNINGS"
	SuiteRunStatusError            SuiteRunStatus = "ERROR"
)

// Values returns all known values for SuiteRunStatus. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (SuiteRunStatus) Values() []SuiteRunStatus {
	return []SuiteRunStatus{
		"PASS",
		"FAIL",
		"CANCELED",
		"PENDING",
		"RUNNING",
		"STOPPING",
		"STOPPED",
		"PASS_WITH_WARNINGS",
		"ERROR",
	}
}

type TestCaseScenarioStatus string

// Enum values for TestCaseScenarioStatus
const (
	TestCaseScenarioStatusPass             TestCaseScenarioStatus = "PASS"
	TestCaseScenarioStatusFail             TestCaseScenarioStatus = "FAIL"
	TestCaseScenarioStatusCanceled         TestCaseScenarioStatus = "CANCELED"
	TestCaseScenarioStatusPending          TestCaseScenarioStatus = "PENDING"
	TestCaseScenarioStatusRunning          TestCaseScenarioStatus = "RUNNING"
	TestCaseScenarioStatusStopping         TestCaseScenarioStatus = "STOPPING"
	TestCaseScenarioStatusStopped          TestCaseScenarioStatus = "STOPPED"
	TestCaseScenarioStatusPassWithWarnings TestCaseScenarioStatus = "PASS_WITH_WARNINGS"
	TestCaseScenarioStatusError            TestCaseScenarioStatus = "ERROR"
)

// Values returns all known values for TestCaseScenarioStatus. Note that this can
// be expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (TestCaseScenarioStatus) Values() []TestCaseScenarioStatus {
	return []TestCaseScenarioStatus{
		"PASS",
		"FAIL",
		"CANCELED",
		"PENDING",
		"RUNNING",
		"STOPPING",
		"STOPPED",
		"PASS_WITH_WARNINGS",
		"ERROR",
	}
}

type TestCaseScenarioType string

// Enum values for TestCaseScenarioType
const (
	TestCaseScenarioTypeAdvanced TestCaseScenarioType = "Advanced"
	TestCaseScenarioTypeBasic    TestCaseScenarioType = "Basic"
)

// Values returns all known values for TestCaseScenarioType. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (TestCaseScenarioType) Values() []TestCaseScenarioType {
	return []TestCaseScenarioType{
		"Advanced",
		"Basic",
	}
}
