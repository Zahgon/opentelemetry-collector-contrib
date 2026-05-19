// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package pathtest // import "github.com/open-telemetry/opentelemetry-collector-contrib/pkg/ottl/contexts/internal/pathtest"

import (
	"context"

	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/ottl"
)

var _ ottl.Path[any] = &Path[any]{}

type Path[K any] struct {
	C        string
	N        string
	KeySlice []ottl.Key[K]
	NextPath *Path[K]
	FullPath string
}

func (p *Path[K]) Name() string { _ = "STUB: not implemented"; return "" }

func (p *Path[K]) Context() string { _ = "STUB: not implemented"; return "" }

func (p *Path[K]) Next() ottl.Path[K] { _ = "STUB: not implemented"; return nil }

func (p *Path[K]) Keys() []ottl.Key[K] { _ = "STUB: not implemented"; return nil }

func (p *Path[K]) String() string { _ = "STUB: not implemented"; return "" }

var _ ottl.Key[any] = &Key[any]{}

type Key[K any] struct {
	S *string
	I *int64
	G ottl.Getter[K]
}

func (k *Key[K]) String(_ context.Context, _ K) (*string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (k *Key[K]) Int(_ context.Context, _ K) (*int64, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (k *Key[K]) ExpressionGetter(_ context.Context, _ K) (ottl.Getter[K], error) {
	_ = "STUB: not implemented"
	return nil, nil
}
