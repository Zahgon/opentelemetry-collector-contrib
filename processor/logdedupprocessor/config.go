// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

// Package logdedupprocessor provides a processor that counts logs as metrics.
package logdedupprocessor // import "github.com/open-telemetry/opentelemetry-collector-contrib/processor/logdedupprocessor"

import (
	"errors"
	"time"

	"go.opentelemetry.io/collector/component"
)

// Config defaults
const (
	// defaultInterval is the default export interval.
	defaultInterval = 10 * time.Second

	// defaultLogCountAttribute is the default log count attribute
	defaultLogCountAttribute = "log_count"

	// defaultTimezone is the default timezone
	defaultTimezone = "UTC"

	// bodyField is the name of the body field
	bodyField = "body"

	// attributeField is the name of the attribute field
	attributeField = "attributes"
)

// Config errors
var (
	errInvalidLogCountAttribute = errors.New("log_count_attribute must be set")
	errInvalidInterval          = errors.New("interval must be greater than 0")
	errCannotExcludeBody        = errors.New("cannot exclude the entire body")
	errCannotIncludeBody        = errors.New("cannot include the entire body")
)

// Config is the config of the processor.
type Config struct {
	LogCountAttribute string        `mapstructure:"log_count_attribute"`
	Interval          time.Duration `mapstructure:"interval"`
	Timezone          string        `mapstructure:"timezone"`
	ExcludeFields     []string      `mapstructure:"exclude_fields"`
	IncludeFields     []string      `mapstructure:"include_fields"`
	Conditions        []string      `mapstructure:"conditions"`
}

// createDefaultConfig returns the default config for the processor.
func createDefaultConfig() component.Config {
	_ = "STUB: not implemented"
	return *new(component.Config)
}

// Validate validates the configuration
func (c Config) Validate() error { _ = "STUB: not implemented"; return nil }

// validateExcludeFields validates that all the exclude fields
func (c Config) validateExcludeFields() error { _ = "STUB: not implemented"; return nil }

// Special check to make sure the entire body is not excluded

// Split and ensure the field starts with `body` or `attributes`

// If a field is valid make sure we haven't already seen it

// validateIncludeFields validates that all the exclude fields
func (c Config) validateIncludeFields() error { _ = "STUB: not implemented"; return nil }

// Special check to make sure the entire body is not included

// Split and ensure the field starts with `body` or `attributes`

// If a field is valid make sure we haven't already seen it
