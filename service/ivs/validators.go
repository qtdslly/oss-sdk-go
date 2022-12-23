// Code generated by smithy-go-codegen DO NOT EDIT.

package ivs

import (
	"context"
	"fmt"
	"oss-sdk-go/service/ivs/types"
	smithy "github.com/aws/smithy-go"
	"github.com/aws/smithy-go/middleware"
)

type validateOpBatchGetChannel struct {
}

func (*validateOpBatchGetChannel) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpBatchGetChannel) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*BatchGetChannelInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpBatchGetChannelInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpBatchGetStreamKey struct {
}

func (*validateOpBatchGetStreamKey) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpBatchGetStreamKey) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*BatchGetStreamKeyInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpBatchGetStreamKeyInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpCreateRecordingConfiguration struct {
}

func (*validateOpCreateRecordingConfiguration) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpCreateRecordingConfiguration) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*CreateRecordingConfigurationInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpCreateRecordingConfigurationInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpCreateStreamKey struct {
}

func (*validateOpCreateStreamKey) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpCreateStreamKey) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*CreateStreamKeyInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpCreateStreamKeyInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpDeleteChannel struct {
}

func (*validateOpDeleteChannel) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpDeleteChannel) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*DeleteChannelInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpDeleteChannelInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpDeletePlaybackKeyPair struct {
}

func (*validateOpDeletePlaybackKeyPair) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpDeletePlaybackKeyPair) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*DeletePlaybackKeyPairInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpDeletePlaybackKeyPairInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpDeleteRecordingConfiguration struct {
}

func (*validateOpDeleteRecordingConfiguration) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpDeleteRecordingConfiguration) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*DeleteRecordingConfigurationInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpDeleteRecordingConfigurationInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpDeleteStreamKey struct {
}

func (*validateOpDeleteStreamKey) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpDeleteStreamKey) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*DeleteStreamKeyInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpDeleteStreamKeyInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpGetChannel struct {
}

func (*validateOpGetChannel) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpGetChannel) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*GetChannelInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpGetChannelInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpGetPlaybackKeyPair struct {
}

func (*validateOpGetPlaybackKeyPair) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpGetPlaybackKeyPair) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*GetPlaybackKeyPairInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpGetPlaybackKeyPairInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpGetRecordingConfiguration struct {
}

func (*validateOpGetRecordingConfiguration) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpGetRecordingConfiguration) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*GetRecordingConfigurationInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpGetRecordingConfigurationInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpGetStream struct {
}

func (*validateOpGetStream) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpGetStream) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*GetStreamInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpGetStreamInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpGetStreamKey struct {
}

func (*validateOpGetStreamKey) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpGetStreamKey) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*GetStreamKeyInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpGetStreamKeyInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpGetStreamSession struct {
}

func (*validateOpGetStreamSession) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpGetStreamSession) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*GetStreamSessionInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpGetStreamSessionInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpImportPlaybackKeyPair struct {
}

func (*validateOpImportPlaybackKeyPair) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpImportPlaybackKeyPair) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*ImportPlaybackKeyPairInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpImportPlaybackKeyPairInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpListStreamKeys struct {
}

func (*validateOpListStreamKeys) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpListStreamKeys) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*ListStreamKeysInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpListStreamKeysInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpListStreamSessions struct {
}

func (*validateOpListStreamSessions) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpListStreamSessions) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*ListStreamSessionsInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpListStreamSessionsInput(input); err != nil {
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

type validateOpPutMetadata struct {
}

func (*validateOpPutMetadata) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpPutMetadata) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*PutMetadataInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpPutMetadataInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpStopStream struct {
}

func (*validateOpStopStream) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpStopStream) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*StopStreamInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpStopStreamInput(input); err != nil {
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

type validateOpUpdateChannel struct {
}

func (*validateOpUpdateChannel) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpUpdateChannel) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*UpdateChannelInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpUpdateChannelInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

func addOpBatchGetChannelValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpBatchGetChannel{}, middleware.After)
}

func addOpBatchGetStreamKeyValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpBatchGetStreamKey{}, middleware.After)
}

func addOpCreateRecordingConfigurationValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpCreateRecordingConfiguration{}, middleware.After)
}

func addOpCreateStreamKeyValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpCreateStreamKey{}, middleware.After)
}

func addOpDeleteChannelValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpDeleteChannel{}, middleware.After)
}

func addOpDeletePlaybackKeyPairValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpDeletePlaybackKeyPair{}, middleware.After)
}

func addOpDeleteRecordingConfigurationValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpDeleteRecordingConfiguration{}, middleware.After)
}

func addOpDeleteStreamKeyValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpDeleteStreamKey{}, middleware.After)
}

func addOpGetChannelValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpGetChannel{}, middleware.After)
}

func addOpGetPlaybackKeyPairValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpGetPlaybackKeyPair{}, middleware.After)
}

func addOpGetRecordingConfigurationValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpGetRecordingConfiguration{}, middleware.After)
}

func addOpGetStreamValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpGetStream{}, middleware.After)
}

func addOpGetStreamKeyValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpGetStreamKey{}, middleware.After)
}

func addOpGetStreamSessionValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpGetStreamSession{}, middleware.After)
}

func addOpImportPlaybackKeyPairValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpImportPlaybackKeyPair{}, middleware.After)
}

func addOpListStreamKeysValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpListStreamKeys{}, middleware.After)
}

func addOpListStreamSessionsValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpListStreamSessions{}, middleware.After)
}

func addOpListTagsForResourceValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpListTagsForResource{}, middleware.After)
}

func addOpPutMetadataValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpPutMetadata{}, middleware.After)
}

func addOpStopStreamValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpStopStream{}, middleware.After)
}

func addOpTagResourceValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpTagResource{}, middleware.After)
}

func addOpUntagResourceValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpUntagResource{}, middleware.After)
}

func addOpUpdateChannelValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpUpdateChannel{}, middleware.After)
}

func validateDestinationConfiguration(v *types.DestinationConfiguration) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "DestinationConfiguration"}
	if v.S3 != nil {
		if err := validateS3DestinationConfiguration(v.S3); err != nil {
			invalidParams.AddNested("S3", err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateS3DestinationConfiguration(v *types.S3DestinationConfiguration) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "S3DestinationConfiguration"}
	if v.BucketName == nil {
		invalidParams.Add(smithy.NewErrParamRequired("BucketName"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpBatchGetChannelInput(v *BatchGetChannelInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "BatchGetChannelInput"}
	if v.Arns == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Arns"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpBatchGetStreamKeyInput(v *BatchGetStreamKeyInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "BatchGetStreamKeyInput"}
	if v.Arns == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Arns"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpCreateRecordingConfigurationInput(v *CreateRecordingConfigurationInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "CreateRecordingConfigurationInput"}
	if v.DestinationConfiguration == nil {
		invalidParams.Add(smithy.NewErrParamRequired("DestinationConfiguration"))
	} else if v.DestinationConfiguration != nil {
		if err := validateDestinationConfiguration(v.DestinationConfiguration); err != nil {
			invalidParams.AddNested("DestinationConfiguration", err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpCreateStreamKeyInput(v *CreateStreamKeyInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "CreateStreamKeyInput"}
	if v.ChannelArn == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ChannelArn"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpDeleteChannelInput(v *DeleteChannelInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "DeleteChannelInput"}
	if v.Arn == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Arn"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpDeletePlaybackKeyPairInput(v *DeletePlaybackKeyPairInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "DeletePlaybackKeyPairInput"}
	if v.Arn == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Arn"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpDeleteRecordingConfigurationInput(v *DeleteRecordingConfigurationInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "DeleteRecordingConfigurationInput"}
	if v.Arn == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Arn"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpDeleteStreamKeyInput(v *DeleteStreamKeyInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "DeleteStreamKeyInput"}
	if v.Arn == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Arn"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpGetChannelInput(v *GetChannelInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "GetChannelInput"}
	if v.Arn == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Arn"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpGetPlaybackKeyPairInput(v *GetPlaybackKeyPairInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "GetPlaybackKeyPairInput"}
	if v.Arn == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Arn"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpGetRecordingConfigurationInput(v *GetRecordingConfigurationInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "GetRecordingConfigurationInput"}
	if v.Arn == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Arn"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpGetStreamInput(v *GetStreamInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "GetStreamInput"}
	if v.ChannelArn == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ChannelArn"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpGetStreamKeyInput(v *GetStreamKeyInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "GetStreamKeyInput"}
	if v.Arn == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Arn"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpGetStreamSessionInput(v *GetStreamSessionInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "GetStreamSessionInput"}
	if v.ChannelArn == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ChannelArn"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpImportPlaybackKeyPairInput(v *ImportPlaybackKeyPairInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "ImportPlaybackKeyPairInput"}
	if v.PublicKeyMaterial == nil {
		invalidParams.Add(smithy.NewErrParamRequired("PublicKeyMaterial"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpListStreamKeysInput(v *ListStreamKeysInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "ListStreamKeysInput"}
	if v.ChannelArn == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ChannelArn"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpListStreamSessionsInput(v *ListStreamSessionsInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "ListStreamSessionsInput"}
	if v.ChannelArn == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ChannelArn"))
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

func validateOpPutMetadataInput(v *PutMetadataInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "PutMetadataInput"}
	if v.ChannelArn == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ChannelArn"))
	}
	if v.Metadata == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Metadata"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpStopStreamInput(v *StopStreamInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "StopStreamInput"}
	if v.ChannelArn == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ChannelArn"))
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

func validateOpUpdateChannelInput(v *UpdateChannelInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "UpdateChannelInput"}
	if v.Arn == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Arn"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}
