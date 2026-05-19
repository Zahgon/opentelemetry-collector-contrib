// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package googlecloudstorageexporter // import "github.com/open-telemetry/opentelemetry-collector-contrib/exporter/googlecloudstorageexporter"

import (
	"context"
	"errors"
	"io"
	"sync"
	"time"

	"cloud.google.com/go/storage"
	"github.com/lestrrat-go/strftime"
	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/config/configcompression"
	"go.opentelemetry.io/collector/consumer"
	"go.opentelemetry.io/collector/exporter"
	"go.opentelemetry.io/collector/pdata/plog"
	"go.opentelemetry.io/collector/pdata/ptrace"
	"go.uber.org/zap"
)

type poolItem interface {
	io.WriteCloser
	Reset(io.Writer)
}
type signalType string

const (
	signalTypeLogs   signalType = "logs"
	signalTypeTraces signalType = "traces"
)

var (
	errNotLogsMarshaler   = errors.New("extension is not a logs marshaler")
	errNotTracesMarshaler = errors.New("extension is not a traces marshaler")
)

var validSignals = []signalType{signalTypeLogs, signalTypeTraces}

type storageExporter struct {
	cfg             *Config
	logsMarshaler   plog.Marshaler
	tracesMarshaler ptrace.Marshaler
	storageClient   *storage.Client
	bucketHandle    *storage.BucketHandle
	logger          *zap.Logger
	partitionFormat *strftime.Strftime
	gzipWriterPool  *sync.Pool
	zstdWriterPool  *sync.Pool
	signal          signalType
}

var (
	_ exporter.Logs   = (*storageExporter)(nil)
	_ exporter.Traces = (*storageExporter)(nil)
)

func newGCSExporter(
	ctx context.Context,
	cfg *Config,
	logger *zap.Logger,
	signal signalType,
) (*storageExporter, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func newStorageExporter(
	ctx context.Context,
	cfg *Config,
	getZone func(context.Context) (string, error),
	getProjectID func(context.Context) (string, error),
	logger *zap.Logger,
	signal signalType,
) (*storageExporter, error) {
	_ = "STUB: not implemented"
	// Validate signal type
	return nil, nil
}

// valid

// should not happen here, prevented by config.Validate

// Create a new gzip writer that writes to a dummy buffer initially
// It will be reset to the actual destination when used

// Create a new zstd writer that writes to a dummy buffer initially
// It will be reset to the actual destination when used

// Return nil on error - sync.Pool will handle this gracefully

func (s *storageExporter) Start(ctx context.Context, host component.Host) error {
	_ = "STUB: not implemented"
	// Initialize default marshalers
	return nil
}

// Load encoding extension if configured

// TODO Add option for authenticator

// Check if bucket exists without attempting to create it

// Return error if it's a permission issue or network failure

// Attempt to create the bucket

func (s *storageExporter) Shutdown(_ context.Context) error { _ = "STUB: not implemented"; return nil }

func (*storageExporter) Capabilities() consumer.Capabilities {
	_ = "STUB: not implemented"
	return *new(consumer.Capabilities)
}

func (s *storageExporter) ConsumeLogs(ctx context.Context, ld plog.Logs) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *storageExporter) ConsumeTraces(ctx context.Context, td ptrace.Traces) error {
	_ = "STUB: not implemented"
	return nil
}

// generateFilename returns the name of the file to be uploaded.
// It starts from a unique ID, and prepends the partitionFormat and the prefix to it.
func generateFilename(
	uniqueID, filePrefix, partitionPrefix string,
	compression configcompression.Type,
	partitionFormat *strftime.Strftime,
	now time.Time,
) string {
	_ = "STUB: not implemented"
	return ""
}

// Add compression extension

func (s *storageExporter) uploadFile(ctx context.Context, content []byte) (err error) {
	_ = "STUB: not implemented"
	return nil
}

// Compress the content if compression is configured

// if we have multiple files coming, we need to make sure the name is unique so they do
// not overwrite each other

func (s *storageExporter) compressContent(raw []byte) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func compress[T poolItem](pool *sync.Pool, raw []byte, newItem func(io.Writer) (T, error)) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Get writer from pool or create new one

// Write the data

// Always close to release resources, but ignore close error on write failure

// Close the writer

// Only return the writer to the pool after successful write and close

// loadExtension tries to load an available extension for the given id.
func loadExtension[T any](host component.Host, id component.ID, _ string, errNotMarshaler error) (T, error) {
	_ = "STUB: not implemented"
	return *new(T), nil
}
