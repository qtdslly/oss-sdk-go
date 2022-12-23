// Code generated by smithy-go-codegen DO NOT EDIT.

package ssmcontacts

import (
	"context"
	"fmt"
	"oss-sdk-go/service/ssmcontacts/types"
	smithy "github.com/aws/smithy-go"
	"github.com/aws/smithy-go/middleware"
)

type validateOpAcceptPage struct {
}

func (*validateOpAcceptPage) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpAcceptPage) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*AcceptPageInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpAcceptPageInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpActivateContactChannel struct {
}

func (*validateOpActivateContactChannel) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpActivateContactChannel) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*ActivateContactChannelInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpActivateContactChannelInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpCreateContactChannel struct {
}

func (*validateOpCreateContactChannel) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpCreateContactChannel) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*CreateContactChannelInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpCreateContactChannelInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpCreateContact struct {
}

func (*validateOpCreateContact) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpCreateContact) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*CreateContactInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpCreateContactInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpDeactivateContactChannel struct {
}

func (*validateOpDeactivateContactChannel) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpDeactivateContactChannel) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*DeactivateContactChannelInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpDeactivateContactChannelInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpDeleteContactChannel struct {
}

func (*validateOpDeleteContactChannel) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpDeleteContactChannel) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*DeleteContactChannelInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpDeleteContactChannelInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpDeleteContact struct {
}

func (*validateOpDeleteContact) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpDeleteContact) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*DeleteContactInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpDeleteContactInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpDescribeEngagement struct {
}

func (*validateOpDescribeEngagement) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpDescribeEngagement) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*DescribeEngagementInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpDescribeEngagementInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpDescribePage struct {
}

func (*validateOpDescribePage) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpDescribePage) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*DescribePageInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpDescribePageInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpGetContactChannel struct {
}

func (*validateOpGetContactChannel) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpGetContactChannel) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*GetContactChannelInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpGetContactChannelInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpGetContact struct {
}

func (*validateOpGetContact) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpGetContact) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*GetContactInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpGetContactInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpGetContactPolicy struct {
}

func (*validateOpGetContactPolicy) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpGetContactPolicy) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*GetContactPolicyInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpGetContactPolicyInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpListContactChannels struct {
}

func (*validateOpListContactChannels) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpListContactChannels) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*ListContactChannelsInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpListContactChannelsInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpListPageReceipts struct {
}

func (*validateOpListPageReceipts) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpListPageReceipts) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*ListPageReceiptsInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpListPageReceiptsInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpListPagesByContact struct {
}

func (*validateOpListPagesByContact) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpListPagesByContact) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*ListPagesByContactInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpListPagesByContactInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpListPagesByEngagement struct {
}

func (*validateOpListPagesByEngagement) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpListPagesByEngagement) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*ListPagesByEngagementInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpListPagesByEngagementInput(input); err != nil {
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

type validateOpPutContactPolicy struct {
}

func (*validateOpPutContactPolicy) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpPutContactPolicy) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*PutContactPolicyInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpPutContactPolicyInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpSendActivationCode struct {
}

func (*validateOpSendActivationCode) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpSendActivationCode) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*SendActivationCodeInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpSendActivationCodeInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpStartEngagement struct {
}

func (*validateOpStartEngagement) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpStartEngagement) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*StartEngagementInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpStartEngagementInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpStopEngagement struct {
}

func (*validateOpStopEngagement) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpStopEngagement) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*StopEngagementInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpStopEngagementInput(input); err != nil {
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

type validateOpUpdateContactChannel struct {
}

func (*validateOpUpdateContactChannel) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpUpdateContactChannel) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*UpdateContactChannelInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpUpdateContactChannelInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpUpdateContact struct {
}

func (*validateOpUpdateContact) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpUpdateContact) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*UpdateContactInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpUpdateContactInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

func addOpAcceptPageValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpAcceptPage{}, middleware.After)
}

func addOpActivateContactChannelValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpActivateContactChannel{}, middleware.After)
}

func addOpCreateContactChannelValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpCreateContactChannel{}, middleware.After)
}

func addOpCreateContactValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpCreateContact{}, middleware.After)
}

func addOpDeactivateContactChannelValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpDeactivateContactChannel{}, middleware.After)
}

func addOpDeleteContactChannelValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpDeleteContactChannel{}, middleware.After)
}

func addOpDeleteContactValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpDeleteContact{}, middleware.After)
}

func addOpDescribeEngagementValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpDescribeEngagement{}, middleware.After)
}

func addOpDescribePageValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpDescribePage{}, middleware.After)
}

func addOpGetContactChannelValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpGetContactChannel{}, middleware.After)
}

func addOpGetContactValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpGetContact{}, middleware.After)
}

func addOpGetContactPolicyValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpGetContactPolicy{}, middleware.After)
}

func addOpListContactChannelsValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpListContactChannels{}, middleware.After)
}

func addOpListPageReceiptsValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpListPageReceipts{}, middleware.After)
}

func addOpListPagesByContactValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpListPagesByContact{}, middleware.After)
}

func addOpListPagesByEngagementValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpListPagesByEngagement{}, middleware.After)
}

func addOpListTagsForResourceValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpListTagsForResource{}, middleware.After)
}

func addOpPutContactPolicyValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpPutContactPolicy{}, middleware.After)
}

func addOpSendActivationCodeValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpSendActivationCode{}, middleware.After)
}

func addOpStartEngagementValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpStartEngagement{}, middleware.After)
}

func addOpStopEngagementValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpStopEngagement{}, middleware.After)
}

func addOpTagResourceValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpTagResource{}, middleware.After)
}

func addOpUntagResourceValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpUntagResource{}, middleware.After)
}

func addOpUpdateContactChannelValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpUpdateContactChannel{}, middleware.After)
}

func addOpUpdateContactValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpUpdateContact{}, middleware.After)
}

