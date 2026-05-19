// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package header // import "github.com/open-telemetry/opentelemetry-collector-contrib/pkg/stanza/fileconsumer/internal/header"

import (
	"context"
	"errors"

	"go.opentelemetry.io/collector/component"

	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/stanza/pipeline"
)

var ErrEndOfHeader = errors.New("end of header")

type Reader struct {
	set      component.TelemetrySettings
	cfg      Config
	pipeline pipeline.Pipeline
	output   *pipelineOutput
}

func NewReader(set component.TelemetrySettings, cfg Config) (*Reader, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Process checks if the given token is a line of the header, and consumes it if it is.
// An EndOfHeaderError is returned if the given line was not a header line.
func (r *Reader) Process(ctx context.Context, token string, fileAttributes map[string]any) error {
	_ = "STUB: not implemented"
	return nil
}

// Do not return yet. An entry was added to the logsChan which must be consumed generically.

// Copy resultant attributes over current set of attributes (upsert)
// fileAttributes is an output parameter

func (r *Reader) Stop() error { _ = "STUB: not implemented"; return nil }
