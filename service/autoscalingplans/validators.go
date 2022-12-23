// Code generated by smithy-go-codegen DO NOT EDIT.

package autoscalingplans

import (
	"context"
	"fmt"
	"oss-sdk-go/service/autoscalingplans/types"
	smithy "github.com/aws/smithy-go"
	"github.com/aws/smithy-go/middleware"
)

type validateOpCreateScalingPlan struct {
}

func (*validateOpCreateScalingPlan) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpCreateScalingPlan) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*CreateScalingPlanInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpCreateScalingPlanInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpDeleteScalingPlan struct {
}

func (*validateOpDeleteScalingPlan) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpDeleteScalingPlan) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*DeleteScalingPlanInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpDeleteScalingPlanInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpDescribeScalingPlanResources struct {
}

func (*validateOpDescribeScalingPlanResources) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpDescribeScalingPlanResources) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*DescribeScalingPlanResourcesInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpDescribeScalingPlanResourcesInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpGetScalingPlanResourceForecastData struct {
}

func (*validateOpGetScalingPlanResourceForecastData) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpGetScalingPlanResourceForecastData) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*GetScalingPlanResourceForecastDataInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpGetScalingPlanResourceForecastDataInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpUpdateScalingPlan struct {
}

func (*validateOpUpdateScalingPlan) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpUpdateScalingPlan) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*UpdateScalingPlanInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpUpdateScalingPlanInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

func addOpCreateScalingPlanValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpCreateScalingPlan{}, middleware.After)
}

func addOpDeleteScalingPlanValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpDeleteScalingPlan{}, middleware.After)
}

func addOpDescribeScalingPlanResourcesValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpDescribeScalingPlanResources{}, middleware.After)
}

func addOpGetScalingPlanResourceForecastDataValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpGetScalingPlanResourceForecastData{}, middleware.After)
}

func addOpUpdateScalingPlanValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpUpdateScalingPlan{}, middleware.After)
}

