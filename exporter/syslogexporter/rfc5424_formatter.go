// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package syslogexporter // import "github.com/open-telemetry/opentelemetry-collector-contrib/exporter/syslogexporter"

import (
	"go.opentelemetry.io/collector/pdata/plog"
)

type rfc5424Formatter struct {
	octetCounting bool
}

func newRFC5424Formatter(octetCounting bool) *rfc5424Formatter {
	_ = "STUB: not implemented"
	return nil
}

func (f *rfc5424Formatter) format(logRecord plog.LogRecord) string {
	_ = "STUB: not implemented"
	return ""
}

func (*rfc5424Formatter) formatPriority(logRecord plog.LogRecord) string {
	_ = "STUB: not implemented"
	return ""
}

func (*rfc5424Formatter) formatVersion(logRecord plog.LogRecord) string {
	_ = "STUB: not implemented"
	return ""
}

func (*rfc5424Formatter) formatTimestamp(logRecord plog.LogRecord) string {
	_ = "STUB: not implemented"
	return ""
}

func (*rfc5424Formatter) formatHostname(logRecord plog.LogRecord) string {
	_ = "STUB: not implemented"
	return ""
}

func (*rfc5424Formatter) formatAppname(logRecord plog.LogRecord) string {
	_ = "STUB: not implemented"
	return ""
}

func (*rfc5424Formatter) formatPid(logRecord plog.LogRecord) string {
	_ = "STUB: not implemented"
	return ""
}

func (*rfc5424Formatter) formatMessageID(logRecord plog.LogRecord) string {
	_ = "STUB: not implemented"
	return ""
}

func (*rfc5424Formatter) formatStructuredData(logRecord plog.LogRecord) string {
	_ = "STUB: not implemented"
	return ""
}

func (*rfc5424Formatter) formatMessage(logRecord plog.LogRecord) string {
	_ = "STUB: not implemented"
	return ""
}
