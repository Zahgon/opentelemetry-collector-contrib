// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package client // import "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/statsdreceiver/internal/transport/client"

import (
	"io"
	"net"
)

// StatsD defines the properties of a StatsD connection.
type StatsD struct {
	transport string
	address   string
	conn      io.Writer
}

// NewStatsD creates a new StatsD instance to support the need for testing
// the statsdreceiver package and is not intended/tested to be used in production.
func NewStatsD(transport, address string) (*StatsD, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// connect populates the StatsD.conn
func (s *StatsD) connect() error { _ = "STUB: not implemented"; return nil }

// Disconnect closes the StatsD.conn.
func (s *StatsD) Disconnect() error { _ = "STUB: not implemented"; return nil }

// SendMetric sends the input metric to the StatsD connection.
func (s *StatsD) SendMetric(metric Metric) error { _ = "STUB: not implemented"; return nil }

func (s *StatsD) ConnectionLocalAddress() (net.Addr, error) {
	_ = "STUB: not implemented"
	return *new(net.Addr), nil
}

// Metric contains the metric fields for a StatsD message.
type Metric struct {
	Name  string
	Value string
	Type  string
}

// String formats a Metric into a StatsD message.
func (m Metric) String() string { _ = "STUB: not implemented"; return "" }
