// Code generated by smithy-go-codegen DO NOT EDIT.

package types_test

import (
	"fmt"
	"oss-sdk-go/service/s3/types"
)

func ExampleAnalyticsFilter_outputUsage() {
	var union types.AnalyticsFilter
	// type switches can be used to check the union value
	switch v := union.(type) {
	case *types.AnalyticsFilterMemberAnd:
		_ = v.Value // Value is types.AnalyticsAndOperator

	case *types.AnalyticsFilterMemberPrefix:
		_ = v.Value // Value is string

	case *types.AnalyticsFilterMemberTag:
		_ = v.Value // Value is types.Tag

	case *types.UnknownUnionMember:
		fmt.Println("unknown tag:", v.Tag)

	default:
		fmt.Println("union is nil or unknown type")

	}
}

var _ *string
var _ *types.Tag
var _ *types.AnalyticsAndOperator

func ExampleLifecycleRuleFilter_outputUsage() {
	var union types.LifecycleRuleFilter
	// type switches can be used to check the union value
	switch v := union.(type) {
	case *types.LifecycleRuleFilterMemberAnd:
		_ = v.Value // Value is types.LifecycleRuleAndOperator

	case *types.LifecycleRuleFilterMemberObjectSizeGreaterThan:
		_ = v.Value // Value is int64

	case *types.LifecycleRuleFilterMemberObjectSizeLessThan:
		_ = v.Value // Value is int64

	case *types.LifecycleRuleFilterMemberPrefix:
		_ = v.Value // Value is string

	case *types.LifecycleRuleFilterMemberTag:
		_ = v.Value // Value is types.Tag

	case *types.UnknownUnionMember:
		fmt.Println("unknown tag:", v.Tag)

	default:
		fmt.Println("union is nil or unknown type")

	}
}

var _ *string
var _ *types.LifecycleRuleAndOperator
var _ int64
var _ int64
var _ *types.Tag

func ExampleMetricsFilter_outputUsage() {
	var union types.MetricsFilter
	// type switches can be used to check the union value
	switch v := union.(type) {
	case *types.MetricsFilterMemberAccessPointArn:
		_ = v.Value // Value is string

	case *types.MetricsFilterMemberAnd:
		_ = v.Value // Value is types.MetricsAndOperator

	case *types.MetricsFilterMemberPrefix:
		_ = v.Value // Value is string

	case *types.MetricsFilterMemberTag:
		_ = v.Value // Value is types.Tag

	case *types.UnknownUnionMember:
		fmt.Println("unknown tag:", v.Tag)

	default:
		fmt.Println("union is nil or unknown type")

	}
}

var _ *string
var _ *string
var _ *types.Tag
var _ *types.MetricsAndOperator

func ExampleReplicationRuleFilter_outputUsage() {
	var union types.ReplicationRuleFilter
	// type switches can be used to check the union value
	switch v := union.(type) {
	case *types.ReplicationRuleFilterMemberAnd:
		_ = v.Value // Value is types.ReplicationRuleAndOperator

	case *types.ReplicationRuleFilterMemberPrefix:
		_ = v.Value // Value is string

	case *types.ReplicationRuleFilterMemberTag:
		_ = v.Value // Value is types.Tag

	case *types.UnknownUnionMember:
		fmt.Println("unknown tag:", v.Tag)

	default:
		fmt.Println("union is nil or unknown type")

	}
}

var _ *string
var _ *types.Tag
var _ *types.ReplicationRuleAndOperator

func ExampleSelectObjectContentEventStream_outputUsage() {
	var union types.SelectObjectContentEventStream
	// type switches can be used to check the union value
	switch v := union.(type) {
	case *types.SelectObjectContentEventStreamMemberCont:
		_ = v.Value // Value is types.ContinuationEvent

	case *types.SelectObjectContentEventStreamMemberEnd:
		_ = v.Value // Value is types.EndEvent

	case *types.SelectObjectContentEventStreamMemberProgress:
		_ = v.Value // Value is types.ProgressEvent

	case *types.SelectObjectContentEventStreamMemberRecords:
		_ = v.Value // Value is types.RecordsEvent

	case *types.SelectObjectContentEventStreamMemberStats:
		_ = v.Value // Value is types.StatsEvent

	case *types.UnknownUnionMember:
		fmt.Println("unknown tag:", v.Tag)

	default:
		fmt.Println("union is nil or unknown type")

	}
}

var _ *types.RecordsEvent
var _ *types.StatsEvent
var _ *types.ContinuationEvent
var _ *types.EndEvent
var _ *types.ProgressEvent
