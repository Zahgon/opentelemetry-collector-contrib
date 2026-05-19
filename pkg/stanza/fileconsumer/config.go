// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

//go:generate make mdatagen

package fileconsumer // import "github.com/open-telemetry/opentelemetry-collector-contrib/pkg/stanza/fileconsumer"

import (
	"bufio"
	"time"

	"go.opentelemetry.io/collector/component"

	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/stanza/fileconsumer/attrs"
	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/stanza/fileconsumer/emit"
	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/stanza/fileconsumer/matcher"
	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/stanza/operator"
	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/stanza/operator/helper"
	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/stanza/split"
	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/stanza/trim"
)

const (
	defaultMaxConcurrentFiles = 1024
	defaultEncoding           = "utf-8"
	defaultPollInterval       = 200 * time.Millisecond

	// MaxLogSizeBehaviorSplit splits oversized log entries into multiple log entries.
	MaxLogSizeBehaviorSplit = "split"
	// MaxLogSizeBehaviorTruncate truncates oversized log entries and drops the remainder.
	MaxLogSizeBehaviorTruncate = "truncate"
)

// OnTruncate defines the behavior when a file with the same fingerprint
// is detected but with a smaller size (indicating a copytruncate rotation).
const (
	// OnTruncateIgnore keeps the current behavior: do not read any data until
	// the file grows past the original offset, then read only the new data.
	OnTruncateIgnore = "ignore"
	// OnTruncateReadWholeFile reads the whole file from the beginning when truncation is detected.
	OnTruncateReadWholeFile = "read_whole_file"
	// OnTruncateReadNew stores the new (lower) offset, so that the next time the file grows,
	// any new data past the new offset is read.
	OnTruncateReadNew = "read_new"
)

const defaultOnTruncate = OnTruncateIgnore

// NewConfig creates a new input config with default values
func NewConfig() *Config { _ = "STUB: not implemented"; return nil }

// Config is the configuration of a file input operator
type Config struct {
	matcher.Criteria        `mapstructure:",squash"`
	attrs.Resolver          `mapstructure:",squash"`
	PollInterval            time.Duration   `mapstructure:"poll_interval,omitempty"`
	MaxConcurrentFiles      int             `mapstructure:"max_concurrent_files,omitempty"`
	MaxBatches              int             `mapstructure:"max_batches,omitempty"`
	StartAt                 string          `mapstructure:"start_at,omitempty"`
	FingerprintSize         helper.ByteSize `mapstructure:"fingerprint_size,omitempty"`
	InitialBufferSize       helper.ByteSize `mapstructure:"initial_buffer_size,omitempty"`
	MaxLogSize              helper.ByteSize `mapstructure:"max_log_size,omitempty"`
	MaxLogSizeBehavior      string          `mapstructure:"max_log_size_behavior,omitempty"`
	Encoding                string          `mapstructure:"encoding,omitempty"`
	SplitConfig             split.Config    `mapstructure:"multiline,omitempty"`
	TrimConfig              trim.Config     `mapstructure:",squash,omitempty"`
	FlushPeriod             time.Duration   `mapstructure:"force_flush_period,omitempty"`
	Header                  *HeaderConfig   `mapstructure:"header,omitempty"`
	DeleteAfterRead         bool            `mapstructure:"delete_after_read,omitempty"`
	IncludeFileRecordNumber bool            `mapstructure:"include_file_record_number,omitempty"`
	IncludeFileRecordOffset bool            `mapstructure:"include_file_record_offset,omitempty"`
	Compression             string          `mapstructure:"compression,omitempty"`
	PollsToArchive          int             `mapstructure:"polls_to_archive,omitempty"`
	AcquireFSLock           bool            `mapstructure:"acquire_fs_lock,omitempty"`
	OnTruncate              string          `mapstructure:"on_truncate,omitempty"`
}

type HeaderConfig struct {
	Pattern           string            `mapstructure:"pattern"`
	MetadataOperators []operator.Config `mapstructure:"metadata_operators"`
}

func (c Config) Build(set component.TelemetrySettings, emit emit.Callback, opts ...Option) (*Manager, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c Config) validate() error { _ = "STUB: not implemented"; return nil }

// Valid values

type options struct {
	splitFunc  bufio.SplitFunc
	noTracking bool
}

type Option func(*options)

// WithSplitFunc overrides the split func which is normally built from other settings on the config
func WithSplitFunc(f bufio.SplitFunc) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithNoTracking forces the readerFactory to not keep track of files in memory. When used, the reader will
// read from the beginning of each file every time it is polled.
func WithNoTracking() Option { _ = "STUB: not implemented"; return *new(Option) }
