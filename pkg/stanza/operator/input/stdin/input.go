// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package stdin // import "github.com/open-telemetry/opentelemetry-collector-contrib/pkg/stanza/operator/input/stdin"

import (
	"context"
	"os"
	"sync"

	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/stanza/operator"
	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/stanza/operator/helper"
)

// Input is an operator that reads input from stdin
type Input struct {
	helper.InputOperator
	wg     sync.WaitGroup
	cancel context.CancelFunc
	stdin  *os.File
}

// Start will start generating log entries.
func (i *Input) Start(_ operator.Persister) error { _ = "STUB: not implemented"; return nil }

// Stop will stop generating logs.
func (i *Input) Stop() error { _ = "STUB: not implemented"; return nil }
