// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package file // import "github.com/open-telemetry/opentelemetry-collector-contrib/pkg/stanza/operator/output/file"

import (
	"context"
	"encoding/json"
	"os"
	"sync"
	"text/template"

	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/stanza/entry"
	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/stanza/operator"
	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/stanza/operator/helper"
)

// Output is an operator that writes logs to a file.
type Output struct {
	helper.OutputOperator

	path    string
	tmpl    *template.Template
	encoder *json.Encoder
	file    *os.File
	mux     sync.Mutex
}

// Start will open the output file.
func (o *Output) Start(_ operator.Persister) error { _ = "STUB: not implemented"; return nil }

// Stop will close the output file.
func (o *Output) Stop() error { _ = "STUB: not implemented"; return nil }

func (o *Output) ProcessBatch(ctx context.Context, entries []*entry.Entry) error {
	_ = "STUB: not implemented"
	return nil
}

// Process will write an entry to the output file.
func (o *Output) Process(_ context.Context, e *entry.Entry) error {
	_ = "STUB: not implemented"
	return nil
}
