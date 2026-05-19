// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package otelarrowexporter // import "github.com/open-telemetry/opentelemetry-collector-contrib/exporter/otelarrowexporter"

import (
	"context"
	"errors"
	"sync"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/consumer/consumererror"
	"go.opentelemetry.io/collector/exporter"
	"go.opentelemetry.io/collector/pdata/plog"
	"go.opentelemetry.io/collector/pdata/pmetric"
	"go.opentelemetry.io/collector/pdata/ptrace"
	"go.opentelemetry.io/otel/attribute"
	"google.golang.org/grpc/metadata"

	"github.com/open-telemetry/opentelemetry-collector-contrib/internal/otelarrow/netstats"
)

// errTooManyExporters is returned when the MetadataCardinalityLimit has been reached.
var errTooManyExporters = consumererror.NewPermanent(errors.New("too many exporter metadata-value combinations"))

type metadataExporter struct {
	config   *Config
	settings exporter.Settings
	scf      streamClientFactory
	host     component.Host

	metadataKeys []string
	exporters    sync.Map
	netReporter  *netstats.NetworkReporter

	userAgent string

	// Guards the size and the storing logic to ensure no more than limit items are stored.
	// If we are willing to allow "some" extra items than the limit this can be removed and size can be made atomic.
	lock sync.Mutex
	size int
}

var _ exp = (*metadataExporter)(nil)

func newMetadataExporter(cfg component.Config, set exporter.Settings, streamClientFactory streamClientFactory) (exp, error) {
	_ = "STUB: not implemented"
	return *new(exp), nil
}

// Ignoring an error because Validate() was called.

// use lower-case, to be consistent with http/2 headers.

func (e *metadataExporter) getSettings() exporter.Settings {
	_ = "STUB: not implemented"
	return *new(exporter.Settings)
}

func (e *metadataExporter) getConfig() component.Config {
	_ = "STUB: not implemented"
	return *new(component.Config)
}

func (e *metadataExporter) start(_ context.Context, host component.Host) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func (e *metadataExporter) shutdown(ctx context.Context) error {
	_ = "STUB: not implemented"
	return nil
}

func (e *metadataExporter) pushTraces(ctx context.Context, td ptrace.Traces) error {
	_ = "STUB: not implemented"
	return nil
}

func (e *metadataExporter) pushMetrics(ctx context.Context, md pmetric.Metrics) error {
	_ = "STUB: not implemented"
	return nil
}

func (e *metadataExporter) pushLogs(ctx context.Context, ld plog.Logs) error {
	_ = "STUB: not implemented"
	return nil
}

func (e *metadataExporter) getOrCreateExporter(ctx context.Context, s attribute.Set, md metadata.MD) (exp, error) {
	_ = "STUB: not implemented"
	return *new(exp), nil
}

// set metadata keys for base exporter to add them to the outgoing context.

// Start the goroutine only if we added the object to the map, otherwise is already started.

// getAttrSet is code taken from the core collector's batchprocessor multibatch logic.
// https://github.com/open-telemetry/opentelemetry-collector/blob/v0.107.0/processor/batchprocessor/batch_processor.go#L298
func (*metadataExporter) getAttrSet(ctx context.Context, keys []string) (attribute.Set, metadata.MD) {
	_ = "STUB: not implemented"
	// Get each metadata key value, form the corresponding
	// attribute set for use as a map lookup key.
	return *new(attribute.Set), *new(metadata.MD)
}

// Lookup the value in the incoming metadata, copy it
// into the outgoing metadata, and create a unique
// value for the attributeSet.
