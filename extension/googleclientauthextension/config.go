// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package googleclientauthextension // import "github.com/open-telemetry/opentelemetry-collector-contrib/extension/googleclientauthextension"

import (
	"github.com/GoogleCloudPlatform/opentelemetry-operations-go/extension/googleclientauthextension"
)

// Config defines configuration for the Google client auth extension.
type Config struct {
	googleclientauthextension.Config `mapstructure:",squash"`
}

func (cfg *Config) Validate() error { _ = "STUB: not implemented"; return nil }
