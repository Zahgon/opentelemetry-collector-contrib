// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package internal // import "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/prometheusreceiver/internal"

import (
	"context"

	"github.com/prometheus/prometheus/model/exemplar"
	"github.com/prometheus/prometheus/model/histogram"
	"github.com/prometheus/prometheus/model/labels"
	"github.com/prometheus/prometheus/model/metadata"
	"github.com/prometheus/prometheus/scrape"
	"github.com/prometheus/prometheus/storage"
	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/consumer"
	"go.opentelemetry.io/collector/pdata/pcommon"
	"go.opentelemetry.io/collector/pdata/pmetric"
	"go.opentelemetry.io/collector/receiver"
	"go.opentelemetry.io/collector/receiver/receiverhelper"
	"go.uber.org/zap"
)

type resourceKey struct {
	job      string
	instance string
}

// The name of the metric family doesn't include magic suffixes (e.g. _bucket),
// so for a classic histgram and a native histogram of the same family, the
// metric family will be the same. To be able to tell them apart, we need to
// store whether the metric is a native histogram or not.
type metricFamilyKey struct {
	isExponentialHistogram bool
	name                   string
}

type transaction struct {
	isNew                 bool
	trimSuffixes          bool
	useMetadata           bool
	addingNativeHistogram bool // true if the last sample was a native histogram.
	addingNHCB            bool // true if the last sample was a NHCB.
	ctx                   context.Context
	families              map[resourceKey]map[scopeID]map[metricFamilyKey]*metricFamily
	mc                    scrape.MetricMetadataStore
	sink                  consumer.Metrics
	externalLabels        labels.Labels
	nodeResources         map[resourceKey]pcommon.Resource
	scopeAttributes       map[resourceKey]map[scopeID]pcommon.Map
	ignoreScopeInfoMetric bool
	logger                *zap.Logger
	buildInfo             component.BuildInfo
	obsrecv               *receiverhelper.ObsReport
	// Used as buffer to calculate series ref hash.
	bufBytes []byte
}

var emptyScopeID scopeID

type scopeID struct {
	name      string
	version   string
	schemaURL string
	attrsHash [16]byte
}

func newTransaction(
	ctx context.Context,
	sink consumer.Metrics,
	externalLabels labels.Labels,
	settings receiver.Settings,
	obsrecv *receiverhelper.ObsReport,
	trimSuffixes bool,
	useMetadata bool,
) *transaction {
	_ = "STUB: not implemented"
	return nil
}

// append returns a stable series reference to enable Prometheus staleness tracking.
func (t *transaction) append(ls labels.Labels, atMs int64, val float64) (storage.SeriesRef, error) {
	_ = "STUB: not implemented"
	return *new(storage.SeriesRef), nil
}

// addSampleDatapoint processes one scraped sample and stores it in the
// appropriate metric family for the resource/scope context. It is shared by
// both V1 and V2 appender paths.
func (t *transaction) addSampleDatapoint(rKey resourceKey, ls labels.Labels, metricName string, atMs int64, val float64, stMs int64) (storage.SeriesRef, error) {
	_ = "STUB: not implemented"
	// See https://www.prometheus.io/docs/concepts/jobs_instances/#automatically-generated-labels-and-time-series
	// up: 1 if the instance is healthy, i.e. reachable, or 0 if the scrape failed.
	// But it can also be a staleNaN, which is inserted when the target goes away.
	return *new(storage.SeriesRef), nil
}

// For the `target_info` metric we need to convert it to resource attributes.

// For the `otel_scope_info` metric we need to convert it to scope attributes.

// never return errors, as that fails the whole scrape
// return ref==0 indicating that the series was not added

// never return errors, as that fails the whole scrape
// return a stable ref so Prometheus can track series staleness

// detectAndStoreNativeHistogramStaleness returns true if it detects
// and stores a native histogram staleness marker.
func (t *transaction) detectAndStoreNativeHistogramStaleness(atMs int64, key resourceKey, scope scopeID, metricName string, ls labels.Labels) bool {
	_ = "STUB: not implemented"
	// Detect the special case of stale native histogram series.
	// Currently Prometheus does not store the histogram type in
	// its staleness tracker.
	return false
}

// Native histograms always have metadata.

// Not a histogram.

// Not a native histogram because it has magic suffixes (e.g. _bucket).

// Store the staleness marker as a native histogram.

// ignore errors here, this is best effort.

// getOrCreateMetricFamily returns the metric family for the given metric name and scope,
// and true if an existing family was found.
func (t *transaction) getOrCreateMetricFamily(key resourceKey, scope scopeID, mn string) *metricFamily {
	_ = "STUB: not implemented"
	return nil
}

// NB (eriksywu): see https://github.com/prometheus/prometheus/issues/14823

// END NB (eriksywu)

func (t *transaction) appendExemplar(l labels.Labels, e exemplar.Exemplar) error {
	_ = "STUB: not implemented"
	return nil
}

func (t *transaction) appendHistogram(ls labels.Labels, atMs int64, h *histogram.Histogram, fh *histogram.FloatHistogram) (storage.SeriesRef, error) {
	_ = "STUB: not implemented"
	return *new(storage.SeriesRef), nil
}

