// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package receivercreator // import "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/receivercreator"

import (
	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/confmap"

	"github.com/open-telemetry/opentelemetry-collector-contrib/extension/observer"
)

const (
	// receiversConfigKey is the config key name used to specify the subreceivers.
	receiversConfigKey = "receivers"
	// endpointConfigKey is the key name mapping to ReceiverSettings.Endpoint.
	endpointConfigKey = "endpoint"
	// configKey is the key name in a subreceiver.
	configKey = "config"
)

// receiverConfig describes a receiver instance with a default config.
type receiverConfig struct {
	// id is the id of the subreceiver (ie <receiver type>/<id>).
	id component.ID
	// config is the map configured by the user in the config file. It is the contents of the map from
	// the "config" section. The keys and values are arbitrarily configured by the user.
	config     userConfigMap
	endpointID observer.EndpointID
}

// userConfigMap is an arbitrary map of string keys to arbitrary values as specified by the user
type userConfigMap map[string]any

type receiverSignals struct {
	metrics  bool
	logs     bool
	traces   bool
	profiles bool
}

// receiverTemplate is the configuration of a single subreceiver.
type receiverTemplate struct {
	receiverConfig

	// Rule is the discovery rule that when matched will create a receiver instance
	// based on receiverTemplate.
	Rule string `mapstructure:"rule"`
	// ResourceAttributes is a map of resource attributes to add to just this receiver's resource metrics.
	// It can contain expr expressions for endpoint env value expansion
	ResourceAttributes map[string]any `mapstructure:"resource_attributes"`
	rule               rule
	signals            receiverSignals
}

// resourceAttributes holds a map of default resource attributes for each Endpoint type.
type resourceAttributes map[observer.EndpointType]map[string]string

// newReceiverTemplate creates a receiverTemplate instance from the full name of a subreceiver
// and its arbitrary config map values.
func newReceiverTemplate(name string, cfg userConfigMap) (receiverTemplate, error) {
	_ = "STUB: not implemented"
	return *new(receiverTemplate), nil
}

var _ confmap.Unmarshaler = (*Config)(nil)

// Config defines configuration for receiver_creator.
type Config struct {
	receiverTemplates map[string]receiverTemplate
	// WatchObservers are the extensions to listen to endpoints from.
	WatchObservers []component.ID `mapstructure:"watch_observers"`
	// ResourceAttributes is a map of default resource attributes to add to each resource
	// object received by this receiver from dynamically created receivers.
	ResourceAttributes resourceAttributes `mapstructure:"resource_attributes"`
	Discovery          DiscoveryConfig    `mapstructure:"discovery"`
}

type DiscoveryConfig struct {
	Enabled              bool              `mapstructure:"enabled"`
	IgnoreReceivers      []string          `mapstructure:"ignore_receivers"`
	DefaultAnnotations   map[string]string `mapstructure:"default_annotations"`
	DefaultFileLogConfig userConfigMap     `mapstructure:"default_file_log_config"`
}

func (cfg *Config) Unmarshal(componentParser *confmap.Conf) error {
	_ = "STUB: not implemented"
	return nil

	// Nothing to do if there is no config given.
}

// Unmarshals receiver_creator configuration like rule.
