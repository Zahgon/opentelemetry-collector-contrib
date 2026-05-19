// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

// Package provider contains the cluster name provider
package provider // import "github.com/open-telemetry/opentelemetry-collector-contrib/internal/datadog/hostmetadata/provider"

import (
	"context"
	"sync"
	"time"

	"github.com/DataDog/datadog-agent/pkg/opentelemetry-mapping-go/otlp/attributes/source"
	"go.uber.org/zap"
)

// SourceAliasesProvider is a source.Provider that also returns host aliases alongside the main source.
type SourceAliasesProvider interface {
	source.Provider
	SourceWithAliases(ctx context.Context) (source.Source, []string, error)
}

var _ SourceAliasesProvider = (*chainProvider)(nil)

type chainProvider struct {
	logger       *zap.Logger
	providers    map[string]source.Provider
	priorityList []string
	aliasedList  []string
	timeout      time.Duration
}

func (p *chainProvider) SourceWithAliases(ctx context.Context) (source.Source, []string, error) {
	_ = "STUB: not implemented"
	// Auxiliary type for storing source provider replies
	return *new(source.Source), nil, nil
}

// Cancel all providers when exiting

// Make a different context for our provider calls, to differentiate between a provider timing out and the entire
// context being cancelled

// Run all providers in parallel

// Capacity required to avoid leaking goroutines / blocking aliasesWg

// Check provider responses in order to ensure priority

func (p *chainProvider) Source(ctx context.Context) (source.Source, error) {
	_ = "STUB: not implemented"
	return *new(source.Source), nil
}

// Chain providers into a single provider that returns the first available hostname.
// aliasedList contains providers whose hostname results are always awaited and added as aliases
// when not chosen as the main source.
func Chain(logger *zap.Logger, providers map[string]source.Provider, priorityList, aliasedList []string, timeout time.Duration) (SourceAliasesProvider, error) {
	_ = "STUB: not implemented"
	return *new(SourceAliasesProvider), nil
}

var _ source.Provider = (*configProvider)(nil)

type configProvider struct {
	hostname string
}

func (p *configProvider) Source(context.Context) (source.Source, error) {
	_ = "STUB: not implemented"
	return *new(source.Source), nil
}

// Config returns fixed hostname.
func Config(hostname string) source.Provider {
	_ = "STUB: not implemented"
	return *new(source.Provider)
}

var _ SourceAliasesProvider = (*onceProvider)(nil)

type onceProvider struct {
	once     sync.Once
	src      source.Source
	aliases  []string
	err      error
	provider SourceAliasesProvider
}

func (c *onceProvider) SourceWithAliases(ctx context.Context) (source.Source, []string, error) {
	_ = "STUB: not implemented"
	return *new(source.Source), nil, nil
}

func (c *onceProvider) Source(ctx context.Context) (source.Source, error) {
	_ = "STUB: not implemented"
	return *new(source.Source), nil
}

// Once wraps a provider to call it only once.
func Once(provider SourceAliasesProvider) SourceAliasesProvider {
	_ = "STUB: not implemented"
	return *new(SourceAliasesProvider)
}
