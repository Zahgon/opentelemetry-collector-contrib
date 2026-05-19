// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package client // import "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/carbonreceiver/internal/client"

import (
	"io"
	"time"
)

// Graphite is a struct that defines the relevant properties of a graphite
// connection.
// This code was initially taken from
// https://github.com/census-ecosystem/opencensus-go-exporter-graphite/tree/master/internal/client
// and modified for the needs of testing the Carbon receiver package and is not
// intended/tested to be used in production.
type Graphite struct {
	Endpoint string
	Timeout  time.Duration
	Conn     io.Writer
}

// Transport is used as an enum to select the type of transport to be used.
type Transport int

// Available transport options: TCP and UDP.
const (
	TCP Transport = iota
	UDP
)

const defaultTimeout = 5

// NewGraphite is a method that's used to create a new Graphite instance.
// This code was initially taken from
// https://github.com/census-ecosystem/opencensus-go-exporter-graphite/tree/master/internal/client
// and modified for the needs of testing the Carbon receiver package and is not
// intended/tested to be used in production.
func NewGraphite(transport Transport, endpoint string) (*Graphite, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// connect populates the Graphite.conn.
func (g *Graphite) connect(transport Transport) error { _ = "STUB: not implemented"; return nil }

// Disconnect closes the Graphite.conn field
func (g *Graphite) Disconnect() (err error) { _ = "STUB: not implemented"; return nil }

// SendMetric method can be used to just pass a metric name and value and
// have it be sent to the Graphite host
func (g *Graphite) SendMetric(metric Metric) error { _ = "STUB: not implemented"; return nil }

// SputterThenSendMetric method sends a bad partial metric, then the whole metric across.
func (g *Graphite) SputterThenSendMetric(metric Metric) error {
	_ = "STUB: not implemented"
	return nil
}

// SendMetrics method can be used to pass a set of metrics and
// have it be sent to the Graphite host
func (g *Graphite) SendMetrics(metrics []Metric) error { _ = "STUB: not implemented"; return nil }

// Metric contains the metric fields expected by Graphite.
type Metric struct {
	Name      string
	Value     float64
	Timestamp time.Time
}

// String formats a Metric to the format expected bt Graphite.
func (m Metric) String() string { _ = "STUB: not implemented"; return "" }
