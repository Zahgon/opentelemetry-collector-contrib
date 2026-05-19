// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package mockdatadogagentexporter // import "github.com/open-telemetry/opentelemetry-collector-contrib/testbed/mockdatasenders/mockdatadogagentexporter"

import (
	"context"
	"net/http"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/config/confighttp"
	"go.opentelemetry.io/collector/pdata/ptrace"
)

type ddExporter struct {
	endpoint       string
	client         *http.Client
	clientSettings *confighttp.ClientConfig
}

func createExporter(c *Config) *ddExporter { _ = "STUB: not implemented"; return nil }

// start creates the http client
func (dd *ddExporter) start(ctx context.Context, host component.Host) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func (dd *ddExporter) pushTraces(ctx context.Context, td ptrace.Traces) error {
	_ = "STUB: not implemented"
	return nil
}
