// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package testutil // import "github.com/open-telemetry/opentelemetry-collector-contrib/pkg/stanza/testutil"

import (
	"context"
	"sync"

	"go.opentelemetry.io/collector/extension/xextension/storage"

	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/stanza/operator"
)

type mockPersister struct {
	data    map[string][]byte
	dataMux sync.Mutex
	errKeys map[string]error
}

func (p *mockPersister) Get(_ context.Context, k string) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (p *mockPersister) Set(_ context.Context, k string, v []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func (p *mockPersister) Delete(_ context.Context, k string) error {
	_ = "STUB: not implemented"
	return nil
}

func (p *mockPersister) Batch(_ context.Context, ops ...*storage.Operation) error {
	_ = "STUB: not implemented"
	return nil
}

// NewUnscopedMockPersister will return a new persister for testing
func NewUnscopedMockPersister() operator.Persister {
	_ = "STUB: not implemented"
	return *new(operator.Persister)
}

func NewMockPersister(scope string) operator.Persister {
	_ = "STUB: not implemented"
	return *new(operator.Persister)
}

// NewErrPersister will return a new persister for testing
// which will return an error if any of the specified keys are used
func NewErrPersister(errKeys map[string]error) operator.Persister {
	_ = "STUB: not implemented"
	return *new(operator.Persister)
}

// Trim removes white space from the lines of a string
func Trim(s string) string { _ = "STUB: not implemented"; return "" }
