// Code generated by smithy-go-codegen DO NOT EDIT.

package route53recoverycontrolconfig

import (
	"context"
	"fmt"
	"oss-sdk-go/service/route53recoverycontrolconfig/types"
	smithy "github.com/aws/smithy-go"
	"github.com/aws/smithy-go/middleware"
)

type validateOpCreateCluster struct {
}

func (*validateOpCreateCluster) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpCreateCluster) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*CreateClusterInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpCreateClusterInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpCreateControlPanel struct {
}

func (*validateOpCreateControlPanel) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpCreateControlPanel) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*CreateControlPanelInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpCreateControlPanelInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpCreateRoutingControl struct {
}

func (*validateOpCreateRoutingControl) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpCreateRoutingControl) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*CreateRoutingControlInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpCreateRoutingControlInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpCreateSafetyRule struct {
}

func (*validateOpCreateSafetyRule) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpCreateSafetyRule) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*CreateSafetyRuleInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpCreateSafetyRuleInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpDeleteCluster struct {
}

func (*validateOpDeleteCluster) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpDeleteCluster) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*DeleteClusterInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpDeleteClusterInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpDeleteControlPanel struct {
}

func (*validateOpDeleteControlPanel) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpDeleteControlPanel) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*DeleteControlPanelInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpDeleteControlPanelInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpDeleteRoutingControl struct {
}

func (*validateOpDeleteRoutingControl) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpDeleteRoutingControl) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*DeleteRoutingControlInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpDeleteRoutingControlInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpDeleteSafetyRule struct {
}

func (*validateOpDeleteSafetyRule) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpDeleteSafetyRule) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*DeleteSafetyRuleInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpDeleteSafetyRuleInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpDescribeCluster struct {
}

func (*validateOpDescribeCluster) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpDescribeCluster) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*DescribeClusterInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpDescribeClusterInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpDescribeControlPanel struct {
}

func (*validateOpDescribeControlPanel) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpDescribeControlPanel) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*DescribeControlPanelInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpDescribeControlPanelInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpDescribeRoutingControl struct {
}

func (*validateOpDescribeRoutingControl) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpDescribeRoutingControl) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*DescribeRoutingControlInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpDescribeRoutingControlInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpDescribeSafetyRule struct {
}

func (*validateOpDescribeSafetyRule) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpDescribeSafetyRule) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*DescribeSafetyRuleInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpDescribeSafetyRuleInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpListAssociatedRoute53HealthChecks struct {
}

func (*validateOpListAssociatedRoute53HealthChecks) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpListAssociatedRoute53HealthChecks) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*ListAssociatedRoute53HealthChecksInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpListAssociatedRoute53HealthChecksInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpListRoutingControls struct {
}

func (*validateOpListRoutingControls) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpListRoutingControls) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*ListRoutingControlsInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpListRoutingControlsInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpListSafetyRules struct {
}

func (*validateOpListSafetyRules) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpListSafetyRules) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*ListSafetyRulesInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpListSafetyRulesInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpListTagsForResource struct {
}

func (*validateOpListTagsForResource) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpListTagsForResource) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*ListTagsForResourceInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpListTagsForResourceInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpTagResource struct {
}

func (*validateOpTagResource) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpTagResource) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*TagResourceInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpTagResourceInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpUntagResource struct {
}

func (*validateOpUntagResource) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpUntagResource) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*UntagResourceInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpUntagResourceInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpUpdateControlPanel struct {
}

func (*validateOpUpdateControlPanel) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpUpdateControlPanel) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*UpdateControlPanelInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpUpdateControlPanelInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpUpdateRoutingControl struct {
}

func (*validateOpUpdateRoutingControl) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpUpdateRoutingControl) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*UpdateRoutingControlInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpUpdateRoutingControlInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpUpdateSafetyRule struct {
}

func (*validateOpUpdateSafetyRule) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpUpdateSafetyRule) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*UpdateSafetyRuleInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpUpdateSafetyRuleInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

func addOpCreateClusterValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpCreateCluster{}, middleware.After)
}

func addOpCreateControlPanelValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpCreateControlPanel{}, middleware.After)
}

func addOpCreateRoutingControlValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpCreateRoutingControl{}, middleware.After)
}

func addOpCreateSafetyRuleValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpCreateSafetyRule{}, middleware.After)
}

func addOpDeleteClusterValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpDeleteCluster{}, middleware.After)
}

func addOpDeleteControlPanelValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpDeleteControlPanel{}, middleware.After)
}

func addOpDeleteRoutingControlValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpDeleteRoutingControl{}, middleware.After)
}

func addOpDeleteSafetyRuleValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpDeleteSafetyRule{}, middleware.After)
}

func addOpDescribeClusterValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpDescribeCluster{}, middleware.After)
}

func addOpDescribeControlPanelValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpDescribeControlPanel{}, middleware.After)
}

func addOpDescribeRoutingControlValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpDescribeRoutingControl{}, middleware.After)
}

func addOpDescribeSafetyRuleValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpDescribeSafetyRule{}, middleware.After)
}

func addOpListAssociatedRoute53HealthChecksValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpListAssociatedRoute53HealthChecks{}, middleware.After)
}

func addOpListRoutingControlsValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpListRoutingControls{}, middleware.After)
}

func addOpListSafetyRulesValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpListSafetyRules{}, middleware.After)
}

func addOpListTagsForResourceValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpListTagsForResource{}, middleware.After)
}

func addOpTagResourceValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpTagResource{}, middleware.After)
}

func addOpUntagResourceValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpUntagResource{}, middleware.After)
}

func addOpUpdateControlPanelValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpUpdateControlPanel{}, middleware.After)
}

func addOpUpdateRoutingControlValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpUpdateRoutingControl{}, middleware.After)
}

func addOpUpdateSafetyRuleValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpUpdateSafetyRule{}, middleware.After)
}

