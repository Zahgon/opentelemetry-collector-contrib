// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package expvarreceiver // import "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/expvarreceiver"

import (
	"context"
	"io"
	"net/http"
	"runtime"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/pdata/pmetric"
	"go.opentelemetry.io/collector/receiver"

	"github.com/open-telemetry/opentelemetry-collector-contrib/receiver/expvarreceiver/internal/metadata"
)

type expVar struct {
	// Use the existing runtime struct for decoding the JSON.
	MemStats *runtime.MemStats `json:"memstats"`
}

type expVarScraper struct {
	cfg    *Config
	set    *receiver.Settings
	client *http.Client
	mb     *metadata.MetricsBuilder
}

func newExpVarScraper(cfg *Config, set receiver.Settings) *expVarScraper {
	_ = "STUB: not implemented"
	return nil
}

func (e *expVarScraper) start(ctx context.Context, host component.Host) error {
	_ = "STUB: not implemented"
	return nil
}

func (e *expVarScraper) scrape(ctx context.Context) (pmetric.Metrics, error) {
	_ = "STUB: not implemented"
	return *new(pmetric.Metrics), nil
}

// Memstats exposes a circular buffer of recent GC stop-the-world pause times.
// The most recent pause is at PauseNs[(NumGC+255)%256].

func decodeResponseBody(body io.ReadCloser) (*expVar, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
