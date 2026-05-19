// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package honeycombmarkerexporter // import "github.com/open-telemetry/opentelemetry-collector-contrib/exporter/honeycombmarkerexporter"

import (
	"context"
	"net/http"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/config/confighttp"
	"go.opentelemetry.io/collector/config/configopaque"
	"go.opentelemetry.io/collector/exporter"
	"go.opentelemetry.io/collector/pdata/plog"

	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/ottl"
	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/ottl/contexts/ottllog"
)

const (
	defaultDatasetSlug = "__all__"
	userAgentHeaderKey = "User-Agent"
	contentType        = "Content-Type"
	honeycombTeam      = "X-Honeycomb-Team"
)

type marker struct {
	Marker
	logBoolExpr *ottl.ConditionSequence[*ottllog.TransformContext]
}

type honeycombLogsExporter struct {
	set                component.TelemetrySettings
	client             *http.Client
	httpClientSettings confighttp.ClientConfig
	apiURL             string
	apiKey             configopaque.String
	markers            []marker
	userAgentHeader    string
}

func newHoneycombLogsExporter(set exporter.Settings, config *Config) (*honeycombLogsExporter, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (e *honeycombLogsExporter) exportMarkers(ctx context.Context, ld plog.Logs) error {
	_ = "STUB: not implemented"
	return nil
}

func (e *honeycombLogsExporter) sendMarker(ctx context.Context, m marker, logRecord plog.LogRecord) error {
	_ = "STUB: not implemented"
	return nil
}

func (e *honeycombLogsExporter) start(ctx context.Context, host component.Host) (err error) {
	_ = "STUB: not implemented"
	return nil
}
