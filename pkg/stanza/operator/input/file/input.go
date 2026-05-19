// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package file // import "github.com/open-telemetry/opentelemetry-collector-contrib/pkg/stanza/operator/input/file"

import (
	"context"

	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/stanza/entry"
	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/stanza/fileconsumer"
	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/stanza/operator"
	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/stanza/operator/helper"
)

type toBodyFunc func([]byte) any

// Input is an operator that monitors files for entries
type Input struct {
	helper.InputOperator

	fileConsumer *fileconsumer.Manager

	toBody                  toBodyFunc
	includeFileRecordNumber bool
	includeFileRecordOffset bool
}

// Start will start the file monitoring process
func (i *Input) Start(persister operator.Persister) error { _ = "STUB: not implemented"; return nil }

// Stop will stop the file monitoring process
func (i *Input) Stop() error { _ = "STUB: not implemented"; return nil }

func (i *Input) emitBatch(ctx context.Context, tokens [][]byte, attributes map[string]any, lastRecordNumber int64, offsets []int64) error {
	_ = "STUB: not implemented"
	return nil
}

func (i *Input) convertTokens(tokens [][]byte, attributes map[string]any, lastRecordNumber int64, offsets []int64) ([]*entry.Entry, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
