package defaults

import (
	"strconv"
	"testing"
	"time"

	"oss-sdk-go/oss"

	"github.com/google/go-cmp/cmp"
)

func TestConfigV1(t *testing.T) {
	cases := []struct {
		Mode     oss.DefaultsMode
		Expected Configuration
	}{
		{
			Mode: oss.DefaultsModeStandard,
			Expected: Configuration{
				ConnectTimeout:        oss.Duration(2000 * time.Millisecond),
				TLSNegotiationTimeout: oss.Duration(2000 * time.Millisecond),
				RetryMode:             oss.RetryModeStandard,
			},
		},
		{
			Mode: oss.DefaultsModeInRegion,
			Expected: Configuration{
				ConnectTimeout:        oss.Duration(1000 * time.Millisecond),
				TLSNegotiationTimeout: oss.Duration(1000 * time.Millisecond),
				RetryMode:             oss.RetryModeStandard,
			},
		},
		{
			Mode: oss.DefaultsModeCrossRegion,
			Expected: Configuration{
				ConnectTimeout:        oss.Duration(2800 * time.Millisecond),
				TLSNegotiationTimeout: oss.Duration(2800 * time.Millisecond),
				RetryMode:             oss.RetryModeStandard,
			},
		},
		{
			Mode: oss.DefaultsModeMobile,
			Expected: Configuration{
				ConnectTimeout:        oss.Duration(10000 * time.Millisecond),
				TLSNegotiationTimeout: oss.Duration(11000 * time.Millisecond),
				RetryMode:             oss.RetryModeAdaptive,
			},
		},
	}

	for i, tt := range cases {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			got, err := v1TestResolver(tt.Mode)
			if err != nil {
				t.Fatalf("expect no error, got %v", err)
			}
			if diff := cmp.Diff(tt.Expected, got); len(diff) > 0 {
				t.Error(diff)
			}
		})
	}
}
