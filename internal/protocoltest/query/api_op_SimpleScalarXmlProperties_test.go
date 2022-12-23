// Code generated by smithy-go-codegen DO NOT EDIT.

package query

import (
	"bytes"
	"context"
	"oss-sdk-go/oss"
	smithydocument "github.com/aws/smithy-go/document"
	"github.com/aws/smithy-go/middleware"
	"github.com/aws/smithy-go/ptr"
	smithyrand "github.com/aws/smithy-go/rand"
	smithytesting "github.com/aws/smithy-go/testing"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"io/ioutil"
	"math"
	"net/http"
	"testing"
)

func TestClient_SimpleScalarXmlProperties_awsAwsqueryDeserialize(t *testing.T) {
	cases := map[string]struct {
		StatusCode    int
		Header        http.Header
		BodyMediaType string
		Body          []byte
		ExpectResult  *SimpleScalarXmlPropertiesOutput
	}{
		// Serializes simple scalar properties
		"QuerySimpleScalarProperties": {
			StatusCode: 200,
			Header: http.Header{
				"Content-Type": []string{"text/xml"},
			},
			BodyMediaType: "application/xml",
			Body: []byte(`<SimpleScalarXmlPropertiesResponse xmlns="https://example.com/">
			    <SimpleScalarXmlPropertiesResult>
			        <stringValue>string</stringValue>
			        <emptyStringValue/>
			        <trueBooleanValue>true</trueBooleanValue>
			        <falseBooleanValue>false</falseBooleanValue>
			        <byteValue>1</byteValue>
			        <shortValue>2</shortValue>
			        <integerValue>3</integerValue>
			        <longValue>4</longValue>
			        <floatValue>5.5</floatValue>
			        <DoubleDribble>6.5</DoubleDribble>
			    </SimpleScalarXmlPropertiesResult>
			</SimpleScalarXmlPropertiesResponse>
			`),
			ExpectResult: &SimpleScalarXmlPropertiesOutput{
				StringValue:       ptr.String("string"),
				EmptyStringValue:  ptr.String(""),
				TrueBooleanValue:  ptr.Bool(true),
				FalseBooleanValue: ptr.Bool(false),
				ByteValue:         ptr.Int8(1),
				ShortValue:        ptr.Int16(2),
				IntegerValue:      ptr.Int32(3),
				LongValue:         ptr.Int64(4),
				FloatValue:        ptr.Float32(5.5),
				DoubleValue:       ptr.Float64(6.5),
			},
		},
		// Supports handling NaN float values.
		"AwsQuerySupportsNaNFloatOutputs": {
			StatusCode: 200,
			Header: http.Header{
				"Content-Type": []string{"text/xml"},
			},
			BodyMediaType: "application/xml",
			Body: []byte(`<SimpleScalarXmlPropertiesResponse xmlns="https://example.com/">
			    <SimpleScalarXmlPropertiesResult>
			        <floatValue>NaN</floatValue>
			        <DoubleDribble>NaN</DoubleDribble>
			    </SimpleScalarXmlPropertiesResult>
			</SimpleScalarXmlPropertiesResponse>
			`),
			ExpectResult: &SimpleScalarXmlPropertiesOutput{
				FloatValue:  ptr.Float32(float32(math.NaN())),
				DoubleValue: ptr.Float64(math.NaN()),
			},
		},
		// Supports handling Infinity float values.
		"AwsQuerySupportsInfinityFloatOutputs": {
			StatusCode: 200,
			Header: http.Header{
				"Content-Type": []string{"text/xml"},
			},
			BodyMediaType: "application/xml",
			Body: []byte(`<SimpleScalarXmlPropertiesResponse xmlns="https://example.com/">
			    <SimpleScalarXmlPropertiesResult>
			        <floatValue>Infinity</floatValue>
			        <DoubleDribble>Infinity</DoubleDribble>
			    </SimpleScalarXmlPropertiesResult>
			</SimpleScalarXmlPropertiesResponse>
			`),
			ExpectResult: &SimpleScalarXmlPropertiesOutput{
				FloatValue:  ptr.Float32(float32(math.Inf(1))),
				DoubleValue: ptr.Float64(math.Inf(1)),
			},
		},
		// Supports handling -Infinity float values.
		"AwsQuerySupportsNegativeInfinityFloatOutputs": {
			StatusCode: 200,
			Header: http.Header{
				"Content-Type": []string{"text/xml"},
			},
			BodyMediaType: "application/xml",
			Body: []byte(`<SimpleScalarXmlPropertiesResponse xmlns="https://example.com/">
			    <SimpleScalarXmlPropertiesResult>
			        <floatValue>-Infinity</floatValue>
			        <DoubleDribble>-Infinity</DoubleDribble>
			    </SimpleScalarXmlPropertiesResult>
			</SimpleScalarXmlPropertiesResponse>
			`),
			ExpectResult: &SimpleScalarXmlPropertiesOutput{
				FloatValue:  ptr.Float32(float32(math.Inf(-1))),
				DoubleValue: ptr.Float64(math.Inf(-1)),
			},
		},
	}
	for name, c := range cases {
		t.Run(name, func(t *testing.T) {
			serverURL := "http://localhost:8888/"
			client := New(Options{
				HTTPClient: smithyhttp.ClientDoFunc(func(r *http.Request) (*http.Response, error) {
					headers := http.Header{}
					for k, vs := range c.Header {
						for _, v := range vs {
							headers.Add(k, v)
						}
					}
					if len(c.BodyMediaType) != 0 && len(headers.Values("Content-Type")) == 0 {
						headers.Set("Content-Type", c.BodyMediaType)
					}
					response := &http.Response{
						StatusCode: c.StatusCode,
						Header:     headers,
						Request:    r,
					}
					if len(c.Body) != 0 {
						response.ContentLength = int64(len(c.Body))
						response.Body = ioutil.NopCloser(bytes.NewReader(c.Body))
					} else {

						response.Body = http.NoBody
					}
					return response, nil
				}),
				APIOptions: []func(*middleware.Stack) error{
					func(s *middleware.Stack) error {
						s.Finalize.Clear()
						s.Initialize.Remove(`OperationInputValidation`)
						return nil
					},
				},
				EndpointResolver: EndpointResolverFunc(func(region string, options EndpointResolverOptions) (e oss.Endpoint, err error) {
					e.URL = serverURL
					e.SigningRegion = "us-west-2"
					return e, err
				}),
				IdempotencyTokenProvider: smithyrand.NewUUIDIdempotencyToken(&smithytesting.ByteLoop{}),
				Region:                   "us-west-2",
			})
			var params SimpleScalarXmlPropertiesInput
			result, err := client.SimpleScalarXmlProperties(context.Background(), &params)
			if err != nil {
				t.Fatalf("expect nil err, got %v", err)
			}
			if result == nil {
				t.Fatalf("expect not nil result")
			}
			opts := cmp.Options{
				cmpopts.IgnoreUnexported(
					middleware.Metadata{},
				),
				cmp.FilterValues(func(x, y float64) bool {
					return math.IsNaN(x) && math.IsNaN(y)
				}, cmp.Comparer(func(_, _ interface{}) bool { return true })),
				cmp.FilterValues(func(x, y float32) bool {
					return math.IsNaN(float64(x)) && math.IsNaN(float64(y))
				}, cmp.Comparer(func(_, _ interface{}) bool { return true })),
				cmpopts.IgnoreTypes(smithydocument.NoSerde{}),
			}
			if err := smithytesting.CompareValues(c.ExpectResult, result, opts...); err != nil {
				t.Errorf("expect c.ExpectResult value match:\n%v", err)
			}
		})
	}
}
