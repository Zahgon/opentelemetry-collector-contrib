// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package generate // import "github.com/open-telemetry/opentelemetry-collector-contrib/pkg/stanza/operator/input/generate"

import (
	"context"
	"sync"

	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/stanza/entry"
	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/stanza/operator"
	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/stanza/operator/helper"
)

// Input is an operator that generates log entries.
type Input struct {
	helper.InputOperator
	entry  entry.Entry
	count  int
	static bool
	wg     sync.WaitGroup
	cancel context.CancelFunc
}

// Start will start generating log entries.
func (i *Input) Start(_ operator.Persister) error { _ = "STUB: not implemented"; return nil }

// Stop will stop generating logs.
func (i *Input) Stop() error { _ = "STUB: not implemented"; return nil }

func recursiveMapInterfaceToMapString(m any) any { _ = "STUB: not implemented"; return *new(any) }
