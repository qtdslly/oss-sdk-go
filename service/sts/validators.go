// Code generated by smithy-go-codegen DO NOT EDIT.

package sts

import (
	"context"
	"fmt"
	"oss-sdk-go/service/sts/types"
	smithy "github.com/aws/smithy-go"
	"github.com/aws/smithy-go/middleware"
)

type validateOpAssumeRole struct {
}

func (*validateOpAssumeRole) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpAssumeRole) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*AssumeRoleInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpAssumeRoleInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpAssumeRoleWithSAML struct {
}

func (*validateOpAssumeRoleWithSAML) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpAssumeRoleWithSAML) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*AssumeRoleWithSAMLInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpAssumeRoleWithSAMLInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpAssumeRoleWithWebIdentity struct {
}

func (*validateOpAssumeRoleWithWebIdentity) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpAssumeRoleWithWebIdentity) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*AssumeRoleWithWebIdentityInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpAssumeRoleWithWebIdentityInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpDecodeAuthorizationMessage struct {
}

func (*validateOpDecodeAuthorizationMessage) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpDecodeAuthorizationMessage) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*DecodeAuthorizationMessageInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpDecodeAuthorizationMessageInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpGetAccessKeyInfo struct {
}

func (*validateOpGetAccessKeyInfo) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpGetAccessKeyInfo) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*GetAccessKeyInfoInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpGetAccessKeyInfoInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpGetFederationToken struct {
}

func (*validateOpGetFederationToken) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpGetFederationToken) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*GetFederationTokenInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpGetFederationTokenInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

func addOpAssumeRoleValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpAssumeRole{}, middleware.After)
}

func addOpAssumeRoleWithSAMLValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpAssumeRoleWithSAML{}, middleware.After)
}

func addOpAssumeRoleWithWebIdentityValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpAssumeRoleWithWebIdentity{}, middleware.After)
}

func addOpDecodeAuthorizationMessageValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpDecodeAuthorizationMessage{}, middleware.After)
}

func addOpGetAccessKeyInfoValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpGetAccessKeyInfo{}, middleware.After)
}

func addOpGetFederationTokenValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpGetFederationToken{}, middleware.After)
}

func validateTag(v *types.Tag) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "Tag"}
	if v.Key == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Key"))
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

func validateTagListType(v []types.Tag) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "TagListType"}
	for i := range v {
		if err := validateTag(&v[i]); err != nil {
			invalidParams.AddNested(fmt.Sprintf("[%d]", i), err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpAssumeRoleInput(v *AssumeRoleInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "AssumeRoleInput"}
	if v.RoleArn == nil {
		invalidParams.Add(smithy.NewErrParamRequired("RoleArn"))
	}
	if v.RoleSessionName == nil {
		invalidParams.Add(smithy.NewErrParamRequired("RoleSessionName"))
	}
	if v.Tags != nil {
		if err := validateTagListType(v.Tags); err != nil {
			invalidParams.AddNested("Tags", err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpAssumeRoleWithSAMLInput(v *AssumeRoleWithSAMLInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "AssumeRoleWithSAMLInput"}
	if v.RoleArn == nil {
		invalidParams.Add(smithy.NewErrParamRequired("RoleArn"))
	}
	if v.PrincipalArn == nil {
		invalidParams.Add(smithy.NewErrParamRequired("PrincipalArn"))
	}
	if v.SAMLAssertion == nil {
		invalidParams.Add(smithy.NewErrParamRequired("SAMLAssertion"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpAssumeRoleWithWebIdentityInput(v *AssumeRoleWithWebIdentityInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "AssumeRoleWithWebIdentityInput"}
	if v.RoleArn == nil {
		invalidParams.Add(smithy.NewErrParamRequired("RoleArn"))
	}
	if v.RoleSessionName == nil {
		invalidParams.Add(smithy.NewErrParamRequired("RoleSessionName"))
	}
	if v.WebIdentityToken == nil {
		invalidParams.Add(smithy.NewErrParamRequired("WebIdentityToken"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpDecodeAuthorizationMessageInput(v *DecodeAuthorizationMessageInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "DecodeAuthorizationMessageInput"}
	if v.EncodedMessage == nil {
		invalidParams.Add(smithy.NewErrParamRequired("EncodedMessage"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpGetAccessKeyInfoInput(v *GetAccessKeyInfoInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "GetAccessKeyInfoInput"}
	if v.AccessKeyId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("AccessKeyId"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpGetFederationTokenInput(v *GetFederationTokenInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "GetFederationTokenInput"}
	if v.Name == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Name"))
	}
	if v.Tags != nil {
		if err := validateTagListType(v.Tags); err != nil {
			invalidParams.AddNested("Tags", err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}
