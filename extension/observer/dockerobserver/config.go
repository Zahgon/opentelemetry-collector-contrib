// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package dockerobserver // import "github.com/open-telemetry/opentelemetry-collector-contrib/extension/observer/dockerobserver"

import (
	"time"

	"go.opentelemetry.io/collector/confmap"

	"github.com/open-telemetry/opentelemetry-collector-contrib/internal/docker"
)

// Config defines configuration for docker observer
type Config struct {
	docker.Config `mapstructure:",squash"`

	// If true, the "Config.Hostname" field (if present) of the docker
	// container will be used as the discovered host that is used to configure
	// receivers.  If false or if no hostname is configured, the field
	// `NetworkSettings.IPAddress` is used instead.
	UseHostnameIfPresent bool `mapstructure:"use_hostname_if_present"`

	// If true, the observer will configure receivers for matching container endpoints
	// using the host bound ip and port.  This is useful if containers exist that are not
	// accessible to an instance of the agent running outside of the docker network stack.
	// If UseHostnameIfPresent and this config are both enabled, this setting will take precedence.
	UseHostBindings bool `mapstructure:"use_host_bindings"`

	// If true, the observer will ignore discovered container endpoints that are not bound
	// to host ports.  This is useful if containers exist that are not accessible
	// to an instance of the agent running outside of the docker network stack.
	IgnoreNonHostBindings bool `mapstructure:"ignore_non_host_bindings"`

	// If true, the observer will emit a port-less endpoint for every running container,
	// alongside any per-port endpoints. This makes every container, including those
	// without exposed ports, discoverable by receiver_creator rules matching
	// `type == "container"`.
	IncludeAllContainers bool `mapstructure:"include_all_containers"`

	// The time to wait before resyncing the list of containers the observer maintains
	// through the docker event listener example: cache_sync_interval: "20m"
	// Default: "60m"
	CacheSyncInterval time.Duration `mapstructure:"cache_sync_interval"`
}

func (config Config) Validate() error { _ = "STUB: not implemented"; return nil }

func (config *Config) Unmarshal(conf *confmap.Conf) error { _ = "STUB: not implemented"; return nil }
