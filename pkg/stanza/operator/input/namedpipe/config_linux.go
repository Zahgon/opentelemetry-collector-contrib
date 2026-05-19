// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

//go:build linux

package namedpipe // import "github.com/open-telemetry/opentelemetry-collector-contrib/pkg/stanza/operator/input/namedpipe"

import (
	"go.opentelemetry.io/collector/component"

	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/stanza/operator"
)

func init() {
	operator.Register(operatorType, func() operator.Builder { return NewConfig() })
}

// Build will build a namedpipe input operator.
func (c *Config) Build(set component.TelemetrySettings) (operator.Operator, error) {
	_ = "STUB: not implemented"
	return *new(operator.Operator), nil
}
