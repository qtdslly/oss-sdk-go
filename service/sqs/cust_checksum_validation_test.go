package sqs

import (
	"context"
	"fmt"
	"strings"
	"testing"

	"oss-sdk-go/oss"
	sqstypes "oss-sdk-go/service/sqs/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

func TestValidateSendMessageChecksum(t *testing.T) {
	cases := map[string]struct {
		input      *SendMessageInput
		output     *SendMessageOutput
		handlerErr error

		expectErr string
	}{
		"success": {
			input: &SendMessageInput{
				MessageBody: oss.String("test"),
			},
			output: &SendMessageOutput{
				MD5OfMessageBody: oss.String("098f6bcd4621d373cade4e832627b4f6"),
				MessageId:        oss.String("id"),
			},
		},
		"no input message": {
			input: &SendMessageInput{},
			output: &SendMessageOutput{
				MD5OfMessageBody: oss.String("098f6bcd4621d373cade4e832627b4f6"),
				MessageId:        oss.String("id"),
			},
		},
		"no md5": {
			input: &SendMessageInput{
				MessageBody: oss.String("test"),
			},
			output: &SendMessageOutput{
				MessageId: oss.String("id"),
			},
		},
		"no message id": {
			input: &SendMessageInput{
				MessageBody: oss.String("test"),
			},
			output: &SendMessageOutput{
				MD5OfMessageBody: oss.String("098f6bcd4621d373cade4e832627b4f6"),
			},
		},
		"invalid checksum": {
			input: &SendMessageInput{
				MessageBody: oss.String("test"),
			},
			output: &SendMessageOutput{
				MD5OfMessageBody: oss.String("01234556"),
				MessageId:        oss.String("id"),
			},
			expectErr: "message id has invalid checksum",
		},
		"response error": {
			input: &SendMessageInput{
				MessageBody: oss.String("test"),
			},
			handlerErr: fmt.Errorf("some error"),
			expectErr:  "some error",
		},
	}

	for name, c := range cases {
		t.Run(name, func(t *testing.T) {
			m := validateMessageChecksumMiddleware{
				validate: validateSendMessageChecksum,
			}

			next := mockInitializeHandler{
				Output: middleware.InitializeOutput{Result: c.output},
				Err:    c.handlerErr,
			}
			input := middleware.InitializeInput{
				Parameters: c.input,
			}
			_, _, err := m.HandleInitialize(context.Background(), input, next)
			if c.expectErr != "" {
				if err == nil {
					t.Fatalf("expect %v error, got none", c.expectErr)
				}
				if e, a := c.expectErr, err.Error(); !strings.Contains(a, e) {
					t.Fatalf("expect %v error, got %v", e, a)
				}
				return
			}
			if err != nil {
				t.Fatalf("expect no error, got %v", err)
			}
		})
	}
}

func TestValidateSendMessageBatchChecksum(t *testing.T) {
	successMD5 := "098f6bcd4621d373cade4e832627b4f6"
	invalidMD5 := "11111111111111111111111111111111"

	cases := map[string]struct {
		input      *SendMessageBatchInput
		output     *SendMessageBatchOutput
		handlerErr error

		expectErrs []string
	}{
		"success": {
			input: &SendMessageBatchInput{
				Entries: []sqstypes.SendMessageBatchRequestEntry{
					{Id: oss.String("1"), MessageBody: oss.String("test")},
					{Id: oss.String("2"), MessageBody: oss.String("test")},
					{Id: oss.String("3"), MessageBody: oss.String("test")},
					{Id: oss.String("4"), MessageBody: oss.String("test")},
				},
			},
			output: &SendMessageBatchOutput{
				Successful: []sqstypes.SendMessageBatchResultEntry{
					{MD5OfMessageBody: &successMD5, MessageId: oss.String("123"), Id: oss.String("1")},
					{MD5OfMessageBody: &successMD5, MessageId: oss.String("456"), Id: oss.String("2")},
					{MD5OfMessageBody: &successMD5, MessageId: oss.String("789"), Id: oss.String("3")},
					{MD5OfMessageBody: &successMD5, MessageId: oss.String("012"), Id: oss.String("4")},
				},
			},
		},
		"no input body": {
			input: &SendMessageBatchInput{
				Entries: []sqstypes.SendMessageBatchRequestEntry{
					{Id: oss.String("1")},
				},
			},
			output: &SendMessageBatchOutput{
				Successful: []sqstypes.SendMessageBatchResultEntry{
					{MD5OfMessageBody: &invalidMD5, MessageId: oss.String("123"), Id: oss.String("1")},
				},
			},
		},
		"no md5": {
			input: &SendMessageBatchInput{
				Entries: []sqstypes.SendMessageBatchRequestEntry{
					{Id: oss.String("1"), MessageBody: oss.String("test")},
				},
			},
			output: &SendMessageBatchOutput{
				Successful: []sqstypes.SendMessageBatchResultEntry{
					{MessageId: oss.String("123"), Id: oss.String("1")},
				},
			},
		},
		"server side failure": {
			input: &SendMessageBatchInput{
				Entries: []sqstypes.SendMessageBatchRequestEntry{
					{Id: oss.String("1"), MessageBody: oss.String("test")},
					{Id: oss.String("2"), MessageBody: oss.String("test")},
					{Id: oss.String("3"), MessageBody: oss.String("test")},
					{Id: oss.String("4"), MessageBody: oss.String("test")},
				},
			},
			output: &SendMessageBatchOutput{
				Successful: []sqstypes.SendMessageBatchResultEntry{
					{MD5OfMessageBody: &successMD5, MessageId: oss.String("123"), Id: oss.String("1")},
					{MD5OfMessageBody: &successMD5, MessageId: oss.String("456"), Id: oss.String("2")},
					{MD5OfMessageBody: &successMD5, MessageId: oss.String("012"), Id: oss.String("4")},
				},
				Failed: []sqstypes.BatchResultErrorEntry{
					{Id: oss.String("3"), Code: oss.String("test"), Message: oss.String("test")},
				},
			},
		},
		"partial invalid checksum": {
			input: &SendMessageBatchInput{
				Entries: []sqstypes.SendMessageBatchRequestEntry{
					{Id: oss.String("1"), MessageBody: oss.String("test")},
					{Id: oss.String("2"), MessageBody: oss.String("test")},
					{Id: oss.String("3"), MessageBody: oss.String("test")},
					{Id: oss.String("4"), MessageBody: oss.String("test")},
				},
			},
			output: &SendMessageBatchOutput{
				Successful: []sqstypes.SendMessageBatchResultEntry{
					{MD5OfMessageBody: &successMD5, MessageId: oss.String("123"), Id: oss.String("1")},
					{MD5OfMessageBody: &successMD5, MessageId: oss.String("456"), Id: oss.String("2")},
					{MD5OfMessageBody: &invalidMD5, MessageId: oss.String("789"), Id: oss.String("3")},
					{MD5OfMessageBody: &successMD5, MessageId: oss.String("012"), Id: oss.String("4")},
				},
			},
			expectErrs: []string{"message 789 has invalid checksum"},
		},
		"complete invalid checksum": {
			input: &SendMessageBatchInput{
				Entries: []sqstypes.SendMessageBatchRequestEntry{
					{Id: oss.String("1"), MessageBody: oss.String("test")},
					{Id: oss.String("2"), MessageBody: oss.String("test")},
					{Id: oss.String("3"), MessageBody: oss.String("test")},
					{Id: oss.String("4"), MessageBody: oss.String("test")},
				},
			},
			output: &SendMessageBatchOutput{
				Successful: []sqstypes.SendMessageBatchResultEntry{
					{MD5OfMessageBody: &invalidMD5, MessageId: oss.String("123"), Id: oss.String("1")},
					{MD5OfMessageBody: &invalidMD5, MessageId: oss.String("456"), Id: oss.String("2")},
					{MD5OfMessageBody: &invalidMD5, MessageId: oss.String("789"), Id: oss.String("3")},
					{MD5OfMessageBody: &invalidMD5, MessageId: oss.String("012"), Id: oss.String("4")},
				},
			},
			expectErrs: []string{
				"message 123 has invalid checksum",
				"message 456 has invalid checksum",
				"message 789 has invalid checksum",
				"message 012 has invalid checksum",
			},
		},
		"invalid checksum no message id": {
			input: &SendMessageBatchInput{
				Entries: []sqstypes.SendMessageBatchRequestEntry{
					{Id: oss.String("1"), MessageBody: oss.String("test")},
				},
			},
			output: &SendMessageBatchOutput{
				Successful: []sqstypes.SendMessageBatchResultEntry{
					{MD5OfMessageBody: &invalidMD5, Id: oss.String("1")},
				},
			},
			expectErrs: []string{"message has invalid checksum"},
		},
		"response error": {
			input: &SendMessageBatchInput{
				Entries: []sqstypes.SendMessageBatchRequestEntry{
					{Id: oss.String("1"), MessageBody: oss.String("test")},
				},
			},
			handlerErr: fmt.Errorf("some error"),
			expectErrs: []string{"some error"},
		},
	}

	for name, c := range cases {
		t.Run(name, func(t *testing.T) {
			m := validateMessageChecksumMiddleware{
				validate: validateSendMessageBatchChecksum,
			}

			next := mockInitializeHandler{
				Output: middleware.InitializeOutput{Result: c.output},
				Err:    c.handlerErr,
			}
			input := middleware.InitializeInput{
				Parameters: c.input,
			}
			_, _, err := m.HandleInitialize(context.Background(), input, next)
			if len(c.expectErrs) != 0 {
				if err == nil {
					t.Fatalf("expect error(s), got none")
				}
				for i, expectErr := range c.expectErrs {
					if e, a := expectErr, err.Error(); !strings.Contains(a, e) {
						t.Errorf("%d expect %v error, got %v", i, e, a)
					}
				}
				return
			}
			if err != nil {
				t.Fatalf("expect no error, got %v", err)
			}
		})
	}
}

func TestValidateReceiveMessageChecksum(t *testing.T) {
	successMD5 := "098f6bcd4621d373cade4e832627b4f6"
	invalidMD5 := "11111111111111111111111111111111"

	cases := map[string]struct {
		output     *ReceiveMessageOutput
		handlerErr error

		expectErrs []string
	}{
		"success": {
			output: &ReceiveMessageOutput{
				Messages: []sqstypes.Message{
					{Body: oss.String("test"), MD5OfBody: &successMD5},
					{Body: oss.String("test"), MD5OfBody: &successMD5},
					{Body: oss.String("test"), MD5OfBody: &successMD5},
					{Body: oss.String("test"), MD5OfBody: &successMD5},
				},
			},
		},
		"no body": {
			output: &ReceiveMessageOutput{
				Messages: []sqstypes.Message{
					{MD5OfBody: &successMD5},
				},
			},
		},
		"no md5": {
			output: &ReceiveMessageOutput{
				Messages: []sqstypes.Message{
					{Body: oss.String("test")},
				},
			},
		},
		"message with no ID partial invalid checksum": {
			output: &ReceiveMessageOutput{
				Messages: []sqstypes.Message{
					{Body: oss.String("test"), MD5OfBody: &successMD5},
					{Body: oss.String("test"), MD5OfBody: &successMD5},
					{Body: oss.String("test"), MD5OfBody: &invalidMD5},
					{Body: oss.String("test"), MD5OfBody: &successMD5},
				},
			},
			expectErrs: []string{"message has invalid checksum"},
		},
		"message with ID partial invalid checksum": {
			output: &ReceiveMessageOutput{
				Messages: []sqstypes.Message{
					{Body: oss.String("test"), MD5OfBody: &successMD5},
					{Body: oss.String("test"), MD5OfBody: &successMD5},
					{Body: oss.String("test"), MD5OfBody: &invalidMD5, MessageId: oss.String("123")},
					{Body: oss.String("test"), MD5OfBody: &successMD5},
				},
			},
			expectErrs: []string{"message 123 has invalid checksum"},
		},
		"complete invalid checksum": {
			output: &ReceiveMessageOutput{
				Messages: []sqstypes.Message{
					{Body: oss.String("test"), MD5OfBody: &invalidMD5},
					{Body: oss.String("test"), MD5OfBody: &invalidMD5, MessageId: oss.String("123")},
					{Body: oss.String("test"), MD5OfBody: &invalidMD5, MessageId: oss.String("456")},
					{Body: oss.String("test"), MD5OfBody: &invalidMD5},
				},
			},
			expectErrs: []string{
				"message has invalid checksum",
				"message 123 has invalid checksum",
				"message 456 has invalid checksum",
				"message has invalid checksum",
			},
		},
		"response error": {
			handlerErr: fmt.Errorf("some error"),
			expectErrs: []string{"some error"},
		},
	}

	for name, c := range cases {
		t.Run(name, func(t *testing.T) {
			m := validateMessageChecksumMiddleware{
				validate: validateReceiveMessageChecksum,
			}

			next := mockInitializeHandler{
				Output: middleware.InitializeOutput{Result: c.output},
				Err:    c.handlerErr,
			}

			input := middleware.InitializeInput{
				Parameters: &ReceiveMessageInput{},
			}
			_, _, err := m.HandleInitialize(context.Background(), input, next)
			if len(c.expectErrs) != 0 {
				if err == nil {
					t.Fatalf("expect error(s), got none")
				}
				for i, expectErr := range c.expectErrs {
					if e, a := expectErr, err.Error(); !strings.Contains(a, e) {
						t.Errorf("%d expect %v error, got %v", i, e, a)
					}
				}
				return
			}
			if err != nil {
				t.Fatalf("expect no error, got %v", err)
			}
		})
	}
}

func TestAddValidateSendMessageChecksum(t *testing.T) {
	cases := map[string]map[string]struct {
		options       Options
		fn            func(*middleware.Stack, Options) error
		expectInStack bool
	}{
		"SendMessage": {
			"enabled": {
				options:       Options{},
				fn:            addValidateSendMessageChecksum,
				expectInStack: true,
			},
			"disabled": {
				options: Options{
					DisableMessageChecksumValidation: true,
				},
				fn:            addValidateSendMessageChecksum,
				expectInStack: false,
			},
		},
		"SendMessageBatch": {
			"enabled": {
				options:       Options{},
				fn:            addValidateSendMessageBatchChecksum,
				expectInStack: true,
			},
			"disabled": {
				options: Options{
					DisableMessageChecksumValidation: true,
				},
				fn:            addValidateSendMessageBatchChecksum,
				expectInStack: false,
			},
		},
		"ReceiveMessage": {
			"enabled": {
				options:       Options{},
				fn:            addValidateReceiveMessageChecksum,
				expectInStack: true,
			},
			"disabled": {
				options: Options{
					DisableMessageChecksumValidation: true,
				},
				fn:            addValidateReceiveMessageChecksum,
				expectInStack: false,
			},
		},
	}

	for opName, opCases := range cases {
		t.Run(opName, func(t *testing.T) {
			for name, c := range opCases {
				t.Run(name, func(t *testing.T) {
					options := c.options.Copy()
					stack := middleware.NewStack("test", smithyhttp.NewStackRequest)

					err := c.fn(stack, options)
					if err != nil {
						t.Fatalf("expect no error, got %v", err)
					}

					expectID := validateMessageChecksumMiddleware{}.ID()
					if e, a := expectID, stack.String(); c.expectInStack != strings.Contains(a, e) {
						t.Errorf("expect %v in stack %v:\n%s", e, c.expectInStack, a)
					}
				})
			}
		})
	}
}

// ******************
// Testing Utilities
// ******************
type mockInitializeHandler struct {
	Output middleware.InitializeOutput
	Err    error
}

func (m mockInitializeHandler) HandleInitialize(
	ctx context.Context, in middleware.InitializeInput,
) (
	out middleware.InitializeOutput, meta middleware.Metadata, err error,
) {
	return m.Output, meta, m.Err
}
