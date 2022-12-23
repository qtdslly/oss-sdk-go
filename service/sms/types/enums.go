// Code generated by smithy-go-codegen DO NOT EDIT.

package types

type AppLaunchConfigurationStatus string

// Enum values for AppLaunchConfigurationStatus
const (
	AppLaunchConfigurationStatusNotConfigured AppLaunchConfigurationStatus = "NOT_CONFIGURED"
	AppLaunchConfigurationStatusConfigured    AppLaunchConfigurationStatus = "CONFIGURED"
)

// Values returns all known values for AppLaunchConfigurationStatus. Note that this
// can be expanded in the future, and so it is only as up to date as the client.
// The ordering of this slice is not guaranteed to be stable across updates.
func (AppLaunchConfigurationStatus) Values() []AppLaunchConfigurationStatus {
	return []AppLaunchConfigurationStatus{
		"NOT_CONFIGURED",
		"CONFIGURED",
	}
}

type AppLaunchStatus string

// Enum values for AppLaunchStatus
const (
	AppLaunchStatusReadyForConfiguration   AppLaunchStatus = "READY_FOR_CONFIGURATION"
	AppLaunchStatusConfigurationInProgress AppLaunchStatus = "CONFIGURATION_IN_PROGRESS"
	AppLaunchStatusConfigurationInvalid    AppLaunchStatus = "CONFIGURATION_INVALID"
	AppLaunchStatusReadyForLaunch          AppLaunchStatus = "READY_FOR_LAUNCH"
	AppLaunchStatusValidationInProgress    AppLaunchStatus = "VALIDATION_IN_PROGRESS"
	AppLaunchStatusLaunchPending           AppLaunchStatus = "LAUNCH_PENDING"
	AppLaunchStatusLaunchInProgress        AppLaunchStatus = "LAUNCH_IN_PROGRESS"
	AppLaunchStatusLaunched                AppLaunchStatus = "LAUNCHED"
	AppLaunchStatusPartiallyLaunched       AppLaunchStatus = "PARTIALLY_LAUNCHED"
	AppLaunchStatusDeltaLaunchInProgress   AppLaunchStatus = "DELTA_LAUNCH_IN_PROGRESS"
	AppLaunchStatusDeltaLaunchFailed       AppLaunchStatus = "DELTA_LAUNCH_FAILED"
	AppLaunchStatusLaunchFailed            AppLaunchStatus = "LAUNCH_FAILED"
	AppLaunchStatusTerminateInProgress     AppLaunchStatus = "TERMINATE_IN_PROGRESS"
	AppLaunchStatusTerminateFailed         AppLaunchStatus = "TERMINATE_FAILED"
	AppLaunchStatusTerminated              AppLaunchStatus = "TERMINATED"
)

// Values returns all known values for AppLaunchStatus. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (AppLaunchStatus) Values() []AppLaunchStatus {
	return []AppLaunchStatus{
		"READY_FOR_CONFIGURATION",
		"CONFIGURATION_IN_PROGRESS",
		"CONFIGURATION_INVALID",
		"READY_FOR_LAUNCH",
		"VALIDATION_IN_PROGRESS",
		"LAUNCH_PENDING",
		"LAUNCH_IN_PROGRESS",
		"LAUNCHED",
		"PARTIALLY_LAUNCHED",
		"DELTA_LAUNCH_IN_PROGRESS",
		"DELTA_LAUNCH_FAILED",
		"LAUNCH_FAILED",
		"TERMINATE_IN_PROGRESS",
		"TERMINATE_FAILED",
		"TERMINATED",
	}
}

type AppReplicationConfigurationStatus string

// Enum values for AppReplicationConfigurationStatus
const (
	AppReplicationConfigurationStatusNotConfigured AppReplicationConfigurationStatus = "NOT_CONFIGURED"
	AppReplicationConfigurationStatusConfigured    AppReplicationConfigurationStatus = "CONFIGURED"
)

// Values returns all known values for AppReplicationConfigurationStatus. Note that
// this can be expanded in the future, and so it is only as up to date as the
// client. The ordering of this slice is not guaranteed to be stable across
// updates.
func (AppReplicationConfigurationStatus) Values() []AppReplicationConfigurationStatus {
	return []AppReplicationConfigurationStatus{
		"NOT_CONFIGURED",
		"CONFIGURED",
	}
}

