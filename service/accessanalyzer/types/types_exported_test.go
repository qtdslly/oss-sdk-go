// Code generated by smithy-go-codegen DO NOT EDIT.

package types_test

import (
	"fmt"
	"oss-sdk-go/service/accessanalyzer/types"
)

func ExampleAclGrantee_outputUsage() {
	var union types.AclGrantee
	// type switches can be used to check the union value
	switch v := union.(type) {
	case *types.AclGranteeMemberId:
		_ = v.Value // Value is string

	case *types.AclGranteeMemberUri:
		_ = v.Value // Value is string

	case *types.UnknownUnionMember:
		fmt.Println("unknown tag:", v.Tag)

	default:
		fmt.Println("union is nil or unknown type")

	}
}

var _ *string
var _ *string

func ExampleConfiguration_outputUsage() {
	var union types.Configuration
	// type switches can be used to check the union value
	switch v := union.(type) {
	case *types.ConfigurationMemberIamRole:
		_ = v.Value // Value is types.IamRoleConfiguration

	case *types.ConfigurationMemberKmsKey:
		_ = v.Value // Value is types.KmsKeyConfiguration

	case *types.ConfigurationMemberS3Bucket:
		_ = v.Value // Value is types.S3BucketConfiguration

	case *types.ConfigurationMemberSecretsManagerSecret:
		_ = v.Value // Value is types.SecretsManagerSecretConfiguration

	case *types.ConfigurationMemberSqsQueue:
		_ = v.Value // Value is types.SqsQueueConfiguration

	case *types.UnknownUnionMember:
		fmt.Println("unknown tag:", v.Tag)

	default:
		fmt.Println("union is nil or unknown type")

	}
}

var _ *types.SqsQueueConfiguration
var _ *types.IamRoleConfiguration
var _ *types.SecretsManagerSecretConfiguration
var _ *types.S3BucketConfiguration
var _ *types.KmsKeyConfiguration

func ExampleNetworkOriginConfiguration_outputUsage() {
	var union types.NetworkOriginConfiguration
	// type switches can be used to check the union value
	switch v := union.(type) {
	case *types.NetworkOriginConfigurationMemberInternetConfiguration:
		_ = v.Value // Value is types.InternetConfiguration

	case *types.NetworkOriginConfigurationMemberVpcConfiguration:
		_ = v.Value // Value is types.VpcConfiguration

	case *types.UnknownUnionMember:
		fmt.Println("unknown tag:", v.Tag)

	default:
		fmt.Println("union is nil or unknown type")

	}
}

var _ *types.VpcConfiguration
var _ *types.InternetConfiguration

func ExamplePathElement_outputUsage() {
	var union types.PathElement
	// type switches can be used to check the union value
	switch v := union.(type) {
	case *types.PathElementMemberIndex:
		_ = v.Value // Value is int32

	case *types.PathElementMemberKey:
		_ = v.Value // Value is string

	case *types.PathElementMemberSubstring:
		_ = v.Value // Value is types.Substring

	case *types.PathElementMemberValue:
		_ = v.Value // Value is string

	case *types.UnknownUnionMember:
		fmt.Println("unknown tag:", v.Tag)

	default:
		fmt.Println("union is nil or unknown type")

	}
}

var _ *string
var _ *int32
var _ *types.Substring
