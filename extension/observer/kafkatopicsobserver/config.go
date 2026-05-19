// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package kafkatopicsobserver // import "github.com/open-telemetry/opentelemetry-collector-contrib/extension/observer/kafkatopicsobserver"

import (
	"time"

	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/kafka/configkafka"
)

// Config defines configuration for docker observer
type Config struct {
	configkafka.ClientConfig `mapstructure:",squash"`
	TopicRegex               string        `mapstructure:"topic_regex"`
	TopicsSyncInterval       time.Duration `mapstructure:"topics_sync_interval"`
}

func (config *Config) Validate() (errs error) { _ = "STUB: not implemented"; return nil }