type AppReplicationStatus string

// Enum values for AppReplicationStatus
const (
	AppReplicationStatusReadyForConfiguration      AppReplicationStatus = "READY_FOR_CONFIGURATION"
	AppReplicationStatusConfigurationInProgress    AppReplicationStatus = "CONFIGURATION_IN_PROGRESS"
	AppReplicationStatusConfigurationInvalid       AppReplicationStatus = "CONFIGURATION_INVALID"
	AppReplicationStatusReadyForReplication        AppReplicationStatus = "READY_FOR_REPLICATION"
	AppReplicationStatusValidationInProgress       AppReplicationStatus = "VALIDATION_IN_PROGRESS"
	AppReplicationStatusReplicationPending         AppReplicationStatus = "REPLICATION_PENDING"
	AppReplicationStatusReplicationInProgress      AppReplicationStatus = "REPLICATION_IN_PROGRESS"
	AppReplicationStatusReplicated                 AppReplicationStatus = "REPLICATED"
	AppReplicationStatusPartiallyReplicated        AppReplicationStatus = "PARTIALLY_REPLICATED"
	AppReplicationStatusDeltaReplicationInProgress AppReplicationStatus = "DELTA_REPLICATION_IN_PROGRESS"
	AppReplicationStatusDeltaReplicated            AppReplicationStatus = "DELTA_REPLICATED"
	AppReplicationStatusDeltaReplicationFailed     AppReplicationStatus = "DELTA_REPLICATION_FAILED"
	AppReplicationStatusReplicationFailed          AppReplicationStatus = "REPLICATION_FAILED"
	AppReplicationStatusReplicationStopping        AppReplicationStatus = "REPLICATION_STOPPING"
	AppReplicationStatusReplicationStopFailed      AppReplicationStatus = "REPLICATION_STOP_FAILED"
	AppReplicationStatusReplicationStopped         AppReplicationStatus = "REPLICATION_STOPPED"
)

// Values returns all known values for AppReplicationStatus. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (AppReplicationStatus) Values() []AppReplicationStatus {
	return []AppReplicationStatus{
		"READY_FOR_CONFIGURATION",
		"CONFIGURATION_IN_PROGRESS",
		"CONFIGURATION_INVALID",
		"READY_FOR_REPLICATION",
		"VALIDATION_IN_PROGRESS",
		"REPLICATION_PENDING",
		"REPLICATION_IN_PROGRESS",
		"REPLICATED",
		"PARTIALLY_REPLICATED",
		"DELTA_REPLICATION_IN_PROGRESS",
		"DELTA_REPLICATED",
		"DELTA_REPLICATION_FAILED",
		"REPLICATION_FAILED",
		"REPLICATION_STOPPING",
		"REPLICATION_STOP_FAILED",
		"REPLICATION_STOPPED",
	}
}

type AppStatus string

// Enum values for AppStatus
const (
	AppStatusCreating     AppStatus = "CREATING"
	AppStatusActive       AppStatus = "ACTIVE"
	AppStatusUpdating     AppStatus = "UPDATING"
	AppStatusDeleting     AppStatus = "DELETING"
	AppStatusDeleted      AppStatus = "DELETED"
	AppStatusDeleteFailed AppStatus = "DELETE_FAILED"
)

// Values returns all known values for AppStatus. Note that this can be expanded in
// the future, and so it is only as up to date as the client. The ordering of this
// slice is not guaranteed to be stable across updates.
func (AppStatus) Values() []AppStatus {
	return []AppStatus{
		"CREATING",
		"ACTIVE",
		"UPDATING",
		"DELETING",
		"DELETED",
		"DELETE_FAILED",
	}
}

type AppValidationStrategy string

// Enum values for AppValidationStrategy
const (
	AppValidationStrategySsm AppValidationStrategy = "SSM"
)

// Values returns all known values for AppValidationStrategy. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (AppValidationStrategy) Values() []AppValidationStrategy {
	return []AppValidationStrategy{
		"SSM",
	}
}

