// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package networkfirewall // import "github.com/open-telemetry/opentelemetry-collector-contrib/extension/encoding/awslogsencodingextension/internal/unmarshaler/network-firewall-log"

import (
	"io"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/pdata/plog"

	"github.com/open-telemetry/opentelemetry-collector-contrib/extension/encoding"
	"github.com/open-telemetry/opentelemetry-collector-contrib/extension/encoding/awslogsencodingextension/internal/unmarshaler"
)

var _ unmarshaler.StreamingLogsUnmarshaler = (*NetworkFirewallLogUnmarshaler)(nil)

type NetworkFirewallLogUnmarshaler struct {
	buildInfo component.BuildInfo
}

func NewNetworkFirewallLogUnmarshaler(buildInfo component.BuildInfo) *NetworkFirewallLogUnmarshaler {
	_ = "STUB: not implemented"
	return nil
}

// See log fields: https://docs.aws.amazon.com/network-firewall/latest/developerguide/firewall-logging-contents.html.
type networkFirewallLog struct {
	FirewallName     string `json:"firewall_name"`
	AvailabilityZone string `json:"availability_zone"`
	EventTimestamp   string `json:"event_timestamp"`
	Event            struct {
		EventType string `json:"event_type"`
		FlowID    int64  `json:"flow_id"`
		Src       string `json:"src_ip"`
		SrcPort   int64  `json:"src_port"`
		Dest      string `json:"dest_ip"`
		DestPort  int64  `json:"dest_port"`
		Proto     string `json:"proto"`
		SNI       string `json:"sni"` // TLS SNI at event level
		Netflow   struct {
			Pkts   int64  `json:"pkts"`
			Bytes  int64  `json:"bytes"`
			Start  string `json:"start"`
			End    string `json:"end"`
			Age    int64  `json:"age"`
			MaxTTL int64  `json:"max_ttl"`
			MinTTL int64  `json:"min_ttl"`
			TxCnt  int64  `json:"tx_cnt"`
		} `json:"netflow"`
		Alert struct {
			Action      string `json:"action"`
			Signature   string `json:"signature"`
			SignatureID int64  `json:"signature_id"`
			Rev         int64  `json:"rev"`
			Category    string `json:"category"`
			Severity    int64  `json:"severity"`
			Gid         int64  `json:"gid"`
			Metadata    struct {
				AffectedProduct   []string `json:"affected_product"`
				AttackTarget      []string `json:"attack_target"`
				Deployment        []string `json:"deployment"`
				FormerCategory    []string `json:"former_category"`
				MalwareFamily     []string `json:"malware_family"`
				PerformanceImpact []string `json:"performance_impact"`
				SignatureSeverity []string `json:"signature_severity"`
				CreatedAt         []string `json:"created_at"`
				UpdatedAt         []string `json:"updated_at"`
			} `json:"metadata"`
		} `json:"alert"`
		RevocationCheck struct {
			LeafCertFpr string `json:"leaf_cert_fpr"`
			Action      string `json:"action"`
			Status      string `json:"status"`
		} `json:"revocation_check"`
		TLSError struct {
			ErrorMessage string `json:"error_message"`
		} `json:"tls_error"`
		TLS struct {
			Subject        string `json:"subject"`
			Issuer         string `json:"issuer"`
			SessionResumed *bool  `json:"session_resumed"`
		} `json:"tls"`
		HTTP struct {
			Hostname        string `json:"hostname"`
			URL             string `json:"url"`
			HTTPUserAgent   string `json:"http_user_agent"`
			HTTPContentType string `json:"http_content_type"`
			Cookie          string `json:"cookie"`
		} `json:"http"`
	} `json:"event"`
}

func (n *NetworkFirewallLogUnmarshaler) UnmarshalAWSLogs(reader io.Reader) (plog.Logs, error) {
	_ = "STUB: not implemented"
	// Decode as a stream but flush all at once using flush options
	return *new(plog.Logs), nil
}

// we must check for EOF with direct comparison and avoid wrapped EOF that can come from stream itself
//nolint:errorlint

// EOF indicates no logs were found, return any logs that's available

// NewLogsDecoder returns a LogsDecoder that processes AWS Network Firewall logs from the provided reader.
// Supports offset-based streaming; offset tracks bytes processed.
func (n *NetworkFirewallLogUnmarshaler) NewLogsDecoder(reader io.Reader, options ...encoding.DecoderOption) (encoding.LogsDecoder, error) {
	_ = "STUB: not implemented"
	return *new(encoding.LogsDecoder), nil
}

// This means there is no log to process, return EOF to indicate no logs.

func setResourceAttributes(resourceLogs plog.ResourceLogs, firewallName, availabilityZone string) error {
	_ = "STUB: not implemented"
	return nil
}

func (*NetworkFirewallLogUnmarshaler) addNetworkFirewallLog(log networkFirewallLog, record plog.LogRecord) error {
	_ = "STUB: not implemented"
	// Parse event timestamp
	return nil
}

// Add common event-level fields

// Add network fields

// Add netflow fields if present (Flow event type)

// Add alert fields if present (Alert event type)

// Add alert metadata if present

// SNI can be at event level

// Revocation check fields (check each field individually)

// TLS error (check for existence)

// TLS details

// TLS session resumed (only if field is present in JSON)

// Set HTTP fields if present
