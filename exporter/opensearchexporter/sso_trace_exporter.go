// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package opensearchexporter // import "github.com/open-telemetry/opentelemetry-collector-contrib/exporter/opensearchexporter"

import (
	"context"
	"net/http"

	"github.com/opensearch-project/opensearch-go/v4/opensearchapi"
	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/config/confighttp"
	"go.opentelemetry.io/collector/exporter"
	"go.opentelemetry.io/collector/pdata/ptrace"
	"go.uber.org/zap"
)

type ssoTracesExporter struct {
	client        *opensearchapi.Client
	Namespace     string
	Dataset       string
	bulkAction    string
	model         mappingModel
	httpSettings  confighttp.ClientConfig
	telemetry     component.TelemetrySettings
	config        *Config
	indexResolver *indexResolver
}

func newSSOTracesExporter(cfg *Config, set exporter.Settings) *ssoTracesExporter {
	_ = "STUB: not implemented"
	return nil
}

func (s *ssoTracesExporter) Start(ctx context.Context, host component.Host) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *ssoTracesExporter) pushTraceData(ctx context.Context, td ptrace.Traces) error {
	_ = "STUB: not implemented"
	return nil
}

// Use timestamp for index resolution

func newOpenSearchClient(endpoint string, httpClient *http.Client, logger *zap.Logger) (*opensearchapi.Client, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// configure connection setup

// configure internal metrics reporting and logging
// TODO
// TODO
