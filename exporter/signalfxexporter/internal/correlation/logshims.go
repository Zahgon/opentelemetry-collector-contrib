// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package correlation // import "github.com/open-telemetry/opentelemetry-collector-contrib/exporter/signalfxexporter/internal/correlation"

import (
	"go.uber.org/zap"

	"github.com/open-telemetry/opentelemetry-collector-contrib/exporter/signalfxexporter/internal/apm/log"
)

type zapShim struct {
	log *zap.Logger
}

func newZapShim(log *zap.Logger) zapShim {
	_ = "STUB: not implemented"
	// Add caller skip so that the shim isn't logged as the caller.
	return *new(zapShim)
}

func (z zapShim) Debug(msg string) { _ = "STUB: not implemented"; return }

func (z zapShim) Warn(msg string) { _ = "STUB: not implemented"; return }

func (z zapShim) Error(msg string) { _ = "STUB: not implemented"; return }

func (z zapShim) Info(msg string) { _ = "STUB: not implemented"; return }

func (z zapShim) Panic(msg string) { _ = "STUB: not implemented"; return }

func (z zapShim) WithFields(fields log.Fields) log.Logger {
	_ = "STUB: not implemented"
	return *new(log.Logger)
}

func (z zapShim) WithError(err error) log.Logger {
	_ = "STUB: not implemented"
	return *new(log.Logger)
}

var _ log.Logger = (*zapShim)(nil)
