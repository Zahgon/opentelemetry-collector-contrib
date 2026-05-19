// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package googlecloudmonitoringreceiver // import "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/googlecloudmonitoringreceiver"

import (
	"time"

	"go.opentelemetry.io/collector/scraper/scraperhelper"
)

const (
	defaultCollectionInterval = 300 * time.Second // Default value for collection interval
	minCollectionInterval     = 60 * time.Second  // Minimum value for collection interval
	defaultFetchDelay         = 60 * time.Second  // Default value for fetch delay
)

type Config struct {
	scraperhelper.ControllerConfig `mapstructure:",squash"`

	ProjectID string `mapstructure:"project_id"`
	// Overrides the default monitoring.googleapis.com:443 endpoint.
	// Use this when targeting non-standard universe domains.
	Endpoint    string         `mapstructure:"endpoint"`
	MetricsList []MetricConfig `mapstructure:"metrics_list"`
}

type MetricConfig struct {
	MetricName string `mapstructure:"metric_name"`
	// Filter for listing metric descriptors. Only support `project` and `metric.type` as filter objects.
	// See https://cloud.google.com/monitoring/api/v3/filters#metric-descriptor-filter for more details.
	MetricDescriptorFilter string `mapstructure:"metric_descriptor_filter"`
}

func (config *Config) Validate() error { _ = "STUB: not implemented"; return nil }

func (metric MetricConfig) Validate() error { _ = "STUB: not implemented"; return nil }
