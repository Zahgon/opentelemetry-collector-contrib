// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package netflowreceiver // import "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/netflowreceiver"

// Config represents the receiver config settings within the collector's config.yaml
type Config struct {
	// The scheme defines the type of flow data that the listener will receive
	// The scheme must be one of sflow, netflow, or flow
	Scheme string `mapstructure:"scheme"`

	// The hostname or IP address that the listener will bind to
	Hostname string `mapstructure:"hostname"`

	// The port that the listener will bind to
	Port int `mapstructure:"port"`

	// The number of sockets that the listener will use
	Sockets int `mapstructure:"sockets"`

	// The number of workers that the listener will use to decode incoming flow messages
	// By default it will be two times the number of sockets
	// Ideally set this to the number of CPU cores
	Workers int `mapstructure:"workers"`

	// The size of the queue that the listener will use
	// This is a buffer that will hold flow messages before they are processed by a worker
	QueueSize int `mapstructure:"queue_size"`

	// SendRaw determines whether to send raw flow messages instead of parsing them
	SendRaw bool `mapstructure:"send_raw"`
}

// Validate checks if the receiver configuration is valid
func (cfg *Config) Validate() error { _ = "STUB: not implemented"; return nil }
