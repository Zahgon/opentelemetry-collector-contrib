// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package parser // import "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/statsdreceiver/internal/parser"

import (
	"errors"
	"net"
	"regexp"
	"time"

	"github.com/lightstep/go-expohisto/structure"
	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/pdata/pcommon"
	"go.opentelemetry.io/collector/pdata/pmetric"
	"go.opentelemetry.io/otel/attribute"

	"github.com/open-telemetry/opentelemetry-collector-contrib/receiver/statsdreceiver/protocol"
)

var (
	errEmptyMetricName  = errors.New("empty metric name")
	errEmptyMetricValue = errors.New("empty metric value")
)

type MetricType string // From the statsd line e.g., "c", "g", "h"

const (
	tagMetricType = "metric_type"

	CounterType      MetricType = "c"
	GaugeType        MetricType = "g"
	HistogramType    MetricType = "h"
	TimingType       MetricType = "ms"
	DistributionType MetricType = "d"

	receiverName = "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/statsdreceiver"
)

type ObserverCategory struct {
	method                protocol.ObserverType
	histogramConfig       structure.Config
	explicitBucketConfigs []explicitBucketConfig
	summaryPercentiles    []float64
}

var defaultObserverCategory = ObserverCategory{
	method: protocol.DefaultObserverType,
}

// StatsDParser supports the Parse method for parsing StatsD messages with Tags.
type StatsDParser struct {
	instrumentsByAddress    map[netAddr]*instruments
	enableMetricType        bool
	enableSimpleTags        bool
	isMonotonicCounter      bool
	enableIPOnlyAggregation bool
	ignoreHost              bool
	counterType             protocol.CounterType
	timerEvents             ObserverCategory
	histogramEvents         ObserverCategory
	lastIntervalTime        time.Time
	BuildInfo               component.BuildInfo
}

type instruments struct {
	addr                   net.Addr
	gauges                 map[statsDMetricDescription]pmetric.ScopeMetrics
	counters               map[statsDMetricDescription]pmetric.ScopeMetrics
	summaries              map[statsDMetricDescription]summaryMetric
	histograms             map[statsDMetricDescription]histogramMetric
	timersAndDistributions []pmetric.ScopeMetrics
}

func newInstruments(addr net.Addr) *instruments { _ = "STUB: not implemented"; return nil }

type sampleValue struct {
	value float64
	count float64
}

type summaryMetric struct {
	points      []float64
	weights     []float64
	percentiles []float64
}

type histogramStructure = structure.Histogram[float64]

type explicitBucketConfig struct {
	re      *regexp.Regexp
	buckets []float64
}

type explicitBucket struct {
	_         struct{}
	bucketMap map[float64]int
	buckets   []float64
	count     uint64
	infCount  uint64
	sum       float64
	min       float64
	max       float64
}

// Init retrieves ascendingly sorted unique buckets
func (e *explicitBucket) Init(buckets []float64) { _ = "STUB: not implemented"; return }

func (e *explicitBucket) UpdateByIncr(value float64, count uint64) {
	_ = "STUB: not implemented"
	return
}

type histogramMetric struct {
	agg            *histogramStructure
	explicitBucket *explicitBucket
}

type statsDMetric struct {
	description statsDMetricDescription
	asFloat     float64
	addition    bool
	unit        string
	sampleRate  float64
	timestamp   uint64
}

type statsDMetricDescription struct {
	name       string
	metricType MetricType
	attrs      attribute.Set
}

func (t MetricType) FullName() protocol.TypeName {
	_ = "STUB: not implemented"
	return *new(protocol.TypeName)
}

func (p *StatsDParser) resetState(when time.Time) { _ = "STUB: not implemented"; return }

func (p *StatsDParser) Initialize(enableMetricType, enableSimpleTags, isMonotonicCounter, enableIPOnlyAggregation, ignoreHost bool, sendTimerHistogram []protocol.TimerHistogramMapping, counterType protocol.CounterType) error {
	_ = "STUB: not implemented"
	return nil
}

// Note: validation occurs in ("../".Config).validate()

func explicitBucketInitializeRegex(opts protocol.HistogramConfig) []explicitBucketConfig {
	_ = "STUB: not implemented"
	return nil
}

func expoHistogramConfig(opts protocol.HistogramConfig) structure.Config {
	_ = "STUB: not implemented"
	return *new(structure.Config)
}

// GetMetrics gets the metrics preparing for flushing and reset the state.
func (p *StatsDParser) GetMetrics() []BatchMetrics { _ = "STUB: not implemented"; return nil }

func (p *StatsDParser) copyMetricAndScope(rm pmetric.ResourceMetrics, metric pmetric.ScopeMetrics) {
	_ = "STUB: not implemented"
	return
}

func (p *StatsDParser) setVersionAndNameScope(ilm pcommon.InstrumentationScope) {
	_ = "STUB: not implemented"
	return
}

var timeNowFunc = time.Now

func (p *StatsDParser) observerCategoryFor(t MetricType) ObserverCategory {
	_ = "STUB: not implemented"
	return *new(ObserverCategory)
}

// Aggregate for each metric line.
func (p *StatsDParser) Aggregate(line string, addr net.Addr) error {
	_ = "STUB: not implemented"
	return nil
}

// Discard NaN and infinite values for all metric types

// Note! Rounding float64 to uint64 here.

// No action.

func parseMessageToMetric(line string, enableMetricType, enableSimpleTags bool) (statsDMetric, error) {
	_ = "STUB: not implemented"
	return *new(statsDMetric), nil
}

// handle an empty tag set
// where the tags part was still sent (some clients do this)

// support both simple tags (w/o value) and dimension tags (w/ value).
// dogstatsd notably allows simple tags.

// As per DogStatD protocol v1.2:
// https://docs.datadoghq.com/developers/dogstatsd/datagram_shell/?tab=metrics#dogstatsd-protocol-v12

// As per DogStatD protocol v1.3:
// https://docs.datadoghq.com/developers/dogstatsd/datagram_shell/?tab=metrics#dogstatsd-protocol-v13

// Convert seconds to nanoseconds

// add metric_type dimension for all metrics

type netAddr struct {
	Network string
	String  string
}

func newNetAddr(addr net.Addr) netAddr { _ = "STUB: not implemented"; return *new(netAddr) }

func newIPOnlyNetAddr(addr net.Addr) netAddr { _ = "STUB: not implemented"; return *new(netAddr) }

// if there is an error, use the original address
