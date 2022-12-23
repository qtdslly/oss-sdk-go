// Code generated by smithy-go-codegen DO NOT EDIT.

package connectcontactlens

import (
	"context"
	"fmt"
	smithy "github.com/aws/smithy-go"
	"github.com/aws/smithy-go/middleware"
)

type validateOpListRealtimeContactAnalysisSegments struct {
}

func (*validateOpListRealtimeContactAnalysisSegments) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpListRealtimeContactAnalysisSegments) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*ListRealtimeContactAnalysisSegmentsInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpListRealtimeContactAnalysisSegmentsInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

func addOpListRealtimeContactAnalysisSegmentsValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpListRealtimeContactAnalysisSegments{}, middleware.After)
}

func validateOpListRealtimeContactAnalysisSegmentsInput(v *ListRealtimeContactAnalysisSegmentsInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "ListRealtimeContactAnalysisSegmentsInput"}
	if v.InstanceId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("InstanceId"))
	}
	if v.ContactId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ContactId"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}
