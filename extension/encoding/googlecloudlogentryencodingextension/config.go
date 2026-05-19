// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package googlecloudlogentryencodingextension // import "github.com/open-telemetry/opentelemetry-collector-contrib/extension/encoding/googlecloudlogentryencodingextension"

import (
	"go.opentelemetry.io/collector/confmap/xconfmap"
)

var _ xconfmap.Validator = (*Config)(nil)

type HandleAs string

const (
	HandleAsProtobuf HandleAs = "protobuf"
	HandleAsJSON     HandleAs = "json"
	HandleAsText     HandleAs = "text"
)

type Config struct {
	// Controls how the json payload of the [LogEntry]  is parsed into the body.
	HandleJSONPayloadAs HandleAs `mapstructure:"handle_json_payload_as"`
	// Controls how the proto payload of the [LogEntry]  is parsed into the body.
	HandleProtoPayloadAs HandleAs `mapstructure:"handle_proto_payload_as"`
	// prevent unkeyed literal initialization
	_ struct{}
}

func (config *Config) Validate() error { _ = "STUB: not implemented"; return nil }
