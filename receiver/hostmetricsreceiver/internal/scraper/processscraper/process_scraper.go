// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package processscraper // import "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/hostmetricsreceiver/internal/scraper/processscraper"

import (
	"context"
	"time"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/pdata/pcommon"
	"go.opentelemetry.io/collector/pdata/pmetric"
	"go.opentelemetry.io/collector/scraper"

	"github.com/open-telemetry/opentelemetry-collector-contrib/internal/filter/filterset"
	"github.com/open-telemetry/opentelemetry-collector-contrib/receiver/hostmetricsreceiver/internal/scraper/processscraper/internal/metadata"
	"github.com/open-telemetry/opentelemetry-collector-contrib/receiver/hostmetricsreceiver/internal/scraper/processscraper/ucal"
)

const (
	cpuMetricsLen               = 1
	memoryMetricsLen            = 2
	memoryUtilizationMetricsLen = 1
	diskMetricsLen              = 1
	pagingMetricsLen            = 1
	threadMetricsLen            = 1
	contextSwitchMetricsLen     = 1
	fileDescriptorMetricsLen    = 1
	handleMetricsLen            = 1
	signalMetricsLen            = 1
	uptimeMetricsLen            = 1

	metricsLen = cpuMetricsLen + memoryMetricsLen + diskMetricsLen + memoryUtilizationMetricsLen + pagingMetricsLen + threadMetricsLen + contextSwitchMetricsLen + fileDescriptorMetricsLen + signalMetricsLen + handleCountMetricsLen + uptimeMetricsLen
)

// scraper for Process Metrics
type processScraper struct {
	settings           scraper.Settings
	config             *Config
	mb                 *metadata.MetricsBuilder
	includeFS          filterset.FilterSet
	excludeFS          filterset.FilterSet
	scrapeProcessDelay time.Duration
	ucals              map[int32]*ucal.CPUUtilizationCalculator
	logicalCores       int

	// for mocking
	getProcessCreateTime func(p processHandle, ctx context.Context) (int64, error)
	getProcessHandles    func(context.Context) (processHandles, error)
}

// newProcessScraper creates a Process Scraper
func newProcessScraper(settings scraper.Settings, cfg *Config) (*processScraper, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *processScraper) start(context.Context, component.Host) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *processScraper) scrape(ctx context.Context) (pmetric.Metrics, error) {
	_ = "STUB: not implemented"
	return *new(pmetric.Metrics), nil
}

// If the boot time cache featuregate is disabled, this will refresh the
// cached boot time value for use in the current scrape. This functionally
// replicates the previous functionality in all but the most extreme
// cases of boot time changing in the middle of a scrape.

// Cleanup any [ucal.CPUUtilizationCalculator]s for PIDs that are no longer present

// getProcessMetadata returns a slice of processMetadata, including handles,
// for all currently running processes. If errors occur obtaining information
// for some processes, an error will be returned, but any processes that were
// successfully obtained will still be returned.
func (s *processScraper) getProcessMetadata(ctx context.Context) ([]*processMetadata, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// filter processes by name

// set the start time to now to avoid including this when a scrape_process_delay is set

func (s *processScraper) scrapeAndAppendCPUTimeMetric(ctx context.Context, now pcommon.Timestamp, handle processHandle, pid int32) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *processScraper) scrapeAndAppendMemoryUsageMetrics(ctx context.Context, now pcommon.Timestamp, handle processHandle) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *processScraper) scrapeAndAppendMemoryUtilizationMetric(ctx context.Context, now pcommon.Timestamp, handle processHandle) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *processScraper) scrapeAndAppendDiskMetrics(ctx context.Context, now pcommon.Timestamp, handle processHandle) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *processScraper) scrapeAndAppendPagingMetric(ctx context.Context, now pcommon.Timestamp, handle processHandle) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *processScraper) scrapeAndAppendThreadsMetrics(ctx context.Context, now pcommon.Timestamp, handle processHandle) error {
	_ = "STUB: not implemented"
	return nil
}

// NOTE: For now, this metric is only supported on Linux because
// that's the only platform gopsutil will collect this value for.
// The implementation makes the assumption that this remains the
// case. If the metric will ever be supported on other platforms,
// the implementation may end up needing to change or be separated
// into individual platform implementations.
func (s *processScraper) scrapeAndAppendContextSwitchMetrics(ctx context.Context, now pcommon.Timestamp, handle processHandle) error {
	_ = "STUB: not implemented"
	return nil
}

// If the task is single threaded, we can just read
// the context switch stat for the task we already have.

// The gopsutil method for enumerating processes on the system uses
// getdents(2). Calling this on procfs will only enumerate process IDs,
// and not the IDs of all their individual tasks. This is good and something
// we actually want. However, the NumCtxSwitchesWithThread method from gopsutil
// will read the context switches from /proc/[pid]/stat, which only has the context
// switches for the given ID, not a sum of the all the subtasks of the process.
//
// Thus if the task is multithreaded, we need to enumerate the tasks
// in the thread and treat them as separate gopsutil process.Process objects.
// This works because while the overall enumeration is done with getdents(2),
// calling NewProcess with each task ID directly reads /proc/[tid], which will
// allow us to separately read the context switches for each task of a process.
// The sum will be reported as the context switches for the "entire process".

func (s *processScraper) scrapeAndAppendOpenFileDescriptorsMetric(ctx context.Context, now pcommon.Timestamp, handle processHandle) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *processScraper) scrapeAndAppendHandlesMetric(ctx context.Context, now pcommon.Timestamp, handle processHandle) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *processScraper) scrapeAndAppendSignalsPendingMetric(ctx context.Context, now pcommon.Timestamp, handle processHandle) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *processScraper) scrapeAndAppendUptimeMetric(ctx context.Context, now pcommon.Timestamp, handle processHandle) error {
	_ = "STUB: not implemented"
	return nil
}

// Since create time is in milliseconds, it needs to be multiplied
// by the constant value so that it can be used as part of the time.Unix function.
