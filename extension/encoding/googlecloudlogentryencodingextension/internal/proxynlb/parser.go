// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

// Package proxynlb contains utilities for parsing Google Cloud Proxy Network Load Balancer logs.
package proxynlb // import "github.com/open-telemetry/opentelemetry-collector-contrib/extension/encoding/googlecloudlogentryencodingextension/internal/proxynlb"

import (
	"errors"
	"time"

	"go.opentelemetry.io/collector/pdata/pcommon"
)

const (
	// ConnectionsLogNameSuffix identifies load balancer connection logs in the logName field.
	ConnectionsLogNameSuffix = "loadbalancing.googleapis.com%2Fconnections"

	loadBalancerLogType = "type.googleapis.com/google.cloud.loadbalancing.type.LoadBalancerLogEntry"

	gcpProxyNLBConnectionStartTime = "gcp.load_balancing.proxy_nlb.connection.start_time"
	gcpProxyNLBConnectionEndTime   = "gcp.load_balancing.proxy_nlb.connection.end_time"
	gcpProxyNLBServerBytesReceived = "gcp.load_balancing.proxy_nlb.server.bytes_received"
	gcpProxyNLBServerBytesSent     = "gcp.load_balancing.proxy_nlb.server.bytes_sent"
)

var (
	ErrUnmarshalPayload    = errors.New("failed to unmarshal Proxy NLB log payload")
	ErrUnexpectedLogType   = errors.New("unexpected log type")
	ErrServerBytesReceived = errors.New("failed to add server bytes received")
	ErrServerBytesSent     = errors.New("failed to add server bytes sent")
)

type loadBalancerLog struct {
	Type       string      `json:"@type"`
	Connection *connection `json:"connection"`
	StartTime  *time.Time  `json:"startTime"`
	EndTime    *time.Time  `json:"endTime"`

	// ServerBytesReceived and ServerBytesSent are string-encoded 64-bit integers.
	// Although the official documentation (https://docs.cloud.google.com/load-balancing/docs/tcp/tcp-ssl-proxy-logging-monitoring)
	// specifies these fields as integers, the actual
	// JSON payload from Cloud Logging seems to stringify int64/uint64 values to
	// preserve precision
	ServerBytesReceived string `json:"serverBytesReceived"`
	ServerBytesSent     string `json:"serverBytesSent"`
}

type connection struct {
	ClientIP   string `json:"clientIp"`
	ClientPort *int64 `json:"clientPort"`
	Protocol   *int64 `json:"protocol"`
	ServerIP   string `json:"serverIp"`
	ServerPort *int64 `json:"serverPort"`
}

// ParsePayloadIntoAttributes unmarshals the provided payload into the supplied attribute map.
func ParsePayloadIntoAttributes(payload []byte, attr pcommon.Map) error {
	_ = "STUB: not implemented"
	return nil
}

func handleConnection(conn *connection, attr pcommon.Map) { _ = "STUB: not implemented"; return }

func handleTimestamps(start, end *time.Time, attr pcommon.Map) { _ = "STUB: not implemented"; return }
