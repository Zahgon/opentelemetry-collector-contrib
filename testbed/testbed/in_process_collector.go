// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package testbed // import "github.com/open-telemetry/opentelemetry-collector-contrib/testbed/testbed"

import (
	"sync"
	"testing"

	"github.com/shirou/gopsutil/v4/process"
	"go.opentelemetry.io/collector/otelcol"
)

// inProcessCollector implements the OtelcolRunner interfaces running a single otelcol as a go routine within the
// same process as the test executor.
type inProcessCollector struct {
	factories  otelcol.Factories
	configStr  string
	svc        *otelcol.Collector
	stopped    bool
	configFile string
	wg         sync.WaitGroup
	t          *testing.T
}

// NewInProcessCollector creates a new inProcessCollector using the supplied component factories.
func NewInProcessCollector(factories otelcol.Factories) OtelcolRunner {
	_ = "STUB: not implemented"
	return *new(OtelcolRunner)
}

func (ipp *inProcessCollector) PrepareConfig(t *testing.T, configStr string) (configCleanup func(), err error) {
	_ = "STUB: not implemented"
	return nil,

		// NoOp
		nil
}

func (ipp *inProcessCollector) Start(StartParams) error { _ = "STUB: not implemented"; return nil }

// TODO: Pass this to the error handler.

func (ipp *inProcessCollector) Stop() (stopped bool, err error) {
	_ = "STUB: not implemented"
	return false, nil
}

// Do not delete temporary files on Windows because it fails too much on scoped tests.
// See https://github.com/open-telemetry/opentelemetry-collector-contrib/issues/42639

func (*inProcessCollector) WatchResourceConsumption() error { _ = "STUB: not implemented"; return nil }

func (*inProcessCollector) GetProcessMon() *process.Process { _ = "STUB: not implemented"; return nil }

func (*inProcessCollector) GetTotalConsumption() *ResourceConsumption {
	_ = "STUB: not implemented"
	return nil
}

func (*inProcessCollector) GetResourceConsumption() string { _ = "STUB: not implemented"; return "" }
