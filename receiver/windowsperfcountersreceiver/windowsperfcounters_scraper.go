// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

//go:build windows

package windowsperfcountersreceiver // import "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/windowsperfcountersreceiver"

import (
	"context"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/pdata/pcommon"
	"go.opentelemetry.io/collector/pdata/pmetric"

	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/winperfcounters"
)

const instanceLabelName = "instance"

type perfCounterMetricWatcher struct {
	winperfcounters.PerfCounterWatcher
	MetricRep
	recreate bool
}

type newWatcherFunc func(string, string, string) (winperfcounters.PerfCounterWatcher, error)

// windowsPerfCountersScraper is the type that scrapes various host metrics.
type windowsPerfCountersScraper struct {
	cfg      *Config
	settings component.TelemetrySettings
	watchers []perfCounterMetricWatcher

	// for mocking
	newWatcher newWatcherFunc
}

func newScraper(cfg *Config, settings component.TelemetrySettings) *windowsPerfCountersScraper {
	_ = "STUB: not implemented"
	return nil
}

func (s *windowsPerfCountersScraper) start(context.Context, component.Host) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *windowsPerfCountersScraper) initWatchers() ([]perfCounterMetricWatcher, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *windowsPerfCountersScraper) shutdown(context.Context) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *windowsPerfCountersScraper) scrape(context.Context) (pmetric.Metrics, error) {
	_ = "STUB: not implemented"
	return *new(pmetric.Metrics), nil
}

// Drop metrics with no datapoints. This happens when configured counters don't exist on the host.
// This may result in a Metrics message with no metrics if all counters are missing.

func initializeMetricDps(metric pmetric.Metric, now pcommon.Timestamp, counterValue winperfcounters.CounterValue,
	attributes map[string]string,
) {
	_ = "STUB: not implemented"
	return
}

func instancesFromConfig(oc ObjectConfig) []string { _ = "STUB: not implemented"; return nil }
