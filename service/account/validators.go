// Code generated by smithy-go-codegen DO NOT EDIT.

package account

import (
	"context"
	"fmt"
	"oss-sdk-go/service/account/types"
	smithy "github.com/aws/smithy-go"
	"github.com/aws/smithy-go/middleware"
)

type validateOpDeleteAlternateContact struct {
}

func (*validateOpDeleteAlternateContact) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpDeleteAlternateContact) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*DeleteAlternateContactInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpDeleteAlternateContactInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpGetAlternateContact struct {
}

func (*validateOpGetAlternateContact) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpGetAlternateContact) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*GetAlternateContactInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpGetAlternateContactInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpPutAlternateContact struct {
}

func (*validateOpPutAlternateContact) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpPutAlternateContact) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*PutAlternateContactInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpPutAlternateContactInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpPutContactInformation struct {
}

func (*validateOpPutContactInformation) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpPutContactInformation) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*PutContactInformationInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpPutContactInformationInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

func addOpDeleteAlternateContactValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpDeleteAlternateContact{}, middleware.After)
}

func addOpGetAlternateContactValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpGetAlternateContact{}, middleware.After)
}

func addOpPutAlternateContactValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpPutAlternateContact{}, middleware.After)
}

func addOpPutContactInformationValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpPutContactInformation{}, middleware.After)
}

func validateContactInformation(v *types.ContactInformation) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "ContactInformation"}
	if v.FullName == nil {
		invalidParams.Add(smithy.NewErrParamRequired("FullName"))
	}
	if v.AddressLine1 == nil {
		invalidParams.Add(smithy.NewErrParamRequired("AddressLine1"))
	}
	if v.City == nil {
		invalidParams.Add(smithy.NewErrParamRequired("City"))
	}
	if v.PostalCode == nil {
		invalidParams.Add(smithy.NewErrParamRequired("PostalCode"))
	}
	if v.CountryCode == nil {
		invalidParams.Add(smithy.NewErrParamRequired("CountryCode"))
	}
	if v.PhoneNumber == nil {
		invalidParams.Add(smithy.NewErrParamRequired("PhoneNumber"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpDeleteAlternateContactInput(v *DeleteAlternateContactInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "DeleteAlternateContactInput"}
	if len(v.AlternateContactType) == 0 {
		invalidParams.Add(smithy.NewErrParamRequired("AlternateContactType"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpGetAlternateContactInput(v *GetAlternateContactInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "GetAlternateContactInput"}
	if len(v.AlternateContactType) == 0 {
		invalidParams.Add(smithy.NewErrParamRequired("AlternateContactType"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpPutAlternateContactInput(v *PutAlternateContactInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "PutAlternateContactInput"}
	if v.Name == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Name"))
	}
	if v.Title == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Title"))
	}
	if v.EmailAddress == nil {
		invalidParams.Add(smithy.NewErrParamRequired("EmailAddress"))
	}
	if v.PhoneNumber == nil {
		invalidParams.Add(smithy.NewErrParamRequired("PhoneNumber"))
	}
	if len(v.AlternateContactType) == 0 {
		invalidParams.Add(smithy.NewErrParamRequired("AlternateContactType"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpPutContactInformationInput(v *PutContactInformationInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "PutContactInformationInput"}
	if v.ContactInformation == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ContactInformation"))
	} else if v.ContactInformation != nil {
		if err := validateContactInformation(v.ContactInformation); err != nil {
			invalidParams.AddNested("ContactInformation", err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}
