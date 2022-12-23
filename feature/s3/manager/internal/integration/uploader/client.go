//go:build integration && perftest
// +build integration,perftest

package uploader

import (
	"net"
	"net/http"
	"time"

	"oss-sdk-go/oss"
	awshttp "oss-sdk-go/oss/transport/http"
)

func NewHTTPClient(cfg ClientConfig) oss.HTTPClient {
	return awshttp.NewBuildableClient().WithTransportOptions(func(transport *http.Transport) {
		*transport = http.Transport{
			Proxy: http.ProxyFromEnvironment,
			DialContext: (&net.Dialer{
				Timeout:   cfg.Timeouts.Connect,
				KeepAlive: 30 * time.Second,
			}).DialContext,
			MaxIdleConns:        cfg.MaxIdleConns,
			MaxIdleConnsPerHost: cfg.MaxIdleConnsPerHost,
			IdleConnTimeout:     90 * time.Second,

			DisableKeepAlives:     !cfg.KeepAlive,
			TLSHandshakeTimeout:   cfg.Timeouts.TLSHandshake,
			ExpectContinueTimeout: cfg.Timeouts.ExpectContinue,
			ResponseHeaderTimeout: cfg.Timeouts.ResponseHeader,
		}
	})
}
