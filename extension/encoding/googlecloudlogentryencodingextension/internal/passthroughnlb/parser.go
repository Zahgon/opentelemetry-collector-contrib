// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

// Package passthroughnlb contains utilities for parsing Google Cloud Passthrough External and Internal Network Load Balancer logs.
package passthroughnlb // import "github.com/open-telemetry/opentelemetry-collector-contrib/extension/encoding/googlecloudlogentryencodingextension/internal/passthroughnlb"

import (
	"errors"
	"time"

	"go.opentelemetry.io/collector/pdata/pcommon"
)

const (
	// ConnectionsLogNameSuffix identifies load balancer connection logs in the logName field.
	ConnectionsLogNameSuffix = "loadbalancing.googleapis.com%2Fflows"

	externalLoadBalancerLogType = "type.googleapis.com/google.cloud.loadbalancing.type.ExternalNetworkLoadBalancerLogEntry"
	internalLoadBalancerLogType = "type.googleapis.com/google.cloud.loadbalancing.type.InternalNetworkLoadBalancerLogEntry"

	gcpPassthroughNLBPacketsStartTime = "gcp.load_balancing.passthrough_nlb.packets.start_time" // #nosec G101
	gcpPassthroughNLBPacketsEndTime   = "gcp.load_balancing.passthrough_nlb.packets.end_time"   // #nosec G101
	gcpPassthroughNLBBytesReceived    = "gcp.load_balancing.passthrough_nlb.bytes_received"     // #nosec G101
	gcpPassthroughNLBBytesSent        = "gcp.load_balancing.passthrough_nlb.bytes_sent"         // #nosec G101
	gcpPassthroughNLBPacketsReceived  = "gcp.load_balancing.passthrough_nlb.packets_received"   // #nosec G101
	gcpPassthroughNLBPacketsSent      = "gcp.load_balancing.passthrough_nlb.packets_sent"       // #nosec G101
	gcpPassthroughNLBRTT              = "gcp.load_balancing.passthrough_nlb.rtt"                // #nosec G101
)

var (
	errUnmarshalPayload  = errors.New("failed to unmarshal Passthrough NLB log payload")
	errUnexpectedLogType = errors.New("unexpected log type")
	errBytesReceived     = errors.New("failed to add bytes received")
	errBytesSent         = errors.New("failed to add bytes sent")
	errPacketsReceived   = errors.New("failed to packets received")
	errPacketsSent       = errors.New("failed to add packets sent")
	errRTT               = errors.New("failed to add RTT")
	errServerAddress     = errors.New("failed to set server address")
)

type loadBalancerLog struct {
	Type       string      `json:"@type"`
	Connection *connection `json:"connection"`
	StartTime  *time.Time  `json:"startTime"`
	EndTime    *time.Time  `json:"endTime"`

	// BytesReceived, BytesSent, PacketsReceived, PacketsSent are string-encoded 64-bit integers.
	// Although the official documentation (https://docs.cloud.google.com/load-balancing/docs/network/networklb-monitoring)
	// specifies these fields as integers, the actual
	// JSON payload from Cloud Logging seems to stringify int64/uint64 values to
	// preserve precision
	BytesReceived   string `json:"bytesReceived"`
	BytesSent       string `json:"bytesSent"`
	PacketsReceived string `json:"packetsReceived"`
	PacketsSent     string `json:"packetsSent"`
	RTT             string `json:"rtt"`
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

func handleConnection(conn *connection, attr pcommon.Map) error {
	_ = "STUB: not implemented"
	return nil
}

func handleTimestamps(start, end *time.Time, attr pcommon.Map) { _ = "STUB: not implemented"; return }