func validateCustomizedLoadMetricSpecification(v *types.CustomizedLoadMetricSpecification) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "CustomizedLoadMetricSpecification"}
	if v.MetricName == nil {
		invalidParams.Add(smithy.NewErrParamRequired("MetricName"))
	}
	if v.Namespace == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Namespace"))
	}
	if v.Dimensions != nil {
		if err := validateMetricDimensions(v.Dimensions); err != nil {
			invalidParams.AddNested("Dimensions", err.(smithy.InvalidParamsError))
		}
	}
	if len(v.Statistic) == 0 {
		invalidParams.Add(smithy.NewErrParamRequired("Statistic"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateCustomizedScalingMetricSpecification(v *types.CustomizedScalingMetricSpecification) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "CustomizedScalingMetricSpecification"}
	if v.MetricName == nil {
		invalidParams.Add(smithy.NewErrParamRequired("MetricName"))
	}
	if v.Namespace == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Namespace"))
	}
	if v.Dimensions != nil {
		if err := validateMetricDimensions(v.Dimensions); err != nil {
			invalidParams.AddNested("Dimensions", err.(smithy.InvalidParamsError))
		}
	}
	if len(v.Statistic) == 0 {
		invalidParams.Add(smithy.NewErrParamRequired("Statistic"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateMetricDimension(v *types.MetricDimension) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "MetricDimension"}
	if v.Name == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Name"))
	}
	if v.Value == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Value"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateMetricDimensions(v []types.MetricDimension) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "MetricDimensions"}
	for i := range v {
		if err := validateMetricDimension(&v[i]); err != nil {
			invalidParams.AddNested(fmt.Sprintf("[%d]", i), err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validatePredefinedLoadMetricSpecification(v *types.PredefinedLoadMetricSpecification) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "PredefinedLoadMetricSpecification"}
	if len(v.PredefinedLoadMetricType) == 0 {
		invalidParams.Add(smithy.NewErrParamRequired("PredefinedLoadMetricType"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validatePredefinedScalingMetricSpecification(v *types.PredefinedScalingMetricSpecification) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "PredefinedScalingMetricSpecification"}
	if len(v.PredefinedScalingMetricType) == 0 {
		invalidParams.Add(smithy.NewErrParamRequired("PredefinedScalingMetricType"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateScalingInstruction(v *types.ScalingInstruction) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "ScalingInstruction"}
	if len(v.ServiceNamespace) == 0 {
		invalidParams.Add(smithy.NewErrParamRequired("ServiceNamespace"))
	}
	if v.ResourceId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ResourceId"))
	}
	if len(v.ScalableDimension) == 0 {
		invalidParams.Add(smithy.NewErrParamRequired("ScalableDimension"))
	}
	if v.MinCapacity == nil {
		invalidParams.Add(smithy.NewErrParamRequired("MinCapacity"))
	}
	if v.MaxCapacity == nil {
		invalidParams.Add(smithy.NewErrParamRequired("MaxCapacity"))
	}
	if v.TargetTrackingConfigurations == nil {
		invalidParams.Add(smithy.NewErrParamRequired("TargetTrackingConfigurations"))
	} else if v.TargetTrackingConfigurations != nil {
		if err := validateTargetTrackingConfigurations(v.TargetTrackingConfigurations); err != nil {
			invalidParams.AddNested("TargetTrackingConfigurations", err.(smithy.InvalidParamsError))
		}
	}
	if v.PredefinedLoadMetricSpecification != nil {
		if err := validatePredefinedLoadMetricSpecification(v.PredefinedLoadMetricSpecification); err != nil {
			invalidParams.AddNested("PredefinedLoadMetricSpecification", err.(smithy.InvalidParamsError))
		}
	}
	if v.CustomizedLoadMetricSpecification != nil {
		if err := validateCustomizedLoadMetricSpecification(v.CustomizedLoadMetricSpecification); err != nil {
			invalidParams.AddNested("CustomizedLoadMetricSpecification", err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateScalingInstructions(v []types.ScalingInstruction) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "ScalingInstructions"}
	for i := range v {
		if err := validateScalingInstruction(&v[i]); err != nil {
			invalidParams.AddNested(fmt.Sprintf("[%d]", i), err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateTargetTrackingConfiguration(v *types.TargetTrackingConfiguration) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "TargetTrackingConfiguration"}
	if v.PredefinedScalingMetricSpecification != nil {
		if err := validatePredefinedScalingMetricSpecification(v.PredefinedScalingMetricSpecification); err != nil {
			invalidParams.AddNested("PredefinedScalingMetricSpecification", err.(smithy.InvalidParamsError))
		}
	}
	if v.CustomizedScalingMetricSpecification != nil {
		if err := validateCustomizedScalingMetricSpecification(v.CustomizedScalingMetricSpecification); err != nil {
			invalidParams.AddNested("CustomizedScalingMetricSpecification", err.(smithy.InvalidParamsError))
		}
	}
	if v.TargetValue == nil {
		invalidParams.Add(smithy.NewErrParamRequired("TargetValue"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateTargetTrackingConfigurations(v []types.TargetTrackingConfiguration) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "TargetTrackingConfigurations"}
	for i := range v {
		if err := validateTargetTrackingConfiguration(&v[i]); err != nil {
			invalidParams.AddNested(fmt.Sprintf("[%d]", i), err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpCreateScalingPlanInput(v *CreateScalingPlanInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "CreateScalingPlanInput"}
	if v.ScalingPlanName == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ScalingPlanName"))
	}
	if v.ApplicationSource == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ApplicationSource"))
	}
	if v.ScalingInstructions == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ScalingInstructions"))
	} else if v.ScalingInstructions != nil {
		if err := validateScalingInstructions(v.ScalingInstructions); err != nil {
			invalidParams.AddNested("ScalingInstructions", err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpDeleteScalingPlanInput(v *DeleteScalingPlanInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "DeleteScalingPlanInput"}
	if v.ScalingPlanName == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ScalingPlanName"))
	}
	if v.ScalingPlanVersion == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ScalingPlanVersion"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpDescribeScalingPlanResourcesInput(v *DescribeScalingPlanResourcesInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "DescribeScalingPlanResourcesInput"}
	if v.ScalingPlanName == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ScalingPlanName"))
	}
	if v.ScalingPlanVersion == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ScalingPlanVersion"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpGetScalingPlanResourceForecastDataInput(v *GetScalingPlanResourceForecastDataInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "GetScalingPlanResourceForecastDataInput"}
	if v.ScalingPlanName == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ScalingPlanName"))
	}
	if v.ScalingPlanVersion == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ScalingPlanVersion"))
	}
	if len(v.ServiceNamespace) == 0 {
		invalidParams.Add(smithy.NewErrParamRequired("ServiceNamespace"))
	}
	if v.ResourceId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ResourceId"))
	}
	if len(v.ScalableDimension) == 0 {
		invalidParams.Add(smithy.NewErrParamRequired("ScalableDimension"))
	}
	if len(v.ForecastDataType) == 0 {
		invalidParams.Add(smithy.NewErrParamRequired("ForecastDataType"))
	}
	if v.StartTime == nil {
		invalidParams.Add(smithy.NewErrParamRequired("StartTime"))
	}
	if v.EndTime == nil {
		invalidParams.Add(smithy.NewErrParamRequired("EndTime"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpUpdateScalingPlanInput(v *UpdateScalingPlanInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "UpdateScalingPlanInput"}
	if v.ScalingPlanName == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ScalingPlanName"))
	}
	if v.ScalingPlanVersion == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ScalingPlanVersion"))
	}
	if v.ScalingInstructions != nil {
		if err := validateScalingInstructions(v.ScalingInstructions); err != nil {
			invalidParams.AddNested("ScalingInstructions", err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}
