// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package datasenders // import "github.com/open-telemetry/opentelemetry-collector-contrib/testbed/datasenders"

import (
	"net/http"

	"github.com/prometheus/client_golang/prometheus"

	"github.com/open-telemetry/opentelemetry-collector-contrib/testbed/testbed"
)

// PrometheusStaticPayloadConfig defines how to populate a prometheus.Registry for benchmarking.
type PrometheusStaticPayloadConfig struct {
	SeriesCount          int
	LabelsPerSeries      int
	WithTargetInfo       bool
	WithNativeHistograms bool
}

// prometheusStaticSender serves a pre-populated Prometheus registry via promhttp.
// The collector's prometheusreceiver scrapes this endpoint.
type prometheusStaticSender struct {
	testbed.DataSenderBase
	server   *http.Server
	registry *prometheus.Registry
}

// NewPrometheusStaticSender builds a registry from the given config and returns
// a DataSender that serves it on the given host:port.
func NewPrometheusStaticSender(host string, port int, cfg PrometheusStaticPayloadConfig) testbed.DataSender {
	_ = "STUB: not implemented"
	return *new(testbed.DataSender)
}

func (s *prometheusStaticSender) Start() error { _ = "STUB: not implemented"; return nil }

func (s *prometheusStaticSender) Stop() error { _ = "STUB: not implemented"; return nil }

func (s *prometheusStaticSender) GenConfigYAMLStr() string { _ = "STUB: not implemented"; return "" }

func (*prometheusStaticSender) ProtocolName() string { _ = "STUB: not implemented"; return "" }

func buildRegistry(cfg PrometheusStaticPayloadConfig) *prometheus.Registry {
	_ = "STUB: not implemented"
	return nil
}
