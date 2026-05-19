// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package schemaprocessor // import "github.com/open-telemetry/opentelemetry-collector-contrib/processor/schemaprocessor"

import (
	"errors"
	"time"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/config/confighttp"
)

var (
	errRequiresTargets          = errors.New("requires schema targets")
	errDuplicateTargets         = errors.New("duplicate targets detected")
	errMigrationTargetNotFound  = errors.New("migration target does not match any configured target")
	errMigrationFamilyMismatch  = errors.New("migration from and target must be in the same schema family")
	errMigrationDuplicateTarget = errors.New("duplicate migration entry for same target")
	errMigrationRequiresFrom    = errors.New("migration entry requires a from schema URL")
)

// Config defines the user provided values for the Schema Processor
type Config struct {
	confighttp.ClientConfig `mapstructure:",squash"`

	// CacheCooldown is the duration to wait before retrying schema fetches
	// after the retry limit has been reached. Defaults to 5 minutes.
	CacheCooldown time.Duration `mapstructure:"cache_cooldown"`

	// CacheRetryLimit is the number of consecutive failed schema fetch
	// attempts allowed before enforcing the cooldown period. Defaults to 5.
	CacheRetryLimit int `mapstructure:"cache_retry_limit"`

	// PreCache is a list of schema URLs that are downloaded
	// and cached at the start of the collector runtime
	// in order to avoid fetching data that later on could
	// block processing of signals. (Optional field)
	Prefetch []string `mapstructure:"prefetch"`

	// Targets define what schema families should be
	// translated to, allowing older and newer formats
	// to conform to the target schema identifier.
	Targets []string `mapstructure:"targets"`

	// StorageID is an optional storage extension used to persist fetched
	// schema files across collector restarts. When set, fetched schemas
	// are saved to persistent storage and loaded on startup, avoiding
	// unnecessary HTTP fetches.
	StorageID *component.ID `mapstructure:"storage"`

	// Migration defines migration entries that preserve original attributes
	// alongside renamed ones during schema translation. Each entry specifies
	// a target and the version being migrated from. Only renames between
	// the from version and the target version are copied.
	Migration []MigrationEntry `mapstructure:"migration"`
}

// MigrationEntry defines a migration for a specific target schema.
type MigrationEntry struct {
	// Target is the schema URL that this migration applies to.
	// Must match one of the configured targets exactly.
	Target string `mapstructure:"target"`

	// From is the schema URL of the version that operators are migrating
	// away from. Must be in the same schema family as Target.
	// Renames between this version and the target are copied
	// (both old and new attribute names are preserved).
	From string `mapstructure:"from"`
}

func (c *Config) Validate() error { _ = "STUB: not implemented"; return nil }

// Not strictly needed since it would just pass on
// any data that doesn't match targets, however defining
// this processor with no targets is wasteful.
