// Code generated by smithy-go-codegen DO NOT EDIT.

package types

import (
	smithydocument "github.com/aws/smithy-go/document"
)

// An assertion rule enforces that, when you change a routing control state, that
// the criteria that you set in the rule configuration is met. Otherwise, the
// change to the routing control is not accepted. For example, the criteria might
// be that at least one routing control state is On after the transation so that
// traffic continues to flow to at least one cell for the application. This ensures
// that you avoid a fail-open scenario.
type AssertionRule struct {

	// The routing controls that are part of transactions that are evaluated to
	// determine if a request to change a routing control state is allowed. For
	// example, you might include three routing controls, one for each of three Amazon
	// Web Services Regions.
	//
	// This member is required.
	AssertedControls []string

	// The Amazon Resource Name (ARN) of the control panel.
	//
	// This member is required.
	ControlPanelArn *string

	// Name of the assertion rule. You can use any non-white space character in the
	// name.
	//
	// This member is required.
	Name *string

	// The criteria that you set for specific assertion routing controls
	// (AssertedControls) that designate how many routing control states must be ON as
	// the result of a transaction. For example, if you have three assertion routing
	// controls, you might specify atleast 2 for your rule configuration. This means
	// that at least two assertion routing control states must be ON, so that at least
	// two Amazon Web Services Regions have traffic flowing to them.
	//
	// This member is required.
	RuleConfig *RuleConfig

	// The Amazon Resource Name (ARN) of the assertion rule.
	//
	// This member is required.
	SafetyRuleArn *string

	// The deployment status of an assertion rule. Status can be one of the following:
	// PENDING, DEPLOYED, PENDING_DELETION.
	//
	// This member is required.
	Status Status

	// An evaluation period, in milliseconds (ms), during which any request against the
	// target routing controls will fail. This helps prevent "flapping" of state. The
	// wait period is 5000 ms by default, but you can choose a custom value.
	//
	// This member is required.
	WaitPeriodMs int32

	noSmithyDocumentSerde
}

// An update to an assertion rule. You can update the name or the evaluation period
// (wait period). If you don't specify one of the items to update, the item is
// unchanged.
type AssertionRuleUpdate struct {

	// The name of the assertion rule. You can use any non-white space character in the
	// name.
	//
	// This member is required.
	Name *string

	// The Amazon Resource Name (ARN) of the assertion rule.
	//
	// This member is required.
	SafetyRuleArn *string

	// An evaluation period, in milliseconds (ms), during which any request against the
	// target routing controls will fail. This helps prevent "flapping" of state. The
	// wait period is 5000 ms by default, but you can choose a custom value.
	//
	// This member is required.
	WaitPeriodMs int32

	noSmithyDocumentSerde
}

// A set of five redundant Regional endpoints against which you can execute API
// calls to update or get the state of routing controls. You can host multiple
// control panels and routing controls on one cluster.
type Cluster struct {

	// The Amazon Resource Name (ARN) of the cluster.
	ClusterArn *string

	// Endpoints for a cluster. Specify one of these endpoints when you want to set or
	// retrieve a routing control state in the cluster. To get or update the routing
	// control state, see the Amazon Route 53 Application Recovery Controller Routing
	// Control Actions.
	ClusterEndpoints []ClusterEndpoint

	// The name of the cluster.
	Name *string

	// Deployment status of a resource. Status can be one of the following: PENDING,
	// DEPLOYED, PENDING_DELETION.
	Status Status

	noSmithyDocumentSerde
}

// A cluster endpoint. Specify an endpoint when you want to set or retrieve a
// routing control state in the cluster.
type ClusterEndpoint struct {

	// A cluster endpoint. Specify an endpoint and Amazon Web Services Region when you
	// want to set or retrieve a routing control state in the cluster. To get or update
	// the routing control state, see the Amazon Route 53 Application Recovery
	// Controller Routing Control Actions.
	Endpoint *string

	// The Amazon Web Services Region for a cluster endpoint.
	Region *string

	noSmithyDocumentSerde
}