func validateChannelTargetInfo(v *types.ChannelTargetInfo) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "ChannelTargetInfo"}
	if v.ContactChannelId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ContactChannelId"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateContactTargetInfo(v *types.ContactTargetInfo) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "ContactTargetInfo"}
	if v.IsEssential == nil {
		invalidParams.Add(smithy.NewErrParamRequired("IsEssential"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validatePlan(v *types.Plan) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "Plan"}
	if v.Stages == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Stages"))
	} else if v.Stages != nil {
		if err := validateStagesList(v.Stages); err != nil {
			invalidParams.AddNested("Stages", err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateStage(v *types.Stage) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "Stage"}
	if v.DurationInMinutes == nil {
		invalidParams.Add(smithy.NewErrParamRequired("DurationInMinutes"))
	}
	if v.Targets == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Targets"))
	} else if v.Targets != nil {
		if err := validateTargetsList(v.Targets); err != nil {
			invalidParams.AddNested("Targets", err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateStagesList(v []types.Stage) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "StagesList"}
	for i := range v {
		if err := validateStage(&v[i]); err != nil {
			invalidParams.AddNested(fmt.Sprintf("[%d]", i), err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateTarget(v *types.Target) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "Target"}
	if v.ChannelTargetInfo != nil {
		if err := validateChannelTargetInfo(v.ChannelTargetInfo); err != nil {
			invalidParams.AddNested("ChannelTargetInfo", err.(smithy.InvalidParamsError))
		}
	}
	if v.ContactTargetInfo != nil {
		if err := validateContactTargetInfo(v.ContactTargetInfo); err != nil {
			invalidParams.AddNested("ContactTargetInfo", err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateTargetsList(v []types.Target) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "TargetsList"}
	for i := range v {
		if err := validateTarget(&v[i]); err != nil {
			invalidParams.AddNested(fmt.Sprintf("[%d]", i), err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpAcceptPageInput(v *AcceptPageInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "AcceptPageInput"}
	if v.PageId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("PageId"))
	}
	if len(v.AcceptType) == 0 {
		invalidParams.Add(smithy.NewErrParamRequired("AcceptType"))
	}
	if v.AcceptCode == nil {
		invalidParams.Add(smithy.NewErrParamRequired("AcceptCode"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpActivateContactChannelInput(v *ActivateContactChannelInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "ActivateContactChannelInput"}
	if v.ContactChannelId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ContactChannelId"))
	}
	if v.ActivationCode == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ActivationCode"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpCreateContactChannelInput(v *CreateContactChannelInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "CreateContactChannelInput"}
	if v.ContactId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ContactId"))
	}
	if v.Name == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Name"))
	}
	if len(v.Type) == 0 {
		invalidParams.Add(smithy.NewErrParamRequired("Type"))
	}
	if v.DeliveryAddress == nil {
		invalidParams.Add(smithy.NewErrParamRequired("DeliveryAddress"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpCreateContactInput(v *CreateContactInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "CreateContactInput"}
	if v.Alias == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Alias"))
	}
	if len(v.Type) == 0 {
		invalidParams.Add(smithy.NewErrParamRequired("Type"))
	}
	if v.Plan == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Plan"))
	} else if v.Plan != nil {
		if err := validatePlan(v.Plan); err != nil {
			invalidParams.AddNested("Plan", err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpDeactivateContactChannelInput(v *DeactivateContactChannelInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "DeactivateContactChannelInput"}
	if v.ContactChannelId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ContactChannelId"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpDeleteContactChannelInput(v *DeleteContactChannelInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "DeleteContactChannelInput"}
	if v.ContactChannelId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ContactChannelId"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpDeleteContactInput(v *DeleteContactInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "DeleteContactInput"}
	if v.ContactId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ContactId"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpDescribeEngagementInput(v *DescribeEngagementInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "DescribeEngagementInput"}
	if v.EngagementId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("EngagementId"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpDescribePageInput(v *DescribePageInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "DescribePageInput"}
	if v.PageId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("PageId"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpGetContactChannelInput(v *GetContactChannelInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "GetContactChannelInput"}
	if v.ContactChannelId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ContactChannelId"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpGetContactInput(v *GetContactInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "GetContactInput"}
	if v.ContactId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ContactId"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpGetContactPolicyInput(v *GetContactPolicyInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "GetContactPolicyInput"}
	if v.ContactArn == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ContactArn"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpListContactChannelsInput(v *ListContactChannelsInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "ListContactChannelsInput"}
	if v.ContactId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ContactId"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpListPageReceiptsInput(v *ListPageReceiptsInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "ListPageReceiptsInput"}
	if v.PageId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("PageId"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpListPagesByContactInput(v *ListPagesByContactInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "ListPagesByContactInput"}
	if v.ContactId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ContactId"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpListPagesByEngagementInput(v *ListPagesByEngagementInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "ListPagesByEngagementInput"}
	if v.EngagementId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("EngagementId"))
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
	if v.ResourceARN == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ResourceARN"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpPutContactPolicyInput(v *PutContactPolicyInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "PutContactPolicyInput"}
	if v.ContactArn == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ContactArn"))
	}
	if v.Policy == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Policy"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpSendActivationCodeInput(v *SendActivationCodeInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "SendActivationCodeInput"}
	if v.ContactChannelId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ContactChannelId"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpStartEngagementInput(v *StartEngagementInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "StartEngagementInput"}
	if v.ContactId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ContactId"))
	}
	if v.Sender == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Sender"))
	}
	if v.Subject == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Subject"))
	}
	if v.Content == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Content"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpStopEngagementInput(v *StopEngagementInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "StopEngagementInput"}
	if v.EngagementId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("EngagementId"))
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
	if v.ResourceARN == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ResourceARN"))
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
	if v.ResourceARN == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ResourceARN"))
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

func validateOpUpdateContactChannelInput(v *UpdateContactChannelInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "UpdateContactChannelInput"}
	if v.ContactChannelId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ContactChannelId"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpUpdateContactInput(v *UpdateContactInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "UpdateContactInput"}
	if v.ContactId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ContactId"))
	}
	if v.Plan != nil {
		if err := validatePlan(v.Plan); err != nil {
			invalidParams.AddNested("Plan", err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}
