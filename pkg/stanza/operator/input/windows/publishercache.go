// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

//go:build windows

package windows // import "github.com/open-telemetry/opentelemetry-collector-contrib/pkg/stanza/operator/input/windows"

type publisherCache struct {
	cache map[string]Publisher
}

func newPublisherCache() publisherCache { _ = "STUB: not implemented"; return *new(publisherCache) }

func (c *publisherCache) get(provider string) (Publisher, error) {
	_ = "STUB: not implemented"
	return *new(Publisher), nil
}

// If the provider is empty, there is nothing to be formatted on the event
// keep the invalid publisher in the cache. See issue #35135

// Always store the publisher even if there was an error opening it.

func (c *publisherCache) evictAll() error { _ = "STUB: not implemented"; return nil }
