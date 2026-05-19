// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

//go:build windows

package windows // import "github.com/open-telemetry/opentelemetry-collector-contrib/pkg/stanza/operator/input/windows"

import (
	"go.opentelemetry.io/collector/component"

	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/stanza/operator"
)

func init() {
	operator.Register(operatorType, func() operator.Builder { return NewConfig() })
}

// Build will build a windows event log operator.
func (c *Config) Build(set component.TelemetrySettings) (operator.Operator, error) {
	_ = "STUB: not implemented"
	return *new(operator.Operator), nil
}

// any not empty
// any empty

func excludeProvidersSet(providers []string) map[string]struct{} {
	_ = "STUB: not implemented"
	return nil
}
