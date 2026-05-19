// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package pipeline // import "github.com/open-telemetry/opentelemetry-collector-contrib/pkg/stanza/pipeline"

import (
	"go.opentelemetry.io/collector/component"

	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/stanza/operator"
)

// Config is the configuration of a pipeline.
type Config struct {
	DefaultOutput operator.Operator
	Operators     []operator.Config
}

// Build will build a pipeline from the config.
func (c Config) Build(set component.TelemetrySettings) (*DirectedPipeline, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Any operator that already has an output will not be changed

// Any operator (except the last) will just output to the next

// The last operator may output to the default output

func dedeplucateIDs(ops []operator.Config) { _ = "STUB: not implemented"; return }
