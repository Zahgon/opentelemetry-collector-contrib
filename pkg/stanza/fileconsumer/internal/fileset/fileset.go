// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package fileset // import "github.com/open-telemetry/opentelemetry-collector-contrib/pkg/stanza/fileconsumer/internal/fileset"

import (
	"errors"

	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/stanza/fileconsumer/internal/fingerprint"
	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/stanza/fileconsumer/internal/reader"
)

var errFilesetEmpty = errors.New("pop() on empty Fileset")

var (
	_ Matchable = (*reader.Reader)(nil)
	_ Matchable = (*reader.Metadata)(nil)
)

type Matchable interface {
	GetFingerprint() *fingerprint.Fingerprint
}

type Fileset[T Matchable] struct {
	readers []T
}

func New[T Matchable](capacity int) *Fileset[T] { _ = "STUB: not implemented"; return nil }

func (set *Fileset[T]) Len() int { _ = "STUB: not implemented"; return 0 }

func (set *Fileset[T]) Get() []T { _ = "STUB: not implemented"; return nil }

func (set *Fileset[T]) Pop() (T, error) {
	_ = "STUB: not implemented"
	// return first element from the array and remove it
	return *new(T), nil
}

func (set *Fileset[T]) Add(readers ...T) {
	_ = "STUB: not implemented"
	// add open readers
	return
}

func (set *Fileset[T]) Match(fp *fingerprint.Fingerprint, cmp func(a, b *fingerprint.Fingerprint) bool) T {
	_ = "STUB: not implemented"
	return *new(T)
}

// comparators
func StartsWith(a, b *fingerprint.Fingerprint) bool { _ = "STUB: not implemented"; return false }

func Equal(a, b *fingerprint.Fingerprint) bool { _ = "STUB: not implemented"; return false }
