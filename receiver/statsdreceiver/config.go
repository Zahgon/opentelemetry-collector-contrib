// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package statsdreceiver // import "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/statsdreceiver"

import (
	"os"
	"time"

	"go.opentelemetry.io/collector/config/confignet"

	"github.com/open-telemetry/opentelemetry-collector-contrib/receiver/statsdreceiver/protocol"
)

// Config defines configuration for StatsD receiver.
type Config struct {
	NetAddr                 confignet.AddrConfig `mapstructure:",squash"`
	AggregationInterval     time.Duration        `mapstructure:"aggregation_interval"`
	EnableIPOnlyAggregation bool                 `mapstructure:"enable_ip_only_aggregation"`
	IgnoreHost              bool                 `mapstructure:"ignore_host"`
	EnableMetricType        bool                 `mapstructure:"enable_metric_type"`
	EnableSimpleTags        bool                 `mapstructure:"enable_simple_tags"`
	IsMonotonicCounter      bool                 `mapstructure:"is_monotonic_counter"`
	// CounterType specifies how counter values are represented in exported metrics.
	// Valid values: "int" (default), "float", "stochastic_int".
	CounterType           protocol.CounterType             `mapstructure:"counter_type"`
	TimerHistogramMapping []protocol.TimerHistogramMapping `mapstructure:"timer_histogram_mapping"`
	// Will only be used when transport set to 'unixgram'.
	SocketPermissions os.FileMode `mapstructure:"socket_permissions"`
}

func (c *Config) Validate() error { _ = "STUB: not implemented"; return nil }

// valid

// do nothing

// do nothing

// Non-histogram observer w/ histogram config

func (*Config) validateExplicitBuckets(explicitBuckets []protocol.ExplicitBucket) error {
	_ = "STUB: not implemented"
	return nil
}
