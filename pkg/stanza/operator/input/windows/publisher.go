// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

//go:build windows

package windows // import "github.com/open-telemetry/opentelemetry-collector-contrib/pkg/stanza/operator/input/windows"

// Publisher is a windows event metadata publisher.
type Publisher struct {
	handle uintptr
}

// Open will open the publisher handle using the supplied provider.
func (p *Publisher) Open(provider string) error { _ = "STUB: not implemented"; return nil }

func (p *Publisher) Valid() bool { _ = "STUB: not implemented"; return false }

// Close will close the publisher handle.
func (p *Publisher) Close() error { _ = "STUB: not implemented"; return nil }

// NewPublisher will create a new publisher with an empty handle.
func NewPublisher() Publisher { _ = "STUB: not implemented"; return *new(Publisher) }