func validateAssertionRuleUpdate(v *types.AssertionRuleUpdate) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "AssertionRuleUpdate"}
	if v.Name == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Name"))
	}
	if v.SafetyRuleArn == nil {
		invalidParams.Add(smithy.NewErrParamRequired("SafetyRuleArn"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateGatingRuleUpdate(v *types.GatingRuleUpdate) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "GatingRuleUpdate"}
	if v.Name == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Name"))
	}
	if v.SafetyRuleArn == nil {
		invalidParams.Add(smithy.NewErrParamRequired("SafetyRuleArn"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateNewAssertionRule(v *types.NewAssertionRule) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "NewAssertionRule"}
	if v.AssertedControls == nil {
		invalidParams.Add(smithy.NewErrParamRequired("AssertedControls"))
	}
	if v.ControlPanelArn == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ControlPanelArn"))
	}
	if v.Name == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Name"))
	}
	if v.RuleConfig == nil {
		invalidParams.Add(smithy.NewErrParamRequired("RuleConfig"))
	} else if v.RuleConfig != nil {
		if err := validateRuleConfig(v.RuleConfig); err != nil {
			invalidParams.AddNested("RuleConfig", err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateNewGatingRule(v *types.NewGatingRule) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "NewGatingRule"}
	if v.ControlPanelArn == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ControlPanelArn"))
	}
	if v.GatingControls == nil {
		invalidParams.Add(smithy.NewErrParamRequired("GatingControls"))
	}
	if v.Name == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Name"))
	}
	if v.RuleConfig == nil {
		invalidParams.Add(smithy.NewErrParamRequired("RuleConfig"))
	} else if v.RuleConfig != nil {
		if err := validateRuleConfig(v.RuleConfig); err != nil {
			invalidParams.AddNested("RuleConfig", err.(smithy.InvalidParamsError))
		}
	}
	if v.TargetControls == nil {
		invalidParams.Add(smithy.NewErrParamRequired("TargetControls"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateRuleConfig(v *types.RuleConfig) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "RuleConfig"}
	if len(v.Type) == 0 {
		invalidParams.Add(smithy.NewErrParamRequired("Type"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpCreateClusterInput(v *CreateClusterInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "CreateClusterInput"}
	if v.ClusterName == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ClusterName"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpCreateControlPanelInput(v *CreateControlPanelInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "CreateControlPanelInput"}
	if v.ClusterArn == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ClusterArn"))
	}
	if v.ControlPanelName == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ControlPanelName"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpCreateRoutingControlInput(v *CreateRoutingControlInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "CreateRoutingControlInput"}
	if v.ClusterArn == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ClusterArn"))
	}
	if v.RoutingControlName == nil {
		invalidParams.Add(smithy.NewErrParamRequired("RoutingControlName"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpCreateSafetyRuleInput(v *CreateSafetyRuleInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "CreateSafetyRuleInput"}
	if v.AssertionRule != nil {
		if err := validateNewAssertionRule(v.AssertionRule); err != nil {
			invalidParams.AddNested("AssertionRule", err.(smithy.InvalidParamsError))
		}
	}
	if v.GatingRule != nil {
		if err := validateNewGatingRule(v.GatingRule); err != nil {
			invalidParams.AddNested("GatingRule", err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpDeleteClusterInput(v *DeleteClusterInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "DeleteClusterInput"}
	if v.ClusterArn == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ClusterArn"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpDeleteControlPanelInput(v *DeleteControlPanelInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "DeleteControlPanelInput"}
	if v.ControlPanelArn == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ControlPanelArn"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpDeleteRoutingControlInput(v *DeleteRoutingControlInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "DeleteRoutingControlInput"}
	if v.RoutingControlArn == nil {
		invalidParams.Add(smithy.NewErrParamRequired("RoutingControlArn"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpDeleteSafetyRuleInput(v *DeleteSafetyRuleInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "DeleteSafetyRuleInput"}
	if v.SafetyRuleArn == nil {
		invalidParams.Add(smithy.NewErrParamRequired("SafetyRuleArn"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpDescribeClusterInput(v *DescribeClusterInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "DescribeClusterInput"}
	if v.ClusterArn == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ClusterArn"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpDescribeControlPanelInput(v *DescribeControlPanelInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "DescribeControlPanelInput"}
	if v.ControlPanelArn == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ControlPanelArn"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpDescribeRoutingControlInput(v *DescribeRoutingControlInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "DescribeRoutingControlInput"}
	if v.RoutingControlArn == nil {
		invalidParams.Add(smithy.NewErrParamRequired("RoutingControlArn"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpDescribeSafetyRuleInput(v *DescribeSafetyRuleInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "DescribeSafetyRuleInput"}
	if v.SafetyRuleArn == nil {
		invalidParams.Add(smithy.NewErrParamRequired("SafetyRuleArn"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpListAssociatedRoute53HealthChecksInput(v *ListAssociatedRoute53HealthChecksInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "ListAssociatedRoute53HealthChecksInput"}
	if v.RoutingControlArn == nil {
		invalidParams.Add(smithy.NewErrParamRequired("RoutingControlArn"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpListRoutingControlsInput(v *ListRoutingControlsInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "ListRoutingControlsInput"}
	if v.ControlPanelArn == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ControlPanelArn"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpListSafetyRulesInput(v *ListSafetyRulesInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "ListSafetyRulesInput"}
	if v.ControlPanelArn == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ControlPanelArn"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpListTagsForResourceInput(v *ListTagsForResourceInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "ListTagsForResourceInput"}
	if v.ResourceArn == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ResourceArn"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpTagResourceInput(v *TagResourceInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "TagResourceInput"}
	if v.ResourceArn == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ResourceArn"))
	}
	if v.Tags == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Tags"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpUntagResourceInput(v *UntagResourceInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "UntagResourceInput"}
	if v.ResourceArn == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ResourceArn"))
	}
	if v.TagKeys == nil {
		invalidParams.Add(smithy.NewErrParamRequired("TagKeys"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpUpdateControlPanelInput(v *UpdateControlPanelInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "UpdateControlPanelInput"}
	if v.ControlPanelArn == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ControlPanelArn"))
	}
	if v.ControlPanelName == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ControlPanelName"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpUpdateRoutingControlInput(v *UpdateRoutingControlInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "UpdateRoutingControlInput"}
	if v.RoutingControlArn == nil {
		invalidParams.Add(smithy.NewErrParamRequired("RoutingControlArn"))
	}
	if v.RoutingControlName == nil {
		invalidParams.Add(smithy.NewErrParamRequired("RoutingControlName"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpUpdateSafetyRuleInput(v *UpdateSafetyRuleInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "UpdateSafetyRuleInput"}
	if v.AssertionRuleUpdate != nil {
		if err := validateAssertionRuleUpdate(v.AssertionRuleUpdate); err != nil {
			invalidParams.AddNested("AssertionRuleUpdate", err.(smithy.InvalidParamsError))
		}
	}
	if v.GatingRuleUpdate != nil {
		if err := validateGatingRuleUpdate(v.GatingRuleUpdate); err != nil {
			invalidParams.AddNested("GatingRuleUpdate", err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}