// The `up`, `target_info`, `otel_scope_info` metrics should never generate native histograms,
// thus we don't check for them here as opposed to the Append function.

// addHistogramDatapoint adds a native histogram or NHCB datapoint to the appropriate metric family.
// When stMs != 0, it also records a creation timestamp on the metric family.
// It is shared by both V1 and V2 appender paths.
func (t *transaction) addHistogramDatapoint(rKey resourceKey, ls labels.Labels, metricName string, atMs int64, h *histogram.Histogram, fh *histogram.FloatHistogram, schema int32, stMs int64) (storage.SeriesRef, error) {
	_ = "STUB: not implemented"
	return *new(storage.SeriesRef), nil
}

// never return errors, as that fails the whole scrape
// return ref==0 indicating that the series was not added

// never return errors, as that fails the whole scrape
// return a stable ref so Prometheus can track series staleness

func (t *transaction) appendSTZeroSample(ls labels.Labels, atMs, stMs int64) (storage.SeriesRef, error) {
	_ = "STUB: not implemented"
	return *new(storage.SeriesRef), nil
}

func (t *transaction) appendHistogramSTZeroSample(ls labels.Labels, atMs, stMs int64, h *histogram.Histogram, fh *histogram.FloatHistogram) (storage.SeriesRef, error) {
	_ = "STUB: not implemented"
	return *new(storage.SeriesRef), nil
}

// prepareLabels merges external labels, initializes the transaction, validates
// labels and extracts the metric name. It is shared by both V1 and V2 appender paths.
func (t *transaction) prepareLabels(ls labels.Labels) (labels.Labels, *resourceKey, string, error) {
	_ = "STUB: not implemented"
	return *new(labels.Labels), nil, "", nil
}

// Any datapoint with duplicate labels MUST be rejected per:
// * https://github.com/open-telemetry/wg-prometheus/issues/44
// * https://github.com/open-telemetry/opentelemetry-collector/issues/3407
// as Prometheus rejects such too as of version 2.16.0, released on 2020-02-13.

func (t *transaction) setStartTimestamp(ls labels.Labels, atMs, stMs int64) (storage.SeriesRef, error) {
	_ = "STUB: not implemented"
	return *new(storage.SeriesRef), nil
}

func (*transaction) SetOptions(_ *storage.AppendOptions) {
	_ = "STUB: not implemented"
	// TODO: implement this func
	return
}

func (t *transaction) getSeriesRef(ls labels.Labels, mtype pmetric.MetricType) uint64 {
	_ = "STUB: not implemented"
	return 0
}

// getMetrics returns all metrics to the given slice.
// The only error returned by this function is errNoDataToBuild.
func (t *transaction) getMetrics() (pmetric.Metrics, error) {
	_ = "STUB: not implemented"
	return *new(pmetric.Metrics), nil
}

// If metrics don't include otel_scope_name or otel_scope_version
// labels, use the receiver name and version.

// Otherwise, use the scope that was provided with the metrics.

// remove the resource if no metrics were added to avoid returning resources with empty data points

func getScopeID(ls labels.Labels) (scopeID, pcommon.Map) {
	_ = "STUB: not implemented"
	return *new(scopeID), *new(pcommon.Map)
}

func (t *transaction) addScopeAttributesFromLabels(key resourceKey, scope scopeID, attrs pcommon.Map) {
	_ = "STUB: not implemented"
	return
}

func (t *transaction) initTransaction(lbs labels.Labels) (*resourceKey, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (t *transaction) getJobAndInstance(labels labels.Labels) (*resourceKey, error) {
	_ = "STUB: not implemented"
	// first, try to get job and instance from the labels
	return nil, nil
}

// if not available in the labels, try to fall back to the scrape job associated
// with the transaction.
// this can be the case for, e.g., aggregated metrics coming from a federate endpoint
// that represent the whole cluster, rather than an individual workload.
// See https://github.com/open-telemetry/opentelemetry-collector-contrib/issues/32555 for reference

func (t *transaction) Commit() error { _ = "STUB: not implemented"; return nil }

func (*transaction) Rollback() error { _ = "STUB: not implemented"; return nil }

func (*transaction) updateMetadata(_ storage.SeriesRef, _ labels.Labels, _ metadata.Metadata) (storage.SeriesRef, error) {
	_ = "STUB: not implemented"
	// TODO: implement this func
	return *new(storage.SeriesRef), nil
}

func (t *transaction) AddTargetInfo(key resourceKey, ls labels.Labels) {
	_ = "STUB: not implemented"
	return
}

func (t *transaction) addScopeInfo(key resourceKey, ls labels.Labels) {
	_ = "STUB: not implemented"
	return
}

func getSeriesRefWithoutScopeLabels(bytes []byte, ls labels.Labels, mtype pmetric.MetricType) (uint64, []byte) {
	_ = "STUB: not implemented"
	return 0, nil
}

// Append implements storage.AppenderV2.
func (t *transaction) Append(_ storage.SeriesRef, ls labels.Labels, stMs, atMs int64, val float64, h *histogram.Histogram, fh *histogram.FloatHistogram, opts storage.AppendV2Options) (storage.SeriesRef, error) {
	_ = "STUB: not implemented"
	return *new(storage.SeriesRef), nil
}

// Append the exemplars, continuing on error to try all exemplars.
