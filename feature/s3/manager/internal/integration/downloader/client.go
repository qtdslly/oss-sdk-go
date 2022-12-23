//go:build integration && perftest
// +build integration,perftest

package downloader

import (
	"net"
	"net/http"
	"time"

	"oss-sdk-go/oss"
	awshttp "oss-sdk-go/oss/transport/http"
)

func NewHTTPClient(cfg ClientConfig) oss.HTTPClient {
	return awshttp.NewBuildableClient().WithTransportOptions(func(tr *http.Transport) {
		*tr = http.Transport{
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

			ReadBufferSize:  cfg.ReadBufferSize,
			WriteBufferSize: cfg.WriteBufferSize,
		}
	})
}
