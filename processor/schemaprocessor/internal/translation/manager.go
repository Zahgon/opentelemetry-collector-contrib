// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package translation // import "github.com/open-telemetry/opentelemetry-collector-contrib/processor/schemaprocessor/internal/translation"

import (
	"context"
	"errors"
	"sync"
	"time"

	"go.uber.org/zap"

	"github.com/open-telemetry/opentelemetry-collector-contrib/processor/schemaprocessor/internal/metadata"
)

var errNilValueProvided = errors.New("nil value provided")

// Manager is responsible for ensuring that schemas are kept up to date
// with the most recent version that are requested.
type Manager interface {
	// RequestTranslation will provide either the defined Translation
	// if it is a known target and able to retrieve from Provider
	// otherwise it will return an error.
	RequestTranslation(ctx context.Context, schemaURL string) (Translation, error)

	// AddProvider will add a provider to the Manager
	AddProvider(p Provider)
}

type manager struct {
	log              *zap.Logger
	telemetryBuilder *metadata.TelemetryBuilder
	migrationMap     map[string]*Version // target schema URL → copyFromVersion
	cooldown         time.Duration
	limit            int

	rw            sync.RWMutex
	providers     []Provider
	match         map[string]*Version
	translatorMap map[string]*translator
}

var _ Manager = (*manager)(nil)

// NewManager creates a manager that will allow for management
// of schema
func NewManager(targetSchemaURLS []string, log *zap.Logger, cooldown time.Duration, limit int, telemetryBuilder *metadata.TelemetryBuilder, migrationMap map[string]*Version, providers ...Provider) (Manager, error) {
	_ = "STUB: not implemented"
	return *new(Manager), nil
}

// wrap providers with cacheable provider

// newCacheableProvider wraps p with a CacheableProvider wired to the manager's telemetry.
func (m *manager) newCacheableProvider(p Provider) Provider {
	_ = "STUB: not implemented"
	return *new(Provider)
}

func (m *manager) RequestTranslation(ctx context.Context, schemaURL string) (Translation, error) {
	_ = "STUB: not implemented"
	return *new(Translation), nil
}

// Always fetch the schema file for the higher version, since it contains the complete
// migration history. For upgrades (signal older than target), fetch the target URL.
// For downgrades (signal newer than target), fetch the signal URL.

// If we fail to retrieve the schema, we should
// try the next provider

// AddProvider will add a provider to the Manager
func (m *manager) AddProvider(p Provider) { _ = "STUB: not implemented"; return }
