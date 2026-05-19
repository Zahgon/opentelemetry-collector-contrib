// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package operator // import "github.com/open-telemetry/opentelemetry-collector-contrib/pkg/stanza/operator"

import (
	"context"

	"go.opentelemetry.io/collector/extension/xextension/storage"
)

// Persister is an interface used to persist data
type Persister interface {
	Get(context.Context, string) ([]byte, error)
	Set(context.Context, string, []byte) error
	Delete(context.Context, string) error
	Batch(ctx context.Context, ops ...*storage.Operation) error
}

type scopedPersister struct {
	Persister
	scope string
}

func NewScopedPersister(s string, p Persister) Persister {
	_ = "STUB: not implemented"
	return *new(Persister)
}

func (p scopedPersister) Get(ctx context.Context, key string) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (p scopedPersister) Set(ctx context.Context, key string, value []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func (p scopedPersister) Delete(ctx context.Context, key string) error {
	_ = "STUB: not implemented"
	return nil
}

func (p scopedPersister) Batch(ctx context.Context, ops ...*storage.Operation) error {
	_ = "STUB: not implemented"
	return nil
}
