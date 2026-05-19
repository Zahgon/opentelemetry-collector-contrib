// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package podmanreceiver // import "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/podmanreceiver"

import (
	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/receiver"
)

const (
	defaultAPIVersion = "3.3.1"
)

func NewFactory() receiver.Factory { _ = "STUB: not implemented"; return *new(receiver.Factory) }

func createDefaultConfig() component.Config {
	_ = "STUB: not implemented"
	return *new(component.Config)
}
