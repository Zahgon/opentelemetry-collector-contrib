// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package integrationtest // import "github.com/open-telemetry/opentelemetry-collector-contrib/exporter/elasticsearchexporter/integrationtest"

import (
	"context"
	"sync"
	"testing"
	"time"

	"github.com/shirou/gopsutil/v4/process"
	"go.opentelemetry.io/collector/otelcol"
	"golang.org/x/sync/errgroup"

	"github.com/open-telemetry/opentelemetry-collector-contrib/testbed/testbed"
)

// createConfigYaml creates a yaml config for an otel collector for testing.
func createConfigYaml(
	tb testing.TB,
	sender testbed.DataSender,
	receiver testbed.DataReceiver,
	processors map[string]string,
	extensions map[string]string,
	pipelineType string,
	debug bool,
) string {
	_ = "STUB: not implemented"
	return ""
}

func createConfigSection(m map[string]string) (sections, list string) {
	_ = "STUB: not implemented"
	return "", ""
}

// recreatableOtelCol creates an otel collector that can be used to simulate
// a crash of the collector. It implements the testbed.OtelColRunner interface.
type recreatableOtelCol struct {
	tempDir   string
	factories otelcol.Factories
	settings  otelcol.CollectorSettings
	configStr string
	errGrp    errgroup.Group
	cancel    context.CancelFunc

	mu  sync.Mutex
	col *otelcol.Collector
}

func newRecreatableOtelCol(tb testing.TB) *recreatableOtelCol {
	_ = "STUB: not implemented"
	return nil
}

func (c *recreatableOtelCol) PrepareConfig(_ *testing.T, configStr string) (func(), error) {
	_ = "STUB: not implemented"
	return nil,

		// NoOp
		nil
}

func (c *recreatableOtelCol) Start(_ testbed.StartParams) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *recreatableOtelCol) Stop() (bool, error) { _ = "STUB: not implemented"; return false, nil }

func (c *recreatableOtelCol) Restart(graceful bool, shutdownFor time.Duration) error {
	_ = "STUB: not implemented"
	return nil
}

func (*recreatableOtelCol) WatchResourceConsumption() error { _ = "STUB: not implemented"; return nil }

func (*recreatableOtelCol) GetProcessMon() *process.Process { _ = "STUB: not implemented"; return nil }

func (*recreatableOtelCol) GetTotalConsumption() *testbed.ResourceConsumption {
	_ = "STUB: not implemented"
	return nil
}

func (*recreatableOtelCol) GetResourceConsumption() string { _ = "STUB: not implemented"; return "" }

func (c *recreatableOtelCol) run() error { _ = "STUB: not implemented"; return nil }

// Ignore context canceled errors