// A control panel represents a group of routing controls that can be changed
// together in a single transaction.
type ControlPanel struct {

	// The Amazon Resource Name (ARN) of the cluster that includes the control panel.
	ClusterArn *string

	// The Amazon Resource Name (ARN) of the control panel.
	ControlPanelArn *string

	// A flag that Amazon Route 53 Application Recovery Controller sets to true to
	// designate the default control panel for a cluster. When you create a cluster,
	// Amazon Route 53 Application Recovery Controller creates a control panel, and
	// sets this flag for that control panel. If you create a control panel yourself,
	// this flag is set to false.
	DefaultControlPanel bool

	// The name of the control panel. You can use any non-white space character in the
	// name.
	Name *string

	// The number of routing controls in the control panel.
	RoutingControlCount int32

	// The deployment status of control panel. Status can be one of the following:
	// PENDING, DEPLOYED, PENDING_DELETION.
	Status Status

	noSmithyDocumentSerde
}

// A gating rule verifies that a gating routing control or set of gating rounting
// controls, evaluates as true, based on a rule configuration that you specify,
// which allows a set of routing control state changes to complete. For example, if
// you specify one gating routing control and you set the Type in the rule
// configuration to OR, that indicates that you must set the gating routing control
// to On for the rule to evaluate as true; that is, for the gating control "switch"
// to be "On". When you do that, then you can update the routing control states for
// the target routing controls that you specify in the gating rule.
type GatingRule struct {

	// The Amazon Resource Name (ARN) of the control panel.
	//
	// This member is required.
	ControlPanelArn *string

	// An array of gating routing control Amazon Resource Names (ARNs). For a simple
	// "on/off" switch, specify the ARN for one routing control. The gating routing
	// controls are evaluated by the rule configuration that you specify to determine
	// if the target routing control states can be changed.
	//
	// This member is required.
	GatingControls []string

	// The name for the gating rule. You can use any non-white space character in the
	// name.
	//
	// This member is required.
	Name *string

	// The criteria that you set for gating routing controls that designates how many
	// of the routing control states must be ON to allow you to update target routing
	// control states.
	//
	// This member is required.
	RuleConfig *RuleConfig

	// The Amazon Resource Name (ARN) of the gating rule.
	//
	// This member is required.
	SafetyRuleArn *string

	// The deployment status of a gating rule. Status can be one of the following:
	// PENDING, DEPLOYED, PENDING_DELETION.
	//
	// This member is required.
	Status Status

	// An array of target routing control Amazon Resource Names (ARNs) for which the
	// states can only be updated if the rule configuration that you specify evaluates
	// to true for the gating routing control. As a simple example, if you have a
	// single gating control, it acts as an overall "on/off" switch for a set of target
	// routing controls. You can use this to manually override automated fail over, for
	// example.
	//
	// This member is required.
	TargetControls []string

	// An evaluation period, in milliseconds (ms), during which any request against the
	// target routing controls will fail. This helps prevent "flapping" of state. The
	// wait period is 5000 ms by default, but you can choose a custom value.
	//
	// This member is required.
	WaitPeriodMs int32

	noSmithyDocumentSerde
}

// Update to a gating rule. You can update the name or the evaluation period (wait
// period). If you don't specify one of the items to update, the item is unchanged.
type GatingRuleUpdate struct {

	// The name for the gating rule. You can use any non-white space character in the
	// name.
	//
	// This member is required.
	Name *string

	// The Amazon Resource Name (ARN) of the gating rule.
	//
	// This member is required.
	SafetyRuleArn *string

	// An evaluation period, in milliseconds (ms), during which any request against the
	// target routing controls will fail. This helps prevent "flapping" of state. The
	// wait period is 5000 ms by default, but you can choose a custom value.
	//
	// This member is required.
	WaitPeriodMs int32

	noSmithyDocumentSerde
}

// A new assertion rule for a control panel.
type NewAssertionRule struct {

	// The routing controls that are part of transactions that are evaluated to
	// determine if a request to change a routing control state is allowed. For
	// example, you might include three routing controls, one for each of three Amazon
	// Web Services Regions.
	//
	// This member is required.
	AssertedControls []string

	// The Amazon Resource Name (ARN) for the control panel.
	//
	// This member is required.
	ControlPanelArn *string

	// The name of the assertion rule. You can use any non-white space character in the
	// name.
	//
	// This member is required.
	Name *string

	// The criteria that you set for specific assertion controls (routing controls)
	// that designate how many control states must be ON as the result of a
	// transaction. For example, if you have three assertion controls, you might
	// specify ATLEAST 2for your rule configuration. This means that at least two
	// assertion controls must be ON, so that at least two Amazon Web Services Regions
	// have traffic flowing to them.
	//
	// This member is required.
	RuleConfig *RuleConfig

	// An evaluation period, in milliseconds (ms), during which any request against the
	// target routing controls will fail. This helps prevent "flapping" of state. The
	// wait period is 5000 ms by default, but you can choose a custom value.
	//
	// This member is required.
	WaitPeriodMs int32

	noSmithyDocumentSerde
}

