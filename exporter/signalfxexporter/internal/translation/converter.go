// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package translation // import "github.com/open-telemetry/opentelemetry-collector-contrib/exporter/signalfxexporter/internal/translation"

import (
	"fmt"

	sfxpb "github.com/signalfx/com_signalfx_metrics_protobuf/model"
	"go.opentelemetry.io/collector/pdata/pcommon"
	"go.opentelemetry.io/collector/pdata/pmetric"
	"go.uber.org/zap"

	"github.com/open-telemetry/opentelemetry-collector-contrib/exporter/signalfxexporter/internal/translation/dpfilters"
	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/translator/signalfx"
)

// Some fields on SignalFx protobuf are pointers, in order to reduce
// allocations create the most used ones.
var (
	// SignalFx metric types used in the conversions.
	sfxMetricTypeGauge             = sfxpb.MetricType_GAUGE
	sfxMetricTypeCumulativeCounter = sfxpb.MetricType_CUMULATIVE_COUNTER
	sfxMetricTypeCounter           = sfxpb.MetricType_COUNTER
)

// MetricsConverter converts MetricsData to sfxpb DataPoints. It holds an optional
// MetricTranslator to translate SFx metrics using translation rules.
type MetricsConverter struct {
	logger               *zap.Logger
	metricTranslator     *MetricTranslator
	filterSet            *dpfilters.FilterSet
	datapointValidator   *datapointValidator
	translator           *signalfx.FromTranslator
	dropHistogramBuckets bool
	processHistograms    bool
}

// NewMetricsConverter creates a MetricsConverter from the passed in logger and
// MetricTranslator. Pass in a nil MetricTranslator to not use translation
// rules.
func NewMetricsConverter(
	logger *zap.Logger,
	t *MetricTranslator,
	excludes []dpfilters.MetricFilter,
	includes []dpfilters.MetricFilter,
	nonAlphanumericDimChars string,
	dropHistogramBuckets bool,
	processHistograms bool,
) (*MetricsConverter, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *MetricsConverter) Start() { _ = "STUB: not implemented"; return }

// MetricsToSignalFxV2 converts the passed in MetricsData to SFx datapoints
// and if processHistograms is set, histogram metrics are not converted to SFx format.
// It returns those datapoints and the number of time series that had to be
// dropped because of errors or warnings.
func (c *MetricsConverter) MetricsToSignalFxV2(md pmetric.Metrics) []*sfxpb.DataPoint {
	_ = "STUB: not implemented"
	return nil
}

// Translate and filter all metrics within the current ScopeMetric

func (c *MetricsConverter) translateAndFilter(dps []*sfxpb.DataPoint) []*sfxpb.DataPoint {
	_ = "STUB: not implemented"
	return nil
}

// resourceToDimensions will return a set of dimension from the
// resource attributes, including a cloud host id (AWSUniqueId, gcp_id, etc.)
// if it can be constructed from the provided metadata.
func resourceToDimensions(res pcommon.Resource) []*sfxpb.Dimension {
	_ = "STUB: not implemented"
	return nil
}

// Never send the SignalFX token

func (c *MetricsConverter) Shutdown() { _ = "STUB: not implemented"; return }

// Values obtained from https://dev.splunk.com/observability/docs/datamodel/ingest#Criteria-for-metric-and-dimension-names-and-values
const (
	maxMetricNameLength     = 256
	maxDimensionNameLength  = 128
	maxDimensionValueLength = 256
	maxNumberOfDimensions   = 36
)

var (
	invalidMetricNameReason = fmt.Sprintf(
		"metric name longer than %d characters", maxMetricNameLength)
	invalidDimensionNameReason = fmt.Sprintf(
		"dimension name longer than %d characters", maxDimensionNameLength)
	invalidDimensionValueReason = fmt.Sprintf(
		"dimension value longer than %d characters", maxDimensionValueLength)
	invalidNumberOfDimensions = fmt.Sprintf(
		"number of dimensions is larger than %d", maxNumberOfDimensions)
)

type datapointValidator struct {
	logger                  *zap.Logger
	nonAlphanumericDimChars string
}

func newDatapointValidator(logger *zap.Logger, nonAlphanumericDimChars string) *datapointValidator {
	_ = "STUB: not implemented"
	return nil
}

// sanitizeDataPoints sanitizes datapoints prior to dispatching them to the backend.
// Datapoints that do not conform to the requirements are removed. This method drops
// datapoints with metric name greater than 256 characters and number of dimensions greater than 36.
func (dpv *datapointValidator) sanitizeDataPoints(dps []*sfxpb.DataPoint) []*sfxpb.DataPoint {
	_ = "STUB: not implemented"
	return nil
}

// Trim datapoints slice to account for any removed datapoints.

// sanitizeDimensions replaces all characters unsupported by SignalFx backend
// in metric label keys and with "_" and drops dimensions when the key is greater
// than 128 characters or when value is greater than 256 characters in length.
func (dpv *datapointValidator) sanitizeDimensions(dims []*sfxpb.Dimension) []*sfxpb.Dimension {
	_ = "STUB: not implemented"
	return nil
}

// Trim dimensions slice to account for any removed dimensions.

func (dpv *datapointValidator) isValidMetricName(name string) bool {
	_ = "STUB: not implemented"
	return false
}

func (dpv *datapointValidator) isValidNumberOfDimension(dp *sfxpb.DataPoint) bool {
	_ = "STUB: not implemented"
	return false
}

func (dpv *datapointValidator) isValidDimension(dimension *sfxpb.Dimension) bool {
	_ = "STUB: not implemented"
	return false
}

func (dpv *datapointValidator) isValidDimensionName(name string) bool {
	_ = "STUB: not implemented"
	return false
}

func (dpv *datapointValidator) isValidDimensionValue(value, name string) bool {
	_ = "STUB: not implemented"
	return false
}
