// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package googlecloudlogentryencodingextension // import "github.com/open-telemetry/opentelemetry-collector-contrib/extension/encoding/googlecloudlogentryencodingextension"

import (
	"context"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/pdata/plog"

	"github.com/open-telemetry/opentelemetry-collector-contrib/extension/encoding"
)

var _ encoding.LogsUnmarshalerExtension = (*ext)(nil)

type ext struct {
	config Config
}

func newExtension(cfg *Config) *ext { _ = "STUB: not implemented"; return nil }

func (*ext) Start(_ context.Context, _ component.Host) error { _ = "STUB: not implemented"; return nil }

func (*ext) Shutdown(context.Context) error { _ = "STUB: not implemented"; return nil }

func (ex *ext) UnmarshalLogs(buf []byte) (plog.Logs, error) {
	_ = "STUB: not implemented"
	return *

	// each line corresponds to a log
	new(plog.Logs), nil
}

func (ex *ext) handleLogLine(logs plog.Logs, logLine []byte) error {
	_ = "STUB: not implemented"
	return nil
}
