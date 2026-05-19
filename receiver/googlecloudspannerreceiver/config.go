// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package googlecloudspannerreceiver // import "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/googlecloudspannerreceiver"

import (
	"go.opentelemetry.io/collector/scraper/scraperhelper"
)

const (
	minCollectionIntervalSeconds = 60
	maxTopMetricsQueryMaxRows    = 100
)

type Config struct {
	scraperhelper.ControllerConfig `mapstructure:",squash"`

	TopMetricsQueryMaxRows            int       `mapstructure:"top_metrics_query_max_rows"`
	BackfillEnabled                   bool      `mapstructure:"backfill_enabled"`
	CardinalityTotalLimit             int       `mapstructure:"cardinality_total_limit"`
	Projects                          []Project `mapstructure:"projects"`
	HideTopnLockstatsRowrangestartkey bool      `mapstructure:"hide_topn_lockstats_rowrangestartkey"`
	TruncateText                      bool      `mapstructure:"truncate_text"`
}

type Project struct {
	ID                string     `mapstructure:"project_id"`
	ServiceAccountKey string     `mapstructure:"service_account_key"`
	Instances         []Instance `mapstructure:"instances"`
}

type Instance struct {
	ID        string   `mapstructure:"instance_id"`
	Databases []string `mapstructure:"databases"`
}

func (config *Config) Validate() error { _ = "STUB: not implemented"; return nil }

func (project Project) Validate() error { _ = "STUB: not implemented"; return nil }

func (instance Instance) Validate() error { _ = "STUB: not implemented"; return nil }
