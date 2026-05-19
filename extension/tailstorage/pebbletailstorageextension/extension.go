// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package pebbletailstorageextension // import "github.com/open-telemetry/opentelemetry-collector-contrib/extension/tailstorage/pebbletailstorageextension"

import (
	"context"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/extension"
	"go.opentelemetry.io/collector/pdata/pcommon"
	"go.opentelemetry.io/collector/pdata/ptrace"
)

type pebbleTailStorageExtension struct {
	settings extension.Settings
	cfg      *Config

	storage *storage
}

var _ extension.Extension = (*pebbleTailStorageExtension)(nil)

func newExtension(settings extension.Settings, cfg *Config) *pebbleTailStorageExtension {
	_ = "STUB: not implemented"
	return nil
}

func (e *pebbleTailStorageExtension) Start(_ context.Context, _ component.Host) error {
	_ = "STUB: not implemented"
	return nil
}

func (e *pebbleTailStorageExtension) Shutdown(_ context.Context) error {
	_ = "STUB: not implemented"
	return nil
}

func (e *pebbleTailStorageExtension) Append(traceID pcommon.TraceID, rss ptrace.ResourceSpans) {
	_ = "STUB: not implemented"
	return
}

func (e *pebbleTailStorageExtension) Take(traceID pcommon.TraceID) (ptrace.Traces, bool) {
	_ = "STUB: not implemented"
	return *new(ptrace.Traces), false
}

func (e *pebbleTailStorageExtension) Delete(traceID pcommon.TraceID) {
	_ = "STUB: not implemented"
	return
}
