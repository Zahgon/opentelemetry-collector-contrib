// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package metadata // import "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/googlecloudspannerreceiver/internal/metadata"

import (
	"time"

	"go.opentelemetry.io/collector/pdata/pmetric"

	"github.com/open-telemetry/opentelemetry-collector-contrib/receiver/googlecloudspannerreceiver/internal/datasource"
	"github.com/open-telemetry/opentelemetry-collector-contrib/receiver/googlecloudspannerreceiver/internal/filter"
)

const (
	projectIDLabelName  = "project_id"
	instanceIDLabelName = "instance_id"
	databaseLabelName   = "database"
)

type MetricsDataPointKey struct {
	MetricName string
	MetricUnit string
	MetricType MetricType
}

type MetricsDataPoint struct {
	metricName  string
	timestamp   time.Time
	databaseID  *datasource.DatabaseID
	labelValues []LabelValue
	metricValue MetricValue
}

// Fields must be exported for hashing purposes
type dataForHashing struct {
	MetricName string
	Labels     []label
}

// Fields must be exported for hashing purposes
type label struct {
	Name  string
	Value any
}

func (mdp *MetricsDataPoint) CopyTo(dataPoint pmetric.NumberDataPoint) {
	_ = "STUB: not implemented"
	return
}

func (mdp *MetricsDataPoint) GroupingKey() MetricsDataPointKey {
	_ = "STUB: not implemented"
	return *new(MetricsDataPointKey)
}

func (mdp *MetricsDataPoint) ToItem() (*filter.Item, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (mdp *MetricsDataPoint) toDataForHashing() dataForHashing {
	_ = "STUB: not implemented"
	// Do not use map here because it has unpredicted order
	// Taking into account 3 default labels: project_id, instance_id, database
	return *new(dataForHashing)
}

// Convert row_range_start_key label of top-lock-stats metric from format "sample(key1, key2)" to "sample(hash1, hash2)"
func parseAndHashRowrangestartkey(key string) string { _ = "STUB: not implemented"; return "" }

// if "(" does not exist or is the last character of the string, then label is of incorrect format

func (mdp *MetricsDataPoint) HideLockStatsRowrangestartkeyPII() { _ = "STUB: not implemented"; return }

func TruncateString(str string, length int) string { _ = "STUB: not implemented"; return "" }

func (mdp *MetricsDataPoint) TruncateQueryText(length int) { _ = "STUB: not implemented"; return }

func (mdp *MetricsDataPoint) hash() (string, error) { _ = "STUB: not implemented"; return "", nil }
