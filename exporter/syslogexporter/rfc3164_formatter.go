// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package syslogexporter // import "github.com/open-telemetry/opentelemetry-collector-contrib/exporter/syslogexporter"

import (
	"go.opentelemetry.io/collector/pdata/plog"
)

type rfc3164Formatter struct{}

func newRFC3164Formatter() *rfc3164Formatter { _ = "STUB: not implemented"; return nil }

func (f *rfc3164Formatter) format(logRecord plog.LogRecord) string {
	_ = "STUB: not implemented"
	return ""
}

func (*rfc3164Formatter) formatPriority(logRecord plog.LogRecord) string {
	_ = "STUB: not implemented"
	return ""
}

func (*rfc3164Formatter) formatTimestamp(logRecord plog.LogRecord) string {
	_ = "STUB: not implemented"
	return ""
}

func (*rfc3164Formatter) formatHostname(logRecord plog.LogRecord) string {
	_ = "STUB: not implemented"
	return ""
}

func (*rfc3164Formatter) formatAppname(logRecord plog.LogRecord) string {
	_ = "STUB: not implemented"
	return ""
}

func (*rfc3164Formatter) formatMessage(logRecord plog.LogRecord) string {
	_ = "STUB: not implemented"
	return ""
}
