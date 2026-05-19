// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package filter // import "github.com/open-telemetry/opentelemetry-collector-contrib/pkg/stanza/fileconsumer/matcher/internal/filter"
import (
	"time"
)

type excludeOlderThanOption struct {
	age time.Duration
}

func (eot excludeOlderThanOption) apply(items []*item) ([]*item, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Keep (include) the file if its age (since last modification)
// is the same or less than the configured age.

// ExcludeOlderThan excludes files whose modification time is older than the specified age.
func ExcludeOlderThan(age time.Duration) Option { _ = "STUB: not implemented"; return *new(Option) }