type ConnectorCapability string

// Enum values for ConnectorCapability
const (
	ConnectorCapabilityVSphere          ConnectorCapability = "VSPHERE"
	ConnectorCapabilityScvmm            ConnectorCapability = "SCVMM"
	ConnectorCapabilityHyperVManager    ConnectorCapability = "HYPERV-MANAGER"
	ConnectorCapabilitySnapshotBatching ConnectorCapability = "SNAPSHOT_BATCHING"
	ConnectorCapabilitySmsOptimized     ConnectorCapability = "SMS_OPTIMIZED"
)

// Values returns all known values for ConnectorCapability. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (ConnectorCapability) Values() []ConnectorCapability {
	return []ConnectorCapability{
		"VSPHERE",
		"SCVMM",
		"HYPERV-MANAGER",
		"SNAPSHOT_BATCHING",
		"SMS_OPTIMIZED",
	}
}

type ConnectorStatus string

// Enum values for ConnectorStatus
const (
	ConnectorStatusHealthy   ConnectorStatus = "HEALTHY"
	ConnectorStatusUnhealthy ConnectorStatus = "UNHEALTHY"
)

// Values returns all known values for ConnectorStatus. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (ConnectorStatus) Values() []ConnectorStatus {
	return []ConnectorStatus{
		"HEALTHY",
		"UNHEALTHY",
	}
}

type LicenseType string

// Enum values for LicenseType
const (
	LicenseTypeAws  LicenseType = "AWS"
	LicenseTypeByol LicenseType = "BYOL"
)

// Values returns all known values for LicenseType. Note that this can be expanded
// in the future, and so it is only as up to date as the client. The ordering of
// this slice is not guaranteed to be stable across updates.
func (LicenseType) Values() []LicenseType {
	return []LicenseType{
		"AWS",
		"BYOL",
	}
}

type OutputFormat string

// Enum values for OutputFormat
const (
	OutputFormatJson OutputFormat = "JSON"
	OutputFormatYaml OutputFormat = "YAML"
)

// Values returns all known values for OutputFormat. Note that this can be expanded
// in the future, and so it is only as up to date as the client. The ordering of
// this slice is not guaranteed to be stable across updates.
func (OutputFormat) Values() []OutputFormat {
	return []OutputFormat{
		"JSON",
		"YAML",
	}
}

type ReplicationJobState string

// Enum values for ReplicationJobState
const (
	ReplicationJobStatePending         ReplicationJobState = "PENDING"
	ReplicationJobStateActive          ReplicationJobState = "ACTIVE"
	ReplicationJobStateFailed          ReplicationJobState = "FAILED"
	ReplicationJobStateDeleting        ReplicationJobState = "DELETING"
	ReplicationJobStateDeleted         ReplicationJobState = "DELETED"
	ReplicationJobStateCompleted       ReplicationJobState = "COMPLETED"
	ReplicationJobStatePausedOnFailure ReplicationJobState = "PAUSED_ON_FAILURE"
	ReplicationJobStateFailing         ReplicationJobState = "FAILING"
)

// Values returns all known values for ReplicationJobState. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (ReplicationJobState) Values() []ReplicationJobState {
	return []ReplicationJobState{
		"PENDING",
		"ACTIVE",
		"FAILED",
		"DELETING",
		"DELETED",
		"COMPLETED",
		"PAUSED_ON_FAILURE",
		"FAILING",
	}
}

type ReplicationRunState string

// Enum values for ReplicationRunState
const (
	ReplicationRunStatePending   ReplicationRunState = "PENDING"
	ReplicationRunStateMissed    ReplicationRunState = "MISSED"
	ReplicationRunStateActive    ReplicationRunState = "ACTIVE"
	ReplicationRunStateFailed    ReplicationRunState = "FAILED"
	ReplicationRunStateCompleted ReplicationRunState = "COMPLETED"
	ReplicationRunStateDeleting  ReplicationRunState = "DELETING"
	ReplicationRunStateDeleted   ReplicationRunState = "DELETED"
)

// Values returns all known values for ReplicationRunState. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (ReplicationRunState) Values() []ReplicationRunState {
	return []ReplicationRunState{
		"PENDING",
		"MISSED",
		"ACTIVE",
		"FAILED",
		"COMPLETED",
		"DELETING",
		"DELETED",
	}
}

