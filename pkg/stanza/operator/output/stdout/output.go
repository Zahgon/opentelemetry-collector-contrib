// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package stdout // import "github.com/open-telemetry/opentelemetry-collector-contrib/pkg/stanza/operator/output/stdout"

import (
	"context"
	"encoding/json"
	"sync"

	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/stanza/entry"
	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/stanza/operator/helper"
)

// Output is an operator that logs entries using stdout.
type Output struct {
	helper.OutputOperator
	encoder *json.Encoder
	mux     sync.Mutex
}

func (o *Output) ProcessBatch(ctx context.Context, entries []*entry.Entry) error {
	_ = "STUB: not implemented"
	return nil
}

// Process will log entries received.
func (o *Output) Process(_ context.Context, e *entry.Entry) error {
	_ = "STUB: not implemented"
	return nil
}
