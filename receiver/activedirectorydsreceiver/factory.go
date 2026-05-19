// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package activedirectorydsreceiver // import "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/activedirectorydsreceiver"

import (
	"time"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/receiver"
)

const (
	defaultCollectionInterval = 10 * time.Second
)

func NewFactory() receiver.Factory { _ = "STUB: not implemented"; return *new(receiver.Factory) }

func createDefaultConfig() component.Config {
	_ = "STUB: not implemented"
	return *new(component.Config)
}
