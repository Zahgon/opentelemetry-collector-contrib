// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

//go:build windows

package loadscraper // import "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/hostmetricsreceiver/internal/scraper/loadscraper"

import (
	"context"
	"math"
	"sync"
	"time"

	"github.com/shirou/gopsutil/v4/load"
	"go.uber.org/zap"

	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/winperfcounters"
)

// Sample processor queue length at a 5s frequency, and calculate exponentially weighted moving averages
// as per https://en.wikipedia.org/wiki/Load_(computing)#Unix-style_load_calculation

const (
	system               = "System"
	processorQueueLength = "Processor Queue Length"
)

var (
	samplingFrequency = 5 * time.Second

	loadAvgFactor1m  = 1 / math.Exp(samplingFrequency.Seconds()/time.Minute.Seconds())
	loadAvgFactor5m  = 1 / math.Exp(samplingFrequency.Seconds()/(5*time.Minute).Seconds())
	loadAvgFactor15m = 1 / math.Exp(samplingFrequency.Seconds()/(15*time.Minute).Seconds())

	// perfCounterFactory is used to facilitate testing
	perfCounterFactory = winperfcounters.NewWatcher
)

var (
	scraperCount int
	startupLock  sync.Mutex

	samplerInstance *sampler
)

type sampler struct {
	done               chan struct{}
	logger             *zap.Logger
	perfCounterWatcher winperfcounters.PerfCounterWatcher
	loadAvg1m          float64
	loadAvg5m          float64
	loadAvg15m         float64
	lock               sync.RWMutex
}

func setSamplingFrequency(freq time.Duration) { _ = "STUB: not implemented"; return }

func startSampling(_ context.Context, logger *zap.Logger) error {
	_ = "STUB: not implemented"
	return nil
}

// startSampling may be called multiple times if multiple scrapers are
// initialized - but we only want to initialize a single load sampler

// To keep the same behavior, as previous versions on Windows, error in this case is just logged
// and the scraper is not started.

func newSampler(logger *zap.Logger) (*sampler, error) { _ = "STUB: not implemented"; return nil, nil }

func (sw *sampler) startSamplingTicker() {
	_ = "STUB: not implemented"
	// Store the sampling frequency in a local variable to avoid race conditions during tests.
	return
}

func (sw *sampler) sampleLoad() { _ = "STUB: not implemented"; return }

func stopSampling(_ context.Context) error { _ = "STUB: not implemented"; return nil }

// no load scraper is running nothing to do

// only stop sampling if all load scrapers have been closed

// no more load scrapers are running, stop the sampler

func getSampledLoadAverages(_ context.Context) (*load.AvgStat, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
