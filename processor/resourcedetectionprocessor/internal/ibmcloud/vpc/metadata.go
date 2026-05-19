// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package vpc // import "github.com/open-telemetry/opentelemetry-collector-contrib/processor/resourcedetectionprocessor/internal/ibmcloud/vpc"

import (
	"go.opentelemetry.io/collector/component"
)

type dummyFactory struct{}

func (dummyFactory) Type() component.Type { _ = "STUB: not implemented"; return *new(component.Type) }

func (dummyFactory) CreateDefaultConfig() component.Config {
	_ = "STUB: not implemented"

	// Necessary to satisfy mdatagen tests
	return *new(component.Config)
}

func NewFactory() component.Factory { _ = "STUB: not implemented"; return *new(component.Factory) }
