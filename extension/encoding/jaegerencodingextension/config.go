// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package jaegerencodingextension // import "github.com/open-telemetry/opentelemetry-collector-contrib/extension/encoding/jaegerencodingextension"

type JaegerProtocol string

const (
	JaegerProtocolProtobuf JaegerProtocol = "protobuf"
	JaegerProtocolJSON     JaegerProtocol = "json"
)

type Config struct {
	Protocol JaegerProtocol `mapstructure:"protocol"`
	// prevent unkeyed literal initialization
	_ struct{}
}

func (c *Config) Validate() error { _ = "STUB: not implemented"; return nil }
