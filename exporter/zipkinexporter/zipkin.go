// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package zipkinexporter // import "github.com/open-telemetry/opentelemetry-collector-contrib/exporter/zipkinexporter"

import (
	"context"
	"net/http"

	zipkinreporter "github.com/openzipkin/zipkin-go/reporter"
	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/config/confighttp"
	"go.opentelemetry.io/collector/pdata/ptrace"

	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/translator/zipkin/zipkinv2"
)

var translator zipkinv2.FromTranslator

// zipkinExporter is a multiplexing exporter that spawns a new OpenCensus-Go Zipkin
// exporter per unique node encountered. This is because serviceNames per node define
// unique services, alongside their IPs. Also it is useful to receive traffic from
// Zipkin servers and then transform them back to the final form when creating an
// OpenCensus spandata.
type zipkinExporter struct {
	defaultServiceName string

	url            string
	client         *http.Client
	serializer     zipkinreporter.SpanSerializer
	clientSettings *confighttp.ClientConfig
	settings       component.TelemetrySettings
}

func createZipkinExporter(cfg *Config, settings component.TelemetrySettings) (*zipkinExporter, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// start creates the http client
func (ze *zipkinExporter) start(ctx context.Context, host component.Host) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func (ze *zipkinExporter) pushTraces(ctx context.Context, td ptrace.Traces) error {
	_ = "STUB: not implemented"
	return nil
}
