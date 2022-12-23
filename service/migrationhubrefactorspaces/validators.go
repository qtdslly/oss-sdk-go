// Code generated by smithy-go-codegen DO NOT EDIT.

package migrationhubrefactorspaces

import (
	"context"
	"fmt"
	"oss-sdk-go/service/migrationhubrefactorspaces/types"
	smithy "github.com/aws/smithy-go"
	"github.com/aws/smithy-go/middleware"
)

type validateOpCreateApplication struct {
}

func (*validateOpCreateApplication) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpCreateApplication) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*CreateApplicationInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpCreateApplicationInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpCreateEnvironment struct {
}

func (*validateOpCreateEnvironment) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpCreateEnvironment) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*CreateEnvironmentInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpCreateEnvironmentInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpCreateRoute struct {
}

func (*validateOpCreateRoute) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpCreateRoute) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*CreateRouteInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpCreateRouteInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpCreateService struct {
}

func (*validateOpCreateService) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpCreateService) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*CreateServiceInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpCreateServiceInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpDeleteApplication struct {
}

func (*validateOpDeleteApplication) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpDeleteApplication) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*DeleteApplicationInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpDeleteApplicationInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpDeleteEnvironment struct {
}

func (*validateOpDeleteEnvironment) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpDeleteEnvironment) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*DeleteEnvironmentInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpDeleteEnvironmentInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpDeleteResourcePolicy struct {
}

func (*validateOpDeleteResourcePolicy) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpDeleteResourcePolicy) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*DeleteResourcePolicyInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpDeleteResourcePolicyInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpDeleteRoute struct {
}

func (*validateOpDeleteRoute) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpDeleteRoute) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*DeleteRouteInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpDeleteRouteInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpDeleteService struct {
}

func (*validateOpDeleteService) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpDeleteService) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*DeleteServiceInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpDeleteServiceInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpGetApplication struct {
}

func (*validateOpGetApplication) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpGetApplication) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*GetApplicationInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpGetApplicationInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpGetEnvironment struct {
}

func (*validateOpGetEnvironment) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpGetEnvironment) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*GetEnvironmentInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpGetEnvironmentInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpGetResourcePolicy struct {
}

func (*validateOpGetResourcePolicy) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpGetResourcePolicy) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*GetResourcePolicyInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpGetResourcePolicyInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpGetRoute struct {
}

func (*validateOpGetRoute) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpGetRoute) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*GetRouteInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpGetRouteInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpGetService struct {
}

func (*validateOpGetService) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpGetService) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*GetServiceInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpGetServiceInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpListApplications struct {
}

func (*validateOpListApplications) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpListApplications) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*ListApplicationsInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpListApplicationsInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpListEnvironmentVpcs struct {
}

func (*validateOpListEnvironmentVpcs) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpListEnvironmentVpcs) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*ListEnvironmentVpcsInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpListEnvironmentVpcsInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpListRoutes struct {
}

func (*validateOpListRoutes) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpListRoutes) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*ListRoutesInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpListRoutesInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpListServices struct {
}

func (*validateOpListServices) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpListServices) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*ListServicesInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpListServicesInput(input); err != nil {
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

type validateOpPutResourcePolicy struct {
}

func (*validateOpPutResourcePolicy) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpPutResourcePolicy) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*PutResourcePolicyInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpPutResourcePolicyInput(input); err != nil {
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

type validateOpUpdateRoute struct {
}

func (*validateOpUpdateRoute) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpUpdateRoute) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*UpdateRouteInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpUpdateRouteInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

func addOpCreateApplicationValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpCreateApplication{}, middleware.After)
}

func addOpCreateEnvironmentValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpCreateEnvironment{}, middleware.After)
}

func addOpCreateRouteValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpCreateRoute{}, middleware.After)
}

func addOpCreateServiceValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpCreateService{}, middleware.After)
}

func addOpDeleteApplicationValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpDeleteApplication{}, middleware.After)
}

func addOpDeleteEnvironmentValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpDeleteEnvironment{}, middleware.After)
}

func addOpDeleteResourcePolicyValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpDeleteResourcePolicy{}, middleware.After)
}

func addOpDeleteRouteValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpDeleteRoute{}, middleware.After)
}

func addOpDeleteServiceValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpDeleteService{}, middleware.After)
}

func addOpGetApplicationValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpGetApplication{}, middleware.After)
}

func addOpGetEnvironmentValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpGetEnvironment{}, middleware.After)
}

func addOpGetResourcePolicyValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpGetResourcePolicy{}, middleware.After)
}

func addOpGetRouteValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpGetRoute{}, middleware.After)
}

func addOpGetServiceValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpGetService{}, middleware.After)
}

func addOpListApplicationsValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpListApplications{}, middleware.After)
}

func addOpListEnvironmentVpcsValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpListEnvironmentVpcs{}, middleware.After)
}

func addOpListRoutesValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpListRoutes{}, middleware.After)
}

func addOpListServicesValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpListServices{}, middleware.After)
}

func addOpListTagsForResourceValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpListTagsForResource{}, middleware.After)
}

func addOpPutResourcePolicyValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpPutResourcePolicy{}, middleware.After)
}

func addOpTagResourceValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpTagResource{}, middleware.After)
}

func addOpUntagResourceValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpUntagResource{}, middleware.After)
}

func addOpUpdateRouteValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpUpdateRoute{}, middleware.After)
}

func validateLambdaEndpointInput(v *types.LambdaEndpointInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "LambdaEndpointInput"}
	if v.Arn == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Arn"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateUriPathRouteInput(v *types.UriPathRouteInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "UriPathRouteInput"}
	if v.SourcePath == nil {
		invalidParams.Add(smithy.NewErrParamRequired("SourcePath"))
	}
	if len(v.ActivationState) == 0 {
		invalidParams.Add(smithy.NewErrParamRequired("ActivationState"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateUrlEndpointInput(v *types.UrlEndpointInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "UrlEndpointInput"}
	if v.Url == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Url"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpCreateApplicationInput(v *CreateApplicationInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "CreateApplicationInput"}
	if v.Name == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Name"))
	}
	if v.EnvironmentIdentifier == nil {
		invalidParams.Add(smithy.NewErrParamRequired("EnvironmentIdentifier"))
	}
	if v.VpcId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("VpcId"))
	}
	if len(v.ProxyType) == 0 {
		invalidParams.Add(smithy.NewErrParamRequired("ProxyType"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpCreateEnvironmentInput(v *CreateEnvironmentInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "CreateEnvironmentInput"}
	if v.Name == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Name"))
	}
	if len(v.NetworkFabricType) == 0 {
		invalidParams.Add(smithy.NewErrParamRequired("NetworkFabricType"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpCreateRouteInput(v *CreateRouteInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "CreateRouteInput"}
	if v.EnvironmentIdentifier == nil {
		invalidParams.Add(smithy.NewErrParamRequired("EnvironmentIdentifier"))
	}
	if v.ApplicationIdentifier == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ApplicationIdentifier"))
	}
	if v.ServiceIdentifier == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ServiceIdentifier"))
	}
	if len(v.RouteType) == 0 {
		invalidParams.Add(smithy.NewErrParamRequired("RouteType"))
	}
	if v.UriPathRoute != nil {
		if err := validateUriPathRouteInput(v.UriPathRoute); err != nil {
			invalidParams.AddNested("UriPathRoute", err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpCreateServiceInput(v *CreateServiceInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "CreateServiceInput"}
	if v.Name == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Name"))
	}
	if v.EnvironmentIdentifier == nil {
		invalidParams.Add(smithy.NewErrParamRequired("EnvironmentIdentifier"))
	}
	if v.ApplicationIdentifier == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ApplicationIdentifier"))
	}
	if len(v.EndpointType) == 0 {
		invalidParams.Add(smithy.NewErrParamRequired("EndpointType"))
	}
	if v.UrlEndpoint != nil {
		if err := validateUrlEndpointInput(v.UrlEndpoint); err != nil {
			invalidParams.AddNested("UrlEndpoint", err.(smithy.InvalidParamsError))
		}
	}
	if v.LambdaEndpoint != nil {
		if err := validateLambdaEndpointInput(v.LambdaEndpoint); err != nil {
			invalidParams.AddNested("LambdaEndpoint", err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpDeleteApplicationInput(v *DeleteApplicationInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "DeleteApplicationInput"}
	if v.EnvironmentIdentifier == nil {
		invalidParams.Add(smithy.NewErrParamRequired("EnvironmentIdentifier"))
	}
	if v.ApplicationIdentifier == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ApplicationIdentifier"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpDeleteEnvironmentInput(v *DeleteEnvironmentInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "DeleteEnvironmentInput"}
	if v.EnvironmentIdentifier == nil {
		invalidParams.Add(smithy.NewErrParamRequired("EnvironmentIdentifier"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpDeleteResourcePolicyInput(v *DeleteResourcePolicyInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "DeleteResourcePolicyInput"}
	if v.Identifier == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Identifier"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpDeleteRouteInput(v *DeleteRouteInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "DeleteRouteInput"}
	if v.EnvironmentIdentifier == nil {
		invalidParams.Add(smithy.NewErrParamRequired("EnvironmentIdentifier"))
	}
	if v.ApplicationIdentifier == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ApplicationIdentifier"))
	}
	if v.RouteIdentifier == nil {
		invalidParams.Add(smithy.NewErrParamRequired("RouteIdentifier"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpDeleteServiceInput(v *DeleteServiceInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "DeleteServiceInput"}
	if v.EnvironmentIdentifier == nil {
		invalidParams.Add(smithy.NewErrParamRequired("EnvironmentIdentifier"))
	}
	if v.ApplicationIdentifier == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ApplicationIdentifier"))
	}
	if v.ServiceIdentifier == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ServiceIdentifier"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpGetApplicationInput(v *GetApplicationInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "GetApplicationInput"}
	if v.EnvironmentIdentifier == nil {
		invalidParams.Add(smithy.NewErrParamRequired("EnvironmentIdentifier"))
	}
	if v.ApplicationIdentifier == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ApplicationIdentifier"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpGetEnvironmentInput(v *GetEnvironmentInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "GetEnvironmentInput"}
	if v.EnvironmentIdentifier == nil {
		invalidParams.Add(smithy.NewErrParamRequired("EnvironmentIdentifier"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpGetResourcePolicyInput(v *GetResourcePolicyInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "GetResourcePolicyInput"}
	if v.Identifier == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Identifier"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpGetRouteInput(v *GetRouteInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "GetRouteInput"}
	if v.EnvironmentIdentifier == nil {
		invalidParams.Add(smithy.NewErrParamRequired("EnvironmentIdentifier"))
	}
	if v.ApplicationIdentifier == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ApplicationIdentifier"))
	}
	if v.RouteIdentifier == nil {
		invalidParams.Add(smithy.NewErrParamRequired("RouteIdentifier"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpGetServiceInput(v *GetServiceInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "GetServiceInput"}
	if v.EnvironmentIdentifier == nil {
		invalidParams.Add(smithy.NewErrParamRequired("EnvironmentIdentifier"))
	}
	if v.ApplicationIdentifier == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ApplicationIdentifier"))
	}
	if v.ServiceIdentifier == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ServiceIdentifier"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpListApplicationsInput(v *ListApplicationsInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "ListApplicationsInput"}
	if v.EnvironmentIdentifier == nil {
		invalidParams.Add(smithy.NewErrParamRequired("EnvironmentIdentifier"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpListEnvironmentVpcsInput(v *ListEnvironmentVpcsInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "ListEnvironmentVpcsInput"}
	if v.EnvironmentIdentifier == nil {
		invalidParams.Add(smithy.NewErrParamRequired("EnvironmentIdentifier"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpListRoutesInput(v *ListRoutesInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "ListRoutesInput"}
	if v.EnvironmentIdentifier == nil {
		invalidParams.Add(smithy.NewErrParamRequired("EnvironmentIdentifier"))
	}
	if v.ApplicationIdentifier == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ApplicationIdentifier"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpListServicesInput(v *ListServicesInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "ListServicesInput"}
	if v.EnvironmentIdentifier == nil {
		invalidParams.Add(smithy.NewErrParamRequired("EnvironmentIdentifier"))
	}
	if v.ApplicationIdentifier == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ApplicationIdentifier"))
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

func validateOpPutResourcePolicyInput(v *PutResourcePolicyInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "PutResourcePolicyInput"}
	if v.ResourceArn == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ResourceArn"))
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

func validateOpUpdateRouteInput(v *UpdateRouteInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "UpdateRouteInput"}
	if v.EnvironmentIdentifier == nil {
		invalidParams.Add(smithy.NewErrParamRequired("EnvironmentIdentifier"))
	}
	if v.ApplicationIdentifier == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ApplicationIdentifier"))
	}
	if v.RouteIdentifier == nil {
		invalidParams.Add(smithy.NewErrParamRequired("RouteIdentifier"))
	}
	if len(v.ActivationState) == 0 {
		invalidParams.Add(smithy.NewErrParamRequired("ActivationState"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}
