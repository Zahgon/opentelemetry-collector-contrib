// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package kafkametricsreceiver // import "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/kafkametricsreceiver"

import (
	"time"

	"go.opentelemetry.io/collector/confmap"
	"go.opentelemetry.io/collector/scraper/scraperhelper"

	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/kafka/configkafka"
	"github.com/open-telemetry/opentelemetry-collector-contrib/receiver/kafkametricsreceiver/internal/metadata"
)

// Config represents user settings for kafkametrics receiver
type Config struct {
	scraperhelper.ControllerConfig `mapstructure:",squash"`
	configkafka.ClientConfig       `mapstructure:",squash"`

	// Alias name of the kafka cluster
	ClusterAlias string `mapstructure:"cluster_alias"`

	// TopicMatch topics to collect metrics on
	TopicMatch string `mapstructure:"topic_match"`

	// GroupMatch consumer groups to collect on
	GroupMatch string `mapstructure:"group_match"`

	// Cluster metadata refresh frequency
	// Configures the refresh frequency to update cached cluster metadata
	//
	// If Metadata.RefreshInterval is set, this will be ignored.
	//
	// Deprecated [v0.122.0]: use Metadata.RefreshInterval instead.
	RefreshFrequency time.Duration `mapstructure:"refresh_frequency"`

	// Scrapers defines which metric data points to be captured from kafka
	Scrapers []string `mapstructure:"scrapers"`

	// MetricsBuilderConfig allows customizing scraped metrics/attributes representation.
	metadata.MetricsBuilderConfig `mapstructure:",squash"`
}

func (c *Config) Unmarshal(conf *confmap.Conf) error { _ = "STUB: not implemented"; return nil }

// User has not explicitly set metadata.refresh_interval,
// but they have set the (deprecated) refresh_frequency,
// so use that.
