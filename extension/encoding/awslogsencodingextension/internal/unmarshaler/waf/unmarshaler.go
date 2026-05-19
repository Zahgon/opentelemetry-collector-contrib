// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package waf // import "github.com/open-telemetry/opentelemetry-collector-contrib/extension/encoding/awslogsencodingextension/internal/unmarshaler/waf"

import (
	"io"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/pdata/plog"

	"github.com/open-telemetry/opentelemetry-collector-contrib/extension/encoding"
	"github.com/open-telemetry/opentelemetry-collector-contrib/extension/encoding/awslogsencodingextension/internal/unmarshaler"
)

// See log fields: https://docs.aws.amazon.com/waf/latest/developerguide/logging-fields.html.
type wafLog struct {
	Timestamp           int64  `json:"timestamp"`
	WebACLID            string `json:"webaclId"`
	TerminatingRuleID   string `json:"terminatingRuleId"`
	TerminatingRuleType string `json:"terminatingRuleType"`
	Action              string `json:"action"`
	HTTPSourceName      string `json:"httpSourceName"`
	HTTPSourceID        string `json:"httpSourceId"`
	HTTPRequest         struct {
		ClientIP string `json:"clientIp"`
		Country  string `json:"country"`
		Headers  []struct {
			Name  string `json:"name"`
			Value string `json:"value"`
		} `json:"headers"`
		URI         string `json:"uri"`
		Args        string `json:"args"`
		HTTPVersion string `json:"httpVersion"`
		HTTPMethod  string `json:"httpMethod"`
		RequestID   string `json:"requestID"`
		Fragment    string `json:"fragment"`
		Scheme      string `json:"scheme"`
		Host        string `json:"host"`
	} `json:"httpRequest"`
	ResponseCodeSent *int64 `json:"responseCodeSent"`
	Ja3Fingerprint   string `json:"ja3Fingerprint"`
	Ja4Fingerprint   string `json:"ja4Fingerprint"`
}

var _ unmarshaler.StreamingLogsUnmarshaler = (*WafLogUnmarshaler)(nil)

type WafLogUnmarshaler struct {
	buildInfo component.BuildInfo
}

func NewWAFLogUnmarshaler(buildInfo component.BuildInfo) *WafLogUnmarshaler {
	_ = "STUB: not implemented"
	return nil
}

func (w *WafLogUnmarshaler) UnmarshalAWSLogs(reader io.Reader) (plog.Logs, error) {
	_ = "STUB: not implemented"
	// Decode as a stream but flush all at once using flush options
	return *new(plog.Logs), nil
}

//nolint:errorlint

// EOF indicates no logs were found, return any logs that's available

// NewLogsDecoder returns a LogsDecoder that processes AWS WAF logs from the provided reader.
// Parses JSON-formatted logs containing WAF events (web ACL evaluations, actions, HTTP request details).
// Supports offset-based streaming; offset tracks bytes processed
func (w *WafLogUnmarshaler) NewLogsDecoder(reader io.Reader, options ...encoding.DecoderOption) (encoding.LogsDecoder, error) {
	_ = "STUB: not implemented"
	return *new(encoding.LogsDecoder), nil
}

func (*WafLogUnmarshaler) addWAFLog(log wafLog, record plog.LogRecord) error {
	_ = "STUB: not implemented"
	// timestamp is in milliseconds, so we need to convert it to ns first
	return nil
}

// setResourceAttributes based on the web ACL ID
func setResourceAttributes(resourceLogs plog.ResourceLogs, webACLID string) error {
	_ = "STUB: not implemented"
	return nil
}