// A new gating rule for a control panel.
type NewGatingRule struct {

	// The Amazon Resource Name (ARN) of the control panel.
	//
	// This member is required.
	ControlPanelArn *string

	// The gating controls for the new gating rule. That is, routing controls that are
	// evaluated by the rule configuration that you specify.
	//
	// This member is required.
	GatingControls []string

	// The name for the new gating rule.
	//
	// This member is required.
	Name *string

	// The criteria that you set for specific gating controls (routing controls) that
	// designates how many control states must be ON to allow you to change (set or
	// unset) the target control states.
	//
	// This member is required.
	RuleConfig *RuleConfig

	// Routing controls that can only be set or unset if the specified RuleConfig
	// evaluates to true for the specified GatingControls. For example, say you have
	// three gating controls, one for each of three Amazon Web Services Regions. Now
	// you specify AtLeast 2 as your RuleConfig. With these settings, you can only
	// change (set or unset) the routing controls that you have specified as
	// TargetControls if that rule evaluates to true. In other words, your ability to
	// change the routing controls that you have specified as TargetControls is gated
	// by the rule that you set for the routing controls in GatingControls.
	//
	// This member is required.
	TargetControls []string

	// An evaluation period, in milliseconds (ms), during which any request against the
	// target routing controls will fail. This helps prevent "flapping" of state. The
	// wait period is 5000 ms by default, but you can choose a custom value.
	//
	// This member is required.
	WaitPeriodMs int32

	noSmithyDocumentSerde
}

// A routing control has one of two states: ON and OFF. You can map the routing
// control state to the state of an Amazon Route 53 health check, which can be used
// to control traffic routing.
type RoutingControl struct {

	// The Amazon Resource Name (ARN) of the control panel that includes the routing
	// control.
	ControlPanelArn *string

	// The name of the routing control.
	Name *string

	// The Amazon Resource Name (ARN) of the routing control.
	RoutingControlArn *string

	// The deployment status of a routing control. Status can be one of the following:
	// PENDING, DEPLOYED, PENDING_DELETION.
	Status Status

	noSmithyDocumentSerde
}

// A safety rule. A safety rule can be an assertion rule or a gating rule.
type Rule struct {

	// An assertion rule enforces that, when a routing control state is changed, the
	// criteria set by the rule configuration is met. Otherwise, the change to the
	// routing control state is not accepted. For example, the criteria might be that
	// at least one routing control state is On after the transation so that traffic
	// continues to flow to at least one cell for the application. This ensures that
	// you avoid a fail-open scenario.
	ASSERTION *AssertionRule

	// A gating rule verifies that a gating routing control or set of gating rounting
	// controls, evaluates as true, based on a rule configuration that you specify,
	// which allows a set of routing control state changes to complete. For example, if
	// you specify one gating routing control and you set the Type in the rule
	// configuration to OR, that indicates that you must set the gating routing control
	// to On for the rule to evaluate as true; that is, for the gating control "switch"
	// to be "On". When you do that, then you can update the routing control states for
	// the target routing controls that you specify in the gating rule.
	GATING *GatingRule

	noSmithyDocumentSerde
}

// The rule configuration for an assertion rule. That is, the criteria that you set
// for specific assertion controls (routing controls) that specify how many control
// states must be ON after a transaction completes.
type RuleConfig struct {

	// Logical negation of the rule. If the rule would usually evaluate true, it's
	// evaluated as false, and vice versa.
	//
	// This member is required.
	Inverted bool

	// The value of N, when you specify an ATLEAST rule type. That is, Threshold is the
	// number of controls that must be set when you specify an ATLEAST type.
	//
	// This member is required.
	Threshold int32

	// A rule can be one of the following: ATLEAST, AND, or OR.
	//
	// This member is required.
	Type RuleType

	noSmithyDocumentSerde
}

type noSmithyDocumentSerde = smithydocument.NoSerde