type ReplicationRunType string

// Enum values for ReplicationRunType
const (
	ReplicationRunTypeOnDemand  ReplicationRunType = "ON_DEMAND"
	ReplicationRunTypeAutomatic ReplicationRunType = "AUTOMATIC"
)

// Values returns all known values for ReplicationRunType. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (ReplicationRunType) Values() []ReplicationRunType {
	return []ReplicationRunType{
		"ON_DEMAND",
		"AUTOMATIC",
	}
}

type ScriptType string

// Enum values for ScriptType
const (
	ScriptTypeShellScript      ScriptType = "SHELL_SCRIPT"
	ScriptTypePowershellScript ScriptType = "POWERSHELL_SCRIPT"
)

// Values returns all known values for ScriptType. Note that this can be expanded
// in the future, and so it is only as up to date as the client. The ordering of
// this slice is not guaranteed to be stable across updates.
func (ScriptType) Values() []ScriptType {
	return []ScriptType{
		"SHELL_SCRIPT",
		"POWERSHELL_SCRIPT",
	}
}

type ServerCatalogStatus string

// Enum values for ServerCatalogStatus
const (
	ServerCatalogStatusNotImported ServerCatalogStatus = "NOT_IMPORTED"
	ServerCatalogStatusImporting   ServerCatalogStatus = "IMPORTING"
	ServerCatalogStatusAvailable   ServerCatalogStatus = "AVAILABLE"
	ServerCatalogStatusDeleted     ServerCatalogStatus = "DELETED"
	ServerCatalogStatusExpired     ServerCatalogStatus = "EXPIRED"
)

// Values returns all known values for ServerCatalogStatus. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (ServerCatalogStatus) Values() []ServerCatalogStatus {
	return []ServerCatalogStatus{
		"NOT_IMPORTED",
		"IMPORTING",
		"AVAILABLE",
		"DELETED",
		"EXPIRED",
	}
}

type ServerType string

// Enum values for ServerType
const (
	ServerTypeVirtualMachine ServerType = "VIRTUAL_MACHINE"
)

// Values returns all known values for ServerType. Note that this can be expanded
// in the future, and so it is only as up to date as the client. The ordering of
// this slice is not guaranteed to be stable across updates.
func (ServerType) Values() []ServerType {
	return []ServerType{
		"VIRTUAL_MACHINE",
	}
}

type ServerValidationStrategy string

// Enum values for ServerValidationStrategy
const (
	ServerValidationStrategyUserdata ServerValidationStrategy = "USERDATA"
)

// Values returns all known values for ServerValidationStrategy. Note that this can
// be expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (ServerValidationStrategy) Values() []ServerValidationStrategy {
	return []ServerValidationStrategy{
		"USERDATA",
	}
}

type ValidationStatus string

// Enum values for ValidationStatus
const (
	ValidationStatusReadyForValidation ValidationStatus = "READY_FOR_VALIDATION"
	ValidationStatusPending            ValidationStatus = "PENDING"
	ValidationStatusInProgress         ValidationStatus = "IN_PROGRESS"
	ValidationStatusSucceeded          ValidationStatus = "SUCCEEDED"
	ValidationStatusFailed             ValidationStatus = "FAILED"
)

// Values returns all known values for ValidationStatus. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (ValidationStatus) Values() []ValidationStatus {
	return []ValidationStatus{
		"READY_FOR_VALIDATION",
		"PENDING",
		"IN_PROGRESS",
		"SUCCEEDED",
		"FAILED",
	}
}

type VmManagerType string

// Enum values for VmManagerType
const (
	VmManagerTypeVSphere       VmManagerType = "VSPHERE"
	VmManagerTypeScvmm         VmManagerType = "SCVMM"
	VmManagerTypeHyperVManager VmManagerType = "HYPERV-MANAGER"
)

// Values returns all known values for VmManagerType. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (VmManagerType) Values() []VmManagerType {
	return []VmManagerType{
		"VSPHERE",
		"SCVMM",
		"HYPERV-MANAGER",
	}
}
