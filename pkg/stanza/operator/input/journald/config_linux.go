// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

//go:build linux

package journald // import "github.com/open-telemetry/opentelemetry-collector-contrib/pkg/stanza/operator/input/journald"

import (
	"context"
	"time"

	"go.opentelemetry.io/collector/component"
	"go.uber.org/zap"

	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/stanza/operator"
)

const waitDuration = 1 * time.Second

func init() {
	operator.Register(operatorType, func() operator.Builder { return NewConfig() })
}

// Build will build a journald input operator from the supplied configuration
func (c Config) Build(set component.TelemetrySettings) (operator.Operator, error) {
	_ = "STUB: not implemented"
	return *new(operator.Operator), nil
}

func (c Config) validate() error { _ = "STUB: not implemented"; return nil }

func (c Config) buildArgs() ([]string, error) { _ = "STUB: not implemented"; return nil, nil }

// Export logs in UTC time
// Export logs as JSON
// Continue watching logs until cancelled

func buildMatchConfig(mc MatchConfig) ([]string, error) { _ = "STUB: not implemented"; return nil, nil }

// Sort keys to be consistent with every run and to be predictable for tests

func (c Config) buildMatchesConfig() ([]string, error) { _ = "STUB: not implemented"; return nil, nil }

func (c Config) buildNewCmdFunc(logger *zap.Logger) (func(ctx context.Context, cursor []byte) cmd, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Copy args and if needed, add the cursor flag

// #nosec - ...
// journalctl is an executable that is required for this operator to function
