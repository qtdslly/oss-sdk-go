// Code generated by smithy-go-codegen DO NOT EDIT.

package types_test

import (
	"fmt"
	"oss-sdk-go/service/transcribestreaming/types"
)

func ExampleAudioStream_outputUsage() {
	var union types.AudioStream
	// type switches can be used to check the union value
	switch v := union.(type) {
	case *types.AudioStreamMemberAudioEvent:
		_ = v.Value // Value is types.AudioEvent

	case *types.UnknownUnionMember:
		fmt.Println("unknown tag:", v.Tag)

	default:
		fmt.Println("union is nil or unknown type")

	}
}

var _ *types.AudioEvent

func ExampleMedicalTranscriptResultStream_outputUsage() {
	var union types.MedicalTranscriptResultStream
	// type switches can be used to check the union value
	switch v := union.(type) {
	case *types.MedicalTranscriptResultStreamMemberTranscriptEvent:
		_ = v.Value // Value is types.MedicalTranscriptEvent

	case *types.UnknownUnionMember:
		fmt.Println("unknown tag:", v.Tag)

	default:
		fmt.Println("union is nil or unknown type")

	}
}

var _ *types.MedicalTranscriptEvent

func ExampleTranscriptResultStream_outputUsage() {
	var union types.TranscriptResultStream
	// type switches can be used to check the union value
	switch v := union.(type) {
	case *types.TranscriptResultStreamMemberTranscriptEvent:
		_ = v.Value // Value is types.TranscriptEvent

	case *types.UnknownUnionMember:
		fmt.Println("unknown tag:", v.Tag)

	default:
		fmt.Println("union is nil or unknown type")

	}
}

var _ *types.TranscriptEvent
